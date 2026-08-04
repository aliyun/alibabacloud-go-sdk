// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrInsertEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetAddress() *string
	SetAlias(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetAlias() *string
	SetCityJsonString(v map[string]interface{}) *UpdateOrInsertEnterpriseInfoRequest
	GetCityJsonString() map[string]interface{}
	SetEnterpriseSize(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetEnterpriseSize() *string
	SetFax(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetFax() *string
	SetName(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetName() *string
	SetPK(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetPK() *string
	SetPhone(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetPhone() *string
	SetProvinceJsonString(v map[string]interface{}) *UpdateOrInsertEnterpriseInfoRequest
	GetProvinceJsonString() map[string]interface{}
	SetYears(v string) *UpdateOrInsertEnterpriseInfoRequest
	GetYears() *string
}

type UpdateOrInsertEnterpriseInfoRequest struct {
	Address            *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	Alias              *string                `json:"Alias,omitempty" xml:"Alias,omitempty"`
	CityJsonString     map[string]interface{} `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	EnterpriseSize     *string                `json:"EnterpriseSize,omitempty" xml:"EnterpriseSize,omitempty"`
	Fax                *string                `json:"Fax,omitempty" xml:"Fax,omitempty"`
	Name               *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	PK                 *string                `json:"PK,omitempty" xml:"PK,omitempty"`
	Phone              *string                `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ProvinceJsonString map[string]interface{} `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
	Years              *string                `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s UpdateOrInsertEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrInsertEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetAlias() *string {
	return s.Alias
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetCityJsonString() map[string]interface{} {
	return s.CityJsonString
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetEnterpriseSize() *string {
	return s.EnterpriseSize
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetFax() *string {
	return s.Fax
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetName() *string {
	return s.Name
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetProvinceJsonString() map[string]interface{} {
	return s.ProvinceJsonString
}

func (s *UpdateOrInsertEnterpriseInfoRequest) GetYears() *string {
	return s.Years
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetAddress(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Address = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetAlias(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Alias = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetCityJsonString(v map[string]interface{}) *UpdateOrInsertEnterpriseInfoRequest {
	s.CityJsonString = v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetEnterpriseSize(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.EnterpriseSize = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetFax(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Fax = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetName(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetPK(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.PK = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetPhone(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Phone = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetProvinceJsonString(v map[string]interface{}) *UpdateOrInsertEnterpriseInfoRequest {
	s.ProvinceJsonString = v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) SetYears(v string) *UpdateOrInsertEnterpriseInfoRequest {
	s.Years = &v
	return s
}

func (s *UpdateOrInsertEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}
