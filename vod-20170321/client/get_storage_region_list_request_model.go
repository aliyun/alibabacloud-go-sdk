// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageRegionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetStorageRegionListRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetStorageRegionListRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetStorageRegionListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetStorageRegionListRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetStorageRegionListRequest
	GetResourceRealOwnerId() *int64
}

type GetStorageRegionListRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s GetStorageRegionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageRegionListRequest) GoString() string {
	return s.String()
}

func (s *GetStorageRegionListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetStorageRegionListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetStorageRegionListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetStorageRegionListRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetStorageRegionListRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetStorageRegionListRequest) SetOwnerAccount(v string) *GetStorageRegionListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetStorageRegionListRequest) SetOwnerId(v string) *GetStorageRegionListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetStorageRegionListRequest) SetResourceOwnerAccount(v string) *GetStorageRegionListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetStorageRegionListRequest) SetResourceOwnerId(v string) *GetStorageRegionListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetStorageRegionListRequest) SetResourceRealOwnerId(v int64) *GetStorageRegionListRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetStorageRegionListRequest) Validate() error {
	return dara.Validate(s)
}
