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
	// The end time in ISO 8601 format (for example, 2024-01-01T00:00:00+Z).
	//
	// 	Notice: StartTime and EndTime must be provided together to define the query time window. An error is returned if either one is missing..
	//
	// example:
	//
	// 2019-12-06T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The site ID. You can obtain this value by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// 	Notice: This parameter is required when you call the ListUploadTasks operation..
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The start time in ISO 8601 format (for example, 2024-01-01T00:00:00+Z).
	//
	// 	Notice: StartTime and EndTime must be provided together to define the query time window. An error is returned if either one is missing..
	//
	// example:
	//
	// 2018-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task type. Valid values:
	//
	// - **file**: URL file purge.
	//
	// - **preload**: resource prefetch.
	//
	// - **directory**: directory purge.
	//
	// - **ignoreparams**: purge with parameters ignored.
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
