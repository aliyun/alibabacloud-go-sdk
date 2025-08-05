// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskEventPayloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDstIP(v string) *DescribeRiskEventPayloadResponseBody
	GetDstIP() *string
	SetDstPort(v int32) *DescribeRiskEventPayloadResponseBody
	GetDstPort() *int32
	SetDstVpcId(v string) *DescribeRiskEventPayloadResponseBody
	GetDstVpcId() *string
	SetHitContentType(v int32) *DescribeRiskEventPayloadResponseBody
	GetHitContentType() *int32
	SetHitTo(v int32) *DescribeRiskEventPayloadResponseBody
	GetHitTo() *int32
	SetParsedContent(v string) *DescribeRiskEventPayloadResponseBody
	GetParsedContent() *string
	SetPayload(v string) *DescribeRiskEventPayloadResponseBody
	GetPayload() *string
	SetPayloadLen(v int32) *DescribeRiskEventPayloadResponseBody
	GetPayloadLen() *int32
	SetProto(v string) *DescribeRiskEventPayloadResponseBody
	GetProto() *string
	SetRealIp(v string) *DescribeRiskEventPayloadResponseBody
	GetRealIp() *string
	SetRequestId(v string) *DescribeRiskEventPayloadResponseBody
	GetRequestId() *string
	SetSrcIP(v string) *DescribeRiskEventPayloadResponseBody
	GetSrcIP() *string
	SetSrcPort(v int32) *DescribeRiskEventPayloadResponseBody
	GetSrcPort() *int32
	SetSrcVpcId(v string) *DescribeRiskEventPayloadResponseBody
	GetSrcVpcId() *string
	SetXForwardFor(v string) *DescribeRiskEventPayloadResponseBody
	GetXForwardFor() *string
}

type DescribeRiskEventPayloadResponseBody struct {
	// The destination IP address of the intrusion event.
	//
	// example:
	//
	// 203.0.113.1
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The destination port of the intrusion event.
	//
	// example:
	//
	// 8080
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The destination VPC ID of the intrusion event.
	//
	// example:
	//
	// vpc-bp10w5nb30r4jzfyc****
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// Type of the hit.
	//
	// example:
	//
	// 1
	HitContentType *int32 `json:"HitContentType,omitempty" xml:"HitContentType,omitempty"`
	// The position where the hit ends.
	//
	// example:
	//
	// 67
	HitTo *int32 `json:"HitTo,omitempty" xml:"HitTo,omitempty"`
	// Hit payload.
	//
	// example:
	//
	// 2f636f6d706f7365722f73656e645f656d61696c3f746f3d6d61667740776f66736f7961792675726c3d687474703a2f2f302e302e302e303a31323334352f692f6431366530312f313664622f673670772f
	ParsedContent *string `json:"ParsedContent,omitempty" xml:"ParsedContent,omitempty"`
	// The attack payload of the intrusion event.
	//
	// example:
	//
	// 302902010004067075626c6963a01c0204036a5f43020100020100300e300c06082b060102010101000500
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The length of the attack payload of the intrusion event.
	//
	// example:
	//
	// 457
	PayloadLen *int32 `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty"`
	// The protocol type of intrusion event. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The HTTP X-Real-IP field.
	//
	// example:
	//
	// 203.0.113.3
	RealIp *string `json:"RealIp,omitempty" xml:"RealIp,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68055BA4-D8BD-5611-AC49-C651E619A12E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source IP address of the intrusion event.
	//
	// example:
	//
	// 203.0.113.2
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The source port of the intrusion event.
	//
	// example:
	//
	// 54360
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The source VPC ID of the intrusion event.
	//
	// example:
	//
	// vpc-t4nlt09olhpazpoeg****
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The HTTP X-Forwarded-For field.
	//
	// example:
	//
	// 203.0.113.4
	XForwardFor *string `json:"XForwardFor,omitempty" xml:"XForwardFor,omitempty"`
}

func (s DescribeRiskEventPayloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskEventPayloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadResponseBody) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeRiskEventPayloadResponseBody) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeRiskEventPayloadResponseBody) GetDstVpcId() *string {
	return s.DstVpcId
}

func (s *DescribeRiskEventPayloadResponseBody) GetHitContentType() *int32 {
	return s.HitContentType
}

func (s *DescribeRiskEventPayloadResponseBody) GetHitTo() *int32 {
	return s.HitTo
}

func (s *DescribeRiskEventPayloadResponseBody) GetParsedContent() *string {
	return s.ParsedContent
}

func (s *DescribeRiskEventPayloadResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *DescribeRiskEventPayloadResponseBody) GetPayloadLen() *int32 {
	return s.PayloadLen
}

func (s *DescribeRiskEventPayloadResponseBody) GetProto() *string {
	return s.Proto
}

func (s *DescribeRiskEventPayloadResponseBody) GetRealIp() *string {
	return s.RealIp
}

func (s *DescribeRiskEventPayloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskEventPayloadResponseBody) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeRiskEventPayloadResponseBody) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *DescribeRiskEventPayloadResponseBody) GetSrcVpcId() *string {
	return s.SrcVpcId
}

func (s *DescribeRiskEventPayloadResponseBody) GetXForwardFor() *string {
	return s.XForwardFor
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstIP(v string) *DescribeRiskEventPayloadResponseBody {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstPort(v int32) *DescribeRiskEventPayloadResponseBody {
	s.DstPort = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstVpcId(v string) *DescribeRiskEventPayloadResponseBody {
	s.DstVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetHitContentType(v int32) *DescribeRiskEventPayloadResponseBody {
	s.HitContentType = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetHitTo(v int32) *DescribeRiskEventPayloadResponseBody {
	s.HitTo = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetParsedContent(v string) *DescribeRiskEventPayloadResponseBody {
	s.ParsedContent = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetPayload(v string) *DescribeRiskEventPayloadResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetPayloadLen(v int32) *DescribeRiskEventPayloadResponseBody {
	s.PayloadLen = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetProto(v string) *DescribeRiskEventPayloadResponseBody {
	s.Proto = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetRealIp(v string) *DescribeRiskEventPayloadResponseBody {
	s.RealIp = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetRequestId(v string) *DescribeRiskEventPayloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcIP(v string) *DescribeRiskEventPayloadResponseBody {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcPort(v int32) *DescribeRiskEventPayloadResponseBody {
	s.SrcPort = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcVpcId(v string) *DescribeRiskEventPayloadResponseBody {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetXForwardFor(v string) *DescribeRiskEventPayloadResponseBody {
	s.XForwardFor = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) Validate() error {
	return dara.Validate(s)
}
