// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainViewTopUrlResponseBody
	GetRequestId() *string
	SetUrlList(v []*DescribeDomainViewTopUrlResponseBodyUrlList) *DescribeDomainViewTopUrlResponseBody
	GetUrlList() []*DescribeDomainViewTopUrlResponseBodyUrlList
}

type DescribeDomainViewTopUrlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the URLs that receive the most requests.
	UrlList []*DescribeDomainViewTopUrlResponseBodyUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewTopUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainViewTopUrlResponseBody) GetUrlList() []*DescribeDomainViewTopUrlResponseBodyUrlList {
	return s.UrlList
}

func (s *DescribeDomainViewTopUrlResponseBody) SetRequestId(v string) *DescribeDomainViewTopUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBody) SetUrlList(v []*DescribeDomainViewTopUrlResponseBodyUrlList) *DescribeDomainViewTopUrlResponseBody {
	s.UrlList = v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainViewTopUrlResponseBodyUrlList struct {
	// The total number of requests.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The URL that is Base64-encoded.
	//
	// example:
	//
	// Lw==
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDomainViewTopUrlResponseBodyUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopUrlResponseBodyUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) GetUrl() *string {
	return s.Url
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetCount(v int64) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Count = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetDomain(v string) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) SetUrl(v string) *DescribeDomainViewTopUrlResponseBodyUrlList {
	s.Url = &v
	return s
}

func (s *DescribeDomainViewTopUrlResponseBodyUrlList) Validate() error {
	return dara.Validate(s)
}
