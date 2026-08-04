// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountProfileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAttribute(v string) *UpdateAccountProfileInfoRequest
	GetAccountAttribute() *string
	SetAddress(v string) *UpdateAccountProfileInfoRequest
	GetAddress() *string
	SetAddress2(v string) *UpdateAccountProfileInfoRequest
	GetAddress2() *string
	SetBindAlipayNo(v string) *UpdateAccountProfileInfoRequest
	GetBindAlipayNo() *string
	SetCertType(v string) *UpdateAccountProfileInfoRequest
	GetCertType() *string
	SetCityJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest
	GetCityJsonString() map[string]interface{}
	SetContactMethod(v string) *UpdateAccountProfileInfoRequest
	GetContactMethod() *string
	SetDistrictJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest
	GetDistrictJsonString() map[string]interface{}
	SetFax(v string) *UpdateAccountProfileInfoRequest
	GetFax() *string
	SetFirstName(v string) *UpdateAccountProfileInfoRequest
	GetFirstName() *string
	SetHead(v string) *UpdateAccountProfileInfoRequest
	GetHead() *string
	SetHeadColor(v string) *UpdateAccountProfileInfoRequest
	GetHeadColor() *string
	SetLastName(v string) *UpdateAccountProfileInfoRequest
	GetLastName() *string
	SetPK(v string) *UpdateAccountProfileInfoRequest
	GetPK() *string
	SetPhone(v string) *UpdateAccountProfileInfoRequest
	GetPhone() *string
	SetPostCode(v string) *UpdateAccountProfileInfoRequest
	GetPostCode() *string
	SetProvinceJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest
	GetProvinceJsonString() map[string]interface{}
	SetSelfServicingBusinessRegNum(v string) *UpdateAccountProfileInfoRequest
	GetSelfServicingBusinessRegNum() *string
	SetSelfServicingIdentificationNum(v string) *UpdateAccountProfileInfoRequest
	GetSelfServicingIdentificationNum() *string
	SetTrueName(v string) *UpdateAccountProfileInfoRequest
	GetTrueName() *string
}

type UpdateAccountProfileInfoRequest struct {
	AccountAttribute               *string                `json:"AccountAttribute,omitempty" xml:"AccountAttribute,omitempty"`
	Address                        *string                `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2                       *string                `json:"Address2,omitempty" xml:"Address2,omitempty"`
	BindAlipayNo                   *string                `json:"BindAlipayNo,omitempty" xml:"BindAlipayNo,omitempty"`
	CertType                       *string                `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CityJsonString                 map[string]interface{} `json:"CityJsonString,omitempty" xml:"CityJsonString,omitempty"`
	ContactMethod                  *string                `json:"ContactMethod,omitempty" xml:"ContactMethod,omitempty"`
	DistrictJsonString             map[string]interface{} `json:"DistrictJsonString,omitempty" xml:"DistrictJsonString,omitempty"`
	Fax                            *string                `json:"Fax,omitempty" xml:"Fax,omitempty"`
	FirstName                      *string                `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	Head                           *string                `json:"Head,omitempty" xml:"Head,omitempty"`
	HeadColor                      *string                `json:"HeadColor,omitempty" xml:"HeadColor,omitempty"`
	LastName                       *string                `json:"LastName,omitempty" xml:"LastName,omitempty"`
	PK                             *string                `json:"PK,omitempty" xml:"PK,omitempty"`
	Phone                          *string                `json:"Phone,omitempty" xml:"Phone,omitempty"`
	PostCode                       *string                `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	ProvinceJsonString             map[string]interface{} `json:"ProvinceJsonString,omitempty" xml:"ProvinceJsonString,omitempty"`
	SelfServicingBusinessRegNum    *string                `json:"SelfServicingBusinessRegNum,omitempty" xml:"SelfServicingBusinessRegNum,omitempty"`
	SelfServicingIdentificationNum *string                `json:"SelfServicingIdentificationNum,omitempty" xml:"SelfServicingIdentificationNum,omitempty"`
	TrueName                       *string                `json:"TrueName,omitempty" xml:"TrueName,omitempty"`
}

