// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubscribeCalendarResponseBody
	GetRequestId() *string
}

type SubscribeCalendarResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SubscribeCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubscribeCalendarResponseBody) SetRequestId(v string) *SubscribeCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
