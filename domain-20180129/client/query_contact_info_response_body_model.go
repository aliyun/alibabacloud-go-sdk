// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *QueryContactInfoResponseBody
	GetAddress() *string
	SetCity(v string) *QueryContactInfoResponseBody
	GetCity() *string
	SetCountry(v string) *QueryContactInfoResponseBody
	GetCountry() *string
	SetCreateDate(v string) *QueryContactInfoResponseBody
	GetCreateDate() *string
	SetEmail(v string) *QueryContactInfoResponseBody
	GetEmail() *string
	SetPostalCode(v string) *QueryContactInfoResponseBody
	GetPostalCode() *string
	SetProvince(v string) *QueryContactInfoResponseBody
	GetProvince() *string
	SetRegistrantName(v string) *QueryContactInfoResponseBody
	GetRegistrantName() *string
	SetRegistrantOrganization(v string) *QueryContactInfoResponseBody
	GetRegistrantOrganization() *string
	SetRequestId(v string) *QueryContactInfoResponseBody
	GetRequestId() *string
	SetTelArea(v string) *QueryContactInfoResponseBody
	GetTelArea() *string
	SetTelExt(v string) *QueryContactInfoResponseBody
	GetTelExt() *string
	SetTelephone(v string) *QueryContactInfoResponseBody
	GetTelephone() *string
	SetZhAddress(v string) *QueryContactInfoResponseBody
	GetZhAddress() *string
	SetZhCity(v string) *QueryContactInfoResponseBody
	GetZhCity() *string
	SetZhProvince(v string) *QueryContactInfoResponseBody
	GetZhProvince() *string
	SetZhRegistrantName(v string) *QueryContactInfoResponseBody
	GetZhRegistrantName() *string
	SetZhRegistrantOrganization(v string) *QueryContactInfoResponseBody
	GetZhRegistrantOrganization() *string
}

type QueryContactInfoResponseBody struct {
	// example:
	//
	// xi hu qu **	- jiedao **	- xiaoqu **	- zhuang 101
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// hang zhou shi
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 2019-03-20 11:37:29
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 310024
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	// example:
	//
	// zhe jiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// zhang san
	RegistrantName *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	// example:
	//
	// zhang san
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	// example:
	//
	// C39ECA8A-BB5E-4F92-B013-6A032FA06B04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 86
	TelArea *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	// example:
	//
	// 1234
	TelExt *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	// example:
	//
	// 1820000****
	Telephone                *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	ZhAddress                *string `json:"ZhAddress,omitempty" xml:"ZhAddress,omitempty"`
	ZhCity                   *string `json:"ZhCity,omitempty" xml:"ZhCity,omitempty"`
	ZhProvince               *string `json:"ZhProvince,omitempty" xml:"ZhProvince,omitempty"`
	ZhRegistrantName         *string `json:"ZhRegistrantName,omitempty" xml:"ZhRegistrantName,omitempty"`
	ZhRegistrantOrganization *string `json:"ZhRegistrantOrganization,omitempty" xml:"ZhRegistrantOrganization,omitempty"`
}

func (s QueryContactInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContactInfoResponseBody) GetAddress() *string {
	return s.Address
}

func (s *QueryContactInfoResponseBody) GetCity() *string {
	return s.City
}

func (s *QueryContactInfoResponseBody) GetCountry() *string {
	return s.Country
}

func (s *QueryContactInfoResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QueryContactInfoResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryContactInfoResponseBody) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryContactInfoResponseBody) GetProvince() *string {
	return s.Province
}

func (s *QueryContactInfoResponseBody) GetRegistrantName() *string {
	return s.RegistrantName
}

func (s *QueryContactInfoResponseBody) GetRegistrantOrganization() *string {
	return s.RegistrantOrganization
}

func (s *QueryContactInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryContactInfoResponseBody) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryContactInfoResponseBody) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryContactInfoResponseBody) GetTelephone() *string {
	return s.Telephone
}

func (s *QueryContactInfoResponseBody) GetZhAddress() *string {
	return s.ZhAddress
}

func (s *QueryContactInfoResponseBody) GetZhCity() *string {
	return s.ZhCity
}

func (s *QueryContactInfoResponseBody) GetZhProvince() *string {
	return s.ZhProvince
}

func (s *QueryContactInfoResponseBody) GetZhRegistrantName() *string {
	return s.ZhRegistrantName
}

func (s *QueryContactInfoResponseBody) GetZhRegistrantOrganization() *string {
	return s.ZhRegistrantOrganization
}

func (s *QueryContactInfoResponseBody) SetAddress(v string) *QueryContactInfoResponseBody {
	s.Address = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCity(v string) *QueryContactInfoResponseBody {
	s.City = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCountry(v string) *QueryContactInfoResponseBody {
	s.Country = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetCreateDate(v string) *QueryContactInfoResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetEmail(v string) *QueryContactInfoResponseBody {
	s.Email = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetPostalCode(v string) *QueryContactInfoResponseBody {
	s.PostalCode = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetProvince(v string) *QueryContactInfoResponseBody {
	s.Province = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRegistrantName(v string) *QueryContactInfoResponseBody {
	s.RegistrantName = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRegistrantOrganization(v string) *QueryContactInfoResponseBody {
	s.RegistrantOrganization = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetRequestId(v string) *QueryContactInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelArea(v string) *QueryContactInfoResponseBody {
	s.TelArea = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelExt(v string) *QueryContactInfoResponseBody {
	s.TelExt = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetTelephone(v string) *QueryContactInfoResponseBody {
	s.Telephone = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhAddress(v string) *QueryContactInfoResponseBody {
	s.ZhAddress = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhCity(v string) *QueryContactInfoResponseBody {
	s.ZhCity = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhProvince(v string) *QueryContactInfoResponseBody {
	s.ZhProvince = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhRegistrantName(v string) *QueryContactInfoResponseBody {
	s.ZhRegistrantName = &v
	return s
}

func (s *QueryContactInfoResponseBody) SetZhRegistrantOrganization(v string) *QueryContactInfoResponseBody {
	s.ZhRegistrantOrganization = &v
	return s
}

func (s *QueryContactInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
