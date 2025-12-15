// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIdsShrink(v string) *ListClustersShrinkRequest
	GetClusterIdsShrink() *string
	SetClusterNamesShrink(v string) *ListClustersShrinkRequest
	GetClusterNamesShrink() *string
	SetPageNumber(v int32) *ListClustersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClustersShrinkRequest
	GetPageSize() *int32
}

type ListClustersShrinkRequest struct {
	// The cluster IDs. You can specify up to 20 IDs.
	ClusterIdsShrink *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// The cluster names. You can specify up to 20 names.
	ClusterNamesShrink *string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListClustersShrinkRequest) GetClusterIdsShrink() *string {
	return s.ClusterIdsShrink
}

func (s *ListClustersShrinkRequest) GetClusterNamesShrink() *string {
	return s.ClusterNamesShrink
}

func (s *ListClustersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersShrinkRequest) SetClusterIdsShrink(v string) *ListClustersShrinkRequest {
	s.ClusterIdsShrink = &v
	return s
}

func (s *ListClustersShrinkRequest) SetClusterNamesShrink(v string) *ListClustersShrinkRequest {
	s.ClusterNamesShrink = &v
	return s
}

func (s *ListClustersShrinkRequest) SetPageNumber(v int32) *ListClustersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersShrinkRequest) SetPageSize(v int32) *ListClustersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
