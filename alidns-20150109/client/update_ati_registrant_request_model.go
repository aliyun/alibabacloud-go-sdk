// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiRegistrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCc(v string) *UpdateAtiRegistrantRequest
	GetCc() *string
	SetCity(v string) *UpdateAtiRegistrantRequest
	GetCity() *string
	SetClientToken(v string) *UpdateAtiRegistrantRequest
	GetClientToken() *string
	SetDocumentCode(v string) *UpdateAtiRegistrantRequest
	GetDocumentCode() *string
	SetDocumentImage(v string) *UpdateAtiRegistrantRequest
	GetDocumentImage() *string
	SetDocumentType(v string) *UpdateAtiRegistrantRequest
	GetDocumentType() *string
	SetEmail(v string) *UpdateAtiRegistrantRequest
	GetEmail() *string
	SetName(v string) *UpdateAtiRegistrantRequest
	GetName() *string
	SetPhone(v string) *UpdateAtiRegistrantRequest
	GetPhone() *string
	SetRegistrantId(v string) *UpdateAtiRegistrantRequest
	GetRegistrantId() *string
	SetState(v string) *UpdateAtiRegistrantRequest
	GetState() *string
	SetStreet(v string) *UpdateAtiRegistrantRequest
	GetStreet() *string
}

type UpdateAtiRegistrantRequest struct {
	// The country.
	//
	// example:
	//
	// 中国
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// The city.
	//
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The document number of the registrant. Maximum length: 50 characters.
	//
	// example:
	//
	// 11011239900307663x
	DocumentCode *string `json:"DocumentCode,omitempty" xml:"DocumentCode,omitempty"`
	// The document image of the registrant (Base64-encoded). The original file size must be between 50 KB and 3 MB.
	//
	// example:
	//
	// 证件图片
	DocumentImage *string `json:"DocumentImage,omitempty" xml:"DocumentImage,omitempty"`
	// The document type of the registrant. For more information, see the appendix on document types.
	//
	// example:
	//
	// SFZ
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// The email address. Maximum length: 300 characters.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the registrant. Maximum length: 255 characters.
	//
	// example:
	//
	// 张xx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone number of the registrant. Maximum length: 128 characters. If the country is China, the area code of a non-mobile phone number must match the city.
	//
	// example:
	//
	// 13112345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The ID of the registrant profile.
	//
	// example:
	//
	// 2072277378616354816
	RegistrantId *string `json:"RegistrantId,omitempty" xml:"RegistrantId,omitempty"`
	// The state or province.
	//
	// example:
	//
	// 浙江省
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The street.
	//
	// example:
	//
	// xx区xx街道
	Street *string `json:"Street,omitempty" xml:"Street,omitempty"`
}

func (s UpdateAtiRegistrantRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiRegistrantRequest) GoString() string {
	return s.String()
}

func (s *UpdateAtiRegistrantRequest) GetCc() *string {
	return s.Cc
}

func (s *UpdateAtiRegistrantRequest) GetCity() *string {
	return s.City
}

func (s *UpdateAtiRegistrantRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAtiRegistrantRequest) GetDocumentCode() *string {
	return s.DocumentCode
}

func (s *UpdateAtiRegistrantRequest) GetDocumentImage() *string {
	return s.DocumentImage
}

func (s *UpdateAtiRegistrantRequest) GetDocumentType() *string {
	return s.DocumentType
}

func (s *UpdateAtiRegistrantRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateAtiRegistrantRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAtiRegistrantRequest) GetPhone() *string {
	return s.Phone
}

func (s *UpdateAtiRegistrantRequest) GetRegistrantId() *string {
	return s.RegistrantId
}

func (s *UpdateAtiRegistrantRequest) GetState() *string {
	return s.State
}

func (s *UpdateAtiRegistrantRequest) GetStreet() *string {
	return s.Street
}

func (s *UpdateAtiRegistrantRequest) SetCc(v string) *UpdateAtiRegistrantRequest {
	s.Cc = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetCity(v string) *UpdateAtiRegistrantRequest {
	s.City = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetClientToken(v string) *UpdateAtiRegistrantRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetDocumentCode(v string) *UpdateAtiRegistrantRequest {
	s.DocumentCode = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetDocumentImage(v string) *UpdateAtiRegistrantRequest {
	s.DocumentImage = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetDocumentType(v string) *UpdateAtiRegistrantRequest {
	s.DocumentType = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetEmail(v string) *UpdateAtiRegistrantRequest {
	s.Email = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetName(v string) *UpdateAtiRegistrantRequest {
	s.Name = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetPhone(v string) *UpdateAtiRegistrantRequest {
	s.Phone = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetRegistrantId(v string) *UpdateAtiRegistrantRequest {
	s.RegistrantId = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetState(v string) *UpdateAtiRegistrantRequest {
	s.State = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) SetStreet(v string) *UpdateAtiRegistrantRequest {
	s.Street = &v
	return s
}

func (s *UpdateAtiRegistrantRequest) Validate() error {
	return dara.Validate(s)
}
