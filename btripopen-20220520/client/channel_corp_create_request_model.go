// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelCorpCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministratorEmail(v string) *ChannelCorpCreateRequest
	GetAdministratorEmail() *string
	SetAdministratorName(v string) *ChannelCorpCreateRequest
	GetAdministratorName() *string
	SetAdministratorPhone(v string) *ChannelCorpCreateRequest
	GetAdministratorPhone() *string
	SetBaseCurrency(v string) *ChannelCorpCreateRequest
	GetBaseCurrency() *string
	SetBtripRegion(v string) *ChannelCorpCreateRequest
	GetBtripRegion() *string
	SetCity(v string) *ChannelCorpCreateRequest
	GetCity() *string
	SetCorpName(v string) *ChannelCorpCreateRequest
	GetCorpName() *string
	SetCorpNameEn(v string) *ChannelCorpCreateRequest
	GetCorpNameEn() *string
	SetExtendField(v string) *ChannelCorpCreateRequest
	GetExtendField() *string
	SetProvince(v string) *ChannelCorpCreateRequest
	GetProvince() *string
	SetScope(v int32) *ChannelCorpCreateRequest
	GetScope() *int32
	SetThirdCorpId(v string) *ChannelCorpCreateRequest
	GetThirdCorpId() *string
	SetUserId(v string) *ChannelCorpCreateRequest
	GetUserId() *string
}

type ChannelCorpCreateRequest struct {
	AdministratorEmail *string `json:"administrator_email,omitempty" xml:"administrator_email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	AdministratorName *string `json:"administrator_name,omitempty" xml:"administrator_name,omitempty"`
	// example:
	//
	// 18378889782
	AdministratorPhone *string `json:"administrator_phone,omitempty" xml:"administrator_phone,omitempty"`
	BaseCurrency       *string `json:"base_currency,omitempty" xml:"base_currency,omitempty"`
	BtripRegion        *string `json:"btrip_region,omitempty" xml:"btrip_region,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 阿里巴巴有限公司
	CorpName    *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CorpNameEn  *string `json:"corp_name_en,omitempty" xml:"corp_name_en,omitempty"`
	ExtendField *string `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// 1
	Scope *int32 `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00001
	ThirdCorpId *string `json:"third_corp_id,omitempty" xml:"third_corp_id,omitempty"`
	// example:
	//
	// user_1234
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ChannelCorpCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ChannelCorpCreateRequest) GoString() string {
	return s.String()
}

func (s *ChannelCorpCreateRequest) GetAdministratorEmail() *string {
	return s.AdministratorEmail
}

func (s *ChannelCorpCreateRequest) GetAdministratorName() *string {
	return s.AdministratorName
}

func (s *ChannelCorpCreateRequest) GetAdministratorPhone() *string {
	return s.AdministratorPhone
}

func (s *ChannelCorpCreateRequest) GetBaseCurrency() *string {
	return s.BaseCurrency
}

func (s *ChannelCorpCreateRequest) GetBtripRegion() *string {
	return s.BtripRegion
}

func (s *ChannelCorpCreateRequest) GetCity() *string {
	return s.City
}

func (s *ChannelCorpCreateRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ChannelCorpCreateRequest) GetCorpNameEn() *string {
	return s.CorpNameEn
}

func (s *ChannelCorpCreateRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *ChannelCorpCreateRequest) GetProvince() *string {
	return s.Province
}

func (s *ChannelCorpCreateRequest) GetScope() *int32 {
	return s.Scope
}

func (s *ChannelCorpCreateRequest) GetThirdCorpId() *string {
	return s.ThirdCorpId
}

func (s *ChannelCorpCreateRequest) GetUserId() *string {
	return s.UserId
}

func (s *ChannelCorpCreateRequest) SetAdministratorEmail(v string) *ChannelCorpCreateRequest {
	s.AdministratorEmail = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetAdministratorName(v string) *ChannelCorpCreateRequest {
	s.AdministratorName = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetAdministratorPhone(v string) *ChannelCorpCreateRequest {
	s.AdministratorPhone = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetBaseCurrency(v string) *ChannelCorpCreateRequest {
	s.BaseCurrency = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetBtripRegion(v string) *ChannelCorpCreateRequest {
	s.BtripRegion = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetCity(v string) *ChannelCorpCreateRequest {
	s.City = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetCorpName(v string) *ChannelCorpCreateRequest {
	s.CorpName = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetCorpNameEn(v string) *ChannelCorpCreateRequest {
	s.CorpNameEn = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetExtendField(v string) *ChannelCorpCreateRequest {
	s.ExtendField = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetProvince(v string) *ChannelCorpCreateRequest {
	s.Province = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetScope(v int32) *ChannelCorpCreateRequest {
	s.Scope = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetThirdCorpId(v string) *ChannelCorpCreateRequest {
	s.ThirdCorpId = &v
	return s
}

func (s *ChannelCorpCreateRequest) SetUserId(v string) *ChannelCorpCreateRequest {
	s.UserId = &v
	return s
}

func (s *ChannelCorpCreateRequest) Validate() error {
	return dara.Validate(s)
}
