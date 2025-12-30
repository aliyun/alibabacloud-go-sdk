// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternDeleteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternDeleteEntryResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternDeleteEntryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternDeleteEntryResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternDeleteEntryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternDeleteEntryResponseBody
	GetSuccess() *bool
}

type ChatBIPatternDeleteEntryResponseBody struct {
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

func (s ChatBIPatternDeleteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternDeleteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternDeleteEntryResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternDeleteEntryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternDeleteEntryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternDeleteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternDeleteEntryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternDeleteEntryResponseBody) SetData(v interface{}) *ChatBIPatternDeleteEntryResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternDeleteEntryResponseBody) SetErrCode(v string) *ChatBIPatternDeleteEntryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternDeleteEntryResponseBody) SetErrMessage(v string) *ChatBIPatternDeleteEntryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternDeleteEntryResponseBody) SetRequestId(v string) *ChatBIPatternDeleteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternDeleteEntryResponseBody) SetSuccess(v bool) *ChatBIPatternDeleteEntryResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternDeleteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
