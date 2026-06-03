// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyCallBackTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *CreateDypnsSmsVerifyCallBackTestRequest
	GetBizType() *string
	SetContent(v string) *CreateDypnsSmsVerifyCallBackTestRequest
	GetContent() *string
	SetMethod(v string) *CreateDypnsSmsVerifyCallBackTestRequest
	GetMethod() *string
	SetOwnerId(v int64) *CreateDypnsSmsVerifyCallBackTestRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyCallBackTestRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyCallBackTestRequest
	GetResourceOwnerId() *int64
	SetUrl(v string) *CreateDypnsSmsVerifyCallBackTestRequest
	GetUrl() *string
}

type CreateDypnsSmsVerifyCallBackTestRequest struct {
	// example:
	//
	// Register
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// {"aa":1,"b":"test"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// POST
	Method               *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDypnsSmsVerifyCallBackTestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyCallBackTestRequest) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetBizType(v string) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.BizType = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetContent(v string) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.Content = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetMethod(v string) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.Method = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetOwnerId(v int64) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) SetUrl(v string) *CreateDypnsSmsVerifyCallBackTestRequest {
	s.Url = &v
	return s
}

func (s *CreateDypnsSmsVerifyCallBackTestRequest) Validate() error {
	return dara.Validate(s)
}
