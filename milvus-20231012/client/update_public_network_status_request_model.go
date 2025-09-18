// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublicNetworkStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *UpdatePublicNetworkStatusRequest
	GetCidr() *string
	SetComponentType(v string) *UpdatePublicNetworkStatusRequest
	GetComponentType() *string
	SetInstanceId(v string) *UpdatePublicNetworkStatusRequest
	GetInstanceId() *string
	SetPublicNetworkEnabled(v bool) *UpdatePublicNetworkStatusRequest
	GetPublicNetworkEnabled() *bool
}

type UpdatePublicNetworkStatusRequest struct {
	// The CIDR blocks.
	//
	// example:
	//
	// ``192.168.**.**``/24,``172.1.**.**``/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The component type. Valid values:
	//
	// 	- Proxy
	//
	// This parameter is required.
	//
	// example:
	//
	// Proxy
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Enable /disable the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	PublicNetworkEnabled *bool `json:"PublicNetworkEnabled,omitempty" xml:"PublicNetworkEnabled,omitempty"`
}

func (s UpdatePublicNetworkStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublicNetworkStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublicNetworkStatusRequest) GetCidr() *string {
	return s.Cidr
}

func (s *UpdatePublicNetworkStatusRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *UpdatePublicNetworkStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdatePublicNetworkStatusRequest) GetPublicNetworkEnabled() *bool {
	return s.PublicNetworkEnabled
}

func (s *UpdatePublicNetworkStatusRequest) SetCidr(v string) *UpdatePublicNetworkStatusRequest {
	s.Cidr = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetComponentType(v string) *UpdatePublicNetworkStatusRequest {
	s.ComponentType = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetInstanceId(v string) *UpdatePublicNetworkStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) SetPublicNetworkEnabled(v bool) *UpdatePublicNetworkStatusRequest {
	s.PublicNetworkEnabled = &v
	return s
}

func (s *UpdatePublicNetworkStatusRequest) Validate() error {
	return dara.Validate(s)
}
