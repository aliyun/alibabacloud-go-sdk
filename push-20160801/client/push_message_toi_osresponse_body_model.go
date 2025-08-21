// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToiOSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushMessageToiOSResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushMessageToiOSResponseBody
	GetRequestId() *string
}

type PushMessageToiOSResponseBody struct {
	// example:
	//
	// 501029
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushMessageToiOSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToiOSResponseBody) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushMessageToiOSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushMessageToiOSResponseBody) SetMessageId(v string) *PushMessageToiOSResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushMessageToiOSResponseBody) SetRequestId(v string) *PushMessageToiOSResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushMessageToiOSResponseBody) Validate() error {
	return dara.Validate(s)
}
