// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTrafficControlTaskItemReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QueryTrafficControlTaskItemReportRequest
	GetEndTime() *string
	SetEnvironment(v string) *QueryTrafficControlTaskItemReportRequest
	GetEnvironment() *string
	SetInstanceId(v string) *QueryTrafficControlTaskItemReportRequest
	GetInstanceId() *string
	SetStartTime(v string) *QueryTrafficControlTaskItemReportRequest
	GetStartTime() *string
}

type QueryTrafficControlTaskItemReportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-02 10:30:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec_123****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-01 10:30:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryTrafficControlTaskItemReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTrafficControlTaskItemReportRequest) GoString() string {
	return s.String()
}

func (s *QueryTrafficControlTaskItemReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryTrafficControlTaskItemReportRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *QueryTrafficControlTaskItemReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryTrafficControlTaskItemReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryTrafficControlTaskItemReportRequest) SetEndTime(v string) *QueryTrafficControlTaskItemReportRequest {
	s.EndTime = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportRequest) SetEnvironment(v string) *QueryTrafficControlTaskItemReportRequest {
	s.Environment = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportRequest) SetInstanceId(v string) *QueryTrafficControlTaskItemReportRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportRequest) SetStartTime(v string) *QueryTrafficControlTaskItemReportRequest {
	s.StartTime = &v
	return s
}

func (s *QueryTrafficControlTaskItemReportRequest) Validate() error {
	return dara.Validate(s)
}
