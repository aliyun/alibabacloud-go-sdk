// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyIdentityRegistrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v int32) *ApplyIdentityRegistrationRequest
	GetAccountType() *int32
	SetCustomerId(v string) *ApplyIdentityRegistrationRequest
	GetCustomerId() *string
	SetDocBackPic(v string) *ApplyIdentityRegistrationRequest
	GetDocBackPic() *string
	SetDocFrontPic(v string) *ApplyIdentityRegistrationRequest
	GetDocFrontPic() *string
	SetDocNum(v string) *ApplyIdentityRegistrationRequest
	GetDocNum() *string
	SetDocType(v string) *ApplyIdentityRegistrationRequest
	GetDocType() *string
	SetEmail(v string) *ApplyIdentityRegistrationRequest
	GetEmail() *string
	SetFullName(v string) *ApplyIdentityRegistrationRequest
	GetFullName() *string
	SetRegisteredAddress(v string) *ApplyIdentityRegistrationRequest
	GetRegisteredAddress() *string
	SetRegisteredCountry(v string) *ApplyIdentityRegistrationRequest
	GetRegisteredCountry() *string
	SetRegisteredNum(v string) *ApplyIdentityRegistrationRequest
	GetRegisteredNum() *string
	SetSource(v string) *ApplyIdentityRegistrationRequest
	GetSource() *string
	SetTel(v string) *ApplyIdentityRegistrationRequest
	GetTel() *string
}

type ApplyIdentityRegistrationRequest struct {
	// This parameter is required.
	AccountType *int32  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	CustomerId  *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// This parameter is required.
	DocBackPic *string `json:"DocBackPic,omitempty" xml:"DocBackPic,omitempty"`
	// This parameter is required.
	DocFrontPic *string `json:"DocFrontPic,omitempty" xml:"DocFrontPic,omitempty"`
	DocNum      *string `json:"DocNum,omitempty" xml:"DocNum,omitempty"`
	DocType     *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// This parameter is required.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	FullName          *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	RegisteredAddress *string `json:"RegisteredAddress,omitempty" xml:"RegisteredAddress,omitempty"`
	RegisteredCountry *string `json:"RegisteredCountry,omitempty" xml:"RegisteredCountry,omitempty"`
	RegisteredNum     *string `json:"RegisteredNum,omitempty" xml:"RegisteredNum,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Tel               *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s ApplyIdentityRegistrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyIdentityRegistrationRequest) GoString() string {
	return s.String()
}

func (s *ApplyIdentityRegistrationRequest) GetAccountType() *int32 {
	return s.AccountType
}

func (s *ApplyIdentityRegistrationRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ApplyIdentityRegistrationRequest) GetDocBackPic() *string {
	return s.DocBackPic
}

func (s *ApplyIdentityRegistrationRequest) GetDocFrontPic() *string {
	return s.DocFrontPic
}

func (s *ApplyIdentityRegistrationRequest) GetDocNum() *string {
	return s.DocNum
}

func (s *ApplyIdentityRegistrationRequest) GetDocType() *string {
	return s.DocType
}

func (s *ApplyIdentityRegistrationRequest) GetEmail() *string {
	return s.Email
}

func (s *ApplyIdentityRegistrationRequest) GetFullName() *string {
	return s.FullName
}

func (s *ApplyIdentityRegistrationRequest) GetRegisteredAddress() *string {
	return s.RegisteredAddress
}

func (s *ApplyIdentityRegistrationRequest) GetRegisteredCountry() *string {
	return s.RegisteredCountry
}

func (s *ApplyIdentityRegistrationRequest) GetRegisteredNum() *string {
	return s.RegisteredNum
}

func (s *ApplyIdentityRegistrationRequest) GetSource() *string {
	return s.Source
}

func (s *ApplyIdentityRegistrationRequest) GetTel() *string {
	return s.Tel
}

func (s *ApplyIdentityRegistrationRequest) SetAccountType(v int32) *ApplyIdentityRegistrationRequest {
	s.AccountType = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetCustomerId(v string) *ApplyIdentityRegistrationRequest {
	s.CustomerId = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetDocBackPic(v string) *ApplyIdentityRegistrationRequest {
	s.DocBackPic = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetDocFrontPic(v string) *ApplyIdentityRegistrationRequest {
	s.DocFrontPic = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetDocNum(v string) *ApplyIdentityRegistrationRequest {
	s.DocNum = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetDocType(v string) *ApplyIdentityRegistrationRequest {
	s.DocType = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetEmail(v string) *ApplyIdentityRegistrationRequest {
	s.Email = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetFullName(v string) *ApplyIdentityRegistrationRequest {
	s.FullName = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetRegisteredAddress(v string) *ApplyIdentityRegistrationRequest {
	s.RegisteredAddress = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetRegisteredCountry(v string) *ApplyIdentityRegistrationRequest {
	s.RegisteredCountry = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetRegisteredNum(v string) *ApplyIdentityRegistrationRequest {
	s.RegisteredNum = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetSource(v string) *ApplyIdentityRegistrationRequest {
	s.Source = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) SetTel(v string) *ApplyIdentityRegistrationRequest {
	s.Tel = &v
	return s
}

func (s *ApplyIdentityRegistrationRequest) Validate() error {
	return dara.Validate(s)
}
