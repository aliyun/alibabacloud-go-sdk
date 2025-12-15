// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *ListClustersRequest
	GetClusterIds() []*string
	SetClusterNames(v []*string) *ListClustersRequest
	GetClusterNames() []*string
	SetPageNumber(v int32) *ListClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClustersRequest
	GetPageSize() *int32
}

type ListClustersRequest struct {
	// The cluster IDs. You can specify up to 20 IDs.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
	// The cluster names. You can specify up to 20 names.
	ClusterNames []*string `json:"ClusterNames,omitempty" xml:"ClusterNames,omitempty" type:"Repeated"`
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

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *ListClustersRequest) GetClusterNames() []*string {
	return s.ClusterNames
}

func (s *ListClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersRequest) SetClusterIds(v []*string) *ListClustersRequest {
	s.ClusterIds = v
	return s
}

func (s *ListClustersRequest) SetClusterNames(v []*string) *ListClustersRequest {
	s.ClusterNames = v
	return s
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}
