// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaPeriodVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifyNum(v string) *Id2MetaPeriodVerifyRequest
	GetIdentifyNum() *string
	SetParamType(v string) *Id2MetaPeriodVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Id2MetaPeriodVerifyRequest
	GetUserName() *string
	SetValidityEndDate(v string) *Id2MetaPeriodVerifyRequest
	GetValidityEndDate() *string
	SetValidityStartDate(v string) *Id2MetaPeriodVerifyRequest
	GetValidityStartDate() *string
}

type Id2MetaPeriodVerifyRequest struct {
	// The ID card number.
	//
	// - If paramType is set to normal, enter the ID card number in plaintext.
	//
	// - If paramType is set to md5, the value is in the following format: first 6 digits of the ID card number (plaintext) + date of birth (ciphertext) + last 4 digits of the ID card number (plaintext).
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
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
	// The name.
	//
	// - paramType = normal: enter the name in plaintext.
	//
	// - paramType = md5: the first character of the name is MD5-encrypted (32-bit lowercase MD5) + the remaining characters of the name in plaintext.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The expiration date of the ID card validity period. Format: YYYYMMDD.
	//
	// example:
	//
	// 20301001
	ValidityEndDate *string `json:"ValidityEndDate,omitempty" xml:"ValidityEndDate,omitempty"`
	// The start date of the ID card validity period. Format: YYYYMMDD.
	//
	// example:
	//
	// 20201001
	ValidityStartDate *string `json:"ValidityStartDate,omitempty" xml:"ValidityStartDate,omitempty"`
}

func (s Id2MetaPeriodVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaPeriodVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaPeriodVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id2MetaPeriodVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id2MetaPeriodVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Id2MetaPeriodVerifyRequest) GetValidityEndDate() *string {
	return s.ValidityEndDate
}

func (s *Id2MetaPeriodVerifyRequest) GetValidityStartDate() *string {
	return s.ValidityStartDate
}

func (s *Id2MetaPeriodVerifyRequest) SetIdentifyNum(v string) *Id2MetaPeriodVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id2MetaPeriodVerifyRequest) SetParamType(v string) *Id2MetaPeriodVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id2MetaPeriodVerifyRequest) SetUserName(v string) *Id2MetaPeriodVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Id2MetaPeriodVerifyRequest) SetValidityEndDate(v string) *Id2MetaPeriodVerifyRequest {
	s.ValidityEndDate = &v
	return s
}

func (s *Id2MetaPeriodVerifyRequest) SetValidityStartDate(v string) *Id2MetaPeriodVerifyRequest {
	s.ValidityStartDate = &v
	return s
}

func (s *Id2MetaPeriodVerifyRequest) Validate() error {
	return dara.Validate(s)
}
