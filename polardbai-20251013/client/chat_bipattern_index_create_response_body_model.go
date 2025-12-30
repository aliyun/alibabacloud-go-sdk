// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternIndexCreateResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternIndexCreateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternIndexCreateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternIndexCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternIndexCreateResponseBody
	GetSuccess() *bool
}

type ChatBIPatternIndexCreateResponseBody struct {
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

func (s ChatBIPatternIndexCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexCreateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternIndexCreateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternIndexCreateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternIndexCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternIndexCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternIndexCreateResponseBody) SetData(v interface{}) *ChatBIPatternIndexCreateResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternIndexCreateResponseBody) SetErrCode(v string) *ChatBIPatternIndexCreateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternIndexCreateResponseBody) SetErrMessage(v string) *ChatBIPatternIndexCreateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternIndexCreateResponseBody) SetRequestId(v string) *ChatBIPatternIndexCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternIndexCreateResponseBody) SetSuccess(v bool) *ChatBIPatternIndexCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternIndexCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
