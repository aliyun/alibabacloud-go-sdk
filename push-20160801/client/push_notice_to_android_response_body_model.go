// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushNoticeToAndroidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushNoticeToAndroidResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushNoticeToAndroidResponseBody
	GetRequestId() *string
}

type PushNoticeToAndroidResponseBody struct {
	// example:
	//
	// 501029
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushNoticeToAndroidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushNoticeToAndroidResponseBody) GoString() string {
	return s.String()
}

func (s *PushNoticeToAndroidResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushNoticeToAndroidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushNoticeToAndroidResponseBody) SetMessageId(v string) *PushNoticeToAndroidResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushNoticeToAndroidResponseBody) SetRequestId(v string) *PushNoticeToAndroidResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushNoticeToAndroidResponseBody) Validate() error {
	return dara.Validate(s)
}
