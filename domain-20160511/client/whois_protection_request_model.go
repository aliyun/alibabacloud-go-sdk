// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhoisProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataContent(v string) *WhoisProtectionRequest
	GetDataContent() *string
	SetDataSource(v int32) *WhoisProtectionRequest
	GetDataSource() *int32
	SetLang(v string) *WhoisProtectionRequest
	GetLang() *string
	SetUserClientIp(v string) *WhoisProtectionRequest
	GetUserClientIp() *string
	SetWhoisProtect(v bool) *WhoisProtectionRequest
	GetWhoisProtect() *bool
}

type WhoisProtectionRequest struct {
	// This parameter is required.
	DataContent *string `json:"DataContent,omitempty" xml:"DataContent,omitempty"`
	// This parameter is required.
	DataSource   *int32  `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// This parameter is required.
	WhoisProtect *bool `json:"WhoisProtect,omitempty" xml:"WhoisProtect,omitempty"`
}

func (s WhoisProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s WhoisProtectionRequest) GoString() string {
	return s.String()
}

func (s *WhoisProtectionRequest) GetDataContent() *string {
	return s.DataContent
}

func (s *WhoisProtectionRequest) GetDataSource() *int32 {
	return s.DataSource
}

func (s *WhoisProtectionRequest) GetLang() *string {
	return s.Lang
}

func (s *WhoisProtectionRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *WhoisProtectionRequest) GetWhoisProtect() *bool {
	return s.WhoisProtect
}

func (s *WhoisProtectionRequest) SetDataContent(v string) *WhoisProtectionRequest {
	s.DataContent = &v
	return s
}

func (s *WhoisProtectionRequest) SetDataSource(v int32) *WhoisProtectionRequest {
	s.DataSource = &v
	return s
}

func (s *WhoisProtectionRequest) SetLang(v string) *WhoisProtectionRequest {
	s.Lang = &v
	return s
}

func (s *WhoisProtectionRequest) SetUserClientIp(v string) *WhoisProtectionRequest {
	s.UserClientIp = &v
	return s
}

func (s *WhoisProtectionRequest) SetWhoisProtect(v bool) *WhoisProtectionRequest {
	s.WhoisProtect = &v
	return s
}

func (s *WhoisProtectionRequest) Validate() error {
	return dara.Validate(s)
}
