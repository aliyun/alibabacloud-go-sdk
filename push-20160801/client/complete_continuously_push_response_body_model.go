// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteContinuouslyPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *CompleteContinuouslyPushResponseBody
	GetMessageId() *string
	SetRequestId(v string) *CompleteContinuouslyPushResponseBody
	GetRequestId() *string
}

type CompleteContinuouslyPushResponseBody struct {
	// example:
	//
	// 4010290149170430
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompleteContinuouslyPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteContinuouslyPushResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteContinuouslyPushResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *CompleteContinuouslyPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteContinuouslyPushResponseBody) SetMessageId(v string) *CompleteContinuouslyPushResponseBody {
	s.MessageId = &v
	return s
}

func (s *CompleteContinuouslyPushResponseBody) SetRequestId(v string) *CompleteContinuouslyPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteContinuouslyPushResponseBody) Validate() error {
	return dara.Validate(s)
}
