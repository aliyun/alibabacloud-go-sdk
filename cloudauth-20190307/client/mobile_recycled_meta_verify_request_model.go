// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecycledMetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *MobileRecycledMetaVerifyRequest
	GetMobile() *string
	SetParamType(v string) *MobileRecycledMetaVerifyRequest
	GetParamType() *string
	SetRegisterDate(v string) *MobileRecycledMetaVerifyRequest
	GetRegisterDate() *string
}

type MobileRecycledMetaVerifyRequest struct {
	// The phone number. Valid values:
	//
	// - If ParamType is set to normal, pass in the phone number in plaintext.
	//
	// - If ParamType is set to md5, pass in the MD5-encrypted phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 明文：186****2055
	//
	// 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: plaintext without encryption
	//
	// - md5: MD5 encryption.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The registration date in the format YYYYMMDD. For example, 19800101 indicates January 1, 1980.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20200505
	RegisterDate *string `json:"RegisterDate,omitempty" xml:"RegisterDate,omitempty"`
}

func (s MobileRecycledMetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileRecycledMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *MobileRecycledMetaVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *MobileRecycledMetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *MobileRecycledMetaVerifyRequest) GetRegisterDate() *string {
	return s.RegisterDate
}

func (s *MobileRecycledMetaVerifyRequest) SetMobile(v string) *MobileRecycledMetaVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) SetParamType(v string) *MobileRecycledMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) SetRegisterDate(v string) *MobileRecycledMetaVerifyRequest {
	s.RegisterDate = &v
	return s
}

func (s *MobileRecycledMetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
