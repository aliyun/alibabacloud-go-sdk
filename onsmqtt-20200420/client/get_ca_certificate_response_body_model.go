// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCaCertificateResponseBodyData) *GetCaCertificateResponseBody
	GetData() *GetCaCertificateResponseBodyData
	SetRequestId(v string) *GetCaCertificateResponseBody
	GetRequestId() *string
}

type GetCaCertificateResponseBody struct {
	// Certificate details.
	Data *GetCaCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponseBody) GetData() *GetCaCertificateResponseBodyData {
	return s.Data
}

func (s *GetCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCaCertificateResponseBody) SetData(v *GetCaCertificateResponseBodyData) *GetCaCertificateResponseBody {
	s.Data = v
	return s
}

func (s *GetCaCertificateResponseBody) SetRequestId(v string) *GetCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCaCertificateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCaCertificateResponseBodyData struct {
	// Content of the CA certificate.
	//
	// > \\n represents a new line.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIDuzCCAqdGVzdC5jbi1xaW5n******\\n-----END CERTIFICATE-----
	CaContent *string `json:"CaContent,omitempty" xml:"CaContent,omitempty"`
	// Name of the CA certificate
	//
	// example:
	//
	// mqtt_ca
	CaName *string `json:"CaName,omitempty" xml:"CaName,omitempty"`
	// Registration code of the CA certificate
	//
	// example:
	//
	// 13274673-8f90-4630-bea1-9cccb25756ad2089******
	RegistrationCode *string `json:"RegistrationCode,omitempty" xml:"RegistrationCode,omitempty"`
	// The SN serial number of the CA certificate, used to uniquely identify a CA certificate. Value range: no more than 128 bytes.
	//
	// example:
	//
	// 00f26900ba87******
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The status of the CA certificate. The values are as follows:
	//
	// - **0**: Indicates that the certificate is in an inactive state. - **1**: Indicates that the certificate is in an active state.
	//
	// > After the CA certificate is registered, it is in an active state by default.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The start time when the CA certificate becomes effective. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1654137303000
	ValidBegin *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	// The end time when the CA certificate becomes effective. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1969497303000
	ValidEnd *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
	// Content of the Verification certificate.
	//
	// > \\n represents a new line.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIID/DCCAu+Y5sRMpp9tnd+4s******\\n-----END CERTIFICATE-----
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s GetCaCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponseBodyData) GetCaContent() *string {
	return s.CaContent
}

func (s *GetCaCertificateResponseBodyData) GetCaName() *string {
	return s.CaName
}

func (s *GetCaCertificateResponseBodyData) GetRegistrationCode() *string {
	return s.RegistrationCode
}

func (s *GetCaCertificateResponseBodyData) GetSn() *string {
	return s.Sn
}

func (s *GetCaCertificateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCaCertificateResponseBodyData) GetValidBegin() *string {
	return s.ValidBegin
}

func (s *GetCaCertificateResponseBodyData) GetValidEnd() *string {
	return s.ValidEnd
}

func (s *GetCaCertificateResponseBodyData) GetVerificationContent() *string {
	return s.VerificationContent
}

func (s *GetCaCertificateResponseBodyData) SetCaContent(v string) *GetCaCertificateResponseBodyData {
	s.CaContent = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetCaName(v string) *GetCaCertificateResponseBodyData {
	s.CaName = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetRegistrationCode(v string) *GetCaCertificateResponseBodyData {
	s.RegistrationCode = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetSn(v string) *GetCaCertificateResponseBodyData {
	s.Sn = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetStatus(v string) *GetCaCertificateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetValidBegin(v string) *GetCaCertificateResponseBodyData {
	s.ValidBegin = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetValidEnd(v string) *GetCaCertificateResponseBodyData {
	s.ValidEnd = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetVerificationContent(v string) *GetCaCertificateResponseBodyData {
	s.VerificationContent = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
