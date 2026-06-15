// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhysicalConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCircuitCode(v string) *ModifyPhysicalConnectionAttributeRequest
	GetCircuitCode() *string
	SetClientToken(v string) *ModifyPhysicalConnectionAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyPhysicalConnectionAttributeRequest
	GetDescription() *string
	SetLineOperator(v string) *ModifyPhysicalConnectionAttributeRequest
	GetLineOperator() *string
	SetName(v string) *ModifyPhysicalConnectionAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest
	GetOwnerId() *int64
	SetPeerLocation(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPeerLocation() *string
	SetPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPhysicalConnectionId() *string
	SetPortType(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPortType() *string
	SetRedundantPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetRedundantPhysicalConnectionId() *string
	SetRegionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *ModifyPhysicalConnectionAttributeRequest
	GetUserCidr() *string
	SetBandwidth(v int32) *ModifyPhysicalConnectionAttributeRequest
	GetBandwidth() *int32
}

type ModifyPhysicalConnectionAttributeRequest struct {
	// The circuit code of the physical connection, provided by the line operator.
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// A client-generated token to ensure the idempotency of the request.
	//
	// The token must be unique across requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the physical connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The line operator that provides the physical connection.
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The new name of the physical connection.
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The peer location of the physical connection.
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The ID of the physical connection.
	//
	// This parameter is required.
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The port type of the physical connection.
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant physical connection.
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The ID of the region where the physical connection is located.
	//
	// To get the latest list of regions, call the `DescribeRegions` operation.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The client-side IP address or CIDR block in your data center.
	//
	// This parameter is used only for outbound traffic shaping.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
	// The bandwidth of the physical connection.
	//
	// Unit: Mbit/s.
	Bandwidth *int32 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
}

func (s ModifyPhysicalConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhysicalConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetLineOperator() *string {
	return s.LineOperator
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPortType() *string {
	return s.PortType
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetCircuitCode(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.CircuitCode = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetClientToken(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetDescription(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetLineOperator(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.LineOperator = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetName(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPeerLocation(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PeerLocation = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPortType(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PortType = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetRedundantPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetRegionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetResourceOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetResourceOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetUserCidr(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.UserCidr = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetBandwidth(v int32) *ModifyPhysicalConnectionAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
