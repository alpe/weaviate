//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package ask

import (
	"reflect"
	"testing"
)

func Test_extractAskFn(t *testing.T) {
	type args struct {
		source map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "should parse properly with only question",
			args: args{
				source: map[string]interface{}{
					"question": "some question",
				},
			},
			want: &AskParams{
				Question: "some question",
			},
		},
		{
			name: "should parse properly with question and certainty",
			args: args{
				source: map[string]interface{}{
					"question":  "some question",
					"certainty": 0.8,
				},
			},
			want: &AskParams{
				Question:  "some question",
				Certainty: 0.8,
			},
		},
		{
			name: "should parse properly without params",
			args: args{
				source: map[string]interface{}{},
			},
			want: &AskParams{},
		},
		{
			name: "should parse properly with question and certainty and properties",
			args: args{
				source: map[string]interface{}{
					"question":   "some question",
					"certainty":  0.8,
					"properties": []interface{}{"prop1", "prop2"},
				},
			},
			want: &AskParams{
				Question:   "some question",
				Certainty:  0.8,
				Properties: []string{"prop1", "prop2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractAskFn(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractAskFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
