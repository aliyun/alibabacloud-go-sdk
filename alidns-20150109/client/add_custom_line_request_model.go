// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddCustomLineRequest
	GetDomainName() *string
	SetIpSegment(v []*AddCustomLineRequestIpSegment) *AddCustomLineRequest
	GetIpSegment() []*AddCustomLineRequestIpSegment
	SetLang(v string) *AddCustomLineRequest
	GetLang() *string
	SetLineName(v string) *AddCustomLineRequest
	GetLineName() *string
}

type AddCustomLineRequest struct {
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The CIDR blocks.
	//
	// This parameter is required.
	IpSegment []*AddCustomLineRequestIpSegment `json:"IpSegment,omitempty" xml:"IpSegment,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the custom line.
	//
	// This parameter is required.
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s AddCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineRequest) GoString() string {
	return s.String()
}

func (s *AddCustomLineRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddCustomLineRequest) GetIpSegment() []*AddCustomLineRequestIpSegment {
	return s.IpSegment
}

func (s *AddCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *AddCustomLineRequest) GetLineName() *string {
	return s.LineName
}

func (s *AddCustomLineRequest) SetDomainName(v string) *AddCustomLineRequest {
	s.DomainName = &v
	return s
}

func (s *AddCustomLineRequest) SetIpSegment(v []*AddCustomLineRequestIpSegment) *AddCustomLineRequest {
	s.IpSegment = v
	return s
}

func (s *AddCustomLineRequest) SetLang(v string) *AddCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *AddCustomLineRequest) SetLineName(v string) *AddCustomLineRequest {
	s.LineName = &v
	return s
}

func (s *AddCustomLineRequest) Validate() error {
	if s.IpSegment != nil {
		for _, item := range s.IpSegment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddCustomLineRequestIpSegment struct {
	// The end IP address of the CIDR block.
	//
	// example:
	//
	// 192.0.2.254
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The start IP address of the CIDR block.
	//
	// example:
	//
	// 192.0.2.0
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s AddCustomLineRequestIpSegment) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLineRequestIpSegment) GoString() string {
	return s.String()
}

func (s *AddCustomLineRequestIpSegment) GetEndIp() *string {
	return s.EndIp
}

func (s *AddCustomLineRequestIpSegment) GetStartIp() *string {
	return s.StartIp
}

func (s *AddCustomLineRequestIpSegment) SetEndIp(v string) *AddCustomLineRequestIpSegment {
	s.EndIp = &v
	return s
}

func (s *AddCustomLineRequestIpSegment) SetStartIp(v string) *AddCustomLineRequestIpSegment {
	s.StartIp = &v
	return s
}

func (s *AddCustomLineRequestIpSegment) Validate() error {
	return dara.Validate(s)
}
