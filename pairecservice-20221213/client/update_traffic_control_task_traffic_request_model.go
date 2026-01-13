// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficControlTaskTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *UpdateTrafficControlTaskTrafficRequest
	GetEnvironment() *string
	SetInstanceId(v string) *UpdateTrafficControlTaskTrafficRequest
	GetInstanceId() *string
	SetTraffics(v []*UpdateTrafficControlTaskTrafficRequestTraffics) *UpdateTrafficControlTaskTrafficRequest
	GetTraffics() []*UpdateTrafficControlTaskTrafficRequestTraffics
	SetNewParam3(v string) *UpdateTrafficControlTaskTrafficRequest
	GetNewParam3() *string
}

type UpdateTrafficControlTaskTrafficRequest struct {
	Environment *string                                           `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Traffics    []*UpdateTrafficControlTaskTrafficRequestTraffics `json:"Traffics,omitempty" xml:"Traffics,omitempty" type:"Repeated"`
	NewParam3   *string                                           `json:"new-param-3,omitempty" xml:"new-param-3,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *UpdateTrafficControlTaskTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateTrafficControlTaskTrafficRequest) GetTraffics() []*UpdateTrafficControlTaskTrafficRequestTraffics {
	return s.Traffics
}

func (s *UpdateTrafficControlTaskTrafficRequest) GetNewParam3() *string {
	return s.NewParam3
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetEnvironment(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.Environment = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetInstanceId(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetTraffics(v []*UpdateTrafficControlTaskTrafficRequestTraffics) *UpdateTrafficControlTaskTrafficRequest {
	s.Traffics = v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) SetNewParam3(v string) *UpdateTrafficControlTaskTrafficRequest {
	s.NewParam3 = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequest) Validate() error {
	if s.Traffics != nil {
		for _, item := range s.Traffics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTrafficControlTaskTrafficRequestTraffics struct {
	ItemOrExperimentId             *string  `json:"ItemOrExperimentId,omitempty" xml:"ItemOrExperimentId,omitempty"`
	RecordTime                     *string  `json:"RecordTime,omitempty" xml:"RecordTime,omitempty"`
	TrafficControlTargetAimTraffic *float64 `json:"TrafficControlTargetAimTraffic,omitempty" xml:"TrafficControlTargetAimTraffic,omitempty"`
	TrafficControlTargetId         *string  `json:"TrafficControlTargetId,omitempty" xml:"TrafficControlTargetId,omitempty"`
	TrafficControlTargetTraffic    *int64   `json:"TrafficControlTargetTraffic,omitempty" xml:"TrafficControlTargetTraffic,omitempty"`
	TrafficControlTaskTraffic      *int64   `json:"TrafficControlTaskTraffic,omitempty" xml:"TrafficControlTaskTraffic,omitempty"`
}

func (s UpdateTrafficControlTaskTrafficRequestTraffics) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficControlTaskTrafficRequestTraffics) GoString() string {
	return s.String()
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetItemOrExperimentId() *string {
	return s.ItemOrExperimentId
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetRecordTime() *string {
	return s.RecordTime
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetTrafficControlTargetAimTraffic() *float64 {
	return s.TrafficControlTargetAimTraffic
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetTrafficControlTargetId() *string {
	return s.TrafficControlTargetId
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetTrafficControlTargetTraffic() *int64 {
	return s.TrafficControlTargetTraffic
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) GetTrafficControlTaskTraffic() *int64 {
	return s.TrafficControlTaskTraffic
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetItemOrExperimentId(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.ItemOrExperimentId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetRecordTime(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.RecordTime = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetAimTraffic(v float64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetAimTraffic = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetId(v string) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetId = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTargetTraffic(v int64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTargetTraffic = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) SetTrafficControlTaskTraffic(v int64) *UpdateTrafficControlTaskTrafficRequestTraffics {
	s.TrafficControlTaskTraffic = &v
	return s
}

func (s *UpdateTrafficControlTaskTrafficRequestTraffics) Validate() error {
	return dara.Validate(s)
}
