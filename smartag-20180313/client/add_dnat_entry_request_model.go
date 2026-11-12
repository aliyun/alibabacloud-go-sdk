// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *AddDnatEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *AddDnatEntryRequest
	GetExternalPort() *string
	SetInternalIp(v string) *AddDnatEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *AddDnatEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *AddDnatEntryRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *AddDnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddDnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddDnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddDnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddDnatEntryRequest
	GetResourceOwnerId() *int64
	SetSagId(v string) *AddDnatEntryRequest
	GetSagId() *string
	SetType(v string) *AddDnatEntryRequest
	GetType() *string
}

type AddDnatEntryRequest struct {
	// The external public IP address.
	//
	// > If the DNAT type is Internet, the system automatically identifies the current public IP address. You do not need to specify this parameter.
	//
	// example:
	//
	// 10.10.1.xx
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The public port.
	//
	// Valid values: **1-65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The destination private IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The destination private port number.
	//
	// Valid values: **1-65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol type. Valid values:
	//
	// - **TCP**: forwards TCP packets.
	//
	// - **UDP**: forwards UDP packets.
	//
	// - **Any**: forwards packets of all protocols.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Smart Access Gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the Smart Access Gateway instance.
	//
	// > Currently, only hardware instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-kdhg*******
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The DNAT type. Valid values:
	//
	// - **Intranet**: private network DNAT. This is the default value.
	//
	// - **Internet**: public network DNAT.
	//
	// example:
	//
	// Intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddDnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnatEntryRequest) GoString() string {
	return s.String()
}

func (s *AddDnatEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *AddDnatEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *AddDnatEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *AddDnatEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *AddDnatEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AddDnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddDnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddDnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddDnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddDnatEntryRequest) GetSagId() *string {
	return s.SagId
}

func (s *AddDnatEntryRequest) GetType() *string {
	return s.Type
}

func (s *AddDnatEntryRequest) SetExternalIp(v string) *AddDnatEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *AddDnatEntryRequest) SetExternalPort(v string) *AddDnatEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *AddDnatEntryRequest) SetInternalIp(v string) *AddDnatEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *AddDnatEntryRequest) SetInternalPort(v string) *AddDnatEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *AddDnatEntryRequest) SetIpProtocol(v string) *AddDnatEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *AddDnatEntryRequest) SetOwnerAccount(v string) *AddDnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddDnatEntryRequest) SetOwnerId(v int64) *AddDnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *AddDnatEntryRequest) SetRegionId(v string) *AddDnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *AddDnatEntryRequest) SetResourceOwnerAccount(v string) *AddDnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddDnatEntryRequest) SetResourceOwnerId(v int64) *AddDnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddDnatEntryRequest) SetSagId(v string) *AddDnatEntryRequest {
	s.SagId = &v
	return s
}

func (s *AddDnatEntryRequest) SetType(v string) *AddDnatEntryRequest {
	s.Type = &v
	return s
}

func (s *AddDnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
