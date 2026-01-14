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
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
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
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListenerCertificatesResponseBodyCertificates struct {
	// The certificate ID.
	//
	// example:
	//
	// 6092**-cn-hangzhou
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The domain name associated with the additional certificate.
	//
	// This parameter is not returned if the certificate is a default one.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether the certificate is a default one.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The status of the certificate.
	//
	// 	- **active**: The certificate is associated with a listener and in effect.
	//
	// 	- **updating**: The additional certificate is being replaced.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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

func (s *ListListenerCertificatesResponseBodyCertificates) GetDomain() *string {
	return s.Domain
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListListenerCertificatesResponseBodyCertificates) GetState() *string {
	return s.State
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetCertificateId(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetDomain(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.Domain = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetIsDefault(v bool) *ListListenerCertificatesResponseBodyCertificates {
	s.IsDefault = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) SetState(v string) *ListListenerCertificatesResponseBodyCertificates {
	s.State = &v
	return s
}

func (s *ListListenerCertificatesResponseBodyCertificates) Validate() error {
	return dara.Validate(s)
}
