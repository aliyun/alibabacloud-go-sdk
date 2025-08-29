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
	SetRecordTime(v int64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue
	GetRecordTime() *int64
}

type TrafficControlTaskTrafficInfoTargetTrafficsDataValue struct {
	Traffic    *float64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
	RecordTime *int64   `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
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

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) GetRecordTime() *int64 {
	return s.RecordTime
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetTraffic(v float64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.Traffic = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) SetRecordTime(v int64) *TrafficControlTaskTrafficInfoTargetTrafficsDataValue {
	s.RecordTime = &v
	return s
}

func (s *TrafficControlTaskTrafficInfoTargetTrafficsDataValue) Validate() error {
	return dara.Validate(s)
}
