// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSslCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSslCertsResponseBody
	GetCode() *string
	SetItems(v []*ListSslCertsResponseBodyItems) *ListSslCertsResponseBody
	GetItems() []*ListSslCertsResponseBodyItems
	SetMaxResults(v int32) *ListSslCertsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSslCertsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSslCertsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSslCertsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSslCertsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListSslCertsResponseBody
	GetTotalCount() *int64
}

type ListSslCertsResponseBody struct {
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Items      []*ListSslCertsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int32                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// req-123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 11
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSslCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSslCertsResponseBody) GetItems() []*ListSslCertsResponseBodyItems {
	return s.Items
}

func (s *ListSslCertsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSslCertsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSslCertsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSslCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSslCertsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSslCertsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSslCertsResponseBody) SetCode(v string) *ListSslCertsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSslCertsResponseBody) SetItems(v []*ListSslCertsResponseBodyItems) *ListSslCertsResponseBody {
	s.Items = v
	return s
}

func (s *ListSslCertsResponseBody) SetMaxResults(v int32) *ListSslCertsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSslCertsResponseBody) SetMessage(v string) *ListSslCertsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSslCertsResponseBody) SetNextToken(v string) *ListSslCertsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSslCertsResponseBody) SetRequestId(v string) *ListSslCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSslCertsResponseBody) SetSuccess(v bool) *ListSslCertsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSslCertsResponseBody) SetTotalCount(v int64) *ListSslCertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSslCertsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSslCertsResponseBodyItems struct {
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 123
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// 22584627-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// example-cert
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// true
	ChainCompleted *bool `json:"ChainCompleted,omitempty" xml:"ChainCompleted,omitempty"`
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// 1893456000000
	NotAfterTimestamp *int64 `json:"NotAfterTimestamp,omitempty" xml:"NotAfterTimestamp,omitempty"`
	// example:
	//
	// 1704067200000
	NotBeforeTimestamp *int64 `json:"NotBeforeTimestamp,omitempty" xml:"NotBeforeTimestamp,omitempty"`
}

func (s ListSslCertsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSslCertsResponseBodyItems) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListSslCertsResponseBodyItems) GetCertId() *int64 {
	return s.CertId
}

func (s *ListSslCertsResponseBodyItems) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListSslCertsResponseBodyItems) GetCertName() *string {
	return s.CertName
}

func (s *ListSslCertsResponseBodyItems) GetChainCompleted() *bool {
	return s.ChainCompleted
}

func (s *ListSslCertsResponseBodyItems) GetCommonName() *string {
	return s.CommonName
}

func (s *ListSslCertsResponseBodyItems) GetDomain() *string {
	return s.Domain
}

func (s *ListSslCertsResponseBodyItems) GetIssuer() *string {
	return s.Issuer
}

func (s *ListSslCertsResponseBodyItems) GetNotAfterTimestamp() *int64 {
	return s.NotAfterTimestamp
}

func (s *ListSslCertsResponseBodyItems) GetNotBeforeTimestamp() *int64 {
	return s.NotBeforeTimestamp
}

func (s *ListSslCertsResponseBodyItems) SetAlgorithm(v string) *ListSslCertsResponseBodyItems {
	s.Algorithm = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetCertId(v int64) *ListSslCertsResponseBodyItems {
	s.CertId = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetCertIdentifier(v string) *ListSslCertsResponseBodyItems {
	s.CertIdentifier = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetCertName(v string) *ListSslCertsResponseBodyItems {
	s.CertName = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetChainCompleted(v bool) *ListSslCertsResponseBodyItems {
	s.ChainCompleted = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetCommonName(v string) *ListSslCertsResponseBodyItems {
	s.CommonName = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetDomain(v string) *ListSslCertsResponseBodyItems {
	s.Domain = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetIssuer(v string) *ListSslCertsResponseBodyItems {
	s.Issuer = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetNotAfterTimestamp(v int64) *ListSslCertsResponseBodyItems {
	s.NotAfterTimestamp = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) SetNotBeforeTimestamp(v int64) *ListSslCertsResponseBodyItems {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *ListSslCertsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
