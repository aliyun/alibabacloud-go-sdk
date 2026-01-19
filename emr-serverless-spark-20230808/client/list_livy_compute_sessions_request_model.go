// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListLivyComputeSessionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListLivyComputeSessionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListLivyComputeSessionsRequest
	GetRegionId() *string
}

type ListLivyComputeSessionsRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLivyComputeSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListLivyComputeSessionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListLivyComputeSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLivyComputeSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLivyComputeSessionsRequest) SetPageNum(v int32) *ListLivyComputeSessionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListLivyComputeSessionsRequest) SetPageSize(v int32) *ListLivyComputeSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLivyComputeSessionsRequest) SetRegionId(v string) *ListLivyComputeSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLivyComputeSessionsRequest) Validate() error {
	return dara.Validate(s)
}
