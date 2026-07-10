// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *DeleteWhitelistSettingRequest
	GetIds() *string
	SetLang(v string) *DeleteWhitelistSettingRequest
	GetLang() *string
	SetServiceCode(v string) *DeleteWhitelistSettingRequest
	GetServiceCode() *string
	SetSourceIp(v string) *DeleteWhitelistSettingRequest
	GetSourceIp() *string
}

type DeleteWhitelistSettingRequest struct {
	// The list of rule IDs to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// [6222001]
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The language of the user information to delete. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The service code of the ID Verification product. Set the value to **antcloudauth**.
	//
	// This parameter is required.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The source IP address of the visitor. CIDR format and IPv4 format are supported. Example: 10.0.3.0/24.
	//
	// example:
	//
	// 120.25.41.25
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteWhitelistSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistSettingRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteWhitelistSettingRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteWhitelistSettingRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DeleteWhitelistSettingRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteWhitelistSettingRequest) SetIds(v string) *DeleteWhitelistSettingRequest {
	s.Ids = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetLang(v string) *DeleteWhitelistSettingRequest {
	s.Lang = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetServiceCode(v string) *DeleteWhitelistSettingRequest {
	s.ServiceCode = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) SetSourceIp(v string) *DeleteWhitelistSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteWhitelistSettingRequest) Validate() error {
	return dara.Validate(s)
}
