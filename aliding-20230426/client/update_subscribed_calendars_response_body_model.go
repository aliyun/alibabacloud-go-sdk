// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSubscribedCalendarsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateSubscribedCalendarsResponseBody
	GetResult() *bool
}

type UpdateSubscribedCalendarsResponseBody struct {
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

func (s UpdateSubscribedCalendarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubscribedCalendarsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateSubscribedCalendarsResponseBody) SetRequestId(v string) *UpdateSubscribedCalendarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubscribedCalendarsResponseBody) SetResult(v bool) *UpdateSubscribedCalendarsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateSubscribedCalendarsResponseBody) Validate() error {
	return dara.Validate(s)
}
