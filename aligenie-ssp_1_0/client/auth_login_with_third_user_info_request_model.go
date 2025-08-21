// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithThirdUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtInfo(v map[string]interface{}) *AuthLoginWithThirdUserInfoRequest
	GetExtInfo() map[string]interface{}
	SetSceneCode(v string) *AuthLoginWithThirdUserInfoRequest
	GetSceneCode() *string
	SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoRequest
	GetThirdUserIdentifier() *string
	SetThirdUserType(v string) *AuthLoginWithThirdUserInfoRequest
	GetThirdUserType() *string
}

type AuthLoginWithThirdUserInfoRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// {}
	ExtInfo map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
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

func (s AuthLoginWithThirdUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithThirdUserInfoRequest) GoString() string {
	return s.String()
}

func (s *AuthLoginWithThirdUserInfoRequest) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *AuthLoginWithThirdUserInfoRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *AuthLoginWithThirdUserInfoRequest) GetThirdUserIdentifier() *string {
	return s.ThirdUserIdentifier
}

func (s *AuthLoginWithThirdUserInfoRequest) GetThirdUserType() *string {
	return s.ThirdUserType
}

func (s *AuthLoginWithThirdUserInfoRequest) SetExtInfo(v map[string]interface{}) *AuthLoginWithThirdUserInfoRequest {
	s.ExtInfo = v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetSceneCode(v string) *AuthLoginWithThirdUserInfoRequest {
	s.SceneCode = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetThirdUserIdentifier(v string) *AuthLoginWithThirdUserInfoRequest {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) SetThirdUserType(v string) *AuthLoginWithThirdUserInfoRequest {
	s.ThirdUserType = &v
	return s
}

func (s *AuthLoginWithThirdUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
