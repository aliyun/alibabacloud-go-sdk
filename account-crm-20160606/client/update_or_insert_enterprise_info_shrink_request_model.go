// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrInsertEnterpriseInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetAddress() *string
	SetAlias(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetAlias() *string
	SetCityJsonStringShrink(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetCityJsonStringShrink() *string
	SetEnterpriseSize(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetEnterpriseSize() *string
	SetFax(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetFax() *string
	SetName(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetName() *string
	SetPK(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetPK() *string
	SetPhone(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetPhone() *string
	SetProvinceJsonStringShrink(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetProvinceJsonStringShrink() *string
	SetYears(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest
	GetYears() *string
}

type UpdateOrInsertEnterpriseInfoShrinkRequest struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Alias                    *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CityJsonStringShrink     *string `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	EnterpriseSize           *string `json:"EnterpriseSize,omitempty" xml:"EnterpriseSize,omitempty"`
	Fax                      *string `json:"Fax,omitempty" xml:"Fax,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PK                       *string `json:"PK,omitempty" xml:"PK,omitempty"`
	Phone                    *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProvinceJsonStringShrink *string `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
	Years                    *string `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s UpdateOrInsertEnterpriseInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrInsertEnterpriseInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetAlias() *string {
	return s.Alias
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetCityJsonStringShrink() *string {
	return s.CityJsonStringShrink
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetEnterpriseSize() *string {
	return s.EnterpriseSize
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetFax() *string {
	return s.Fax
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetProvinceJsonStringShrink() *string {
	return s.ProvinceJsonStringShrink
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) GetYears() *string {
	return s.Years
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetAddress(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetAlias(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Alias = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetCityJsonStringShrink(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.CityJsonStringShrink = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetEnterpriseSize(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.EnterpriseSize = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetFax(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Fax = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetName(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetPK(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.PK = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetPhone(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Phone = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetProvinceJsonStringShrink(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.ProvinceJsonStringShrink = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) SetYears(v string) *UpdateOrInsertEnterpriseInfoShrinkRequest {
	s.Years = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
