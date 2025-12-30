// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ChatBISchemaIndexDeleteRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBISchemaIndexDeleteRequest
	GetInstanceName() *string
	SetTableName(v string) *ChatBISchemaIndexDeleteRequest
	GetTableName() *string
}

type ChatBISchemaIndexDeleteRequest struct {
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
	// pc-2ze2ak7avn731y760
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// schema_index, schema_index_1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ChatBISchemaIndexDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexDeleteRequest) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexDeleteRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBISchemaIndexDeleteRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBISchemaIndexDeleteRequest) GetTableName() *string {
	return s.TableName
}

func (s *ChatBISchemaIndexDeleteRequest) SetDbName(v string) *ChatBISchemaIndexDeleteRequest {
	s.DbName = &v
	return s
}

func (s *ChatBISchemaIndexDeleteRequest) SetInstanceName(v string) *ChatBISchemaIndexDeleteRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBISchemaIndexDeleteRequest) SetTableName(v string) *ChatBISchemaIndexDeleteRequest {
	s.TableName = &v
	return s
}

func (s *ChatBISchemaIndexDeleteRequest) Validate() error {
	return dara.Validate(s)
}
