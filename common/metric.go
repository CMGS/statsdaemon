package common

type Metric struct {
	Bucket   string
	Value    float64
	Modifier string
	Sampling float32
	//optional, only for "archive" metrics
	Time uint32
}
