package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gsttests = []struct {
	name string
	raw  string
	err  string
	msg  GST
}{
	{
		name: "good sentance",
		raw:  "$GNGST,135637.60,6.800,,,,0.360,0.300,0.810*65",
		msg: GST{
			Time:        Time{true, 13, 56, 37, 600},
			RMS:         6.8,
			Major:       0,
			Minor:       0,
			Orientation: 0,
			Latitude:    0.36,
			Longitude:   0.3,
			Height:      0.81,
		},
	},
}

func TestGST(t *testing.T) {
	for _, tt := range gsttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gst := m.(GST)
				gst.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gst)
			}
		})
	}
}
