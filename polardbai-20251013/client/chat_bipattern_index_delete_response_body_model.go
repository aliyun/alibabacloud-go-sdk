// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternIndexDeleteResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternIndexDeleteResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternIndexDeleteResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternIndexDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternIndexDeleteResponseBody
	GetSuccess() *bool
}

type ChatBIPatternIndexDeleteResponseBody struct {
	// example:
	//
	// true/false
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
	// 983863E2-****-1D15-A4AE-3468CD75383D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBIPatternIndexDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexDeleteResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternIndexDeleteResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternIndexDeleteResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternIndexDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternIndexDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternIndexDeleteResponseBody) SetData(v interface{}) *ChatBIPatternIndexDeleteResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternIndexDeleteResponseBody) SetErrCode(v string) *ChatBIPatternIndexDeleteResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternIndexDeleteResponseBody) SetErrMessage(v string) *ChatBIPatternIndexDeleteResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternIndexDeleteResponseBody) SetRequestId(v string) *ChatBIPatternIndexDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternIndexDeleteResponseBody) SetSuccess(v bool) *ChatBIPatternIndexDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternIndexDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
