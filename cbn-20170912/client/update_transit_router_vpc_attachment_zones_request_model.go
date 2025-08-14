// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVpcAttachmentZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddZoneMappings(v []*UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetAddZoneMappings() []*UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings
	SetClientToken(v string) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetOwnerId() *int64
	SetRemoveZoneMappings(v []*UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetRemoveZoneMappings() []*UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings
	SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentZonesRequest
	GetTransitRouterAttachmentId() *string
}

type UpdateTransitRouterVpcAttachmentZonesRequest struct {
	// The zones and vSwitches that you want to add to the VPC connection.
	AddZoneMappings []*UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings `json:"AddZoneMappings,omitempty" xml:"AddZoneMappings,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The zones and vSwitches that you want to remove from the VPC connection.
	RemoveZoneMappings   []*UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings `json:"RemoveZoneMappings,omitempty" xml:"RemoveZoneMappings,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-9bbqyygouv4cpn****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetAddZoneMappings() []*UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings {
	return s.AddZoneMappings
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetRemoveZoneMappings() []*UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings {
	return s.RemoveZoneMappings
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetAddZoneMappings(v []*UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.AddZoneMappings = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetClientToken(v string) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetDryRun(v bool) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetOwnerId(v int64) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetRemoveZoneMappings(v []*UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.RemoveZoneMappings = v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) SetTransitRouterAttachmentId(v string) *UpdateTransitRouterVpcAttachmentZonesRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings struct {
	// The ID of the vSwitch that you want to add to the VPC connection.
	//
	// You can specify at most 10 vSwitches in each call.
	//
	// 	- If the VPC connection belongs to the current Alibaba Cloud account, you can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the IDs of the vSwitches and zones of the VPC.
	//
	// 	- If the VPC connection belongs to another Alibaba Cloud account, you can call the [ListGrantVSwitchesToCen](https://help.aliyun.com/document_detail/427599.html) operation to query the IDs of the vSwitches and zones of the VPC.
	//
	// example:
	//
	// vsw-wz988dda8ldm4uvmx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where the vSwitch that you want to add to the VPC connection is deployed.
	//
	// You can specify at most 10 vSwitches in each call.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) SetVSwitchId(v string) *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) SetZoneId(v string) *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestAddZoneMappings) Validate() error {
	return dara.Validate(s)
}

type UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings struct {
	// The ID of the vSwitch that you want to remove from the VPC connection.
	//
	// You can remove at most 10 vSwitches from a VPC in each call.
	//
	// example:
	//
	// vsw-wz9f5izl6wshndmta****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where the vSwitch that you want to remove from the VPC connection is deployed.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) SetVSwitchId(v string) *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) SetZoneId(v string) *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *UpdateTransitRouterVpcAttachmentZonesRequestRemoveZoneMappings) Validate() error {
	return dara.Validate(s)
}
