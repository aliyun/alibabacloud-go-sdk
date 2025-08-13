// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseCreatedByRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *CloseCreatedByRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CloseCreatedByRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CloseCreatedByRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CloseCreatedByRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *CloseCreatedByRequest
	GetResourceOwnerId() *string
}

type CloseCreatedByRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloseCreatedByRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseCreatedByRequest) GoString() string {
	return s.String()
}

func (s *CloseCreatedByRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CloseCreatedByRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloseCreatedByRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseCreatedByRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloseCreatedByRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *CloseCreatedByRequest) SetOwnerAccount(v string) *CloseCreatedByRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseCreatedByRequest) SetOwnerId(v int64) *CloseCreatedByRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseCreatedByRequest) SetRegionId(v string) *CloseCreatedByRequest {
	s.RegionId = &v
	return s
}

func (s *CloseCreatedByRequest) SetResourceOwnerAccount(v string) *CloseCreatedByRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseCreatedByRequest) SetResourceOwnerId(v string) *CloseCreatedByRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseCreatedByRequest) Validate() error {
	return dara.Validate(s)
}
