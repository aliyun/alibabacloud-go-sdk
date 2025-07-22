// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v []*string) *ForwardInfo
	GetAccessType() []*string
	SetContainerName(v string) *ForwardInfo
	GetContainerName() *string
	SetEipAllocationId(v string) *ForwardInfo
	GetEipAllocationId() *string
	SetEnable(v bool) *ForwardInfo
	GetEnable() *bool
	SetExternalPort(v string) *ForwardInfo
	GetExternalPort() *string
	SetForwardPort(v string) *ForwardInfo
	GetForwardPort() *string
	SetName(v string) *ForwardInfo
	GetName() *string
	SetNatGatewayId(v string) *ForwardInfo
	GetNatGatewayId() *string
	SetSSHPublicKey(v string) *ForwardInfo
	GetSSHPublicKey() *string
}

type ForwardInfo struct {
	AccessType []*string `json:"AccessType,omitempty" xml:"AccessType,omitempty" type:"Repeated"`
	// example:
	//
	// dsw-notebook
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// eip-25877c70gddh****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 10086
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// example:
	//
	// 22
	ForwardPort *string `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// example:
	//
	// ssh
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SSHPublicKey *string `json:"SSHPublicKey,omitempty" xml:"SSHPublicKey,omitempty"`
}

func (s ForwardInfo) String() string {
	return dara.Prettify(s)
}

func (s ForwardInfo) GoString() string {
	return s.String()
}

func (s *ForwardInfo) GetAccessType() []*string {
	return s.AccessType
}

func (s *ForwardInfo) GetContainerName() *string {
	return s.ContainerName
}

func (s *ForwardInfo) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *ForwardInfo) GetEnable() *bool {
	return s.Enable
}

func (s *ForwardInfo) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ForwardInfo) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *ForwardInfo) GetName() *string {
	return s.Name
}

func (s *ForwardInfo) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ForwardInfo) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *ForwardInfo) SetAccessType(v []*string) *ForwardInfo {
	s.AccessType = v
	return s
}

func (s *ForwardInfo) SetContainerName(v string) *ForwardInfo {
	s.ContainerName = &v
	return s
}

func (s *ForwardInfo) SetEipAllocationId(v string) *ForwardInfo {
	s.EipAllocationId = &v
	return s
}

func (s *ForwardInfo) SetEnable(v bool) *ForwardInfo {
	s.Enable = &v
	return s
}

func (s *ForwardInfo) SetExternalPort(v string) *ForwardInfo {
	s.ExternalPort = &v
	return s
}

func (s *ForwardInfo) SetForwardPort(v string) *ForwardInfo {
	s.ForwardPort = &v
	return s
}

func (s *ForwardInfo) SetName(v string) *ForwardInfo {
	s.Name = &v
	return s
}

func (s *ForwardInfo) SetNatGatewayId(v string) *ForwardInfo {
	s.NatGatewayId = &v
	return s
}

func (s *ForwardInfo) SetSSHPublicKey(v string) *ForwardInfo {
	s.SSHPublicKey = &v
	return s
}

func (s *ForwardInfo) Validate() error {
	return dara.Validate(s)
}
