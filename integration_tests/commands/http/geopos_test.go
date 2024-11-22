package http

import (
	"strconv"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGeoPosHttp(t *testing.T) {
	exec := NewHTTPCommandExecutor()

	testCases := []struct {
		name 		string
		commands 	[]HTTPCommand
		expected 	[]interface{}
		delays 		time.Duration
	}{
		{
			name: "GEOPOS for existing points",
			commands: []HTTPCommand{
				{Command: "GEOADD", Body: map[string]interface{}{"key": "points", "values": []interface{}{"13.361389", "38.115556", "Palermo"}}},
				{Command: "GEOPOS", Body: map[string]interface{}{"key": "points", "values": []interface{}{"Palermo"}}},
			},
			expected: []interface{}{"13.361389", "38.115556"},
		},
		{
			name: "GEOPOS for non existing points",
			commands: []HTTPCommand{
				{Command: "GEOPOS", Body: map[string]interface{}{"key": "points", "values": []interface{}{"NonExisting"}}},
			},
			expected: []interface{}{nil},
		},
		{
			name: "GEOPOS for non existing index",
			commands: []HTTPCommand{
				{Command: "GEOPOS", Body: map[string]interface{}{"key": "NonExisting", "values": []interface{}{"Palermo"}}},
			},
			expected: clientio.RespNIL,
		},
	}

	for _, tc := range testCases {
		t.run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.commands {
				if tc.delays[i] > 0 {
					time.Sleep(tc.delays[i])
				}
				result, err := exec.FireCommand(cmd)
				assert.Equal(t, tc.expected[i], result)
			}
		})
	}
}