// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVMsInServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeVMsInServiceMeshRequest
	GetServiceMeshId() *string
}

type DescribeVMsInServiceMeshRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccb37ff104caf419fbf48fb38e6f3****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeVMsInServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVMsInServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *DescribeVMsInServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeVMsInServiceMeshRequest) SetServiceMeshId(v string) *DescribeVMsInServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeVMsInServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
