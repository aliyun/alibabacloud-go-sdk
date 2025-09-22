// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserVipsByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodUserVipsByDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeVodUserVipsByDomainResponseBody
	GetRequestId() *string
	SetVips(v *DescribeVodUserVipsByDomainResponseBodyVips) *DescribeVodUserVipsByDomainResponseBody
	GetVips() *DescribeVodUserVipsByDomainResponseBodyVips
}

type DescribeVodUserVipsByDomainResponseBody struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 6730AC93-7B12-4B*****7F-49EE1FE8BC49
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Vips      *DescribeVodUserVipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeVodUserVipsByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserVipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserVipsByDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodUserVipsByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserVipsByDomainResponseBody) GetVips() *DescribeVodUserVipsByDomainResponseBodyVips {
	return s.Vips
}

func (s *DescribeVodUserVipsByDomainResponseBody) SetDomainName(v string) *DescribeVodUserVipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserVipsByDomainResponseBody) SetRequestId(v string) *DescribeVodUserVipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserVipsByDomainResponseBody) SetVips(v *DescribeVodUserVipsByDomainResponseBodyVips) *DescribeVodUserVipsByDomainResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeVodUserVipsByDomainResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserVipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeVodUserVipsByDomainResponseBodyVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserVipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeVodUserVipsByDomainResponseBodyVips) GetVip() []*string {
	return s.Vip
}

func (s *DescribeVodUserVipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeVodUserVipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

func (s *DescribeVodUserVipsByDomainResponseBodyVips) Validate() error {
	return dara.Validate(s)
}
