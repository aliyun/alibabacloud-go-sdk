// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v int64) *GetSmsSignRequest
	GetCustomerId() *int64
	SetOwnerId(v int64) *GetSmsSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsSignRequest
	GetResourceOwnerId() *int64
	SetSmsSignName(v string) *GetSmsSignRequest
	GetSmsSignName() *string
}

type GetSmsSignRequest struct {
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
}

func (s GetSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignRequest) GoString() string {
	return s.String()
}

func (s *GetSmsSignRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *GetSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsSignRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *GetSmsSignRequest) SetCustomerId(v int64) *GetSmsSignRequest {
	s.CustomerId = &v
	return s
}

func (s *GetSmsSignRequest) SetOwnerId(v int64) *GetSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsSignRequest) SetResourceOwnerAccount(v string) *GetSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsSignRequest) SetResourceOwnerId(v int64) *GetSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsSignRequest) SetSmsSignName(v string) *GetSmsSignRequest {
	s.SmsSignName = &v
	return s
}

func (s *GetSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
