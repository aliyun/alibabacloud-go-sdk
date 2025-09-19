// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainListResponseList(v []*DescribeDomainListResponseBodyDomainListResponseList) *DescribeDomainListResponseBody
	GetDomainListResponseList() []*DescribeDomainListResponseBodyDomainListResponseList
	SetPageInfo(v *DescribeDomainListResponseBodyPageInfo) *DescribeDomainListResponseBody
	GetPageInfo() *DescribeDomainListResponseBodyPageInfo
	SetRequestId(v string) *DescribeDomainListResponseBody
	GetRequestId() *string
}

type DescribeDomainListResponseBody struct {
	// An array that consists of the information about domain names.
	DomainListResponseList []*DescribeDomainListResponseBodyDomainListResponseList `json:"DomainListResponseList,omitempty" xml:"DomainListResponseList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeDomainListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBody) GetDomainListResponseList() []*DescribeDomainListResponseBodyDomainListResponseList {
	return s.DomainListResponseList
}

func (s *DescribeDomainListResponseBody) GetPageInfo() *DescribeDomainListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainListResponseBody) SetDomainListResponseList(v []*DescribeDomainListResponseBodyDomainListResponseList) *DescribeDomainListResponseBody {
	s.DomainListResponseList = v
	return s
}

func (s *DescribeDomainListResponseBody) SetPageInfo(v *DescribeDomainListResponseBodyPageInfo) *DescribeDomainListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeDomainListResponseBody) SetRequestId(v string) *DescribeDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainListResponseBodyDomainListResponseList struct {
	// The name of the domain or website.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP addresses of the domain name.
	//
	// example:
	//
	// 1.2.XX.XX,1.0.XX.XX
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s DescribeDomainListResponseBodyDomainListResponseList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainListResponseBodyDomainListResponseList) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) GetIpList() *string {
	return s.IpList
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) SetDomain(v string) *DescribeDomainListResponseBodyDomainListResponseList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) SetIpList(v string) *DescribeDomainListResponseBodyDomainListResponseList {
	s.IpList = &v
	return s
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 27
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDomainListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDomainListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDomainListResponseBodyPageInfo) SetCount(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetPageSize(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
