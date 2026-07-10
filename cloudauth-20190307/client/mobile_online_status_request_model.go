// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *MobileOnlineStatusRequest
	GetMobile() *string
	SetParamType(v string) *MobileOnlineStatusRequest
	GetParamType() *string
}

type MobileOnlineStatusRequest struct {
	// The phone number. Valid values:
	//
	// - If paramType is set to normal, pass in the phone number in plaintext.
	//
	// - If paramType is set to md5, pass in the MD5-encrypted phone number.
	//
	// example:
	//
	// 13665148158
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The parameter type. Valid values:
	//
	// - normal: not encrypted.
	//
	// - md5: MD5-encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileOnlineStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineStatusRequest) GoString() string {
	return s.String()
}

func (s *MobileOnlineStatusRequest) GetMobile() *string {
	return s.Mobile
}

func (s *MobileOnlineStatusRequest) GetParamType() *string {
	return s.ParamType
}

func (s *MobileOnlineStatusRequest) SetMobile(v string) *MobileOnlineStatusRequest {
	s.Mobile = &v
	return s
}

func (s *MobileOnlineStatusRequest) SetParamType(v string) *MobileOnlineStatusRequest {
	s.ParamType = &v
	return s
}

func (s *MobileOnlineStatusRequest) Validate() error {
	return dara.Validate(s)
}
