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
	SetEnableJumboFrame(v bool) *ModifyVpcPeerConnectionRequest
	GetEnableJumboFrame() *bool
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
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the VPC peering connection.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newdescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the information. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the jumbo frame feature is enabled for the VPC peering connection. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// > Note: To specify this parameter, make sure that you create a cross-region peering connection.
	//
	// example:
	//
	// true
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// The instance ID of the VPC peering connection whose name or description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The link type.
	//
	// Valid values: Platinum and Gold. Default value: Gold.
	//
	// > To specify this parameter, make sure that you create a cross-region peering connection.
	//
	// example:
	//
	// Gold
	LinkType *string `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// The new name of the VPC peering connection.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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

func (s *ModifyVpcPeerConnectionRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
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

func (s *ModifyVpcPeerConnectionRequest) SetEnableJumboFrame(v bool) *ModifyVpcPeerConnectionRequest {
	s.EnableJumboFrame = &v
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
