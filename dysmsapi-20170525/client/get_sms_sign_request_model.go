// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSmsSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *GetSmsSignRequest
	GetSignName() *string
}

type GetSmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature name. The signature must be applied for by your account.
	//
	// - After you call the [CreateSmsSign](https://help.aliyun.com/document_detail/2807427.html) operation, obtain the signature name from the response.
	//
	// - View the signature on the [Signatures](https://dysms.console.aliyun.com/domestic/text/sign) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s GetSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignRequest) GoString() string {
	return s.String()
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

func (s *GetSmsSignRequest) GetSignName() *string {
	return s.SignName
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

func (s *GetSmsSignRequest) SetSignName(v string) *GetSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
