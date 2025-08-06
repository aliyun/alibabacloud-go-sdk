// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetCustomerConfigRequest
	GetAppId() *string
	SetOwnerId(v int64) *GetCustomerConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCustomerConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCustomerConfigRequest
	GetResourceOwnerId() *int64
}

type GetCustomerConfigRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCustomerConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetCustomerConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCustomerConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCustomerConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCustomerConfigRequest) SetAppId(v string) *GetCustomerConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetCustomerConfigRequest) SetOwnerId(v int64) *GetCustomerConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCustomerConfigRequest) SetResourceOwnerAccount(v string) *GetCustomerConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCustomerConfigRequest) SetResourceOwnerId(v int64) *GetCustomerConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCustomerConfigRequest) Validate() error {
	return dara.Validate(s)
}
