// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodServiceRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetVodServiceRegionRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetVodServiceRegionRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetVodServiceRegionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetVodServiceRegionRequest
	GetResourceOwnerId() *string
}

type GetVodServiceRegionRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetVodServiceRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVodServiceRegionRequest) GoString() string {
	return s.String()
}

func (s *GetVodServiceRegionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVodServiceRegionRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetVodServiceRegionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVodServiceRegionRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetVodServiceRegionRequest) SetOwnerAccount(v string) *GetVodServiceRegionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVodServiceRegionRequest) SetOwnerId(v string) *GetVodServiceRegionRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVodServiceRegionRequest) SetResourceOwnerAccount(v string) *GetVodServiceRegionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVodServiceRegionRequest) SetResourceOwnerId(v string) *GetVodServiceRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVodServiceRegionRequest) Validate() error {
	return dara.Validate(s)
}
