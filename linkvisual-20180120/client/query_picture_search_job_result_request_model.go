// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPictureSearchJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *QueryPictureSearchJobResultRequest
	GetAppInstanceId() *string
	SetCurrentPage(v int32) *QueryPictureSearchJobResultRequest
	GetCurrentPage() *int32
	SetJobId(v string) *QueryPictureSearchJobResultRequest
	GetJobId() *string
	SetPageSize(v int32) *QueryPictureSearchJobResultRequest
	GetPageSize() *int32
}

type QueryPictureSearchJobResultRequest struct {
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
	// The ID of the integrated image search job. You can call the [QueryPictureSearchJob](https://help.aliyun.com/document_detail/206778.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b7e501d0d2524a4caeb50df3282a****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of entries to return on each page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryPictureSearchJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPictureSearchJobResultRequest) GoString() string {
	return s.String()
}

func (s *QueryPictureSearchJobResultRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *QueryPictureSearchJobResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryPictureSearchJobResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryPictureSearchJobResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPictureSearchJobResultRequest) SetAppInstanceId(v string) *QueryPictureSearchJobResultRequest {
	s.AppInstanceId = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetCurrentPage(v int32) *QueryPictureSearchJobResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetJobId(v string) *QueryPictureSearchJobResultRequest {
	s.JobId = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) SetPageSize(v int32) *QueryPictureSearchJobResultRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPictureSearchJobResultRequest) Validate() error {
	return dara.Validate(s)
}
