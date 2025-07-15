// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoomRealTimeStreamAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoomRealTimeStreamAddressResponseBody
	GetRequestId() *string
	SetRtmpAddress(v string) *CreateRoomRealTimeStreamAddressResponseBody
	GetRtmpAddress() *string
}

type CreateRoomRealTimeStreamAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F8DB7E25-6A35-161A-AA41-B7A658AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The RTMP ingest URL.
	RtmpAddress *string `json:"RtmpAddress,omitempty" xml:"RtmpAddress,omitempty"`
}

func (s CreateRoomRealTimeStreamAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoomRealTimeStreamAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoomRealTimeStreamAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoomRealTimeStreamAddressResponseBody) GetRtmpAddress() *string {
	return s.RtmpAddress
}

func (s *CreateRoomRealTimeStreamAddressResponseBody) SetRequestId(v string) *CreateRoomRealTimeStreamAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressResponseBody) SetRtmpAddress(v string) *CreateRoomRealTimeStreamAddressResponseBody {
	s.RtmpAddress = &v
	return s
}

func (s *CreateRoomRealTimeStreamAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
