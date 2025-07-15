// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSnatEntryRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteSnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSnatEntryRequest
	GetResourceOwnerId() *int64
	SetSnatEntryId(v string) *DeleteSnatEntryRequest
	GetSnatEntryId() *string
	SetSnatTableId(v string) *DeleteSnatEntryRequest
	GetSnatTableId() *string
}

type DeleteSnatEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SNAT entry that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-bp1vcgcf8tm0plqcg****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The ID of the SNAT table to which the SNAT entry belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// stb-bp190wu8io1vgev80****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
}

func (s DeleteSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSnatEntryRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DeleteSnatEntryRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DeleteSnatEntryRequest) SetClientToken(v string) *DeleteSnatEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetOwnerAccount(v string) *DeleteSnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetOwnerId(v int64) *DeleteSnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetRegionId(v string) *DeleteSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetResourceOwnerAccount(v string) *DeleteSnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetResourceOwnerId(v int64) *DeleteSnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetSnatEntryId(v string) *DeleteSnatEntryRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetSnatTableId(v string) *DeleteSnatEntryRequest {
	s.SnatTableId = &v
	return s
}

func (s *DeleteSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
