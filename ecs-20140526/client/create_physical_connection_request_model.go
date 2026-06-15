// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreatePhysicalConnectionRequest
	GetAccessPointId() *string
	SetCircuitCode(v string) *CreatePhysicalConnectionRequest
	GetCircuitCode() *string
	SetClientToken(v string) *CreatePhysicalConnectionRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePhysicalConnectionRequest
	GetDescription() *string
	SetLineOperator(v string) *CreatePhysicalConnectionRequest
	GetLineOperator() *string
	SetName(v string) *CreatePhysicalConnectionRequest
	GetName() *string
	SetOwnerAccount(v string) *CreatePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPeerLocation(v string) *CreatePhysicalConnectionRequest
	GetPeerLocation() *string
	SetPortType(v string) *CreatePhysicalConnectionRequest
	GetPortType() *string
	SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionRequest
	GetRedundantPhysicalConnectionId() *string
	SetRegionId(v string) *CreatePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhysicalConnectionRequest
	GetResourceOwnerId() *int64
	SetType(v string) *CreatePhysicalConnectionRequest
	GetType() *string
	SetUserCidr(v string) *CreatePhysicalConnectionRequest
	GetUserCidr() *string
	SetBandwidth(v int32) *CreatePhysicalConnectionRequest
	GetBandwidth() *int32
}

type CreatePhysicalConnectionRequest struct {
	// The access point ID. You can call the `DescribeAccessPoints` operation to obtain a list of available access points.
	//
	// This parameter is required.
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The circuit code provided by the carrier.
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// A client-generated token that you can use to ensure the idempotency of the request. This token must be unique across requests, contain only ASCII characters, and be no more than 64 characters long.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the physical connection. The description must be 2 to 256 characters long and cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The carrier that provides the physical connection. Valid values: `CT` (China Telecom), `CU` (China Unicom), `CM` (China Mobile), `CO` (other Chinese carriers), and `AL` (Alibaba Cloud).
	//
	// This parameter is required.
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The name of the physical connection. The name must be 2 to 128 characters long. It must start with a letter and can contain letters, digits, underscores (`_`), and hyphens (`-`).
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The physical location of your on-premises data center.
	//
	// This parameter is required.
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The port type of the physical connection. You cannot change this parameter after the physical connection is created. Valid values: `1000Base-LX` (1 Gbit/s), `10GBase-LR` (10 Gbit/s), and `40GBase-LR` (40 Gbit/s).
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant physical connection. The redundant physical connection must be in the `Allocated`, `Confirmed`, or `Enabled` state.
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The ID of the region for the physical connection. You can call the `DescribeRegions` operation to obtain the latest list of regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the physical connection. Valid values: `VPC` and `VBR`. The default value is `VPC`. This parameter is available only to whitelisted users.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user CIDR block. This parameter is required when `Type` is set to `VPC`. The CIDR block must be a private IPv4 block. Valid CIDR blocks include the following blocks and their subnets: `10.0.0.0/8`, `172.16.0.0/12`, and `192.168.0.0/16`.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The bandwidth of the physical connection in Mbit/s. The value must be an integer that ranges from 1 to 10,240.
	Bandwidth *int32 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
}

func (s CreatePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreatePhysicalConnectionRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *CreatePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePhysicalConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePhysicalConnectionRequest) GetLineOperator() *string {
	return s.LineOperator
}

func (s *CreatePhysicalConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreatePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhysicalConnectionRequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *CreatePhysicalConnectionRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreatePhysicalConnectionRequest) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *CreatePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhysicalConnectionRequest) GetType() *string {
	return s.Type
}

func (s *CreatePhysicalConnectionRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *CreatePhysicalConnectionRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreatePhysicalConnectionRequest) SetAccessPointId(v string) *CreatePhysicalConnectionRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetCircuitCode(v string) *CreatePhysicalConnectionRequest {
	s.CircuitCode = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetClientToken(v string) *CreatePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetDescription(v string) *CreatePhysicalConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetLineOperator(v string) *CreatePhysicalConnectionRequest {
	s.LineOperator = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetName(v string) *CreatePhysicalConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetOwnerAccount(v string) *CreatePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetOwnerId(v int64) *CreatePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetPeerLocation(v string) *CreatePhysicalConnectionRequest {
	s.PeerLocation = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetPortType(v string) *CreatePhysicalConnectionRequest {
	s.PortType = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionRequest {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetRegionId(v string) *CreatePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *CreatePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetResourceOwnerId(v int64) *CreatePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetType(v string) *CreatePhysicalConnectionRequest {
	s.Type = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetUserCidr(v string) *CreatePhysicalConnectionRequest {
	s.UserCidr = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetBandwidth(v int32) *CreatePhysicalConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) Validate() error {
	return dara.Validate(s)
}
