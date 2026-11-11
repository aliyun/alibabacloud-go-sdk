// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenerCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertType(v string) *ListListenerCertificatesRequest
	GetCertType() *string
	SetCertificateIds(v []*string) *ListListenerCertificatesRequest
	GetCertificateIds() []*string
	SetListenerId(v string) *ListListenerCertificatesRequest
	GetListenerId() *string
	SetMaxResults(v int32) *ListListenerCertificatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenerCertificatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListListenerCertificatesRequest
	GetRegionId() *string
}

type ListListenerCertificatesRequest struct {
	// The certificate type. Valid values:
	//
	// - **Ca**: CA certificate.
	//
	// - **Server**: Server certificate.
	//
	// example:
	//
	// Ca
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The server certificates. You can add only one server certificate.
	//
	// > This parameter takes effect only for TCPSSL listeners.
	//
	// if can be null:
	// true
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The listener ID. This parameter is valid only for TCPSSL listeners.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-j49ht1jxxqyg45****@80
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries to return per page. Valid values: **1*	- to **50**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// - Leave this parameter empty for the first query or when no subsequent query is required.
	//
	// - For subsequent queries, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the Network Load Balancer (NLB) instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) GetCertType() *string {
	return s.CertType
}

func (s *ListListenerCertificatesRequest) GetCertificateIds() []*string {
	return s.CertificateIds
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

func (s *ListListenerCertificatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListListenerCertificatesRequest) SetCertType(v string) *ListListenerCertificatesRequest {
	s.CertType = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetCertificateIds(v []*string) *ListListenerCertificatesRequest {
	s.CertificateIds = v
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

func (s *ListListenerCertificatesRequest) SetRegionId(v string) *ListListenerCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenerCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
