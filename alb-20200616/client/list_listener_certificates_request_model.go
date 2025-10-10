// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenerCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateIds(v []*string) *ListListenerCertificatesRequest
	GetCertificateIds() []*string
	SetCertificateType(v string) *ListListenerCertificatesRequest
	GetCertificateType() *string
	SetListenerId(v string) *ListListenerCertificatesRequest
	GetListenerId() *string
	SetMaxResults(v int32) *ListListenerCertificatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenerCertificatesRequest
	GetNextToken() *string
}

type ListListenerCertificatesRequest struct {
	// The certificates.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The type of the certificate. Valid values: **Ca*	- and **Server**.
	//
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The listener ID. You must specify the ID of an HTTPS listener or a QUIC listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries to return in each call. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *ListListenerCertificatesRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListListenerCertificatesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListListenerCertificatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenerCertificatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenerCertificatesRequest) SetCertificateIds(v []*string) *ListListenerCertificatesRequest {
	s.CertificateIds = v
	return s
}

func (s *ListListenerCertificatesRequest) SetCertificateType(v string) *ListListenerCertificatesRequest {
	s.CertificateType = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
