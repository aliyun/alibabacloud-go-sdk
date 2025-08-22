// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListJobsRequest
	GetAppName() *string
	SetClusterId(v string) *ListJobsRequest
	GetClusterId() *string
	SetDescription(v string) *ListJobsRequest
	GetDescription() *string
	SetJobHandler(v string) *ListJobsRequest
	GetJobHandler() *string
	SetJobId(v int64) *ListJobsRequest
	GetJobId() *int64
	SetJobName(v string) *ListJobsRequest
	GetJobName() *string
	SetPageNum(v int32) *ListJobsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListJobsRequest
	GetStatus() *string
}

type ListJobsRequest struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// jobDemoHandler
	JobHandler *string `json:"JobHandler,omitempty" xml:"JobHandler,omitempty"`
	// example:
	//
	// 10
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// job01
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListJobsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListJobsRequest) GetJobHandler() *string {
	return s.JobHandler
}

func (s *ListJobsRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequest) SetAppName(v string) *ListJobsRequest {
	s.AppName = &v
	return s
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetDescription(v string) *ListJobsRequest {
	s.Description = &v
	return s
}

func (s *ListJobsRequest) SetJobHandler(v string) *ListJobsRequest {
	s.JobHandler = &v
	return s
}

func (s *ListJobsRequest) SetJobId(v int64) *ListJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsRequest) SetJobName(v string) *ListJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListJobsRequest) SetPageNum(v int32) *ListJobsRequest {
	s.PageNum = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
