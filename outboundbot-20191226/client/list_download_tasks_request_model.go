// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownloadTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDownloadTasksRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDownloadTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDownloadTasksRequest
	GetPageSize() *int32
}

type ListDownloadTasksRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDownloadTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDownloadTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDownloadTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDownloadTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownloadTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownloadTasksRequest) SetInstanceId(v string) *ListDownloadTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDownloadTasksRequest) SetPageNumber(v int32) *ListDownloadTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDownloadTasksRequest) SetPageSize(v int32) *ListDownloadTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDownloadTasksRequest) Validate() error {
	return dara.Validate(s)
}
