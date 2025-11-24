// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeUpgradeVersionRequest
	GetServiceMeshId() *string
}

type DescribeUpgradeVersionRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeUpgradeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeUpgradeVersionRequest) SetServiceMeshId(v string) *DescribeUpgradeVersionRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeUpgradeVersionRequest) Validate() error {
	return dara.Validate(s)
}
