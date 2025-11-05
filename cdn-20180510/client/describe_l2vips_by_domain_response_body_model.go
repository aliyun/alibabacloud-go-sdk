// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL2VipsByDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeL2VipsByDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeL2VipsByDomainResponseBody
	GetRequestId() *string
	SetVips(v *DescribeL2VipsByDomainResponseBodyVips) *DescribeL2VipsByDomainResponseBody
	GetVips() *DescribeL2VipsByDomainResponseBodyVips
}

type DescribeL2VipsByDomainResponseBody struct {
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
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of VIPs.
	Vips *DescribeL2VipsByDomainResponseBodyVips `json:"Vips,omitempty" xml:"Vips,omitempty" type:"Struct"`
}

func (s DescribeL2VipsByDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeL2VipsByDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeL2VipsByDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeL2VipsByDomainResponseBody) GetVips() *DescribeL2VipsByDomainResponseBodyVips {
	return s.Vips
}

func (s *DescribeL2VipsByDomainResponseBody) SetDomainName(v string) *DescribeL2VipsByDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeL2VipsByDomainResponseBody) SetRequestId(v string) *DescribeL2VipsByDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeL2VipsByDomainResponseBody) SetVips(v *DescribeL2VipsByDomainResponseBodyVips) *DescribeL2VipsByDomainResponseBody {
	s.Vips = v
	return s
}

func (s *DescribeL2VipsByDomainResponseBody) Validate() error {
	if s.Vips != nil {
		if err := s.Vips.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeL2VipsByDomainResponseBodyVips struct {
	Vip []*string `json:"Vip,omitempty" xml:"Vip,omitempty" type:"Repeated"`
}

func (s DescribeL2VipsByDomainResponseBodyVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeL2VipsByDomainResponseBodyVips) GoString() string {
	return s.String()
}

func (s *DescribeL2VipsByDomainResponseBodyVips) GetVip() []*string {
	return s.Vip
}

func (s *DescribeL2VipsByDomainResponseBodyVips) SetVip(v []*string) *DescribeL2VipsByDomainResponseBodyVips {
	s.Vip = v
	return s
}

func (s *DescribeL2VipsByDomainResponseBodyVips) Validate() error {
	return dara.Validate(s)
}
