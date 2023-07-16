package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
    testArg := args{
        headers: http.Header{},
    }
    testArg.headers.Set("Authorization", "ApiKey fakeapikey")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "abc", args: testArg, want: "fakeapikey", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
