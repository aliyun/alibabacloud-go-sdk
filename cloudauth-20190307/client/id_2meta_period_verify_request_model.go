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
	// ID number:
	//
	// - When `paramType` is `normal`: Enter the plain text of the ID number.
	//
	// - When `paramType` is `md5`:
	//
	// The first 6 digits (plain text) + date of birth (encrypted) + last 4 digits (plain text).
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: MD5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Name.
	//
	// - When `paramType` = `normal`: Enter the plain text of the name.
	//
	// - When `paramType` = `md5`: The first character of the name MD5 encrypted (32 lowercase MD5) + the rest of the name in plain text.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// End date of ID validity, format: YYYYMMDD
	//
	// example:
	//
	// 20301001
	ValidityEndDate *string `json:"ValidityEndDate,omitempty" xml:"ValidityEndDate,omitempty"`
	// Start date of ID validity, format: YYYYMMDD
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
