// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVipsByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeUserVipsByDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeUserVipsByDomainResponseBody
	GetRequestId() *string
	SetVips(v *DescribeUserVipsByDomainResponseBodyVips) *DescribeUserVipsByDomainResponseBody
	GetVips() *DescribeUserVipsByDomainResponseBodyVips
}

type DescribeUserVipsByDomainResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 820E7900-5CA9-4AEF-B0DD-20ED5F64BE55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of VIPs.
	Vips *DescribeUserVipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeUserVipsByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeUserVipsByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserVipsByDomainResponseBody) GetVips() *DescribeUserVipsByDomainResponseBodyVips {
	return s.Vips
}

func (s *DescribeUserVipsByDomainResponseBody) SetDomainName(v string) *DescribeUserVipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeUserVipsByDomainResponseBody) SetRequestId(v string) *DescribeUserVipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserVipsByDomainResponseBody) SetVips(v *DescribeUserVipsByDomainResponseBodyVips) *DescribeUserVipsByDomainResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeUserVipsByDomainResponseBody) Validate() error {
	if s.Vips != nil {
		if err := s.Vips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserVipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeUserVipsByDomainResponseBodyVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeUserVipsByDomainResponseBodyVips) GetVip() []*string {
	return s.Vip
}

func (s *DescribeUserVipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeUserVipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

func (s *DescribeUserVipsByDomainResponseBodyVips) Validate() error {
	return dara.Validate(s)
}
