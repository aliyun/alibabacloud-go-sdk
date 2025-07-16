// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveMeetingRoomsResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemoveMeetingRoomsResponseBody
	GetResult() *bool
}

type RemoveMeetingRoomsResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveMeetingRoomsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveMeetingRoomsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemoveMeetingRoomsResponseBody) SetRequestId(v string) *RemoveMeetingRoomsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveMeetingRoomsResponseBody) SetResult(v bool) *RemoveMeetingRoomsResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveMeetingRoomsResponseBody) Validate() error {
	return dara.Validate(s)
}
