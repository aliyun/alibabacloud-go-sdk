// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsSignRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *QuerySmsSignRequest
	GetSignName() *string
}

type QuerySmsSignRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s QuerySmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsSignRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *QuerySmsSignRequest) SetOwnerId(v int64) *QuerySmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsSignRequest) SetResourceOwnerAccount(v string) *QuerySmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsSignRequest) SetResourceOwnerId(v int64) *QuerySmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsSignRequest) SetSignName(v string) *QuerySmsSignRequest {
	s.SignName = &v
	return s
}

func (s *QuerySmsSignRequest) Validate() error {
	return dara.Validate(s)
}
