// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNetworkInterfaceRequest
	GetDescription() *string
	SetName(v string) *CreateNetworkInterfaceRequest
	GetName() *string
	SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest
	GetSecurityGroupIds() []*string
	SetVSwitchId(v string) *CreateNetworkInterfaceRequest
	GetVSwitchId() *string
}

type CreateNetworkInterfaceRequest struct {
	// Description of the ENI.
	//
	// example:
	//
	// example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ENI name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Security group ID.
	//
	// This parameter is required.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// vSwitch ID.
	//
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkInterfaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkInterfaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNetworkInterfaceRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateNetworkInterfaceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNetworkInterfaceRequest) SetDescription(v string) *CreateNetworkInterfaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetName(v string) *CreateNetworkInterfaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateNetworkInterfaceRequest) SetVSwitchId(v string) *CreateNetworkInterfaceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
