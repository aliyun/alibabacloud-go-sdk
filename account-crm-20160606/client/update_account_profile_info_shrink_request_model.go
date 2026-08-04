// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountProfileInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAttribute(v string) *UpdateAccountProfileInfoShrinkRequest
	GetAccountAttribute() *string
	SetAddress(v string) *UpdateAccountProfileInfoShrinkRequest
	GetAddress() *string
	SetAddress2(v string) *UpdateAccountProfileInfoShrinkRequest
	GetAddress2() *string
	SetBindAlipayNo(v string) *UpdateAccountProfileInfoShrinkRequest
	GetBindAlipayNo() *string
	SetCertType(v string) *UpdateAccountProfileInfoShrinkRequest
	GetCertType() *string
	SetCityJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest
	GetCityJsonStringShrink() *string
	SetContactMethod(v string) *UpdateAccountProfileInfoShrinkRequest
	GetContactMethod() *string
	SetDistrictJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest
	GetDistrictJsonStringShrink() *string
	SetFax(v string) *UpdateAccountProfileInfoShrinkRequest
	GetFax() *string
	SetFirstName(v string) *UpdateAccountProfileInfoShrinkRequest
	GetFirstName() *string
	SetHead(v string) *UpdateAccountProfileInfoShrinkRequest
	GetHead() *string
	SetHeadColor(v string) *UpdateAccountProfileInfoShrinkRequest
	GetHeadColor() *string
	SetLastName(v string) *UpdateAccountProfileInfoShrinkRequest
	GetLastName() *string
	SetPK(v string) *UpdateAccountProfileInfoShrinkRequest
	GetPK() *string
	SetPhone(v string) *UpdateAccountProfileInfoShrinkRequest
	GetPhone() *string
	SetPostCode(v string) *UpdateAccountProfileInfoShrinkRequest
	GetPostCode() *string
	SetProvinceJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest
	GetProvinceJsonStringShrink() *string
	SetSelfServicingBusinessRegNum(v string) *UpdateAccountProfileInfoShrinkRequest
	GetSelfServicingBusinessRegNum() *string
	SetSelfServicingIdentificationNum(v string) *UpdateAccountProfileInfoShrinkRequest
	GetSelfServicingIdentificationNum() *string
	SetTrueName(v string) *UpdateAccountProfileInfoShrinkRequest
	GetTrueName() *string
}

type UpdateAccountProfileInfoShrinkRequest struct {
	AccountAttribute               *string `json:"AccountAttribute,omitempty" xml:"AccountAttribute,omitempty"`
	Address                        *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2                       *string `json:"Address2,omitempty" xml:"Address2,omitempty"`
	BindAlipayNo                   *string `json:"BindAlipayNo,omitempty" xml:"BindAlipayNo,omitempty"`
	CertType                       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CityJsonStringShrink           *string `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	ContactMethod                  *string `json:"ContactMethod,omitempty" xml:"ContactMethod,omitempty"`
	DistrictJsonStringShrink       *string `json:"DistrictJsonString,omitempty" xml:"DistrictJsonString,omitempty"`
	Fax                            *string `json:"Fax,omitempty" xml:"Fax,omitempty"`
	FirstName                      *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	Head                           *string `json:"Head,omitempty" xml:"Head,omitempty"`
	HeadColor                      *string `json:"HeadColor,omitempty" xml:"HeadColor,omitempty"`
	LastName                       *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	PK                             *string `json:"PK,omitempty" xml:"PK,omitempty"`
	Phone                          *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PostCode                       *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	ProvinceJsonStringShrink       *string `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
	SelfServicingBusinessRegNum    *string `json:"SelfServicingBusinessRegNum,omitempty" xml:"SelfServicingBusinessRegNum,omitempty"`
	SelfServicingIdentificationNum *string `json:"SelfServicingIdentificationNum,omitempty" xml:"SelfServicingIdentificationNum,omitempty"`
	TrueName                       *string `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
}

func (s UpdateAccountProfileInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountProfileInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetAccountAttribute() *string {
	return s.AccountAttribute
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetAddress2() *string {
	return s.Address2
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetBindAlipayNo() *string {
	return s.BindAlipayNo
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetCertType() *string {
	return s.CertType
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetCityJsonStringShrink() *string {
	return s.CityJsonStringShrink
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetContactMethod() *string {
	return s.ContactMethod
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetDistrictJsonStringShrink() *string {
	return s.DistrictJsonStringShrink
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetFax() *string {
	return s.Fax
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetFirstName() *string {
	return s.FirstName
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetHead() *string {
	return s.Head
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetHeadColor() *string {
	return s.HeadColor
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetLastName() *string {
	return s.LastName
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetProvinceJsonStringShrink() *string {
	return s.ProvinceJsonStringShrink
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetSelfServicingBusinessRegNum() *string {
	return s.SelfServicingBusinessRegNum
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetSelfServicingIdentificationNum() *string {
	return s.SelfServicingIdentificationNum
}

func (s *UpdateAccountProfileInfoShrinkRequest) GetTrueName() *string {
	return s.TrueName
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetAccountAttribute(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.AccountAttribute = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetAddress(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.Address = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetAddress2(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.Address2 = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetBindAlipayNo(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.BindAlipayNo = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetCertType(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.CertType = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetCityJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.CityJsonStringShrink = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetContactMethod(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.ContactMethod = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetDistrictJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.DistrictJsonStringShrink = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetFax(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.Fax = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetFirstName(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.FirstName = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetHead(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.Head = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetHeadColor(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.HeadColor = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetLastName(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.LastName = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetPK(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.PK = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetPhone(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.Phone = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetPostCode(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetProvinceJsonStringShrink(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.ProvinceJsonStringShrink = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetSelfServicingBusinessRegNum(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.SelfServicingBusinessRegNum = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetSelfServicingIdentificationNum(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.SelfServicingIdentificationNum = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) SetTrueName(v string) *UpdateAccountProfileInfoShrinkRequest {
	s.TrueName = &v
	return s
}

func (s *UpdateAccountProfileInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
