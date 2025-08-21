// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithThirdUserInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtInfoShrink(v string) *AuthLoginWithThirdUserInfoShrinkRequest
	GetExtInfoShrink() *string
	SetSceneCode(v string) *AuthLoginWithThirdUserInfoShrinkRequest
	GetSceneCode() *string
	SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoShrinkRequest
	GetThirdUserIdentifier() *string
	SetThirdUserType(v string) *AuthLoginWithThirdUserInfoShrinkRequest
	GetThirdUserType() *string
}

type AuthLoginWithThirdUserInfoShrinkRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// {}
	ExtInfoShrink *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// wechat_ecology_openness
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// o5qxy6EItZBasv2VZAf-MGwllHL4
	ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// weChatUser
	ThirdUserType *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
}

func (s AuthLoginWithThirdUserInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) GetExtInfoShrink() *string {
	return s.ExtInfoShrink
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) GetThirdUserIdentifier() *string {
	return s.ThirdUserIdentifier
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) GetThirdUserType() *string {
	return s.ThirdUserType
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetExtInfoShrink(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ExtInfoShrink = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetSceneCode(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.SceneCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) SetThirdUserType(v string) *AuthLoginWithThirdUserInfoShrinkRequest {
	s.ThirdUserType = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
