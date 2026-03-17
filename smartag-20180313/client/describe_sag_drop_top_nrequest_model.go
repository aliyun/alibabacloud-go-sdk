// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagDropTopNRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagDropTopNRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagDropTopNRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagDropTopNRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagDropTopNRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagDropTopNRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *DescribeSagDropTopNRequest
	GetSize() *int32
}

type DescribeSagDropTopNRequest struct {
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

func (s DescribeSagDropTopNRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagDropTopNRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagDropTopNRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagDropTopNRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagDropTopNRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagDropTopNRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagDropTopNRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagDropTopNRequest) GetSize() *int32 {
	return s.Size
}

func (s *DescribeSagDropTopNRequest) SetOwnerAccount(v string) *DescribeSagDropTopNRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagDropTopNRequest) SetOwnerId(v int64) *DescribeSagDropTopNRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagDropTopNRequest) SetRegionId(v string) *DescribeSagDropTopNRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagDropTopNRequest) SetResourceOwnerAccount(v string) *DescribeSagDropTopNRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagDropTopNRequest) SetResourceOwnerId(v int64) *DescribeSagDropTopNRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagDropTopNRequest) SetSize(v int32) *DescribeSagDropTopNRequest {
	s.Size = &v
	return s
}

func (s *DescribeSagDropTopNRequest) Validate() error {
	return dara.Validate(s)
}
