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
	// 	- AUTH_CODE: A verification code is used for authorization.
	//
	// 	- RESOURCE_DIRECTORY: A resource directory is used for authorization.
	//
	// Default value: AUTH_CODE.
	//
	// example:
	//
	// AUTH_CODE
	AuthChannel *string `json:"AuthChannel,omitempty" xml:"AuthChannel,omitempty"`
	// The verification code.
	//
	// >
	//
	// 	- The specified authentication code is used if the value of AuthChannel is left empty or is set to AUTH_CODE.
	//
	// 	- In other cases, a random 6-digit number is used. Example: 123456.
	//
	// example:
	//
	// 123456
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization
	//
	// 	- CLOUD_PRODUCT: cloud service-related authorization
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account to which the permissions on the resources are granted.
	//
	// >  You can set an effective scope across accounts only by using an Alibaba Cloud account instead of a RAM user. You can set an effective scope across accounts registered on the same site. For example, you can perform the operation across accounts that are both registered on the Alibaba Cloud China site or Alibaba Cloud international site. You cannot set an effective scope across accounts registered on different sites. For example, you cannot perform the operation across accounts that are separately registered on the Alibaba Cloud China site and Alibaba Cloud international site.
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
