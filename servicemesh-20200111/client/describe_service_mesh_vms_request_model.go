// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshVMsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeServiceMeshVMsRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshVMsRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccb37ff104caf419fbf48fb38e6f****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshVMsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshVMsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshVMsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshVMsRequest) SetServiceMeshId(v string) *DescribeServiceMeshVMsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshVMsRequest) Validate() error {
	return dara.Validate(s)
}
