// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainCountResponseBody
	GetRequestId() *string
	SetRootDomainsCount(v int32) *DescribeDomainCountResponseBody
	GetRootDomainsCount() *int32
	SetSubDomainsCount(v int32) *DescribeDomainCountResponseBody
	GetSubDomainsCount() *int32
	SetTotalDomainsCount(v int32) *DescribeDomainCountResponseBody
	GetTotalDomainsCount() *int32
}

type DescribeDomainCountResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C286491D-4A2F-589A-B63B-D2AD3DA9BD71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of root domains.
	//
	// example:
	//
	// 5
	RootDomainsCount *int32 `json:"RootDomainsCount,omitempty" xml:"RootDomainsCount,omitempty"`
	// The number of subdomains.
	//
	// example:
	//
	// 5
	SubDomainsCount *int32 `json:"SubDomainsCount,omitempty" xml:"SubDomainsCount,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalDomainsCount *int32 `json:"TotalDomainsCount,omitempty" xml:"TotalDomainsCount,omitempty"`
}

func (s DescribeDomainCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCountResponseBody) GetRootDomainsCount() *int32 {
	return s.RootDomainsCount
}

func (s *DescribeDomainCountResponseBody) GetSubDomainsCount() *int32 {
	return s.SubDomainsCount
}

func (s *DescribeDomainCountResponseBody) GetTotalDomainsCount() *int32 {
	return s.TotalDomainsCount
}

func (s *DescribeDomainCountResponseBody) SetRequestId(v string) *DescribeDomainCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCountResponseBody) SetRootDomainsCount(v int32) *DescribeDomainCountResponseBody {
	s.RootDomainsCount = &v
	return s
}

func (s *DescribeDomainCountResponseBody) SetSubDomainsCount(v int32) *DescribeDomainCountResponseBody {
	s.SubDomainsCount = &v
	return s
}

func (s *DescribeDomainCountResponseBody) SetTotalDomainsCount(v int32) *DescribeDomainCountResponseBody {
	s.TotalDomainsCount = &v
	return s
}

func (s *DescribeDomainCountResponseBody) Validate() error {
	return dara.Validate(s)
}
