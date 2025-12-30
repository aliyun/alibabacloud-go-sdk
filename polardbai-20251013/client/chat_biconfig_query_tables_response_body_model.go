// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIConfigQueryTablesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIConfigQueryTablesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIConfigQueryTablesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIConfigQueryTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIConfigQueryTablesResponseBody
	GetSuccess() *bool
}

type ChatBIConfigQueryTablesResponseBody struct {
	// example:
	//
	// [], [{"table_name":"polar4ai_nl2sql_llm_config", "skip_validation": 0}]
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

func (s ChatBIConfigQueryTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryTablesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIConfigQueryTablesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIConfigQueryTablesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIConfigQueryTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIConfigQueryTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIConfigQueryTablesResponseBody) SetData(v interface{}) *ChatBIConfigQueryTablesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIConfigQueryTablesResponseBody) SetErrCode(v string) *ChatBIConfigQueryTablesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIConfigQueryTablesResponseBody) SetErrMessage(v string) *ChatBIConfigQueryTablesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIConfigQueryTablesResponseBody) SetRequestId(v string) *ChatBIConfigQueryTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIConfigQueryTablesResponseBody) SetSuccess(v bool) *ChatBIConfigQueryTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIConfigQueryTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
