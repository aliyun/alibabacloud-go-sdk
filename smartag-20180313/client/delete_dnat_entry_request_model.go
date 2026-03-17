// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnatEntryId(v string) *DeleteDnatEntryRequest
	GetDnatEntryId() *string
	SetOwnerAccount(v string) *DeleteDnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteDnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDnatEntryRequest
	GetResourceOwnerId() *int64
	SetSagId(v string) *DeleteDnatEntryRequest
	GetSagId() *string
}

type DeleteDnatEntryRequest struct {
	// The ID of the DNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// fwd-kxe4fq3xuzczze****
	DnatEntryId  *string `json:"DnatEntryId,omitempty" xml:"DnatEntryId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
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
	// The ID of the SAG instance.
	//
	// >  Only SAG instances used to manage SAG devices support DNAT.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-jfh*********
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
}

func (s DeleteDnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDnatEntryRequest) GetDnatEntryId() *string {
	return s.DnatEntryId
}

func (s *DeleteDnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDnatEntryRequest) GetSagId() *string {
	return s.SagId
}

func (s *DeleteDnatEntryRequest) SetDnatEntryId(v string) *DeleteDnatEntryRequest {
	s.DnatEntryId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetOwnerAccount(v string) *DeleteDnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetOwnerId(v int64) *DeleteDnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetRegionId(v string) *DeleteDnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetResourceOwnerAccount(v string) *DeleteDnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetResourceOwnerId(v int64) *DeleteDnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDnatEntryRequest) SetSagId(v string) *DeleteDnatEntryRequest {
	s.SagId = &v
	return s
}

func (s *DeleteDnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
