// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPredictSseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPredictSseResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPredictSseResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPredictSseResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPredictSseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPredictSseResponseBody
	GetSuccess() *bool
}

type ChatBIPredictSseResponseBody struct {
	// example:
	//
	// sql，取数，总结，图表
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 错误码
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 737A7E8B-9C9E-5EE6-8A04-D61400A5D083
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBIPredictSseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPredictSseResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPredictSseResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPredictSseResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPredictSseResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPredictSseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPredictSseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPredictSseResponseBody) SetData(v interface{}) *ChatBIPredictSseResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPredictSseResponseBody) SetErrCode(v string) *ChatBIPredictSseResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPredictSseResponseBody) SetErrMessage(v string) *ChatBIPredictSseResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPredictSseResponseBody) SetRequestId(v string) *ChatBIPredictSseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPredictSseResponseBody) SetSuccess(v bool) *ChatBIPredictSseResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPredictSseResponseBody) Validate() error {
	return dara.Validate(s)
}
