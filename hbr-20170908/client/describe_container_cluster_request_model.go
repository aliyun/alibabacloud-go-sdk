// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerClusterRequest
	GetClusterId() *string
	SetIdentifier(v string) *DescribeContainerClusterRequest
	GetIdentifier() *string
	SetPageNumber(v int32) *DescribeContainerClusterRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContainerClusterRequest
	GetPageSize() *int32
}

type DescribeContainerClusterRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-000*************hg9
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The identifier of the container cluster. For a Container Service for Kubernetes (ACK) cluster, specify the cluster ID.
	//
	// example:
	//
	// cca*******************************87a
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeContainerClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerClusterRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeContainerClusterRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContainerClusterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerClusterRequest) SetClusterId(v string) *DescribeContainerClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerClusterRequest) SetIdentifier(v string) *DescribeContainerClusterRequest {
	s.Identifier = &v
	return s
}

func (s *DescribeContainerClusterRequest) SetPageNumber(v int32) *DescribeContainerClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerClusterRequest) SetPageSize(v int32) *DescribeContainerClusterRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerClusterRequest) Validate() error {
	return dara.Validate(s)
}
