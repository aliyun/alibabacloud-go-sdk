// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserIPSWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v int64) *ModifyUserIPSWhitelistRequest
	GetDirection() *int64
	SetIpVersion(v string) *ModifyUserIPSWhitelistRequest
	GetIpVersion() *string
	SetLang(v string) *ModifyUserIPSWhitelistRequest
	GetLang() *string
	SetListType(v int64) *ModifyUserIPSWhitelistRequest
	GetListType() *int64
	SetListValue(v string) *ModifyUserIPSWhitelistRequest
	GetListValue() *string
	SetSourceIp(v string) *ModifyUserIPSWhitelistRequest
	GetSourceIp() *string
	SetWhiteType(v int64) *ModifyUserIPSWhitelistRequest
	GetWhiteType() *int64
}

type ModifyUserIPSWhitelistRequest struct {
	// example:
	//
	// 1
	Direction *int64 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 2
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// example:
	//
	// 115.236.36.114/32
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 2
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ModifyUserIPSWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistRequest) GetDirection() *int64 {
	return s.Direction
}

func (s *ModifyUserIPSWhitelistRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ModifyUserIPSWhitelistRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyUserIPSWhitelistRequest) GetListType() *int64 {
	return s.ListType
}

func (s *ModifyUserIPSWhitelistRequest) GetListValue() *string {
	return s.ListValue
}

func (s *ModifyUserIPSWhitelistRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyUserIPSWhitelistRequest) GetWhiteType() *int64 {
	return s.WhiteType
}

func (s *ModifyUserIPSWhitelistRequest) SetDirection(v int64) *ModifyUserIPSWhitelistRequest {
	s.Direction = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetIpVersion(v string) *ModifyUserIPSWhitelistRequest {
	s.IpVersion = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetLang(v string) *ModifyUserIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetListType(v int64) *ModifyUserIPSWhitelistRequest {
	s.ListType = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetListValue(v string) *ModifyUserIPSWhitelistRequest {
	s.ListValue = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetSourceIp(v string) *ModifyUserIPSWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetWhiteType(v int64) *ModifyUserIPSWhitelistRequest {
	s.WhiteType = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
