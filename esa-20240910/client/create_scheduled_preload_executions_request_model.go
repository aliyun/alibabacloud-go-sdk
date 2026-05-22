// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutions(v []*CreateScheduledPreloadExecutionsRequestExecutions) *CreateScheduledPreloadExecutionsRequest
	GetExecutions() []*CreateScheduledPreloadExecutionsRequestExecutions
	SetId(v string) *CreateScheduledPreloadExecutionsRequest
	GetId() *string
}

type CreateScheduledPreloadExecutionsRequest struct {
	// This parameter is required.
	Executions []*CreateScheduledPreloadExecutionsRequestExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	Id         *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateScheduledPreloadExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsRequest) GetExecutions() []*CreateScheduledPreloadExecutionsRequestExecutions {
	return s.Executions
}

func (s *CreateScheduledPreloadExecutionsRequest) GetId() *string {
	return s.Id
}

func (s *CreateScheduledPreloadExecutionsRequest) SetExecutions(v []*CreateScheduledPreloadExecutionsRequestExecutions) *CreateScheduledPreloadExecutionsRequest {
	s.Executions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequest) SetId(v string) *CreateScheduledPreloadExecutionsRequest {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequest) Validate() error {
	if s.Executions != nil {
		for _, item := range s.Executions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScheduledPreloadExecutionsRequestExecutions struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateScheduledPreloadExecutionsRequestExecutions) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsRequestExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsRequestExecutions {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsRequestExecutions) Validate() error {
	return dara.Validate(s)
}
