// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int64) *DescribeServiceMeshClustersRequest
	GetLimit() *int64
	SetOffset(v int64) *DescribeServiceMeshClustersRequest
	GetOffset() *int64
	SetServiceMeshId(v string) *DescribeServiceMeshClustersRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshClustersRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 30
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The position where the query starts.
	//
	// example:
	//
	// 20
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeServiceMeshClustersRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *DescribeServiceMeshClustersRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshClustersRequest) SetLimit(v int64) *DescribeServiceMeshClustersRequest {
	s.Limit = &v
	return s
}

func (s *DescribeServiceMeshClustersRequest) SetOffset(v int64) *DescribeServiceMeshClustersRequest {
	s.Offset = &v
	return s
}

func (s *DescribeServiceMeshClustersRequest) SetServiceMeshId(v string) *DescribeServiceMeshClustersRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshClustersRequest) Validate() error {
	return dara.Validate(s)
}
