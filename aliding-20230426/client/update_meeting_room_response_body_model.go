// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResult(v bool) *UpdateMeetingRoomResponseBody
	GetResult() *bool
	SetRequestId(v string) *UpdateMeetingRoomResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *UpdateMeetingRoomResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *UpdateMeetingRoomResponseBody
	GetVendorType() *string
}

type UpdateMeetingRoomResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s UpdateMeetingRoomResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateMeetingRoomResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMeetingRoomResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *UpdateMeetingRoomResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *UpdateMeetingRoomResponseBody) SetResult(v bool) *UpdateMeetingRoomResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateMeetingRoomResponseBody) SetRequestId(v string) *UpdateMeetingRoomResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMeetingRoomResponseBody) SetVendorRequestId(v string) *UpdateMeetingRoomResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *UpdateMeetingRoomResponseBody) SetVendorType(v string) *UpdateMeetingRoomResponseBody {
	s.VendorType = &v
	return s
}

func (s *UpdateMeetingRoomResponseBody) Validate() error {
	return dara.Validate(s)
}
