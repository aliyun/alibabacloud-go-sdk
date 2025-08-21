// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToAndroidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushMessageToAndroidResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushMessageToAndroidResponseBody
	GetRequestId() *string
}

type PushMessageToAndroidResponseBody struct {
	// example:
	//
	// 501029
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMessageToAndroidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToAndroidResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushMessageToAndroidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushMessageToAndroidResponseBody) SetMessageId(v string) *PushMessageToAndroidResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushMessageToAndroidResponseBody) SetRequestId(v string) *PushMessageToAndroidResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMessageToAndroidResponseBody) Validate() error {
	return dara.Validate(s)
}
