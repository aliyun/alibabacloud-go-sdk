// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckReportStepByJobIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListDataCheckReportStepByJobIdRequest
	GetJobId() *string
	SetPageIndex(v int32) *ListDataCheckReportStepByJobIdRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckReportStepByJobIdRequest
	GetPageSize() *int32
}

type ListDataCheckReportStepByJobIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataCheckReportStepByJobIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckReportStepByJobIdRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckReportStepByJobIdRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListDataCheckReportStepByJobIdRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckReportStepByJobIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckReportStepByJobIdRequest) SetJobId(v string) *ListDataCheckReportStepByJobIdRequest {
	s.JobId = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdRequest) SetPageIndex(v int32) *ListDataCheckReportStepByJobIdRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdRequest) SetPageSize(v int32) *ListDataCheckReportStepByJobIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckReportStepByJobIdRequest) Validate() error {
	return dara.Validate(s)
}
