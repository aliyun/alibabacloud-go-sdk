// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeTraceDiagnoseReportRequest
	GetTaskId() *string
	SetTraceId(v string) *DescribeTraceDiagnoseReportRequest
	GetTraceId() *string
}

type DescribeTraceDiagnoseReportRequest struct {
	// The diagnostic task ID. You must specify at least one of TraceId and TaskId. If neither is specified, the API returns an error.
	//
	// example:
	//
	// xxxxxxxxx-x-x-xxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The diagnostic trace ID. You must specify at least one of TraceId and TaskId. If neither is specified, the API returns an error.
	//
	// example:
	//
	// 0000xxxxxxxxxxxxxxxxxxxxxx75e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeTraceDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceDiagnoseReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTraceDiagnoseReportRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeTraceDiagnoseReportRequest) SetTaskId(v string) *DescribeTraceDiagnoseReportRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportRequest) SetTraceId(v string) *DescribeTraceDiagnoseReportRequest {
	s.TraceId = &v
	return s
}

func (s *DescribeTraceDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
