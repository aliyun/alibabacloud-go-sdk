// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexQueryTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBIPatternIndexQueryTablesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBIPatternIndexQueryTablesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBIPatternIndexQueryTablesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBIPatternIndexQueryTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBIPatternIndexQueryTablesResponseBody
	GetSuccess() *bool
}

type ChatBIPatternIndexQueryTablesResponseBody struct {
	// example:
	//
	// [{"tableName": "pattern_index_1"，"patternTableName": "polar4ai_nl2sql_pattern_1", "tableStatus": 0},
	//
	// {"tableName": "pattern_index_2"，"patternTableName": "polar4ai_nl2sql_pattern_2", "tableStatus": 1},
	//
	// {"tableName": "pattern_index_3"，"patternTableName": "polar4ai_nl2sql_pattern_3, "tableStatus": 2}]
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

func (s ChatBIPatternIndexQueryTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexQueryTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) SetData(v interface{}) *ChatBIPatternIndexQueryTablesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) SetErrCode(v string) *ChatBIPatternIndexQueryTablesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) SetErrMessage(v string) *ChatBIPatternIndexQueryTablesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) SetRequestId(v string) *ChatBIPatternIndexQueryTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) SetSuccess(v bool) *ChatBIPatternIndexQueryTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
