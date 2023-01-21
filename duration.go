package jsdur

import (
	"encoding"
	"encoding/json"
	"fmt"
	"time"
)

var (
	_ encoding.TextUnmarshaler = (*Duration)(nil)
	_ encoding.TextMarshaler   = (*Duration)(nil)
	_ json.Marshaler           = (*Duration)(nil)
	_ json.Unmarshaler         = (*Duration)(nil)
)

// Duration struct embeds time.Duration and implements interfaces
// for json and text marshaling/unmarshalling.
type Duration struct {
	time.Duration
}

func NewDuration(d time.Duration) Duration {
	return Duration{Duration: d}
}

func (d Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *Duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = fromString(string(text))
	if err != nil {
		return fmt.Errorf("from string: %v", err)
	}

	return nil
}
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Duration) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return fmt.Errorf("json unmarshal: %v", err)
	}

	d.Duration, err = fromString(s)
	if err != nil {
		return fmt.Errorf("from string: %v", err)
	}

	return nil
}

func (d Duration) String() string {
	if d.Duration == 0 {
		return ""
	}

	return d.Duration.String()
}

func fromString(s string) (time.Duration, error) {
	if s == "0" || s == "" {
		return 0, nil
	}

	duration, err := time.ParseDuration(s)
	if err != nil {
		return 0, fmt.Errorf("parse duration: %v", err)
	}

	return duration, nil
}
