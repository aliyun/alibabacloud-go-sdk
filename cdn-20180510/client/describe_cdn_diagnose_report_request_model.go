// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeCdnDiagnoseReportRequest
	GetTaskId() *string
	SetTraceId(v string) *DescribeCdnDiagnoseReportRequest
	GetTraceId() *string
}

type DescribeCdnDiagnoseReportRequest struct {
	// example:
	//
	// xxxxxxxxx-x-x-xxxxxxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0000xxxxxxxxxxxxxxxxxxxxxx75e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeCdnDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCdnDiagnoseReportRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeCdnDiagnoseReportRequest) SetTaskId(v string) *DescribeCdnDiagnoseReportRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportRequest) SetTraceId(v string) *DescribeCdnDiagnoseReportRequest {
	s.TraceId = &v
	return s
}

func (s *DescribeCdnDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
