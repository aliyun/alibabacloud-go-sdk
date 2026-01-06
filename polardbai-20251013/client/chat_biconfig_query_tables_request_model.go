// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIConfigQueryTablesRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIConfigQueryTablesRequest
	GetAuthType() *string
	SetDbName(v string) *ChatBIConfigQueryTablesRequest
	GetDbName() *string
	SetInputField(v string) *ChatBIConfigQueryTablesRequest
	GetInputField() *string
	SetInstanceName(v string) *ChatBIConfigQueryTablesRequest
	GetInstanceName() *string
}

type ChatBIConfigQueryTablesRequest struct {
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
	// polar
	InputField *string `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ChatBIConfigQueryTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIConfigQueryTablesRequest) GoString() string {
	return s.String()
}

func (s *ChatBIConfigQueryTablesRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIConfigQueryTablesRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIConfigQueryTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigQueryTablesRequest) GetInputField() *string {
	return s.InputField
}

func (s *ChatBIConfigQueryTablesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIConfigQueryTablesRequest) SetAuthMessage(v string) *ChatBIConfigQueryTablesRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIConfigQueryTablesRequest) SetAuthType(v string) *ChatBIConfigQueryTablesRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIConfigQueryTablesRequest) SetDbName(v string) *ChatBIConfigQueryTablesRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIConfigQueryTablesRequest) SetInputField(v string) *ChatBIConfigQueryTablesRequest {
	s.InputField = &v
	return s
}

func (s *ChatBIConfigQueryTablesRequest) SetInstanceName(v string) *ChatBIConfigQueryTablesRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIConfigQueryTablesRequest) Validate() error {
	return dara.Validate(s)
}
