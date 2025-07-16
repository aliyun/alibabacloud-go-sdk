// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMeetingRoomResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateMeetingRoomResponseBody
	GetResult() *string
	SetVendorRequestId(v string) *CreateMeetingRoomResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateMeetingRoomResponseBody
	GetVendorType() *string
}

type CreateMeetingRoomResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0ffb718xxxxx
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateMeetingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMeetingRoomResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateMeetingRoomResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateMeetingRoomResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateMeetingRoomResponseBody) SetRequestId(v string) *CreateMeetingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMeetingRoomResponseBody) SetResult(v string) *CreateMeetingRoomResponseBody {
	s.Result = &v
	return s
}

func (s *CreateMeetingRoomResponseBody) SetVendorRequestId(v string) *CreateMeetingRoomResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateMeetingRoomResponseBody) SetVendorType(v string) *CreateMeetingRoomResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateMeetingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