func (s UpdateAccountProfileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountProfileInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountProfileInfoRequest) GetAccountAttribute() *string {
	return s.AccountAttribute
}

func (s *UpdateAccountProfileInfoRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateAccountProfileInfoRequest) GetAddress2() *string {
	return s.Address2
}

func (s *UpdateAccountProfileInfoRequest) GetBindAlipayNo() *string {
	return s.BindAlipayNo
}

func (s *UpdateAccountProfileInfoRequest) GetCertType() *string {
	return s.CertType
}

func (s *UpdateAccountProfileInfoRequest) GetCityJsonString() map[string]interface{} {
	return s.CityJsonString
}

func (s *UpdateAccountProfileInfoRequest) GetContactMethod() *string {
	return s.ContactMethod
}

func (s *UpdateAccountProfileInfoRequest) GetDistrictJsonString() map[string]interface{} {
	return s.DistrictJsonString
}

func (s *UpdateAccountProfileInfoRequest) GetFax() *string {
	return s.Fax
}

func (s *UpdateAccountProfileInfoRequest) GetFirstName() *string {
	return s.FirstName
}

func (s *UpdateAccountProfileInfoRequest) GetHead() *string {
	return s.Head
}

func (s *UpdateAccountProfileInfoRequest) GetHeadColor() *string {
	return s.HeadColor
}

func (s *UpdateAccountProfileInfoRequest) GetLastName() *string {
	return s.LastName
}

func (s *UpdateAccountProfileInfoRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateAccountProfileInfoRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateAccountProfileInfoRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateAccountProfileInfoRequest) GetProvinceJsonString() map[string]interface{} {
	return s.ProvinceJsonString
}

func (s *UpdateAccountProfileInfoRequest) GetSelfServicingBusinessRegNum() *string {
	return s.SelfServicingBusinessRegNum
}

func (s *UpdateAccountProfileInfoRequest) GetSelfServicingIdentificationNum() *string {
	return s.SelfServicingIdentificationNum
}

func (s *UpdateAccountProfileInfoRequest) GetTrueName() *string {
	return s.TrueName
}

func (s *UpdateAccountProfileInfoRequest) SetAccountAttribute(v string) *UpdateAccountProfileInfoRequest {
	s.AccountAttribute = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetAddress(v string) *UpdateAccountProfileInfoRequest {
	s.Address = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetAddress2(v string) *UpdateAccountProfileInfoRequest {
	s.Address2 = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetBindAlipayNo(v string) *UpdateAccountProfileInfoRequest {
	s.BindAlipayNo = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetCertType(v string) *UpdateAccountProfileInfoRequest {
	s.CertType = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetCityJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest {
	s.CityJsonString = v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetContactMethod(v string) *UpdateAccountProfileInfoRequest {
	s.ContactMethod = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetDistrictJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest {
	s.DistrictJsonString = v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetFax(v string) *UpdateAccountProfileInfoRequest {
	s.Fax = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetFirstName(v string) *UpdateAccountProfileInfoRequest {
	s.FirstName = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetHead(v string) *UpdateAccountProfileInfoRequest {
	s.Head = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetHeadColor(v string) *UpdateAccountProfileInfoRequest {
	s.HeadColor = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetLastName(v string) *UpdateAccountProfileInfoRequest {
	s.LastName = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetPK(v string) *UpdateAccountProfileInfoRequest {
	s.PK = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetPhone(v string) *UpdateAccountProfileInfoRequest {
	s.Phone = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetPostCode(v string) *UpdateAccountProfileInfoRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetProvinceJsonString(v map[string]interface{}) *UpdateAccountProfileInfoRequest {
	s.ProvinceJsonString = v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetSelfServicingBusinessRegNum(v string) *UpdateAccountProfileInfoRequest {
	s.SelfServicingBusinessRegNum = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetSelfServicingIdentificationNum(v string) *UpdateAccountProfileInfoRequest {
	s.SelfServicingIdentificationNum = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) SetTrueName(v string) *UpdateAccountProfileInfoRequest {
	s.TrueName = &v
	return s
}

func (s *UpdateAccountProfileInfoRequest) Validate() error {
	return dara.Validate(s)
}
