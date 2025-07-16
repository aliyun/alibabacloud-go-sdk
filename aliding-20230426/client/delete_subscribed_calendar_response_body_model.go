// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscribedCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResult(v bool) *DeleteSubscribedCalendarResponseBody
	GetResult() *bool
	SetRequestId(v string) *DeleteSubscribedCalendarResponseBody
	GetRequestId() *string
}

type DeleteSubscribedCalendarResponseBody struct {
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
}

func (s DeleteSubscribedCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteSubscribedCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSubscribedCalendarResponseBody) SetResult(v bool) *DeleteSubscribedCalendarResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSubscribedCalendarResponseBody) SetRequestId(v string) *DeleteSubscribedCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscribedCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
