// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodesInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeNodesInstanceTypeRequest
	GetServiceMeshId() *string
}

type DescribeNodesInstanceTypeRequest struct {
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeNodesInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodesInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodesInstanceTypeRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeNodesInstanceTypeRequest) SetServiceMeshId(v string) *DescribeNodesInstanceTypeRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeNodesInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
