// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternIndexQueryTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPatternIndexQueryTablesRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPatternIndexQueryTablesRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPatternIndexQueryTablesRequest
	GetDbName() *string
	SetInputField(v string) *ChatBIPatternIndexQueryTablesRequest
	GetInputField() *string
	SetInstanceName(v string) *ChatBIPatternIndexQueryTablesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ChatBIPatternIndexQueryTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ChatBIPatternIndexQueryTablesRequest
	GetPageSize() *int32
}

type ChatBIPatternIndexQueryTablesRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// polar4ai_nl2sql_pattern
	InputField *string `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ChatBIPatternIndexQueryTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternIndexQueryTablesRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetInputField() *string {
	return s.InputField
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ChatBIPatternIndexQueryTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetAuthMessage(v string) *ChatBIPatternIndexQueryTablesRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetAuthType(v string) *ChatBIPatternIndexQueryTablesRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetDbName(v string) *ChatBIPatternIndexQueryTablesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetInputField(v string) *ChatBIPatternIndexQueryTablesRequest {
	s.InputField = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetInstanceName(v string) *ChatBIPatternIndexQueryTablesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetPageNumber(v int32) *ChatBIPatternIndexQueryTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) SetPageSize(v int32) *ChatBIPatternIndexQueryTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ChatBIPatternIndexQueryTablesRequest) Validate() error {
	return dara.Validate(s)
}
