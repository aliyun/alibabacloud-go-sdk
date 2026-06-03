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
}

type QueryContactInfoResponseBody struct {
	Address                *string `json:"Address,omitempty" xml:"Address,omitempty"`
	City                   *string `json:"City,omitempty" xml:"City,omitempty"`
	Country                *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CreateDate             *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Email                  *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PostalCode             *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Province               *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RegistrantName         *string `json:"RegistrantName,omitempty" xml:"RegistrantName,omitempty"`
	RegistrantOrganization *string `json:"RegistrantOrganization,omitempty" xml:"RegistrantOrganization,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TelArea                *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt                 *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	Telephone              *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
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

func (s *QueryContactInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
