// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexQueryTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBISchemaIndexQueryTablesRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBISchemaIndexQueryTablesRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBISchemaIndexQueryTablesRequest
	GetDbName() *string
	SetInputField(v string) *ChatBISchemaIndexQueryTablesRequest
	GetInputField() *string
	SetInstanceName(v string) *ChatBISchemaIndexQueryTablesRequest
	GetInstanceName() *string
	SetPageNumber(v string) *ChatBISchemaIndexQueryTablesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ChatBISchemaIndexQueryTablesRequest
	GetPageSize() *string
}

type ChatBISchemaIndexQueryTablesRequest struct {
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
	// 过滤字段，例如 schema_index
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
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ChatBISchemaIndexQueryTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexQueryTablesRequest) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetInputField() *string {
	return s.InputField
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ChatBISchemaIndexQueryTablesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetAuthMessage(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetAuthType(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetDbName(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetInputField(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.InputField = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetInstanceName(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetPageNumber(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) SetPageSize(v string) *ChatBISchemaIndexQueryTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ChatBISchemaIndexQueryTablesRequest) Validate() error {
	return dara.Validate(s)
}
