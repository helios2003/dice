package websocket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoAdd(t *testing.T) {
	exec := NewWebsocketCommandExecutor()
	conn := exec.ConnectToServer()

	testCases := []struct {
		name   string
		cmds   []string
		expect []interface{}
	}{
		{
			name:   "GeoAdd With Wrong Number of Arguments",
			cmds:   []string{"GEOADD mygeo 1 2"},
			expect: []interface{}{"ERR wrong number of arguments for 'geoadd' command"},
		},
		{
			name:   "GeoAdd With Adding New Member And Updating it",
			cmds:   []string{"GEOADD mygeo 1.21 1.44 NJ", "GEOADD mygeo 1.22 1.54 NJ"},
			expect: []interface{}{float64(1), float64(0)},
		},
		{
			name:   "GeoAdd With Adding New Member And Updating it with NX",
			cmds:   []string{"GEOADD mygeo NX 1.21 1.44 MD", "GEOADD mygeo 1.22 1.54 MD"},
			expect: []interface{}{float64(1), float64(0)},
		},
		{
			name:   "GEOADD with both NX and XX options",
			cmds:   []string{"GEOADD mygeo NX XX 1.21 1.44 DEL"},
			expect: []interface{}{"ERR XX and NX options at the same time are not compatible"},
		},
		{
			name:   "GEOADD invalid longitude",
			cmds:   []string{"GEOADD mygeo 181.0 1.44 MD"},
			expect: []interface{}{"ERR invalid longitude"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.cmds {
				result, err := exec.FireCommandAndReadResponse(conn, cmd)
				assert.Nil(t, err)
				assert.Equal(t, tc.expect[i], result, "Value mismatch for cmd %s", cmd)
			}
		})
	}
}

func TestGeoDist(t *testing.T) {
	exec := NewWebsocketCommandExecutor()
	conn := exec.ConnectToServer()
	defer conn.Close()

	testCases := []struct {
		name   string
		cmds   []string
		expect []interface{}
	}{
		{
			name: "GEODIST b/w existing points",
			cmds: []string{
				"GEOADD points 13.361389 38.115556 Palermo",
				"GEOADD points 15.087269 37.502669 Catania",
				"GEODIST points Palermo Catania",
				"GEODIST points Palermo Catania km",
			},
			expect: []interface{}{float64(1), float64(1), float64(166274.144), float64(166.2741)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.cmds {
				result, err := exec.FireCommandAndReadResponse(conn, cmd)
				assert.Nil(t, err)
				assert.Equal(t, tc.expect[i], result, "Value mismatch for cmd %s", cmd)
			}
		})
	}
}

func TestGeoPos(t *testing.T) {
	exec := NewWebsocketCommandExecutor()
	conn := exec.ConnectToServer()
	defer conn.Close()

	testCases := []struct {
		name   string
		cmds   []string
		expect []interface{}
	}{
		{
			name: "GEOPOS b/w existing points",
			cmds: []string{
				"GEOADD index 13.361389 38.115556 Palermo",
				"GEOPOS index Palermo",
			},
			expect: []interface{}{
				float64(1),
				[]interface{}{float64(13.361387), float64(38.115556)},
			},
		},
		{
			name: "GEOPOS for non existing points",
			cmds: []string{
				"GEOPOS index NonExisting",
			},
			expect: []interface{}{[]interface{}{nil}},
		},
		{
			name: "GEOPOS for non existing index",
			cmds: []string{
				"GEOPOS NonExisting Palermo",
			},
			expect: []interface{}{nil},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i, cmd := range tc.cmds {
				result, err := exec.FireCommandAndReadResponse(conn, cmd)
				assert.Nil(t, err)
				assert.Equal(t, tc.expect[i], result, "Value mismatch for cmd %s", cmd)
			}
		})
	}
}

