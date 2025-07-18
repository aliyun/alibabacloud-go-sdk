// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpLoginConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMobileLoginType(v string) *IdpLoginConfig
	GetMobileLoginType() *string
	SetMobileMfaTypes(v string) *IdpLoginConfig
	GetMobileMfaTypes() *string
	SetPcLoginType(v string) *IdpLoginConfig
	GetPcLoginType() *string
	SetPcMfaTypes(v string) *IdpLoginConfig
	GetPcMfaTypes() *string
	SetTotpCorpVerifyAesKey(v string) *IdpLoginConfig
	GetTotpCorpVerifyAesKey() *string
	SetTotpCorpVerifyToken(v string) *IdpLoginConfig
	GetTotpCorpVerifyToken() *string
	SetTotpCorpVerifyUrl(v string) *IdpLoginConfig
	GetTotpCorpVerifyUrl() *string
}

type IdpLoginConfig struct {
	MobileLoginType      *string `json:"MobileLoginType,omitempty" xml:"MobileLoginType,omitempty"`
	MobileMfaTypes       *string `json:"MobileMfaTypes,omitempty" xml:"MobileMfaTypes,omitempty"`
	PcLoginType          *string `json:"PcLoginType,omitempty" xml:"PcLoginType,omitempty"`
	PcMfaTypes           *string `json:"PcMfaTypes,omitempty" xml:"PcMfaTypes,omitempty"`
	TotpCorpVerifyAesKey *string `json:"TotpCorpVerifyAesKey,omitempty" xml:"TotpCorpVerifyAesKey,omitempty"`
	TotpCorpVerifyToken  *string `json:"TotpCorpVerifyToken,omitempty" xml:"TotpCorpVerifyToken,omitempty"`
	TotpCorpVerifyUrl    *string `json:"TotpCorpVerifyUrl,omitempty" xml:"TotpCorpVerifyUrl,omitempty"`
}

func (s IdpLoginConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpLoginConfig) GoString() string {
	return s.String()
}

func (s *IdpLoginConfig) GetMobileLoginType() *string {
	return s.MobileLoginType
}

func (s *IdpLoginConfig) GetMobileMfaTypes() *string {
	return s.MobileMfaTypes
}

func (s *IdpLoginConfig) GetPcLoginType() *string {
	return s.PcLoginType
}

func (s *IdpLoginConfig) GetPcMfaTypes() *string {
	return s.PcMfaTypes
}

func (s *IdpLoginConfig) GetTotpCorpVerifyAesKey() *string {
	return s.TotpCorpVerifyAesKey
}

func (s *IdpLoginConfig) GetTotpCorpVerifyToken() *string {
	return s.TotpCorpVerifyToken
}

func (s *IdpLoginConfig) GetTotpCorpVerifyUrl() *string {
	return s.TotpCorpVerifyUrl
}

func (s *IdpLoginConfig) SetMobileLoginType(v string) *IdpLoginConfig {
	s.MobileLoginType = &v
	return s
}

func (s *IdpLoginConfig) SetMobileMfaTypes(v string) *IdpLoginConfig {
	s.MobileMfaTypes = &v
	return s
}

func (s *IdpLoginConfig) SetPcLoginType(v string) *IdpLoginConfig {
	s.PcLoginType = &v
	return s
}

func (s *IdpLoginConfig) SetPcMfaTypes(v string) *IdpLoginConfig {
	s.PcMfaTypes = &v
	return s
}

func (s *IdpLoginConfig) SetTotpCorpVerifyAesKey(v string) *IdpLoginConfig {
	s.TotpCorpVerifyAesKey = &v
	return s
}

func (s *IdpLoginConfig) SetTotpCorpVerifyToken(v string) *IdpLoginConfig {
	s.TotpCorpVerifyToken = &v
	return s
}

func (s *IdpLoginConfig) SetTotpCorpVerifyUrl(v string) *IdpLoginConfig {
	s.TotpCorpVerifyUrl = &v
	return s
}

func (s *IdpLoginConfig) Validate() error {
	return dara.Validate(s)
}
