// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListRayClusterRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListRayClusterRequest
	GetPageSize() *int32
}

type ListRayClusterRequest struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterRequest) GoString() string {
	return s.String()
}

func (s *ListRayClusterRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRayClusterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRayClusterRequest) SetPageNum(v int32) *ListRayClusterRequest {
	s.PageNum = &v
	return s
}

func (s *ListRayClusterRequest) SetPageSize(v int32) *ListRayClusterRequest {
	s.PageSize = &v
	return s
}

func (s *ListRayClusterRequest) Validate() error {
	return dara.Validate(s)
}
