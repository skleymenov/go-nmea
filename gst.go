package nmea

const (
	// TypeGST type for GST sentences
	TypeGST = "GST"
)

// GST represents position error statistics
// https://www.trimble.com/oem_receiverhelp/v4.44/en/NMEA-0183messages_GST.html
type GST struct {
	BaseSentence
	Time        Time    // UTC of position fix
	RMS         float64 // RMS value of the pseudorange residuals; includes carrier phase residuals during periods of RTK (float) and RTK (fixed) processing
	Major       float64 // Error ellipse semi-major axis 1 sigma error, in meters
	Minor       float64 // Error ellipse semi-minor axis 1 sigma error, in meters
	Orientation float64 // Error ellipse orientation, degrees from true north
	Latitude    float64 // Latitude 1 sigma error, in meters
	Longitude   float64 // Longitude 1 sigma error, in meters
	Height      float64 // Height 1 sigma error, in meters
}

// newGST constuctor
func newGST(s BaseSentence) (GST, error) {
	p := NewParser(s)
	p.AssertType(TypeGST)
	m := GST{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		RMS:          p.Float64(1, "rms"),
		Major:        p.Float64(2, "major"),
		Minor:        p.Float64(3, "minor"),
		Orientation:  p.Float64(4, "orientation"),
		Latitude:     p.Float64(5, "latitude"),
		Longitude:    p.Float64(6, "longitude"),
		Height:       p.Float64(7, "height"),
	}

	return m, p.Err()
}
