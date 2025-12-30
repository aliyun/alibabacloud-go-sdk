// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexQueryTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ChatBISchemaIndexQueryTablesResponseBody
	GetData() interface{}
	SetErrCode(v string) *ChatBISchemaIndexQueryTablesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ChatBISchemaIndexQueryTablesResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ChatBISchemaIndexQueryTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatBISchemaIndexQueryTablesResponseBody
	GetSuccess() *bool
}

type ChatBISchemaIndexQueryTablesResponseBody struct {
	// example:
	//
	// [{"tableName": "schema_index_1"，"tablesIncluded": "", "columnsExcluded": "", "toSample": 0, "tableStatus": 0},
	//
	// {"tableName": "schema_index_2"，"tablesIncluded": "", "columnsExcluded": "", "toSample": 0, "tableStatus": 1},
	//
	// {"tableName": "schema_index_3"，"tablesIncluded": "", "columnsExcluded": "", "toSample": 0, "tableStatus": 2}]
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// FC388CBF-F12C-57E8-832F-61A18131B106
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatBISchemaIndexQueryTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexQueryTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) SetData(v interface{}) *ChatBISchemaIndexQueryTablesResponseBody {
	s.Data = v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) SetErrCode(v string) *ChatBISchemaIndexQueryTablesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) SetErrMessage(v string) *ChatBISchemaIndexQueryTablesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) SetRequestId(v string) *ChatBISchemaIndexQueryTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) SetSuccess(v bool) *ChatBISchemaIndexQueryTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
