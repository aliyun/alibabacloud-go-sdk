// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileOnlineTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *MobileOnlineTimeRequest
	GetMobile() *string
	SetParamType(v string) *MobileOnlineTimeRequest
	GetParamType() *string
}

type MobileOnlineTimeRequest struct {
	// Mobile number:
	//
	// - When `paramType` is `normal`: provide the plaintext mobile number.
	//
	// - When `paramType` is `md5`: provide the encrypted mobile number.
	//
	// example:
	//
	// 明文：186****2055
	//
	// 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Parameter type:
	//
	// - normal: unencrypted.
	//
	// - md5: md5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s MobileOnlineTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *MobileOnlineTimeRequest) GetMobile() *string {
	return s.Mobile
}

func (s *MobileOnlineTimeRequest) GetParamType() *string {
	return s.ParamType
}

func (s *MobileOnlineTimeRequest) SetMobile(v string) *MobileOnlineTimeRequest {
	s.Mobile = &v
	return s
}

func (s *MobileOnlineTimeRequest) SetParamType(v string) *MobileOnlineTimeRequest {
	s.ParamType = &v
	return s
}

func (s *MobileOnlineTimeRequest) Validate() error {
	return dara.Validate(s)
}
