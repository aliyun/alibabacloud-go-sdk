// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrusteeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v int64) *ListTrusteeOrderRequest
	GetCertificateId() *int64
	SetMaxResults(v int32) *ListTrusteeOrderRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrusteeOrderRequest
	GetNextToken() *string
	SetOrderId(v int64) *ListTrusteeOrderRequest
	GetOrderId() *int64
}

type ListTrusteeOrderRequest struct {
	// The certificate ID. You must specify either CertificateId or OrderId. Both cannot be empty at the same time.
	//
	// example:
	//
	// 23787679
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. If NextToken is empty, no more results are available.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The order ID. You must specify either CertificateId or OrderId. Both cannot be empty at the same time.
	//
	// example:
	//
	// 14933279
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ListTrusteeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrusteeOrderRequest) GoString() string {
	return s.String()
}

func (s *ListTrusteeOrderRequest) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *ListTrusteeOrderRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrusteeOrderRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrusteeOrderRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListTrusteeOrderRequest) SetCertificateId(v int64) *ListTrusteeOrderRequest {
	s.CertificateId = &v
	return s
}

func (s *ListTrusteeOrderRequest) SetMaxResults(v int32) *ListTrusteeOrderRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrusteeOrderRequest) SetNextToken(v string) *ListTrusteeOrderRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrusteeOrderRequest) SetOrderId(v int64) *ListTrusteeOrderRequest {
	s.OrderId = &v
	return s
}

func (s *ListTrusteeOrderRequest) Validate() error {
	return dara.Validate(s)
}
