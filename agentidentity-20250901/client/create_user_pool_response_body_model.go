// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserPoolResponseBody
	GetRequestId() *string
	SetUserPool(v *CreateUserPoolResponseBodyUserPool) *CreateUserPoolResponseBody
	GetUserPool() *CreateUserPoolResponseBodyUserPool
}

type CreateUserPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the newly created user pool.
	UserPool *CreateUserPoolResponseBodyUserPool `json:"UserPool,omitempty" xml:"UserPool,omitempty" type:"Struct"`
}

func (s CreateUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserPoolResponseBody) GetUserPool() *CreateUserPoolResponseBodyUserPool {
	return s.UserPool
}

func (s *CreateUserPoolResponseBody) SetRequestId(v string) *CreateUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserPoolResponseBody) SetUserPool(v *CreateUserPoolResponseBodyUserPool) *CreateUserPoolResponseBody {
	s.UserPool = v
	return s
}

func (s *CreateUserPoolResponseBody) Validate() error {
	if s.UserPool != nil {
		if err := s.UserPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserPoolResponseBodyUserPool struct {
	// The creation time in ISO 8601 format (UTC).
	//
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user pool.
	//
	// example:
	//
	// Sample user pool
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The discovery URL for the OpenID Connect (OIDC) protocol.
	//
	// example:
	//
	// https://signin-cn-beijing.aliyunagentid.com/up_7b6908bcb9ba4c****c8e187/.well-known/openid-configuration
	OIDCDiscoveryURL *string `json:"OIDCDiscoveryURL,omitempty" xml:"OIDCDiscoveryURL,omitempty"`
	// The base URL for logon.
	//
	// example:
	//
	// https://pre-signin-cn-hangzhou.aliyunagentid.com/up_e14d889*****716841
	SigninBaseURL *string `json:"SigninBaseURL,omitempty" xml:"SigninBaseURL,omitempty"`
	// The creation source.
	//
	// example:
	//
	// Manual
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The update time in ISO 8601 format (UTC).
	//
	// example:
	//
	// 2026-05-07T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user pool.
	//
	// example:
	//
	// up_xxxxxxxxxxxxxxxxxxxx
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserPoolResponseBodyUserPool) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolResponseBodyUserPool) GoString() string {
	return s.String()
}

func (s *CreateUserPoolResponseBodyUserPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateUserPoolResponseBodyUserPool) GetDescription() *string {
	return s.Description
}

func (s *CreateUserPoolResponseBodyUserPool) GetOIDCDiscoveryURL() *string {
	return s.OIDCDiscoveryURL
}

func (s *CreateUserPoolResponseBodyUserPool) GetSigninBaseURL() *string {
	return s.SigninBaseURL
}

func (s *CreateUserPoolResponseBodyUserPool) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreateUserPoolResponseBodyUserPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateUserPoolResponseBodyUserPool) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *CreateUserPoolResponseBodyUserPool) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolResponseBodyUserPool) SetCreateTime(v string) *CreateUserPoolResponseBodyUserPool {
	s.CreateTime = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetDescription(v string) *CreateUserPoolResponseBodyUserPool {
	s.Description = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetOIDCDiscoveryURL(v string) *CreateUserPoolResponseBodyUserPool {
	s.OIDCDiscoveryURL = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetSigninBaseURL(v string) *CreateUserPoolResponseBodyUserPool {
	s.SigninBaseURL = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetSourcePlatform(v string) *CreateUserPoolResponseBodyUserPool {
	s.SourcePlatform = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUpdateTime(v string) *CreateUserPoolResponseBodyUserPool {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUserPoolId(v string) *CreateUserPoolResponseBodyUserPool {
	s.UserPoolId = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUserPoolName(v string) *CreateUserPoolResponseBodyUserPool {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) Validate() error {
	return dara.Validate(s)
}
