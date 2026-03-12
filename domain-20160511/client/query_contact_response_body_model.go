// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCCity(v string) *QueryContactResponseBody
	GetCCity() *string
	SetCCompany(v string) *QueryContactResponseBody
	GetCCompany() *string
	SetCCountry(v string) *QueryContactResponseBody
	GetCCountry() *string
	SetCName(v string) *QueryContactResponseBody
	GetCName() *string
	SetCProvince(v string) *QueryContactResponseBody
	GetCProvince() *string
	SetCVenu(v string) *QueryContactResponseBody
	GetCVenu() *string
	SetCreateDate(v string) *QueryContactResponseBody
	GetCreateDate() *string
	SetECity(v string) *QueryContactResponseBody
	GetECity() *string
	SetECompany(v string) *QueryContactResponseBody
	GetECompany() *string
	SetEName(v string) *QueryContactResponseBody
	GetEName() *string
	SetEProvince(v string) *QueryContactResponseBody
	GetEProvince() *string
	SetEVenu(v string) *QueryContactResponseBody
	GetEVenu() *string
	SetEmail(v string) *QueryContactResponseBody
	GetEmail() *string
	SetPostalCode(v string) *QueryContactResponseBody
	GetPostalCode() *string
	SetRegType(v string) *QueryContactResponseBody
	GetRegType() *string
	SetRequestId(v string) *QueryContactResponseBody
	GetRequestId() *string
	SetTelArea(v string) *QueryContactResponseBody
	GetTelArea() *string
	SetTelExt(v string) *QueryContactResponseBody
	GetTelExt() *string
	SetTelMain(v string) *QueryContactResponseBody
	GetTelMain() *string
	SetUpdateDate(v string) *QueryContactResponseBody
	GetUpdateDate() *string
}

type QueryContactResponseBody struct {
	CCity      *string `json:"CCity,omitempty" xml:"CCity,omitempty"`
	CCompany   *string `json:"CCompany,omitempty" xml:"CCompany,omitempty"`
	CCountry   *string `json:"CCountry,omitempty" xml:"CCountry,omitempty"`
	CName      *string `json:"CName,omitempty" xml:"CName,omitempty"`
	CProvince  *string `json:"CProvince,omitempty" xml:"CProvince,omitempty"`
	CVenu      *string `json:"CVenu,omitempty" xml:"CVenu,omitempty"`
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	ECity      *string `json:"ECity,omitempty" xml:"ECity,omitempty"`
	ECompany   *string `json:"ECompany,omitempty" xml:"ECompany,omitempty"`
	EName      *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EProvince  *string `json:"EProvince,omitempty" xml:"EProvince,omitempty"`
	EVenu      *string `json:"EVenu,omitempty" xml:"EVenu,omitempty"`
	Email      *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PostalCode *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	RegType    *string `json:"RegType,omitempty" xml:"RegType,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TelArea    *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt     *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	TelMain    *string `json:"TelMain,omitempty" xml:"TelMain,omitempty"`
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s QueryContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryContactResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContactResponseBody) GetCCity() *string {
	return s.CCity
}

func (s *QueryContactResponseBody) GetCCompany() *string {
	return s.CCompany
}

func (s *QueryContactResponseBody) GetCCountry() *string {
	return s.CCountry
}

func (s *QueryContactResponseBody) GetCName() *string {
	return s.CName
}

func (s *QueryContactResponseBody) GetCProvince() *string {
	return s.CProvince
}

func (s *QueryContactResponseBody) GetCVenu() *string {
	return s.CVenu
}

func (s *QueryContactResponseBody) GetCreateDate() *string {
	return s.CreateDate
}

func (s *QueryContactResponseBody) GetECity() *string {
	return s.ECity
}

func (s *QueryContactResponseBody) GetECompany() *string {
	return s.ECompany
}

func (s *QueryContactResponseBody) GetEName() *string {
	return s.EName
}

func (s *QueryContactResponseBody) GetEProvince() *string {
	return s.EProvince
}

func (s *QueryContactResponseBody) GetEVenu() *string {
	return s.EVenu
}

func (s *QueryContactResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryContactResponseBody) GetPostalCode() *string {
	return s.PostalCode
}

func (s *QueryContactResponseBody) GetRegType() *string {
	return s.RegType
}

func (s *QueryContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryContactResponseBody) GetTelArea() *string {
	return s.TelArea
}

func (s *QueryContactResponseBody) GetTelExt() *string {
	return s.TelExt
}

func (s *QueryContactResponseBody) GetTelMain() *string {
	return s.TelMain
}

func (s *QueryContactResponseBody) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *QueryContactResponseBody) SetCCity(v string) *QueryContactResponseBody {
	s.CCity = &v
	return s
}

func (s *QueryContactResponseBody) SetCCompany(v string) *QueryContactResponseBody {
	s.CCompany = &v
	return s
}

func (s *QueryContactResponseBody) SetCCountry(v string) *QueryContactResponseBody {
	s.CCountry = &v
	return s
}

func (s *QueryContactResponseBody) SetCName(v string) *QueryContactResponseBody {
	s.CName = &v
	return s
}

func (s *QueryContactResponseBody) SetCProvince(v string) *QueryContactResponseBody {
	s.CProvince = &v
	return s
}

func (s *QueryContactResponseBody) SetCVenu(v string) *QueryContactResponseBody {
	s.CVenu = &v
	return s
}

func (s *QueryContactResponseBody) SetCreateDate(v string) *QueryContactResponseBody {
	s.CreateDate = &v
	return s
}

func (s *QueryContactResponseBody) SetECity(v string) *QueryContactResponseBody {
	s.ECity = &v
	return s
}

func (s *QueryContactResponseBody) SetECompany(v string) *QueryContactResponseBody {
	s.ECompany = &v
	return s
}

func (s *QueryContactResponseBody) SetEName(v string) *QueryContactResponseBody {
	s.EName = &v
	return s
}

func (s *QueryContactResponseBody) SetEProvince(v string) *QueryContactResponseBody {
	s.EProvince = &v
	return s
}

func (s *QueryContactResponseBody) SetEVenu(v string) *QueryContactResponseBody {
	s.EVenu = &v
	return s
}

func (s *QueryContactResponseBody) SetEmail(v string) *QueryContactResponseBody {
	s.Email = &v
	return s
}

func (s *QueryContactResponseBody) SetPostalCode(v string) *QueryContactResponseBody {
	s.PostalCode = &v
	return s
}

func (s *QueryContactResponseBody) SetRegType(v string) *QueryContactResponseBody {
	s.RegType = &v
	return s
}

func (s *QueryContactResponseBody) SetRequestId(v string) *QueryContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContactResponseBody) SetTelArea(v string) *QueryContactResponseBody {
	s.TelArea = &v
	return s
}

func (s *QueryContactResponseBody) SetTelExt(v string) *QueryContactResponseBody {
	s.TelExt = &v
	return s
}

func (s *QueryContactResponseBody) SetTelMain(v string) *QueryContactResponseBody {
	s.TelMain = &v
	return s
}

func (s *QueryContactResponseBody) SetUpdateDate(v string) *QueryContactResponseBody {
	s.UpdateDate = &v
	return s
}

func (s *QueryContactResponseBody) Validate() error {
	return dara.Validate(s)
}
