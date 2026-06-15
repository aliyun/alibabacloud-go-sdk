// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *CreateForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *CreateForwardEntryRequest
	GetExternalPort() *string
	SetForwardTableId(v string) *CreateForwardEntryRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *CreateForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *CreateForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *CreateForwardEntryRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *CreateForwardEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateForwardEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateForwardEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateForwardEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateForwardEntryRequest
	GetResourceOwnerId() *int64
}

type CreateForwardEntryRequest struct {
	// The public IP address for the DNAT entry. This must be a public IP address of the NAT gateway associated with the DNAT table.
	//
	// This parameter is required.
	ExternalIp *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	// The external port used for DNAT. Valid values: 1 to 65535.
	//
	// This parameter is required.
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The ID of the DNAT table.
	//
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	// The private IP address to which traffic is forwarded.
	//
	// This parameter is required.
	InternalIp *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	// The internal port to which traffic is forwarded. Valid values: 1 to 65535.
	//
	// This parameter is required.
	InternalPort *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	// The protocol. Valid values: TCP, UDP, and Any. If you set this parameter to Any, the DNAT entry applies to all protocols.
	//
	// This parameter is required.
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the DNAT table is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to retrieve a list of available regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *CreateForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *CreateForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *CreateForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *CreateForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *CreateForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateForwardEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateForwardEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateForwardEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateForwardEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateForwardEntryRequest) SetExternalIp(v string) *CreateForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetExternalPort(v string) *CreateForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetForwardTableId(v string) *CreateForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalIp(v string) *CreateForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *CreateForwardEntryRequest) SetInternalPort(v string) *CreateForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *CreateForwardEntryRequest) SetIpProtocol(v string) *CreateForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateForwardEntryRequest) SetOwnerAccount(v string) *CreateForwardEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateForwardEntryRequest) SetOwnerId(v int64) *CreateForwardEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetRegionId(v string) *CreateForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardEntryRequest) SetResourceOwnerAccount(v string) *CreateForwardEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateForwardEntryRequest) SetResourceOwnerId(v int64) *CreateForwardEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
