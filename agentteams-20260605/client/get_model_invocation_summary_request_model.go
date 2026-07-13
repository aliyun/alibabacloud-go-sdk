// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelInvocationSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetModelInvocationSummaryRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetModelInvocationSummaryRequest
	GetInstanceId() *string
	SetStartTime(v string) *GetModelInvocationSummaryRequest
	GetStartTime() *string
}

type GetModelInvocationSummaryRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetModelInvocationSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelInvocationSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetModelInvocationSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetModelInvocationSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetModelInvocationSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetModelInvocationSummaryRequest) SetEndTime(v string) *GetModelInvocationSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetModelInvocationSummaryRequest) SetInstanceId(v string) *GetModelInvocationSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetModelInvocationSummaryRequest) SetStartTime(v string) *GetModelInvocationSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetModelInvocationSummaryRequest) Validate() error {
	return dara.Validate(s)
}
