// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshKubeconfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddress(v bool) *DescribeServiceMeshKubeconfigRequest
	GetPrivateIpAddress() *bool
	SetServiceMeshId(v string) *DescribeServiceMeshKubeconfigRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshKubeconfigRequest struct {
	// Specifies whether to query the kubeconfig file that is used for Internet access or internal network access.
	//
	// example:
	//
	// false
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshKubeconfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshKubeconfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *DescribeServiceMeshKubeconfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshKubeconfigRequest) SetPrivateIpAddress(v bool) *DescribeServiceMeshKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigRequest) SetServiceMeshId(v string) *DescribeServiceMeshKubeconfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshKubeconfigRequest) Validate() error {
	return dara.Validate(s)
}
