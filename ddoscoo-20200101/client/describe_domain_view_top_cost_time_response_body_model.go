// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopCostTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainViewTopCostTimeResponseBody
	GetRequestId() *string
	SetUrlList(v []*DescribeDomainViewTopCostTimeResponseBodyUrlList) *DescribeDomainViewTopCostTimeResponseBody
	GetUrlList() []*DescribeDomainViewTopCostTimeResponseBodyUrlList
}

type DescribeDomainViewTopCostTimeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The URLs which require the longest time to respond to requests.
	UrlList []*DescribeDomainViewTopCostTimeResponseBodyUrlList `json:"UrlList,omitempty" xml:"UrlList,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewTopCostTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainViewTopCostTimeResponseBody) GetUrlList() []*DescribeDomainViewTopCostTimeResponseBodyUrlList {
	return s.UrlList
}

func (s *DescribeDomainViewTopCostTimeResponseBody) SetRequestId(v string) *DescribeDomainViewTopCostTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBody) SetUrlList(v []*DescribeDomainViewTopCostTimeResponseBodyUrlList) *DescribeDomainViewTopCostTimeResponseBody {
	s.UrlList = v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainViewTopCostTimeResponseBodyUrlList struct {
	// The response duration. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	CostTime *float32 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
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

func (s DescribeDomainViewTopCostTimeResponseBodyUrlList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeResponseBodyUrlList) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) GetCostTime() *float32 {
	return s.CostTime
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) GetUrl() *string {
	return s.Url
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetCostTime(v float32) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.CostTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetDomain(v string) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) SetUrl(v string) *DescribeDomainViewTopCostTimeResponseBodyUrlList {
	s.Url = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeResponseBodyUrlList) Validate() error {
	return dara.Validate(s)
}
