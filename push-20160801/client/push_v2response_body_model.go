// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *PushV2ResponseBody
	GetMessageId() *string
	SetRequestId(v string) *PushV2ResponseBody
	GetRequestId() *string
}

type PushV2ResponseBody struct {
	// example:
	//
	// 11747540****88320
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 159E4422-6624-****-8943-DFD98D34858C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushV2ResponseBody) GoString() string {
	return s.String()
}

func (s *PushV2ResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PushV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushV2ResponseBody) SetMessageId(v string) *PushV2ResponseBody {
	s.MessageId = &v
	return s
}

func (s *PushV2ResponseBody) SetRequestId(v string) *PushV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
