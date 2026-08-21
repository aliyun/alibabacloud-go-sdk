// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserVpcAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthChannel(v string) *AddUserVpcAuthorizationRequest
	GetAuthChannel() *string
	SetAuthCode(v string) *AddUserVpcAuthorizationRequest
	GetAuthCode() *string
	SetAuthType(v string) *AddUserVpcAuthorizationRequest
	GetAuthType() *string
	SetAuthorizedUserId(v int64) *AddUserVpcAuthorizationRequest
	GetAuthorizedUserId() *int64
}

type AddUserVpcAuthorizationRequest struct {
	// The authorization channel. Valid values:
	//
	// - AUTH_CODE: verification code authorization.
	//
	// - RESOURCE_DIRECTORY: resource directory authorization.
	//
	// Default value: AUTH_CODE.
	//
	// example:
	//
	// AUTH_CODE
	AuthChannel *string `json:"AuthChannel,omitempty" xml:"AuthChannel,omitempty"`
	// The verification code.
	//
	// > - If AuthChannel is empty or set to AUTH_CODE, specify the verification code.
	//
	// > - In other cases, specify a random 6-digit number, such as 123456.
	//
	// example:
	//
	// 123456
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The authorization type. Valid values:
	//
	// - NORMAL: normal authorization.
	//
	// - CLOUD_PRODUCT: cloud product authorization.
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account that owns the authorized resource.
	//
	// >Cross-account authorization only supports Alibaba Cloud accounts (primary accounts) and does not support RAM users. Only accounts within the same site can be associated, such as between Alibaba Cloud China Website (www.aliyun.com) accounts or between Alibaba Cloud International Website (www.alibabacloud.com) accounts. Cross-site association is not supported, such as between a China Website account and an International Website account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 141339776561****
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
}

func (s AddUserVpcAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserVpcAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *AddUserVpcAuthorizationRequest) GetAuthChannel() *string {
	return s.AuthChannel
}

func (s *AddUserVpcAuthorizationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *AddUserVpcAuthorizationRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *AddUserVpcAuthorizationRequest) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *AddUserVpcAuthorizationRequest) SetAuthChannel(v string) *AddUserVpcAuthorizationRequest {
	s.AuthChannel = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthCode(v string) *AddUserVpcAuthorizationRequest {
	s.AuthCode = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthType(v string) *AddUserVpcAuthorizationRequest {
	s.AuthType = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) SetAuthorizedUserId(v int64) *AddUserVpcAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *AddUserVpcAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
