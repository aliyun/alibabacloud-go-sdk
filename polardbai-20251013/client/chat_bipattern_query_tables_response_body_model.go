// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternQueryTablesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternQueryTablesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternQueryTablesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternQueryTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternQueryTablesResponseBody
	GetSuccess() *bool
}

type ChatBIPatternQueryTablesResponseBody struct {
	// example:
	//
	// ["polar4ai_nl2sql_pattern", "polar4ai_nl2sql_pattern_1"]
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

func (s ChatBIPatternQueryTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryTablesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternQueryTablesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternQueryTablesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternQueryTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternQueryTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternQueryTablesResponseBody) SetData(v interface{}) *ChatBIPatternQueryTablesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternQueryTablesResponseBody) SetErrCode(v string) *ChatBIPatternQueryTablesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternQueryTablesResponseBody) SetErrMessage(v string) *ChatBIPatternQueryTablesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternQueryTablesResponseBody) SetRequestId(v string) *ChatBIPatternQueryTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternQueryTablesResponseBody) SetSuccess(v bool) *ChatBIPatternQueryTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternQueryTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
