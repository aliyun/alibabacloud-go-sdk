// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIUpdateTableValidationColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBIUpdateTableValidationColumnsRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBIUpdateTableValidationColumnsRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBIUpdateTableValidationColumnsRequest
	GetTableName() *string
	SetTableType(v string) *ChatBIUpdateTableValidationColumnsRequest
	GetTableType() *string
}

type ChatBIUpdateTableValidationColumnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// polar4ai_nl2sql_pattern
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pattern/config
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ChatBIUpdateTableValidationColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIUpdateTableValidationColumnsRequest) GoString() string {
	return s.String()
}

func (s *ChatBIUpdateTableValidationColumnsRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBIUpdateTableValidationColumnsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIUpdateTableValidationColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBIUpdateTableValidationColumnsRequest) GetTableType() *string {
	return s.TableType
}

func (s *ChatBIUpdateTableValidationColumnsRequest) SetDbName(v string) *ChatBIUpdateTableValidationColumnsRequest {
	s.DbName = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsRequest) SetInstanceName(v string) *ChatBIUpdateTableValidationColumnsRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsRequest) SetTableName(v string) *ChatBIUpdateTableValidationColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsRequest) SetTableType(v string) *ChatBIUpdateTableValidationColumnsRequest {
	s.TableType = &v
	return s
}

func (s *ChatBIUpdateTableValidationColumnsRequest) Validate() error {
	return dara.Validate(s)
}
