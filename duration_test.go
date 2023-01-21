package jsdur

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDuration_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   Duration
		want    string
		wantErr bool
	}{
		{
			name:    "simple",
			input:   NewDuration(10 * time.Hour),
			want:    "10h0m0s",
			wantErr: false,
		},
		{
			name:    "complex",
			input:   NewDuration(5*time.Hour + 12*time.Minute + 47*time.Second),
			want:    "5h12m47s",
			wantErr: false,
		},
		{
			name:    "zero",
			input:   Duration{},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestDuration_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Duration
		wantErr bool
	}{
		{
			name:    "simple",
			input:   "24h",
			want:    NewDuration(24 * time.Hour),
			wantErr: false,
		},
		{
			name:    "complex",
			input:   "2h15m3s99ms",
			want:    NewDuration(2*time.Hour + 15*time.Minute + 3*time.Second + 99*time.Millisecond),
			wantErr: false,
		},
		{
			name:    "zero",
			input:   "0",
			want:    NewDuration(0),
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    NewDuration(0),
			wantErr: false,
		},
		{
			name:    "error",
			input:   "invalid",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Duration{}
			err := got.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Log(err)
			}
		})
	}
}

func TestDuration_MarshalJSON(t *testing.T) {
	type testStruct struct {
		Duration Duration `json:"duration"`
	}
	tests := []struct {
		name    string
		input   testStruct
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			input: testStruct{
				Duration: NewDuration(10 * time.Minute),
			},
			want:    `{"duration":"10m0s"}`,
			wantErr: false,
		},
		{
			name: "complex",
			input: testStruct{
				Duration: NewDuration(1000*time.Hour + 10*time.Minute + 5*time.Second),
			},
			want:    `{"duration":"1000h10m5s"}`,
			wantErr: false,
		},
		{
			name: "zero",
			input: testStruct{
				Duration: NewDuration(0),
			},
			want:    `{"duration":""}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	type testStruct struct {
		Duration Duration `json:"duration"`
	}
	tests := []struct {
		name    string
		input   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "simple",
			input: `{"duration":"24h"}`,
			want: testStruct{
				Duration: NewDuration(24 * time.Hour),
			},
			wantErr: false,
		},
		{
			name:  "complex",
			input: `{"duration":"2h15m3s99ms"}`,
			want: testStruct{
				Duration: NewDuration(2*time.Hour + 15*time.Minute + 3*time.Second + 99*time.Millisecond),
			},
			wantErr: false,
		},
		{
			name:  "zero",
			input: `{"duration":"0"}`,
			want: testStruct{
				Duration: NewDuration(0),
			},
			wantErr: false,
		},
		{
			name:  "empty string",
			input: `{}`,
			want: testStruct{
				Duration: NewDuration(0),
			},
			wantErr: false,
		},
		{
			name:    "error",
			input:   `{"duration":"invalid"}`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := testStruct{}
			err := json.Unmarshal([]byte(tt.input), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Log(err)
			}
		})
	}
}
