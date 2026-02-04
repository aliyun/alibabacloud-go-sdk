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
	SetInstanceId(v string) *GenerateOauthTokenRequest
	GetInstanceId() *string
	SetScopeValues(v []*string) *GenerateOauthTokenRequest
	GetScopeValues() []*string
}

type GenerateOauthTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com
	Audience *string `json:"Audience,omitempty" xml:"Audience,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ScopeValues []*string `json:"ScopeValues,omitempty" xml:"ScopeValues,omitempty" type:"Repeated"`
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

func (s *GenerateOauthTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateOauthTokenRequest) GetScopeValues() []*string {
	return s.ScopeValues
}

func (s *GenerateOauthTokenRequest) SetApplicationId(v string) *GenerateOauthTokenRequest {
	s.ApplicationId = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetAudience(v string) *GenerateOauthTokenRequest {
	s.Audience = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetInstanceId(v string) *GenerateOauthTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateOauthTokenRequest) SetScopeValues(v []*string) *GenerateOauthTokenRequest {
	s.ScopeValues = v
	return s
}

func (s *GenerateOauthTokenRequest) Validate() error {
	return dara.Validate(s)
}
