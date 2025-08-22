// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserVipsByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnUserVipsByDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeDcdnUserVipsByDomainResponseBody
	GetRequestId() *string
	SetVips(v *DescribeDcdnUserVipsByDomainResponseBodyVips) *DescribeDcdnUserVipsByDomainResponseBody
	GetVips() *DescribeDcdnUserVipsByDomainResponseBodyVips
}

type DescribeDcdnUserVipsByDomainResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 820E7900-5CA9-4AEF-B0DD-20ED5F64BE55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of VIPs.
	Vips *DescribeDcdnUserVipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeDcdnUserVipsByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserVipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) GetVips() *DescribeDcdnUserVipsByDomainResponseBodyVips {
	return s.Vips
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) SetDomainName(v string) *DescribeDcdnUserVipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) SetRequestId(v string) *DescribeDcdnUserVipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) SetVips(v *DescribeDcdnUserVipsByDomainResponseBodyVips) *DescribeDcdnUserVipsByDomainResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserVipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserVipsByDomainResponseBodyVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserVipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserVipsByDomainResponseBodyVips) GetVip() []*string {
	return s.Vip
}

func (s *DescribeDcdnUserVipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeDcdnUserVipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

func (s *DescribeDcdnUserVipsByDomainResponseBodyVips) Validate() error {
	return dara.Validate(s)
}
