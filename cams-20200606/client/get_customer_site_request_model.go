// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetCustomerSiteRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCustomerSiteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCustomerSiteRequest
	GetResourceOwnerId() *int64
}

type GetCustomerSiteRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCustomerSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerSiteRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerSiteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCustomerSiteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCustomerSiteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCustomerSiteRequest) SetOwnerId(v int64) *GetCustomerSiteRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCustomerSiteRequest) SetResourceOwnerAccount(v string) *GetCustomerSiteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCustomerSiteRequest) SetResourceOwnerId(v int64) *GetCustomerSiteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCustomerSiteRequest) Validate() error {
	return dara.Validate(s)
}
