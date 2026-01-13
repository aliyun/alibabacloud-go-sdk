// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrafficControlTaskTrafficInfoTargetTrafficsDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetTraffic(v float64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue
	GetTraffic() *float64
	SetRecorfTime(v int64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue
	GetRecorfTime() *int64
}

type TrafficControlTaskTrafficInfoTargetTrafficsDataValue struct {
	// example:
	//
	// 80
	Traffic    *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
	RecorfTime *int64   `json:"RecorfTime,omitempty" xml:"RecorfTime,omitempty"`
}

func (s TrafficControlTaskTrafficInfoTargetTrafficsDataValue) String() string {
	return dara.Prettify(s)
}

func (s TrafficControlTaskTrafficInfoTargetTrafficsDataValue) GoString() string {
	return s.String()
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) GetTraffic() *float64 {
	return s.Traffic
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) GetRecorfTime() *int64 {
	return s.RecorfTime
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetTraffic(v float64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.Traffic = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetRecorfTime(v int64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.RecorfTime = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) Validate() error {
	return dara.Validate(s)
}
