// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenerCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody
	GetCertificates() []*ListListenerCertificatesResponseBodyCertificates
	SetMaxResults(v int32) *ListListenerCertificatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenerCertificatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListListenerCertificatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListListenerCertificatesResponseBody
	GetTotalCount() *int32
}

type ListListenerCertificatesResponseBody struct {
	// The certificates.
	Certificates []*ListListenerCertificatesResponseBodyCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListListenerCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListenerCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBody) GetCertificates() []*ListListenerCertificatesResponseBodyCertificates {
	return s.Certificates
}

func (s *ListListenerCertificatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenerCertificatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenerCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListenerCertificatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListenerCertificatesResponseBody) SetCertificates(v []*ListListenerCertificatesResponseBodyCertificates) *ListListenerCertificatesResponseBody {
	s.Certificates = v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetMaxResults(v int32) *ListListenerCertificatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetNextToken(v string) *ListListenerCertificatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetRequestId(v string) *ListListenerCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) SetTotalCount(v int32) *ListListenerCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListenerCertificatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListListenerCertificatesResponseBodyCertificates struct {
	// The certificate ID. Only server certificates are supported.
	//
	// example:
	//
	// 12315790343_166f8204689_1714763408_70998****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// Indicates whether the certificate is the default certificate of the listener. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Indicates whether the certificate is associated with the listener. Valid values:
	//
	// 	- **Associating**
	//
	// 	- **Associated**
	//
	// 	- **Diassociating**
	//
	// example:
	//
	// Associating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListListenerCertificatesResponseBodyCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListListenerCertificatesResponseBodyCertificates) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetStatus() *string {
	return s.Status
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateType(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateType = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetStatus(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Status = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}
