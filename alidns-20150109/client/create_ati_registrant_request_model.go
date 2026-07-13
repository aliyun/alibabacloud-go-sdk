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
	// The country or region of the registrant. Specify a 2-character country or region code in compliance with GB/T 2659.1-2022.
	//
	// This parameter is required.
	//
	// example:
	//
	// 中国
	Cc *string `json:"Cc,omitempty" xml:"Cc,omitempty"`
	// The city of the registrant. The value cannot exceed 255 characters in length. If the country is China, the value must comply with GB/T 2260-2007.
	//
	// This parameter is required.
	//
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests. The ClientToken value supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The document number of the registrant. The value cannot exceed 50 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110123456789122341
	DocumentCode *string `json:"DocumentCode,omitempty" xml:"DocumentCode,omitempty"`
	// The document image of the registrant in Base64 encoding. The original file size must be between 50 KB and 3 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 证件图片
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
	// | BJWSXX	| Beijing operating license for schools for children of foreign embassy staff |
	//
	// | JWJG	| Overseas institution certificate |
	//
	// | JWFZFDBJ	| Overseas non-governmental organization representative office registration certificate |
	//
	// | WGCZJG | Foreign enterprise permanent representative office registration certificate |
	//
	// | WGZHWH	| Foreign cultural center registration certificate in China |
	//
	// | WGZHXWJG	| Foreign news agency certificate in China |
	//
	// | WJLSFZ| Foreigner permanent residence ID card |
	//
	// | WLCZJG	| Approval registration certificate for permanent representative offices of foreign government tourism departments |
	//
	// | QT     | Other |
	//
	// This parameter is required.
	//
	// example:
	//
	// SFZ
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// The email address. The value cannot exceed 300 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the registrant. The value cannot exceed 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张xx
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
	// 浙江省
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The address of the registrant. The value cannot exceed 255 characters in length.
	//
	// example:
	//
	// xx区xx街道
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
