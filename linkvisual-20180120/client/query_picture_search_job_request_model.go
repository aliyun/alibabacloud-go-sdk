// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *QueryPictureSearchJobRequest
	GetAppInstanceId() *string
	SetCurrentPage(v int32) *QueryPictureSearchJobRequest
	GetCurrentPage() *int32
	SetJobStatus(v int32) *QueryPictureSearchJobRequest
	GetJobStatus() *int32
	SetPageSize(v int32) *QueryPictureSearchJobRequest
	GetPageSize() *int32
}

type QueryPictureSearchJobRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fdc576d03c094faea9fa10fbf48c****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The page number of the page to return. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The job status. Valid values:
	//
	// - 0: Initialized.
	//
	// - 1: No data.
	//
	// - 2: Running.
	//
	// - 3: Succeeded.
	//
	// - 4: Failed.
	//
	// If you do not specify this parameter, jobs in all statuses are queried.
	//
	// example:
	//
	// 3
	JobStatus *int32 `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The number of entries to return on each page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *QueryPictureSearchJobRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchJobRequest) GetJobStatus() *int32 {
	return s.JobStatus
}

func (s *QueryPictureSearchJobRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchJobRequest) SetAppInstanceId(v string) *QueryPictureSearchJobRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetCurrentPage(v int32) *QueryPictureSearchJobRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetJobStatus(v int32) *QueryPictureSearchJobRequest {
	s.JobStatus = &v
	return s
}

func (s *QueryPictureSearchJobRequest) SetPageSize(v int32) *QueryPictureSearchJobRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchJobRequest) Validate() error {
	return dara.Validate(s)
}
