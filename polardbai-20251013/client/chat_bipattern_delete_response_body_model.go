// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternDeleteResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternDeleteResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternDeleteResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternDeleteResponseBody
	GetSuccess() *bool
}

type ChatBIPatternDeleteResponseBody struct {
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

func (s ChatBIPatternDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternDeleteResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternDeleteResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternDeleteResponseBody) SetData(v interface{}) *ChatBIPatternDeleteResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternDeleteResponseBody) SetErrCode(v string) *ChatBIPatternDeleteResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternDeleteResponseBody) SetErrMessage(v string) *ChatBIPatternDeleteResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternDeleteResponseBody) SetRequestId(v string) *ChatBIPatternDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternDeleteResponseBody) SetSuccess(v bool) *ChatBIPatternDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
