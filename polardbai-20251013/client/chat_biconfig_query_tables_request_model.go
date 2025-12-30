// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIConfigQueryTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIConfigQueryTablesRequest
	GetDbName() *string
	SetInputField(v string) *ChatBIConfigQueryTablesRequest
	GetInputField() *string
	SetInstanceName(v string) *ChatBIConfigQueryTablesRequest
	GetInstanceName() *string
}

type ChatBIConfigQueryTablesRequest struct {
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

func (s *ChatBIConfigQueryTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIConfigQueryTablesRequest) GetInputField() *string {
	return s.InputField
}

func (s *ChatBIConfigQueryTablesRequest) GetInstanceName() *string {
	return s.InstanceName
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
