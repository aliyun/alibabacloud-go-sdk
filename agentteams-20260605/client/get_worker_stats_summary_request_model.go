// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerStatsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetWorkerStatsSummaryRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetWorkerStatsSummaryRequest
	GetInstanceId() *string
	SetStartTime(v string) *GetWorkerStatsSummaryRequest
	GetStartTime() *string
}

type GetWorkerStatsSummaryRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetWorkerStatsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerStatsSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerStatsSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetWorkerStatsSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerStatsSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetWorkerStatsSummaryRequest) SetEndTime(v string) *GetWorkerStatsSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetWorkerStatsSummaryRequest) SetInstanceId(v string) *GetWorkerStatsSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerStatsSummaryRequest) SetStartTime(v string) *GetWorkerStatsSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetWorkerStatsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
