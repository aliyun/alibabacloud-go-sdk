// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpSegment(v []*UpdateCustomLineRequestIpSegment) *UpdateCustomLineRequest
	GetIpSegment() []*UpdateCustomLineRequestIpSegment
	SetLang(v string) *UpdateCustomLineRequest
	GetLang() *string
	SetLineId(v int64) *UpdateCustomLineRequest
	GetLineId() *int64
	SetLineName(v string) *UpdateCustomLineRequest
	GetLineName() *string
}

type UpdateCustomLineRequest struct {
	// The CIDR blocks. Separate IP addresses with a hyphen (-). Enter a CIDR block in each row. You can enter 1 to 50 CIDR blocks at a time. If a CIDR block contains only one IP address, enter the IP address in the format of IP1-IP1. Different CIDR blocks cannot be overlapped.
	IpSegment []*UpdateCustomLineRequestIpSegment `json:"IpSegment,omitempty" xml:"IpSegment,omitempty" type:"Repeated"`
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
	// The unique ID of the custom line. You can call the [DescribeCustomLines](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describecustomlines?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LineId *int64 `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The name of the custom line. The name must be 1 to 20 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s UpdateCustomLineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineRequest) GetIpSegment() []*UpdateCustomLineRequestIpSegment {
	return s.IpSegment
}

func (s *UpdateCustomLineRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateCustomLineRequest) GetLineId() *int64 {
	return s.LineId
}

func (s *UpdateCustomLineRequest) GetLineName() *string {
	return s.LineName
}

func (s *UpdateCustomLineRequest) SetIpSegment(v []*UpdateCustomLineRequestIpSegment) *UpdateCustomLineRequest {
	s.IpSegment = v
	return s
}

func (s *UpdateCustomLineRequest) SetLang(v string) *UpdateCustomLineRequest {
	s.Lang = &v
	return s
}

func (s *UpdateCustomLineRequest) SetLineId(v int64) *UpdateCustomLineRequest {
	s.LineId = &v
	return s
}

func (s *UpdateCustomLineRequest) SetLineName(v string) *UpdateCustomLineRequest {
	s.LineName = &v
	return s
}

func (s *UpdateCustomLineRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateCustomLineRequestIpSegment struct {
	// The end IP address of the CIDR block.
	//
	// example:
	//
	// 2.2.2.2
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The start IP address of the CIDR block.
	//
	// example:
	//
	// 1.1.1.1
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s UpdateCustomLineRequestIpSegment) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLineRequestIpSegment) GoString() string {
	return s.String()
}

func (s *UpdateCustomLineRequestIpSegment) GetEndIp() *string {
	return s.EndIp
}

func (s *UpdateCustomLineRequestIpSegment) GetStartIp() *string {
	return s.StartIp
}

func (s *UpdateCustomLineRequestIpSegment) SetEndIp(v string) *UpdateCustomLineRequestIpSegment {
	s.EndIp = &v
	return s
}

func (s *UpdateCustomLineRequestIpSegment) SetStartIp(v string) *UpdateCustomLineRequestIpSegment {
	s.StartIp = &v
	return s
}

func (s *UpdateCustomLineRequestIpSegment) Validate() error {
	return dara.Validate(s)
}
