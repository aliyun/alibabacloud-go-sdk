// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListRenderingInstancesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListRenderingInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingInstancesRequest
	GetPageSize() *int32
	SetRenderingInstanceId(v string) *ListRenderingInstancesRequest
	GetRenderingInstanceId() *string
	SetRenderingSpec(v string) *ListRenderingInstancesRequest
	GetRenderingSpec() *string
	SetStartTime(v string) *ListRenderingInstancesRequest
	GetStartTime() *string
	SetStorageSize(v int32) *ListRenderingInstancesRequest
	GetStorageSize() *int32
}

type ListRenderingInstancesRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StorageSize   *int32  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s ListRenderingInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingInstancesRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingInstancesRequest) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *ListRenderingInstancesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingInstancesRequest) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ListRenderingInstancesRequest) SetEndTime(v string) *ListRenderingInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetPageNumber(v int32) *ListRenderingInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetPageSize(v int32) *ListRenderingInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetRenderingInstanceId(v string) *ListRenderingInstancesRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetRenderingSpec(v string) *ListRenderingInstancesRequest {
	s.RenderingSpec = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetStartTime(v string) *ListRenderingInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingInstancesRequest) SetStorageSize(v int32) *ListRenderingInstancesRequest {
	s.StorageSize = &v
	return s
}

func (s *ListRenderingInstancesRequest) Validate() error {
	return dara.Validate(s)
}
