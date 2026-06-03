// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageCallBackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreateDypnsSmsVerifyMessageCallBackRequest
	GetBizType() *string
	SetOwnerId(v int64) *CreateDypnsSmsVerifyMessageCallBackRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyMessageCallBackRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyMessageCallBackRequest
	GetResourceOwnerId() *int64
	SetUrl(v string) *CreateDypnsSmsVerifyMessageCallBackRequest
	GetUrl() *string
}

type CreateDypnsSmsVerifyMessageCallBackRequest struct {
	// example:
	//
	// Register
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDypnsSmsVerifyMessageCallBackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageCallBackRequest) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) SetBizType(v string) *CreateDypnsSmsVerifyMessageCallBackRequest {
	s.BizType = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) SetOwnerId(v int64) *CreateDypnsSmsVerifyMessageCallBackRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyMessageCallBackRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyMessageCallBackRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) SetUrl(v string) *CreateDypnsSmsVerifyMessageCallBackRequest {
	s.Url = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageCallBackRequest) Validate() error {
	return dara.Validate(s)
}
