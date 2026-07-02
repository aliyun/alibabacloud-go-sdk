// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlLastPayloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDstPortList(v string) *DescribeSdlLastPayloadResponseBody
	GetDstPortList() *string
	SetPayload(v string) *DescribeSdlLastPayloadResponseBody
	GetPayload() *string
	SetProtoList(v string) *DescribeSdlLastPayloadResponseBody
	GetProtoList() *string
	SetRequestId(v string) *DescribeSdlLastPayloadResponseBody
	GetRequestId() *string
	SetSrcPortList(v string) *DescribeSdlLastPayloadResponseBody
	GetSrcPortList() *string
}

type DescribeSdlLastPayloadResponseBody struct {
	// The destination port.
	//
	// example:
	//
	// 80
	DstPortList *string `json:"DstPortList,omitempty" xml:"DstPortList,omitempty"`
	// The attack payload of the intrusion prevention event.
	//
	// example:
	//
	// 302902010004067075626c6963a01c0204036a5f43020100020100300e300c06082b060102010101000500
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The protocol.
	//
	// example:
	//
	// tcp
	ProtoList *string `json:"ProtoList,omitempty" xml:"ProtoList,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source port.
	//
	// example:
	//
	// 1586
	SrcPortList *string `json:"SrcPortList,omitempty" xml:"SrcPortList,omitempty"`
}

func (s DescribeSdlLastPayloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlLastPayloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdlLastPayloadResponseBody) GetDstPortList() *string {
	return s.DstPortList
}

func (s *DescribeSdlLastPayloadResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *DescribeSdlLastPayloadResponseBody) GetProtoList() *string {
	return s.ProtoList
}

func (s *DescribeSdlLastPayloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSdlLastPayloadResponseBody) GetSrcPortList() *string {
	return s.SrcPortList
}

func (s *DescribeSdlLastPayloadResponseBody) SetDstPortList(v string) *DescribeSdlLastPayloadResponseBody {
	s.DstPortList = &v
	return s
}

func (s *DescribeSdlLastPayloadResponseBody) SetPayload(v string) *DescribeSdlLastPayloadResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeSdlLastPayloadResponseBody) SetProtoList(v string) *DescribeSdlLastPayloadResponseBody {
	s.ProtoList = &v
	return s
}

func (s *DescribeSdlLastPayloadResponseBody) SetRequestId(v string) *DescribeSdlLastPayloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdlLastPayloadResponseBody) SetSrcPortList(v string) *DescribeSdlLastPayloadResponseBody {
	s.SrcPortList = &v
	return s
}

func (s *DescribeSdlLastPayloadResponseBody) Validate() error {
	return dara.Validate(s)
}
