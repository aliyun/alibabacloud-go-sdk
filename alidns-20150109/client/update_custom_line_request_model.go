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
	// The list of IP ranges. Use a hyphen (-) to separate the start and end IP addresses. Specify one IP segment per line. You can specify 1 to 50 IP ranges. To specify a single IP address, use the format IP1-IP1. The IP ranges cannot overlap.
	IpSegment []*UpdateCustomLineRequestIpSegment `json:"IpSegment,omitempty" xml:"IpSegment,omitempty" type:"Repeated"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique ID of the custom line. You can call [DescribeCustomLines](https://help.aliyun.com/document_detail/2355671.html) to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1*******
	LineId *int64 `json:"LineId,omitempty" xml:"LineId,omitempty"`
	// The name of the custom line. The name must be 1 to 20 characters long and can contain Chinese characters, letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// 望京线路
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

type UpdateCustomLineRequestIpSegment struct {
	// The end IP address of the IP range.
	//
	// example:
	//
	// 2.2.XX.XX
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The start IP address of the IP range.
	//
	// example:
	//
	// 1.1.XX.XX
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
