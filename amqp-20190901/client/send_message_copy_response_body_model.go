// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageCopyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SendMessageCopyResponseBody
	GetCode() *int32
	SetData(v bool) *SendMessageCopyResponseBody
	GetData() *bool
	SetMessage(v string) *SendMessageCopyResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendMessageCopyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendMessageCopyResponseBody
	GetSuccess() *bool
}

type SendMessageCopyResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageCopyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageCopyResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageCopyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SendMessageCopyResponseBody) GetData() *bool {
	return s.Data
}

func (s *SendMessageCopyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendMessageCopyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageCopyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendMessageCopyResponseBody) SetCode(v int32) *SendMessageCopyResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageCopyResponseBody) SetData(v bool) *SendMessageCopyResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageCopyResponseBody) SetMessage(v string) *SendMessageCopyResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageCopyResponseBody) SetRequestId(v string) *SendMessageCopyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageCopyResponseBody) SetSuccess(v bool) *SendMessageCopyResponseBody {
	s.Success = &v
	return s
}

func (s *SendMessageCopyResponseBody) Validate() error {
	return dara.Validate(s)
}
