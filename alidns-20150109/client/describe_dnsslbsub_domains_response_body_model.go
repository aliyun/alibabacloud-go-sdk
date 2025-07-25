// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDNSSLBSubDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeDNSSLBSubDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDNSSLBSubDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDNSSLBSubDomainsResponseBody
	GetRequestId() *string
	SetSlbSubDomains(v *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) *DescribeDNSSLBSubDomainsResponseBody
	GetSlbSubDomains() *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains
	SetTotalCount(v int64) *DescribeDNSSLBSubDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDNSSLBSubDomainsResponseBody struct {
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The subdomains for which weighted round-robin is enabled.
	SlbSubDomains *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains `json:"SlbSubDomains,omitempty" xml:"SlbSubDomains,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDNSSLBSubDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDNSSLBSubDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDNSSLBSubDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDNSSLBSubDomainsResponseBody) GetSlbSubDomains() *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	return s.SlbSubDomains
}

func (s *DescribeDNSSLBSubDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetPageNumber(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetPageSize(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetRequestId(v string) *DescribeDNSSLBSubDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetSlbSubDomains(v *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) *DescribeDNSSLBSubDomainsResponseBody {
	s.SlbSubDomains = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) SetTotalCount(v int64) *DescribeDNSSLBSubDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDNSSLBSubDomainsResponseBodySlbSubDomains struct {
	SlbSubDomain []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain `json:"SlbSubDomain,omitempty" xml:"SlbSubDomain,omitempty" type:"Repeated"`
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) GetSlbSubDomain() []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	return s.SlbSubDomain
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) SetSlbSubDomain(v []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains {
	s.SlbSubDomain = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain struct {
	// The lines for which weighted round-robin is enabled.
	LineAlgorithms *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms `json:"LineAlgorithms,omitempty" xml:"LineAlgorithms,omitempty" type:"Struct"`
	// Indicates whether weighted round-robin is enabled for the subdomain.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The number of DNS records added for the subdomain.
	//
	// example:
	//
	// 3
	RecordCount *int64 `json:"RecordCount,omitempty" xml:"RecordCount,omitempty"`
	// The name of the subdomain.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The type of the Domain Name System (DNS) record that supports weighted round-robin. Valid values: A, AAAA, and CNAME.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GetLineAlgorithms() *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms {
	return s.LineAlgorithms
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GetOpen() *bool {
	return s.Open
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) GetType() *string {
	return s.Type
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) SetLineAlgorithms(v *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	s.LineAlgorithms = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) SetOpen(v bool) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	s.Open = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) SetRecordCount(v int64) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	s.RecordCount = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) SetSubDomain(v string) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	s.SubDomain = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) SetType(v string) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain {
	s.Type = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomain) Validate() error {
	return dara.Validate(s)
}

type DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms struct {
	LineAlgorithm []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm `json:"LineAlgorithm,omitempty" xml:"LineAlgorithm,omitempty" type:"Repeated"`
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) GetLineAlgorithm() []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm {
	return s.LineAlgorithm
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) SetLineAlgorithm(v []*DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms {
	s.LineAlgorithm = v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithms) Validate() error {
	return dara.Validate(s)
}

type DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm struct {
	// The DNS resolution line. The line can be China Telecom, China Mobile, and China Unicom.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// Indicates whether weighted round-robin is enabled for the line. Valid values:
	//
	// 	- **true*	- (default): Weighted round-robin is enabled.
	//
	// 	- **false**: Weighted round-robin is disabled.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) String() string {
	return dara.Prettify(s)
}

func (s DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) GoString() string {
	return s.String()
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) GetLine() *string {
	return s.Line
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) GetOpen() *bool {
	return s.Open
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) SetLine(v string) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm {
	s.Line = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) SetOpen(v bool) *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm {
	s.Open = &v
	return s
}

func (s *DescribeDNSSLBSubDomainsResponseBodySlbSubDomainsSlbSubDomainLineAlgorithmsLineAlgorithm) Validate() error {
	return dara.Validate(s)
}
