// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagTrafficTopNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagTrafficTopNRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagTrafficTopNRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagTrafficTopNRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagTrafficTopNRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagTrafficTopNRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *DescribeSagTrafficTopNRequest
	GetSize() *int32
}

type DescribeSagTrafficTopNRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of SAG instances to be queried. Default value:**10**. This value cannot be modified.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeSagTrafficTopNRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagTrafficTopNRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagTrafficTopNRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagTrafficTopNRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagTrafficTopNRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagTrafficTopNRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagTrafficTopNRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagTrafficTopNRequest) GetSize() *int32 {
	return s.Size
}

func (s *DescribeSagTrafficTopNRequest) SetOwnerAccount(v string) *DescribeSagTrafficTopNRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) SetOwnerId(v int64) *DescribeSagTrafficTopNRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) SetRegionId(v string) *DescribeSagTrafficTopNRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) SetResourceOwnerAccount(v string) *DescribeSagTrafficTopNRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) SetResourceOwnerId(v int64) *DescribeSagTrafficTopNRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) SetSize(v int32) *DescribeSagTrafficTopNRequest {
	s.Size = &v
	return s
}

func (s *DescribeSagTrafficTopNRequest) Validate() error {
	return dara.Validate(s)
}
