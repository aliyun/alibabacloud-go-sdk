// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSSLCertResponseBody
	GetCode() *int32
	SetData(v []*ListSSLCertResponseBodyData) *ListSSLCertResponseBody
	GetData() []*ListSSLCertResponseBodyData
	SetHttpStatusCode(v int32) *ListSSLCertResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSSLCertResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSSLCertResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSSLCertResponseBody
	GetSuccess() *bool
}

type ListSSLCertResponseBody struct {
	// The status code returned. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListSSLCertResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CAA9A229-141D-5FBA-AC5C-516C02026A11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSSLCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSSLCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListSSLCertResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSSLCertResponseBody) GetData() []*ListSSLCertResponseBodyData {
	return s.Data
}

func (s *ListSSLCertResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSSLCertResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSSLCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSSLCertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSSLCertResponseBody) SetCode(v int32) *ListSSLCertResponseBody {
	s.Code = &v
	return s
}

func (s *ListSSLCertResponseBody) SetData(v []*ListSSLCertResponseBodyData) *ListSSLCertResponseBody {
	s.Data = v
	return s
}

func (s *ListSSLCertResponseBody) SetHttpStatusCode(v int32) *ListSSLCertResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSSLCertResponseBody) SetMessage(v string) *ListSSLCertResponseBody {
	s.Message = &v
	return s
}

func (s *ListSSLCertResponseBody) SetRequestId(v string) *ListSSLCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSSLCertResponseBody) SetSuccess(v bool) *ListSSLCertResponseBody {
	s.Success = &v
	return s
}

func (s *ListSSLCertResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSSLCertResponseBodyData struct {
	// The time when the certificate expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The algorithm.
	//
	// example:
	//
	// test
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The time when the certificate took effect. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 1234
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// test.com
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The domain name with which the certificate is associated.
	//
	// example:
	//
	// *.test.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the certificate expires. This value is a GMT timestamp.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	GmtAfter *string `json:"GmtAfter,omitempty" xml:"GmtAfter,omitempty"`
	// The time when the certificate took effect. This value is a GMT timestamp.
	//
	// example:
	//
	// 2031-03-30 02:35:12
	GmtBefore *string `json:"GmtBefore,omitempty" xml:"GmtBefore,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// test
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The SSL certificate.
	//
	// example:
	//
	// SSL
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
}

func (s ListSSLCertResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSSLCertResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSSLCertResponseBodyData) GetAfterDate() *string {
	return s.AfterDate
}

func (s *ListSSLCertResponseBodyData) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListSSLCertResponseBodyData) GetBeforeDate() *string {
	return s.BeforeDate
}

func (s *ListSSLCertResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListSSLCertResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *ListSSLCertResponseBodyData) GetCommonName() *string {
	return s.CommonName
}

func (s *ListSSLCertResponseBodyData) GetGmtAfter() *string {
	return s.GmtAfter
}

func (s *ListSSLCertResponseBodyData) GetGmtBefore() *string {
	return s.GmtBefore
}

func (s *ListSSLCertResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *ListSSLCertResponseBodyData) GetSans() *string {
	return s.Sans
}

func (s *ListSSLCertResponseBodyData) SetAfterDate(v string) *ListSSLCertResponseBodyData {
	s.AfterDate = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetAlgorithm(v string) *ListSSLCertResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetBeforeDate(v string) *ListSSLCertResponseBodyData {
	s.BeforeDate = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetCertIdentifier(v string) *ListSSLCertResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetCertName(v string) *ListSSLCertResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetCommonName(v string) *ListSSLCertResponseBodyData {
	s.CommonName = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetGmtAfter(v string) *ListSSLCertResponseBodyData {
	s.GmtAfter = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetGmtBefore(v string) *ListSSLCertResponseBodyData {
	s.GmtBefore = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetIssuer(v string) *ListSSLCertResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *ListSSLCertResponseBodyData) SetSans(v string) *ListSSLCertResponseBodyData {
	s.Sans = &v
	return s
}

func (s *ListSSLCertResponseBodyData) Validate() error {
	return dara.Validate(s)
}
