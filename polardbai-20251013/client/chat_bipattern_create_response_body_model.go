// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternCreateResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternCreateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternCreateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternCreateResponseBody
	GetSuccess() *bool
}

type ChatBIPatternCreateResponseBody struct {
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

func (s ChatBIPatternCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternCreateResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternCreateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternCreateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternCreateResponseBody) SetData(v interface{}) *ChatBIPatternCreateResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternCreateResponseBody) SetErrCode(v string) *ChatBIPatternCreateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternCreateResponseBody) SetErrMessage(v string) *ChatBIPatternCreateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternCreateResponseBody) SetRequestId(v string) *ChatBIPatternCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternCreateResponseBody) SetSuccess(v bool) *ChatBIPatternCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
