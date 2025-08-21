// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAuthCodeBindForExtShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CheckAuthCodeBindForExtShrinkRequest
	GetAuthCode() *string
	SetEncodeKey(v string) *CheckAuthCodeBindForExtShrinkRequest
	GetEncodeKey() *string
	SetEncodeType(v string) *CheckAuthCodeBindForExtShrinkRequest
	GetEncodeType() *string
	SetUserInfoShrink(v string) *CheckAuthCodeBindForExtShrinkRequest
	GetUserInfoShrink() *string
}

type CheckAuthCodeBindForExtShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Aexfgc
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CheckAuthCodeBindForExtShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtShrinkRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CheckAuthCodeBindForExtShrinkRequest) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *CheckAuthCodeBindForExtShrinkRequest) GetEncodeType() *string {
	return s.EncodeType
}

func (s *CheckAuthCodeBindForExtShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetAuthCode(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.AuthCode = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetEncodeKey(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.EncodeKey = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetEncodeType(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.EncodeType = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) SetUserInfoShrink(v string) *CheckAuthCodeBindForExtShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CheckAuthCodeBindForExtShrinkRequest) Validate() error {
	return dara.Validate(s)
}
