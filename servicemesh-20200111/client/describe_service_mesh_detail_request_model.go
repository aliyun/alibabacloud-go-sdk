// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeServiceMeshDetailRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshDetailRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshDetailRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshDetailRequest) SetServiceMeshId(v string) *DescribeServiceMeshDetailRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshDetailRequest) Validate() error {
	return dara.Validate(s)
}
