// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternUpdateEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternUpdateEntryResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternUpdateEntryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternUpdateEntryResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternUpdateEntryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternUpdateEntryResponseBody
	GetSuccess() *bool
}

type ChatBIPatternUpdateEntryResponseBody struct {
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

func (s ChatBIPatternUpdateEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternUpdateEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternUpdateEntryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternUpdateEntryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternUpdateEntryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternUpdateEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternUpdateEntryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternUpdateEntryResponseBody) SetData(v interface{}) *ChatBIPatternUpdateEntryResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternUpdateEntryResponseBody) SetErrCode(v string) *ChatBIPatternUpdateEntryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternUpdateEntryResponseBody) SetErrMessage(v string) *ChatBIPatternUpdateEntryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternUpdateEntryResponseBody) SetRequestId(v string) *ChatBIPatternUpdateEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternUpdateEntryResponseBody) SetSuccess(v bool) *ChatBIPatternUpdateEntryResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternUpdateEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
