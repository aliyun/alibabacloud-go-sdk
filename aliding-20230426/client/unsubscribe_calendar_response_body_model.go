// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnsubscribeCalendarResponseBody
	GetRequestId() *string
	SetResult(v bool) *UnsubscribeCalendarResponseBody
	GetResult() *bool
}

type UnsubscribeCalendarResponseBody struct {
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

func (s UnsubscribeCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnsubscribeCalendarResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UnsubscribeCalendarResponseBody) SetRequestId(v string) *UnsubscribeCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeCalendarResponseBody) SetResult(v bool) *UnsubscribeCalendarResponseBody {
	s.Result = &v
	return s
}

func (s *UnsubscribeCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
