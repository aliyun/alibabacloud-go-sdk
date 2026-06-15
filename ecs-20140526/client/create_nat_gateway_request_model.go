// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackage(v []*CreateNatGatewayRequestBandwidthPackage) *CreateNatGatewayRequest
	GetBandwidthPackage() []*CreateNatGatewayRequestBandwidthPackage
	SetClientToken(v string) *CreateNatGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *CreateNatGatewayRequest
	GetDescription() *string
	SetName(v string) *CreateNatGatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateNatGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNatGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNatGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNatGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNatGatewayRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *CreateNatGatewayRequest
	GetVpcId() *string
}

type CreateNatGatewayRequest struct {
	// Configurations for the bandwidth packages to create and associate with the nat gateway.
	//
	// This parameter is required.
	BandwidthPackage []*CreateNatGatewayRequestBandwidthPackage `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty" type:"Repeated"`
	// A client token to ensure the idempotence of the request.
	//
	// This token is client-generated and must be unique for each request. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// A description of the nat gateway.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter or a Chinese character but cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A name for the nat gateway.
	//
	// The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character but cannot start with `http://` or `https://`.
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which to create the nat gateway.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC in which to create the nat gateway.
	//
	// This parameter is required.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequest) GetBandwidthPackage() []*CreateNatGatewayRequestBandwidthPackage {
	return s.BandwidthPackage
}

func (s *CreateNatGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNatGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNatGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateNatGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNatGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNatGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNatGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNatGatewayRequest) SetBandwidthPackage(v []*CreateNatGatewayRequestBandwidthPackage) *CreateNatGatewayRequest {
	s.BandwidthPackage = v
	return s
}

func (s *CreateNatGatewayRequest) SetClientToken(v string) *CreateNatGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNatGatewayRequest) SetDescription(v string) *CreateNatGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateNatGatewayRequest) SetName(v string) *CreateNatGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateNatGatewayRequest) SetOwnerAccount(v string) *CreateNatGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNatGatewayRequest) SetOwnerId(v int64) *CreateNatGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetRegionId(v string) *CreateNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetResourceOwnerAccount(v string) *CreateNatGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNatGatewayRequest) SetResourceOwnerId(v int64) *CreateNatGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetVpcId(v string) *CreateNatGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNatGatewayRequest) Validate() error {
	if s.BandwidthPackage != nil {
		for _, item := range s.BandwidthPackage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNatGatewayRequestBandwidthPackage struct {
	// The peak bandwidth for the EIPs in the bandwidth package. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of EIPs to create in the bandwidth package. Valid values: 1 to 10.
	IpCount *int32 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The ID of the zone in which to create the EIPs. If you do not specify a zone, the system randomly selects one.
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s CreateNatGatewayRequestBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequestBandwidthPackage) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequestBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateNatGatewayRequestBandwidthPackage) GetIpCount() *int32 {
	return s.IpCount
}

func (s *CreateNatGatewayRequestBandwidthPackage) GetZone() *string {
	return s.Zone
}

func (s *CreateNatGatewayRequestBandwidthPackage) SetBandwidth(v int32) *CreateNatGatewayRequestBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *CreateNatGatewayRequestBandwidthPackage) SetIpCount(v int32) *CreateNatGatewayRequestBandwidthPackage {
	s.IpCount = &v
	return s
}

func (s *CreateNatGatewayRequestBandwidthPackage) SetZone(v string) *CreateNatGatewayRequestBandwidthPackage {
	s.Zone = &v
	return s
}

func (s *CreateNatGatewayRequestBandwidthPackage) Validate() error {
	return dara.Validate(s)
}
