// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserPoolResponseBody
	GetRequestId() *string
	SetUserPool(v *GetUserPoolResponseBodyUserPool) *GetUserPoolResponseBody
	GetUserPool() *GetUserPoolResponseBodyUserPool
}

type GetUserPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user pool information.
	UserPool *GetUserPoolResponseBodyUserPool `json:"UserPool,omitempty" xml:"UserPool,omitempty" type:"Struct"`
}

func (s GetUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPoolResponseBody) GetUserPool() *GetUserPoolResponseBodyUserPool {
	return s.UserPool
}

func (s *GetUserPoolResponseBody) SetRequestId(v string) *GetUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPoolResponseBody) SetUserPool(v *GetUserPoolResponseBodyUserPool) *GetUserPoolResponseBody {
	s.UserPool = v
	return s
}

func (s *GetUserPoolResponseBody) Validate() error {
	if s.UserPool != nil {
		if err := s.UserPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPoolResponseBodyUserPool struct {
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
	// The source platform.
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

func (s GetUserPoolResponseBodyUserPool) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolResponseBodyUserPool) GoString() string {
	return s.String()
}

func (s *GetUserPoolResponseBodyUserPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserPoolResponseBodyUserPool) GetDescription() *string {
	return s.Description
}

func (s *GetUserPoolResponseBodyUserPool) GetOIDCDiscoveryURL() *string {
	return s.OIDCDiscoveryURL
}

func (s *GetUserPoolResponseBodyUserPool) GetSigninBaseURL() *string {
	return s.SigninBaseURL
}

func (s *GetUserPoolResponseBodyUserPool) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *GetUserPoolResponseBodyUserPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserPoolResponseBodyUserPool) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *GetUserPoolResponseBodyUserPool) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolResponseBodyUserPool) SetCreateTime(v string) *GetUserPoolResponseBodyUserPool {
	s.CreateTime = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetDescription(v string) *GetUserPoolResponseBodyUserPool {
	s.Description = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetOIDCDiscoveryURL(v string) *GetUserPoolResponseBodyUserPool {
	s.OIDCDiscoveryURL = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetSigninBaseURL(v string) *GetUserPoolResponseBodyUserPool {
	s.SigninBaseURL = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetSourcePlatform(v string) *GetUserPoolResponseBodyUserPool {
	s.SourcePlatform = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUpdateTime(v string) *GetUserPoolResponseBodyUserPool {
	s.UpdateTime = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUserPoolId(v string) *GetUserPoolResponseBodyUserPool {
	s.UserPoolId = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUserPoolName(v string) *GetUserPoolResponseBodyUserPool {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) Validate() error {
	return dara.Validate(s)
}
