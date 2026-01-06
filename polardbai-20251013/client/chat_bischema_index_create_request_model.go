// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBISchemaIndexCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBISchemaIndexCreateRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBISchemaIndexCreateRequest
	GetAuthType() *string
	SetColumnsExcluded(v string) *ChatBISchemaIndexCreateRequest
	GetColumnsExcluded() *string
	SetDbName(v string) *ChatBISchemaIndexCreateRequest
	GetDbName() *string
	SetInstanceName(v string) *ChatBISchemaIndexCreateRequest
	GetInstanceName() *string
	SetTableNameSuffix(v string) *ChatBISchemaIndexCreateRequest
	GetTableNameSuffix() *string
	SetTablesIncluded(v string) *ChatBISchemaIndexCreateRequest
	GetTablesIncluded() *string
	SetToSample(v int32) *ChatBISchemaIndexCreateRequest
	GetToSample() *int32
}

type ChatBISchemaIndexCreateRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// 空字符串, \"graph_info.time,text_info.ext\"
	ColumnsExcluded *string `json:"ColumnsExcluded,omitempty" xml:"ColumnsExcluded,omitempty"`
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
	// example:
	//
	// 空字符串
	TableNameSuffix *string `json:"TableNameSuffix,omitempty" xml:"TableNameSuffix,omitempty"`
	// example:
	//
	// 空字符串, \"graph_info,image_info,text_info\"
	TablesIncluded *string `json:"TablesIncluded,omitempty" xml:"TablesIncluded,omitempty"`
	// example:
	//
	// 0,1
	ToSample *int32 `json:"ToSample,omitempty" xml:"ToSample,omitempty"`
}

func (s ChatBISchemaIndexCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBISchemaIndexCreateRequest) GoString() string {
	return s.String()
}

func (s *ChatBISchemaIndexCreateRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBISchemaIndexCreateRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBISchemaIndexCreateRequest) GetColumnsExcluded() *string {
	return s.ColumnsExcluded
}

func (s *ChatBISchemaIndexCreateRequest) GetDbName() *string {
	return s.DbName
}

func (s *ChatBISchemaIndexCreateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBISchemaIndexCreateRequest) GetTableNameSuffix() *string {
	return s.TableNameSuffix
}

func (s *ChatBISchemaIndexCreateRequest) GetTablesIncluded() *string {
	return s.TablesIncluded
}

func (s *ChatBISchemaIndexCreateRequest) GetToSample() *int32 {
	return s.ToSample
}

func (s *ChatBISchemaIndexCreateRequest) SetAuthMessage(v string) *ChatBISchemaIndexCreateRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetAuthType(v string) *ChatBISchemaIndexCreateRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetColumnsExcluded(v string) *ChatBISchemaIndexCreateRequest {
	s.ColumnsExcluded = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetDbName(v string) *ChatBISchemaIndexCreateRequest {
	s.DbName = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetInstanceName(v string) *ChatBISchemaIndexCreateRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetTableNameSuffix(v string) *ChatBISchemaIndexCreateRequest {
	s.TableNameSuffix = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetTablesIncluded(v string) *ChatBISchemaIndexCreateRequest {
	s.TablesIncluded = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) SetToSample(v int32) *ChatBISchemaIndexCreateRequest {
	s.ToSample = &v
	return s
}

func (s *ChatBISchemaIndexCreateRequest) Validate() error {
	return dara.Validate(s)
}
