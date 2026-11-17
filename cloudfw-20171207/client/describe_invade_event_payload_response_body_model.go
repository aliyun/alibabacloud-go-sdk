// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventPayloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInvadeEventPayloadResponseBody
	GetAssetsInstanceId() *string
	SetEventUuid(v string) *DescribeInvadeEventPayloadResponseBody
	GetEventUuid() *string
	SetHighlightEnd(v int32) *DescribeInvadeEventPayloadResponseBody
	GetHighlightEnd() *int32
	SetHighlightStart(v int32) *DescribeInvadeEventPayloadResponseBody
	GetHighlightStart() *int32
	SetPayload(v string) *DescribeInvadeEventPayloadResponseBody
	GetPayload() *string
	SetPayloadLen(v int32) *DescribeInvadeEventPayloadResponseBody
	GetPayloadLen() *int32
	SetRequestId(v string) *DescribeInvadeEventPayloadResponseBody
	GetRequestId() *string
}

type DescribeInvadeEventPayloadResponseBody struct {
	// The instance ID of the asset to which the event belongs.
	//
	// example:
	//
	// i-uf6faknmuby7ezht****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// fadd-dfdd-****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The end position (character index) of the matched content in the payload.
	//
	// example:
	//
	// 0
	HighlightEnd *int32 `json:"HighlightEnd,omitempty" xml:"HighlightEnd,omitempty"`
	// The start position (character index) of the matched content in the payload.
	//
	// example:
	//
	// 0
	HighlightStart *int32 `json:"HighlightStart,omitempty" xml:"HighlightStart,omitempty"`
	// The original packet payload of the intrusion event.
	//
	// example:
	//
	// 302902010004067075626c6963a01c0204036a5f43020100020100300e300c06082b060102010101000500
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The length of the payload in bytes.
	//
	// example:
	//
	// 457
	PayloadLen *int32 `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558FE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvadeEventPayloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventPayloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventPayloadResponseBody) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventPayloadResponseBody) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventPayloadResponseBody) GetHighlightEnd() *int32 {
	return s.HighlightEnd
}

func (s *DescribeInvadeEventPayloadResponseBody) GetHighlightStart() *int32 {
	return s.HighlightStart
}

func (s *DescribeInvadeEventPayloadResponseBody) GetPayload() *string {
	return s.Payload
}

func (s *DescribeInvadeEventPayloadResponseBody) GetPayloadLen() *int32 {
	return s.PayloadLen
}

func (s *DescribeInvadeEventPayloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEventPayloadResponseBody) SetAssetsInstanceId(v string) *DescribeInvadeEventPayloadResponseBody {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetEventUuid(v string) *DescribeInvadeEventPayloadResponseBody {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetHighlightEnd(v int32) *DescribeInvadeEventPayloadResponseBody {
	s.HighlightEnd = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetHighlightStart(v int32) *DescribeInvadeEventPayloadResponseBody {
	s.HighlightStart = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetPayload(v string) *DescribeInvadeEventPayloadResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetPayloadLen(v int32) *DescribeInvadeEventPayloadResponseBody {
	s.PayloadLen = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) SetRequestId(v string) *DescribeInvadeEventPayloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEventPayloadResponseBody) Validate() error {
	return dara.Validate(s)
}
