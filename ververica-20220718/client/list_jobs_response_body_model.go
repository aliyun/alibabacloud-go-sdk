// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListJobsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*Job) *ListJobsResponseBody
	GetData() []*Job
	SetErrorCode(v string) *ListJobsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListJobsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListJobsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListJobsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListJobsResponseBody
	GetTotalSize() *int32
}

type ListJobsResponseBody struct {
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// 	- If the value of success was true, all jobs that meet the condition were returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Job `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListJobsResponseBody) GetData() []*Job {
	return s.Data
}

func (s *ListJobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListJobsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListJobsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListJobsResponseBody) SetAccessDeniedDetail(v string) *ListJobsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListJobsResponseBody) SetData(v []*Job) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetErrorCode(v string) *ListJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobsResponseBody) SetErrorMessage(v string) *ListJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListJobsResponseBody) SetHttpCode(v int32) *ListJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobsResponseBody) SetPageIndex(v int32) *ListJobsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalSize(v int32) *ListJobsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
