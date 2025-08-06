// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVodServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *OpenVodServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *OpenVodServiceRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *OpenVodServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *OpenVodServiceRequest
	GetResourceOwnerId() *string
}

type OpenVodServiceRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenVodServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenVodServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenVodServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenVodServiceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *OpenVodServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenVodServiceRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *OpenVodServiceRequest) SetOwnerAccount(v string) *OpenVodServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenVodServiceRequest) SetOwnerId(v string) *OpenVodServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenVodServiceRequest) SetResourceOwnerAccount(v string) *OpenVodServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenVodServiceRequest) SetResourceOwnerId(v string) *OpenVodServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenVodServiceRequest) Validate() error {
	return dara.Validate(s)
}
