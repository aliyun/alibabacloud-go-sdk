// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushResponseBody
	GetRequestId() *string
}

type PushResponseBody struct {
	// example:
	//
	// 501029
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushResponseBody) GoString() string {
	return s.String()
}

func (s *PushResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushResponseBody) SetMessageId(v string) *PushResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushResponseBody) SetRequestId(v string) *PushResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushResponseBody) Validate() error {
	return dara.Validate(s)
}
