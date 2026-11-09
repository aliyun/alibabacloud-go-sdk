// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCaCertificateResponseBodyData) *ListCaCertificateResponseBody
	GetData() *ListCaCertificateResponseBodyData
	SetRequestId(v string) *ListCaCertificateResponseBody
	GetRequestId() *string
}

type ListCaCertificateResponseBody struct {
	// Query result.
	Data *ListCaCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBody) GetData() *ListCaCertificateResponseBodyData {
	return s.Data
}

func (s *ListCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCaCertificateResponseBody) SetData(v *ListCaCertificateResponseBodyData) *ListCaCertificateResponseBody {
	s.Data = v
	return s
}

func (s *ListCaCertificateResponseBody) SetRequestId(v string) *ListCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCaCertificateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCaCertificateResponseBodyData struct {
	// Details of the CA certificate
	CaCertificateVOS []*ListCaCertificateResponseBodyDataCaCertificateVOS `json:"CaCertificateVOS,omitempty" xml:"CaCertificateVOS,omitempty" type:"Repeated"`
	// The current page number of the returned query records.
	//
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of results to display per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Maximum number of pages in the query result.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCaCertificateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCaCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBodyData) GetCaCertificateVOS() []*ListCaCertificateResponseBodyDataCaCertificateVOS {
	return s.CaCertificateVOS
}

func (s *ListCaCertificateResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListCaCertificateResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCaCertificateResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListCaCertificateResponseBodyData) SetCaCertificateVOS(v []*ListCaCertificateResponseBodyDataCaCertificateVOS) *ListCaCertificateResponseBodyData {
	s.CaCertificateVOS = v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetPageNo(v int32) *ListCaCertificateResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetPageSize(v int32) *ListCaCertificateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetTotal(v int32) *ListCaCertificateResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListCaCertificateResponseBodyData) Validate() error {
	if s.CaCertificateVOS != nil {
		for _, item := range s.CaCertificateVOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCaCertificateResponseBodyDataCaCertificateVOS struct {
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
	// SN serial number of the CA certificate
	//
	// example:
	//
	// 007269004887******
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
	// Verify the content of the certificate.
	//
	// > \\n represents a new line.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIID/DCCAu+Y5sRMpp9tnd+4s******\\n-----END CERTIFICATE-----
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s ListCaCertificateResponseBodyDataCaCertificateVOS) String() string {
	return dara.Prettify(s)
}

func (s ListCaCertificateResponseBodyDataCaCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetCaContent() *string {
	return s.CaContent
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetCaName() *string {
	return s.CaName
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetRegistrationCode() *string {
	return s.RegistrationCode
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetSn() *string {
	return s.Sn
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetStatus() *string {
	return s.Status
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetValidBegin() *string {
	return s.ValidBegin
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetValidEnd() *string {
	return s.ValidEnd
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) GetVerificationContent() *string {
	return s.VerificationContent
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetCaContent(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.CaContent = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetCaName(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.CaName = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetRegistrationCode(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.RegistrationCode = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetSn(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.Sn = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetStatus(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetValidBegin(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetValidEnd(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.ValidEnd = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetVerificationContent(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.VerificationContent = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) Validate() error {
	return dara.Validate(s)
}
