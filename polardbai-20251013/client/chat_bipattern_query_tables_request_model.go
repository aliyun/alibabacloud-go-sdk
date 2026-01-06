// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIPatternQueryTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIPatternQueryTablesRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIPatternQueryTablesRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIPatternQueryTablesRequest
	GetDbName() *string
	SetInputField(v string) *ChatBIPatternQueryTablesRequest
	GetInputField() *string
	SetInstanceName(v string) *ChatBIPatternQueryTablesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ChatBIPatternQueryTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ChatBIPatternQueryTablesRequest
	GetPageSize() *int32
}

type ChatBIPatternQueryTablesRequest struct {
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
	// pattern
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

func (s ChatBIPatternQueryTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIPatternQueryTablesRequest) GoString() string {
	return s.String()
}

func (s *ChatBIPatternQueryTablesRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIPatternQueryTablesRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIPatternQueryTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIPatternQueryTablesRequest) GetInputField() *string {
	return s.InputField
}

func (s *ChatBIPatternQueryTablesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIPatternQueryTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ChatBIPatternQueryTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChatBIPatternQueryTablesRequest) SetAuthMessage(v string) *ChatBIPatternQueryTablesRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetAuthType(v string) *ChatBIPatternQueryTablesRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetDbName(v string) *ChatBIPatternQueryTablesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetInputField(v string) *ChatBIPatternQueryTablesRequest {
	s.InputField = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetInstanceName(v string) *ChatBIPatternQueryTablesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetPageNumber(v int32) *ChatBIPatternQueryTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) SetPageSize(v int32) *ChatBIPatternQueryTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ChatBIPatternQueryTablesRequest) Validate() error {
	return dara.Validate(s)
}
