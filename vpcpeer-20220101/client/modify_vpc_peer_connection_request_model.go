// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifyVpcPeerConnectionRequest
	GetBandwidth() *int32
	SetClientToken(v string) *ModifyVpcPeerConnectionRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyVpcPeerConnectionRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyVpcPeerConnectionRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyVpcPeerConnectionRequest
	GetInstanceId() *string
	SetLinkType(v string) *ModifyVpcPeerConnectionRequest
	GetLinkType() *string
	SetName(v string) *ModifyVpcPeerConnectionRequest
	GetName() *string
}

type ModifyVpcPeerConnectionRequest struct {
	// The new bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be an integer greater than 0.
	//
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the VPC peering connection.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newdescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system prechecks the required parameters, request syntax, and limits. If the request fails the precheck, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC peering connection whose name or description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Type of connection. Valid values:
	//
	// - Platinum.
	//
	// - Gold: default value.
	//
	// >
	//
	// > - If you need to specify this parameter, ensure that the VPC peering connection is an inter-region connection.
	//
	// example:
	//
	// Gold
	LinkType *string `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// The new name of the VPC peering connection.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vpcpeername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyVpcPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcPeerConnectionRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyVpcPeerConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyVpcPeerConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcPeerConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyVpcPeerConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyVpcPeerConnectionRequest) GetLinkType() *string {
	return s.LinkType
}

func (s *ModifyVpcPeerConnectionRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVpcPeerConnectionRequest) SetBandwidth(v int32) *ModifyVpcPeerConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetClientToken(v string) *ModifyVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetDescription(v string) *ModifyVpcPeerConnectionRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetDryRun(v bool) *ModifyVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetInstanceId(v string) *ModifyVpcPeerConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetLinkType(v string) *ModifyVpcPeerConnectionRequest {
	s.LinkType = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) SetName(v string) *ModifyVpcPeerConnectionRequest {
	s.Name = &v
	return s
}

func (s *ModifyVpcPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}
