// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuouslyPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *ContinuouslyPushResponseBody
	GetMessageId() *string
	SetRequestId(v string) *ContinuouslyPushResponseBody
	GetRequestId() *string
}

type ContinuouslyPushResponseBody struct {
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 500131
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuouslyPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinuouslyPushResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuouslyPushResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ContinuouslyPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinuouslyPushResponseBody) SetMessageId(v string) *ContinuouslyPushResponseBody {
	s.MessageId = &v
	return s
}

func (s *ContinuouslyPushResponseBody) SetRequestId(v string) *ContinuouslyPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinuouslyPushResponseBody) Validate() error {
	return dara.Validate(s)
}
