// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginStateInfo interface {
	dara.Model
	String() string
	GoString() string
	SetSceneCode(v string) *LoginStateInfo
	GetSceneCode() *string
	SetThirdUserIdentifier(v string) *LoginStateInfo
	GetThirdUserIdentifier() *string
	SetThirdUserType(v string) *LoginStateInfo
	GetThirdUserType() *string
	SetUserId(v string) *LoginStateInfo
	GetUserId() *string
}

type LoginStateInfo struct {
	SceneCode           *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ThirdUserIdentifier *string `json:"ThirdUserIdentifier,omitempty" xml:"ThirdUserIdentifier,omitempty"`
	ThirdUserType       *string `json:"ThirdUserType,omitempty" xml:"ThirdUserType,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LoginStateInfo) String() string {
	return dara.Prettify(s)
}

func (s LoginStateInfo) GoString() string {
	return s.String()
}

func (s *LoginStateInfo) GetSceneCode() *string {
	return s.SceneCode
}

func (s *LoginStateInfo) GetThirdUserIdentifier() *string {
	return s.ThirdUserIdentifier
}

func (s *LoginStateInfo) GetThirdUserType() *string {
	return s.ThirdUserType
}

func (s *LoginStateInfo) GetUserId() *string {
	return s.UserId
}

func (s *LoginStateInfo) SetSceneCode(v string) *LoginStateInfo {
	s.SceneCode = &v
	return s
}

func (s *LoginStateInfo) SetThirdUserIdentifier(v string) *LoginStateInfo {
	s.ThirdUserIdentifier = &v
	return s
}

func (s *LoginStateInfo) SetThirdUserType(v string) *LoginStateInfo {
	s.ThirdUserType = &v
	return s
}

func (s *LoginStateInfo) SetUserId(v string) *LoginStateInfo {
	s.UserId = &v
	return s
}

func (s *LoginStateInfo) Validate() error {
	return dara.Validate(s)
}
