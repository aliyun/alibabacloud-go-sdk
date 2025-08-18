// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUploadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListUploadTasksRequest
	GetEndTime() *string
	SetSiteId(v int64) *ListUploadTasksRequest
	GetSiteId() *int64
	SetStartTime(v string) *ListUploadTasksRequest
	GetStartTime() *string
	SetType(v string) *ListUploadTasksRequest
	GetType() *string
}

type ListUploadTasksRequest struct {
	// The time when the task ends. Specify the time in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2019-12-06T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The time when the task starts. Specify the time in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task type. Valid values:
	//
	// 	- **file**: purges the cache by file URL.
	//
	// 	- **preload**: prefetches files.
	//
	// 	- **directory**: purges the cache by directory.
	//
	// 	- **ignoreparams**: purges the cache by URL with specified parameters ignored.
	//
	// example:
	//
	// file
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUploadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUploadTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUploadTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListUploadTasksRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListUploadTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListUploadTasksRequest) GetType() *string {
	return s.Type
}

func (s *ListUploadTasksRequest) SetEndTime(v string) *ListUploadTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetSiteId(v int64) *ListUploadTasksRequest {
	s.SiteId = &v
	return s
}

func (s *ListUploadTasksRequest) SetStartTime(v string) *ListUploadTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListUploadTasksRequest) SetType(v string) *ListUploadTasksRequest {
	s.Type = &v
	return s
}

func (s *ListUploadTasksRequest) Validate() error {
	return dara.Validate(s)
}
