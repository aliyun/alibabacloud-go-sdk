// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKubeconfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateIpAddress(v bool) *RevokeKubeconfigRequest
	GetPrivateIpAddress() *bool
	SetServiceMeshId(v string) *RevokeKubeconfigRequest
	GetServiceMeshId() *string
}

type RevokeKubeconfigRequest struct {
	// Specifies whether to return the kubeconfig file for private access.
	//
	// 	- `true`: returns the kubeconfig file for private access.
	//
	// 	- `false`: returns the kubeconfig file for public access.
	//
	// example:
	//
	// false
	PrivateIpAddress *bool `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the ASM instance for which you want to revoke its kubeconfig file.
	//
	// This parameter is required.
	//
	// example:
	//
	// cf08a11940e8c46c48bc791fcdb3****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s RevokeKubeconfigRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeKubeconfigRequest) GoString() string {
	return s.String()
}

func (s *RevokeKubeconfigRequest) GetPrivateIpAddress() *bool {
	return s.PrivateIpAddress
}

func (s *RevokeKubeconfigRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *RevokeKubeconfigRequest) SetPrivateIpAddress(v bool) *RevokeKubeconfigRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RevokeKubeconfigRequest) SetServiceMeshId(v string) *RevokeKubeconfigRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *RevokeKubeconfigRequest) Validate() error {
	return dara.Validate(s)
}
