// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetTaskStatsSummaryRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetTaskStatsSummaryRequest
	GetInstanceId() *string
	SetStartTime(v string) *GetTaskStatsSummaryRequest
	GetStartTime() *string
}

type GetTaskStatsSummaryRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetTaskStatsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatsSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatsSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskStatsSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTaskStatsSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskStatsSummaryRequest) SetEndTime(v string) *GetTaskStatsSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetTaskStatsSummaryRequest) SetInstanceId(v string) *GetTaskStatsSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTaskStatsSummaryRequest) SetStartTime(v string) *GetTaskStatsSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetTaskStatsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
