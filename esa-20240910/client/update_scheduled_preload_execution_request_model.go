// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPreloadExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateScheduledPreloadExecutionRequest
	GetEndTime() *string
	SetId(v string) *UpdateScheduledPreloadExecutionRequest
	GetId() *string
	SetInterval(v int32) *UpdateScheduledPreloadExecutionRequest
	GetInterval() *int32
	SetSliceLen(v int32) *UpdateScheduledPreloadExecutionRequest
	GetSliceLen() *int32
	SetStartTime(v string) *UpdateScheduledPreloadExecutionRequest
	GetStartTime() *string
}

type UpdateScheduledPreloadExecutionRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateScheduledPreloadExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateScheduledPreloadExecutionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateScheduledPreloadExecutionRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateScheduledPreloadExecutionRequest) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *UpdateScheduledPreloadExecutionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateScheduledPreloadExecutionRequest) SetEndTime(v string) *UpdateScheduledPreloadExecutionRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetId(v string) *UpdateScheduledPreloadExecutionRequest {
	s.Id = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetInterval(v int32) *UpdateScheduledPreloadExecutionRequest {
	s.Interval = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetSliceLen(v int32) *UpdateScheduledPreloadExecutionRequest {
	s.SliceLen = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) SetStartTime(v string) *UpdateScheduledPreloadExecutionRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionRequest) Validate() error {
	return dara.Validate(s)
}
