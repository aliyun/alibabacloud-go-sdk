// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAtiRegistrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCc(v string) *CreateAtiRegistrantRequest
	GetCc() *string
	SetCity(v string) *CreateAtiRegistrantRequest
	GetCity() *string
	SetClientToken(v string) *CreateAtiRegistrantRequest
	GetClientToken() *string
	SetDocumentCode(v string) *CreateAtiRegistrantRequest
	GetDocumentCode() *string
	SetDocumentImage(v string) *CreateAtiRegistrantRequest
	GetDocumentImage() *string
	SetDocumentType(v string) *CreateAtiRegistrantRequest
	GetDocumentType() *string
	SetEmail(v string) *CreateAtiRegistrantRequest
	GetEmail() *string
	SetName(v string) *CreateAtiRegistrantRequest
	GetName() *string
	SetPhone(v string) *CreateAtiRegistrantRequest
	GetPhone() *string
	SetState(v string) *CreateAtiRegistrantRequest
	GetState() *string
	SetStreet(v string) *CreateAtiRegistrantRequest
	GetStreet() *string
}

type CreateAtiRegistrantRequest struct {
	// The country or region of the registrant. Specify a 2-character country or region code (refer to GB/T 2659.1-2022).
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// The city of the registrant. The value cannot exceed 255 characters in length. If the country is China, the value must comply with GB/T 2260-2007.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Ensures the idempotency of the request. Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The document number of the registrant. The number cannot exceed 50 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110123456789122341
	DocumentCode *string `json:"DocumentCode,omitempty" xml:"DocumentCode,omitempty"`
	// The document image of the registrant (base64-encoded). The original file size must be between 50 KB and 3 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// Document image
	DocumentImage *string `json:"DocumentImage,omitempty" xml:"DocumentImage,omitempty"`
	// The document type of the registrant.
	//
	// | Field | Description             |
	//
	// | ---- | ---------------- |
	//
	// | SFZ   | ID card  |
	//
	// | HZ    | Passport  |
	//
	// | ORG   | Organization code certificate  |
	//
	// | YYZZ	| Business license |
	//
	// | BDDM | Military unit code |
	//
	// | CHNSHXYD	| Unified Social Credit Code certificate of the People\\"s Republic of China |
	//
	// | GAJMTX | Mainland Travel Permit for Hong Kong and Macao Residents |
	//
	// | GAJZZ	| Residence Permit for Hong Kong and Macao Residents |
	//
	// | GATLYCZD	| Registration certificate for permanent representative offices of tourism departments in Hong Kong, Macao, and Taiwan |
	//
	// | GAXWZNDJ	| Registration certificate for permanent mainland journalist stations of Hong Kong and Macao news agencies |
	//
	// | GZJGZY	| Notary office practice certificate |
	//
	// | JDDWFW | Military unit paid service license |
	//
	// | JGZ	| Military officer certificate |
	//
	// | JJHFR	| Foundation legal person registration certificate |
	//
	// | LSZY	| Law firm practice license |
	//
	// | MBFQY	| Private non-enterprise unit registration certificate |
	//
	// | MBXXBX	| Private school operating license |
	//
	// | NCJTJJZZ	| Rural collective economic organization registration certificate |
	//
	// | QTTYDM | Other - Unified Social Credit Code  |
	//
	// | SFJD	| Judicial appraisal license |
	//
	// | SHTTFR | Social organization legal person registration certificate |
	//
	// | SHFWJG	| Social service institution registration certificate |
	//
	// | SYDWFR | Public institution legal person certificate |
	//
	// | TYDM  | Unified Social Credit Code certificate  |
	//
	// | YLJGZY	| Medical institution practice license |
	//
	// | ZCWYHDJZ	| Arbitration commission registration certificate |
	//
	// | ZJCS	| Religious activity venue registration certificate |
	//
	// | BJWSXX	| Operating license for schools for children of foreign embassy staff in Beijing |
	//
	// | JWJG	| Certificate of overseas institution |
	//
	// | JWFZFDBJ	| Registration certificate for representative offices of overseas non-governmental organizations |
	//
	// | WGCZJG | Registration certificate for permanent representative offices of foreign enterprises |
	//
	// | WGZHWH	| Registration certificate for foreign cultural centers in China |
	//
	// | WGZHXWJG	| Certificate for foreign news agencies in China |
	//
	// | WJLSFZ| Foreigner permanent residence ID card |
	//
	// | WLCZJG	| Registration certificate for permanent representative offices of foreign government tourism departments |
	//
	// | QT     | Other |
	//
	// This parameter is required.
	//
	// example:
	//
	// SFZ
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// The email address. The address cannot exceed 300 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the registrant. The name cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhang XX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone number of the registrant. The value cannot exceed 128 characters in length. If the country is China and the number is not a mobile phone number, the area code must match the city.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13112345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The province of the registrant. The value cannot exceed 255 characters in length. If the country is China, the value must comply with GB/T 2260-2007.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The address of the registrant. The value cannot exceed 255 characters in length.
	//
	// example:
	//
	// XX District, XX Street
	Street *string `json:"Street,omitempty" xml:"Street,omitempty"`
}

func (s CreateAtiRegistrantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAtiRegistrantRequest) GoString() string {
	return s.String()
}

func (s *CreateAtiRegistrantRequest) GetCc() *string {
	return s.Cc
}

func (s *CreateAtiRegistrantRequest) GetCity() *string {
	return s.City
}

func (s *CreateAtiRegistrantRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAtiRegistrantRequest) GetDocumentCode() *string {
	return s.DocumentCode
}

func (s *CreateAtiRegistrantRequest) GetDocumentImage() *string {
	return s.DocumentImage
}

func (s *CreateAtiRegistrantRequest) GetDocumentType() *string {
	return s.DocumentType
}

func (s *CreateAtiRegistrantRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateAtiRegistrantRequest) GetName() *string {
	return s.Name
}

func (s *CreateAtiRegistrantRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateAtiRegistrantRequest) GetState() *string {
	return s.State
}

func (s *CreateAtiRegistrantRequest) GetStreet() *string {
	return s.Street
}

func (s *CreateAtiRegistrantRequest) SetCc(v string) *CreateAtiRegistrantRequest {
	s.Cc = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetCity(v string) *CreateAtiRegistrantRequest {
	s.City = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetClientToken(v string) *CreateAtiRegistrantRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetDocumentCode(v string) *CreateAtiRegistrantRequest {
	s.DocumentCode = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetDocumentImage(v string) *CreateAtiRegistrantRequest {
	s.DocumentImage = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetDocumentType(v string) *CreateAtiRegistrantRequest {
	s.DocumentType = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetEmail(v string) *CreateAtiRegistrantRequest {
	s.Email = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetName(v string) *CreateAtiRegistrantRequest {
	s.Name = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetPhone(v string) *CreateAtiRegistrantRequest {
	s.Phone = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetState(v string) *CreateAtiRegistrantRequest {
	s.State = &v
	return s
}

func (s *CreateAtiRegistrantRequest) SetStreet(v string) *CreateAtiRegistrantRequest {
	s.Street = &v
	return s
}

func (s *CreateAtiRegistrantRequest) Validate() error {
	return dara.Validate(s)
}
