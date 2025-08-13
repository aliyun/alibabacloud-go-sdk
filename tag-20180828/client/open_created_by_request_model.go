// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCreatedByRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *OpenCreatedByRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenCreatedByRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenCreatedByRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenCreatedByRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *OpenCreatedByRequest
	GetResourceOwnerId() *string
}

type OpenCreatedByRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenCreatedByRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenCreatedByRequest) GoString() string {
	return s.String()
}

func (s *OpenCreatedByRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenCreatedByRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenCreatedByRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenCreatedByRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenCreatedByRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *OpenCreatedByRequest) SetOwnerAccount(v string) *OpenCreatedByRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenCreatedByRequest) SetOwnerId(v int64) *OpenCreatedByRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCreatedByRequest) SetRegionId(v string) *OpenCreatedByRequest {
	s.RegionId = &v
	return s
}

func (s *OpenCreatedByRequest) SetResourceOwnerAccount(v string) *OpenCreatedByRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenCreatedByRequest) SetResourceOwnerId(v string) *OpenCreatedByRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenCreatedByRequest) Validate() error {
	return dara.Validate(s)
}
