// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOauthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GenerateOauthTokenRequest
	GetApplicationId() *string
	SetAudience(v string) *GenerateOauthTokenRequest
	GetAudience() *string
	SetGrantType(v string) *GenerateOauthTokenRequest
	GetGrantType() *string
	SetInstanceId(v string) *GenerateOauthTokenRequest
	GetInstanceId() *string
	SetRequestedTokenType(v string) *GenerateOauthTokenRequest
	GetRequestedTokenType() *string
	SetScopeValues(v []*string) *GenerateOauthTokenRequest
	GetScopeValues() []*string
	SetSubjectToken(v string) *GenerateOauthTokenRequest
	GetSubjectToken() *string
	SetSubjectTokenType(v string) *GenerateOauthTokenRequest
	GetSubjectTokenType() *string
}

type GenerateOauthTokenRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The audience identifier of the resource server.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com
	Audience  *string `json:"Audience,omitempty" xml:"Audience,omitempty"`
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestedTokenType *string `json:"RequestedTokenType,omitempty" xml:"RequestedTokenType,omitempty"`
	// The permission scopes.
	//
	// This parameter is required.
	ScopeValues      []*string `json:"ScopeValues,omitempty" xml:"ScopeValues,omitempty" type:"Repeated"`
	SubjectToken     *string   `json:"SubjectToken,omitempty" xml:"SubjectToken,omitempty"`
	SubjectTokenType *string   `json:"SubjectTokenType,omitempty" xml:"SubjectTokenType,omitempty"`
}

func (s GenerateOauthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOauthTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateOauthTokenRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GenerateOauthTokenRequest) GetAudience() *string {
	return s.Audience
}

func (s *GenerateOauthTokenRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GenerateOauthTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateOauthTokenRequest) GetRequestedTokenType() *string {
	return s.RequestedTokenType
}

func (s *GenerateOauthTokenRequest) GetScopeValues() []*string {
	return s.ScopeValues
}

func (s *GenerateOauthTokenRequest) GetSubjectToken() *string {
	return s.SubjectToken
}

func (s *GenerateOauthTokenRequest) GetSubjectTokenType() *string {
	return s.SubjectTokenType
}

func (s *GenerateOauthTokenRequest) SetApplicationId(v string) *GenerateOauthTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetAudience(v string) *GenerateOauthTokenRequest {
	s.Audience = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetGrantType(v string) *GenerateOauthTokenRequest {
	s.GrantType = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetInstanceId(v string) *GenerateOauthTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetRequestedTokenType(v string) *GenerateOauthTokenRequest {
	s.RequestedTokenType = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetScopeValues(v []*string) *GenerateOauthTokenRequest {
	s.ScopeValues = v
	return s
}

func (s *GenerateOauthTokenRequest) SetSubjectToken(v string) *GenerateOauthTokenRequest {
	s.SubjectToken = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetSubjectTokenType(v string) *GenerateOauthTokenRequest {
	s.SubjectTokenType = &v
	return s
}

func (s *GenerateOauthTokenRequest) Validate() error {
	return dara.Validate(s)
}
