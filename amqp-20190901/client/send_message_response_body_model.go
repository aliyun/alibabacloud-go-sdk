// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SendMessageResponseBody
	GetCode() *int32
	SetData(v bool) *SendMessageResponseBody
	GetData() *bool
	SetMessage(v string) *SendMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendMessageResponseBody
	GetSuccess() *bool
}

type SendMessageResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SendMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *SendMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendMessageResponseBody) SetCode(v int32) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetData(v bool) *SendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) SetSuccess(v bool) *SendMessageResponseBody {
	s.Success = &v
	return s
}

func (s *SendMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
