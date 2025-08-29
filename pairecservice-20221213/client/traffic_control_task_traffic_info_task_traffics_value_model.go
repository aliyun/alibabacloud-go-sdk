// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrafficControlTaskTrafficInfoTaskTrafficsValue interface {
	dara.Model
	String() string
	GoString() string
	SetTraffic(v float64) *TrafficControlTaskTrafficInfoTaskTrafficsValue
	GetTraffic() *float64
}

type TrafficControlTaskTrafficInfoTaskTrafficsValue struct {
	Traffic *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s TrafficControlTaskTrafficInfoTaskTrafficsValue) String() string {
	return dara.Prettify(s)
}

func (s TrafficControlTaskTrafficInfoTaskTrafficsValue) GoString() string {
	return s.String()
}

func (s *TrafficControlTaskTrafficInfoTaskTrafficsValue) GetTraffic() *float64 {
	return s.Traffic
}

func (s *TrafficControlTaskTrafficInfoTaskTrafficsValue) SetTraffic(v float64) *TrafficControlTaskTrafficInfoTaskTrafficsValue {
	s.Traffic = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTaskTrafficsValue) Validate() error {
	return dara.Validate(s)
}
