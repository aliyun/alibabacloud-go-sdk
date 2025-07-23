// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListClustersRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListClustersRequest
	GetRegionId() *string
}

type ListClustersRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClustersRequest) SetCurrentPage(v int32) *ListClustersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}
