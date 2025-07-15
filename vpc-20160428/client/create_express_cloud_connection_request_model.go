// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressCloudConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *CreateExpressCloudConnectionRequest
	GetBandwidth() *int32
	SetContactMail(v string) *CreateExpressCloudConnectionRequest
	GetContactMail() *string
	SetContactTel(v string) *CreateExpressCloudConnectionRequest
	GetContactTel() *string
	SetDescription(v string) *CreateExpressCloudConnectionRequest
	GetDescription() *string
	SetIDCardNo(v string) *CreateExpressCloudConnectionRequest
	GetIDCardNo() *string
	SetIdcSP(v string) *CreateExpressCloudConnectionRequest
	GetIdcSP() *string
	SetName(v string) *CreateExpressCloudConnectionRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateExpressCloudConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateExpressCloudConnectionRequest
	GetOwnerId() *int64
	SetPeerCity(v string) *CreateExpressCloudConnectionRequest
	GetPeerCity() *string
	SetPeerLocation(v string) *CreateExpressCloudConnectionRequest
	GetPeerLocation() *string
	SetPortType(v string) *CreateExpressCloudConnectionRequest
	GetPortType() *string
	SetRedundantEccId(v string) *CreateExpressCloudConnectionRequest
	GetRedundantEccId() *string
	SetRegionId(v string) *CreateExpressCloudConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateExpressCloudConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateExpressCloudConnectionRequest
	GetResourceOwnerId() *int64
}

type CreateExpressCloudConnectionRequest struct {
	// The bandwidth for ECC, which corresponds to the bandwidth for the underlying circuit.
	//
	// Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The email address of the contact who applies for ECC.
	//
	// example:
	//
	// XX@example.com
	ContactMail *string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty"`
	// The phone number of the contact who applies for ECC.
	//
	// example:
	//
	// 132*********
	ContactTel *string `json:"ContactTel,omitempty" xml:"ContactTel,omitempty"`
	// The description of ECC.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ECC
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID card number of the contact who applies for ECC.
	//
	// example:
	//
	// 32*****************
	IDCardNo *string `json:"IDCardNo,omitempty" xml:"IDCardNo,omitempty"`
	// The Internet service provider (ISP) for the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// CU
	IdcSP *string `json:"IdcSP,omitempty" xml:"IdcSP,omitempty"`
	// The name of the ECC instance.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The city where the data center is located.
	//
	// example:
	//
	// Hangzhou
	PeerCity *string `json:"PeerCity,omitempty" xml:"PeerCity,omitempty"`
	// The geographical location of the data center.
	//
	// > It must be accurate to house number-floor-room number-server rack number.
	//
	// This parameter is required.
	//
	// example:
	//
	// \\*\\*city\\*\\*district/county\\*\\*road\\*\\*number\\*\\*property or building name\\*\\*building\\*\\*floor\\*\\*room number\\*\\*server rack number\\*\\*server rack name\\*\\*device\\*\\*port
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The port of the Express Connect circuit. Valid values:
	//
	// 	- 100Base-T
	//
	// 	- 1000Base-T
	//
	// 	- 1000Base-LX
	//
	// 	- 10GBase-T
	//
	// 	- 10GBase-LR
	//
	// example:
	//
	// 100Base-T
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the standby Express Connect circuit.
	//
	// example:
	//
	// ecc-d****
	RedundantEccId *string `json:"RedundantEccId,omitempty" xml:"RedundantEccId,omitempty"`
	// The region ID of the ECC instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateExpressCloudConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressCloudConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressCloudConnectionRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateExpressCloudConnectionRequest) GetContactMail() *string {
	return s.ContactMail
}

func (s *CreateExpressCloudConnectionRequest) GetContactTel() *string {
	return s.ContactTel
}

func (s *CreateExpressCloudConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExpressCloudConnectionRequest) GetIDCardNo() *string {
	return s.IDCardNo
}

func (s *CreateExpressCloudConnectionRequest) GetIdcSP() *string {
	return s.IdcSP
}

func (s *CreateExpressCloudConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateExpressCloudConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateExpressCloudConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateExpressCloudConnectionRequest) GetPeerCity() *string {
	return s.PeerCity
}

func (s *CreateExpressCloudConnectionRequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *CreateExpressCloudConnectionRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreateExpressCloudConnectionRequest) GetRedundantEccId() *string {
	return s.RedundantEccId
}

func (s *CreateExpressCloudConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExpressCloudConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateExpressCloudConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateExpressCloudConnectionRequest) SetBandwidth(v int32) *CreateExpressCloudConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetContactMail(v string) *CreateExpressCloudConnectionRequest {
	s.ContactMail = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetContactTel(v string) *CreateExpressCloudConnectionRequest {
	s.ContactTel = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetDescription(v string) *CreateExpressCloudConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetIDCardNo(v string) *CreateExpressCloudConnectionRequest {
	s.IDCardNo = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetIdcSP(v string) *CreateExpressCloudConnectionRequest {
	s.IdcSP = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetName(v string) *CreateExpressCloudConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetOwnerAccount(v string) *CreateExpressCloudConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetOwnerId(v int64) *CreateExpressCloudConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetPeerCity(v string) *CreateExpressCloudConnectionRequest {
	s.PeerCity = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetPeerLocation(v string) *CreateExpressCloudConnectionRequest {
	s.PeerLocation = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetPortType(v string) *CreateExpressCloudConnectionRequest {
	s.PortType = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetRedundantEccId(v string) *CreateExpressCloudConnectionRequest {
	s.RedundantEccId = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetRegionId(v string) *CreateExpressCloudConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetResourceOwnerAccount(v string) *CreateExpressCloudConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) SetResourceOwnerId(v int64) *CreateExpressCloudConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateExpressCloudConnectionRequest) Validate() error {
	return dara.Validate(s)
}
