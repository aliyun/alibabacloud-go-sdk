// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToiOSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushNoticeToiOSResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushNoticeToiOSResponseBody
	GetRequestId() *string
}

type PushNoticeToiOSResponseBody struct {
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 501029
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushNoticeToiOSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToiOSResponseBody) GoString() string {
	return s.String()
}

func (s *PushNoticeToiOSResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushNoticeToiOSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushNoticeToiOSResponseBody) SetMessageId(v string) *PushNoticeToiOSResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushNoticeToiOSResponseBody) SetRequestId(v string) *PushNoticeToiOSResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushNoticeToiOSResponseBody) Validate() error {
	return dara.Validate(s)
}
