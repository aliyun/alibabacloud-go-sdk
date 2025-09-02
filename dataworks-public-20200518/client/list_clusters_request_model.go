// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *ListClustersRequest
	GetClusterType() *string
	SetPageNumber(v int32) *ListClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClustersRequest
	GetPageSize() *int32
}

type ListClustersRequest struct {
	// The type of the cluster. Valid values:
	//
	// 	- CDH: CDH cluster
	//
	// 	- EMR: EMR cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
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

func (s *ListClustersRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersRequest) SetClusterType(v string) *ListClustersRequest {
	s.ClusterType = &v
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
