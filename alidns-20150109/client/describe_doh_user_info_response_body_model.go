// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int32) *DescribeDohUserInfoResponseBody
	GetDomainCount() *int32
	SetPdnsId(v int64) *DescribeDohUserInfoResponseBody
	GetPdnsId() *int64
	SetRequestId(v string) *DescribeDohUserInfoResponseBody
	GetRequestId() *string
	SetSubDomainCount(v int32) *DescribeDohUserInfoResponseBody
	GetSubDomainCount() *int32
}

type DescribeDohUserInfoResponseBody struct {
	// The number of accessed domains.
	//
	// example:
	//
	// 123
	DomainCount *int32 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The ID of the Alibaba Cloud public DNS user.
	//
	// example:
	//
	// 12345678
	PdnsId *int64 `json:"PdnsId,omitempty" xml:"PdnsId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of accessed subdomains.
	//
	// example:
	//
	// 123
	SubDomainCount *int32 `json:"SubDomainCount,omitempty" xml:"SubDomainCount,omitempty"`
}

func (s DescribeDohUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohUserInfoResponseBody) GetDomainCount() *int32 {
	return s.DomainCount
}

func (s *DescribeDohUserInfoResponseBody) GetPdnsId() *int64 {
	return s.PdnsId
}

func (s *DescribeDohUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohUserInfoResponseBody) GetSubDomainCount() *int32 {
	return s.SubDomainCount
}

func (s *DescribeDohUserInfoResponseBody) SetDomainCount(v int32) *DescribeDohUserInfoResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetPdnsId(v int64) *DescribeDohUserInfoResponseBody {
	s.PdnsId = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetRequestId(v string) *DescribeDohUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) SetSubDomainCount(v int32) *DescribeDohUserInfoResponseBody {
	s.SubDomainCount = &v
	return s
}

func (s *DescribeDohUserInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
