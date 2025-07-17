// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *GetTableKnowledgeInfoRequest
	GetDbId() *int32
	SetTableName(v string) *GetTableKnowledgeInfoRequest
	GetTableName() *string
	SetTableSchemaName(v string) *GetTableKnowledgeInfoRequest
	GetTableSchemaName() *string
}

type GetTableKnowledgeInfoRequest struct {
	// This parameter is required.
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	TableName       *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
}

func (s GetTableKnowledgeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeInfoRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetTableKnowledgeInfoRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableKnowledgeInfoRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *GetTableKnowledgeInfoRequest) SetDbId(v int32) *GetTableKnowledgeInfoRequest {
	s.DbId = &v
	return s
}

func (s *GetTableKnowledgeInfoRequest) SetTableName(v string) *GetTableKnowledgeInfoRequest {
	s.TableName = &v
	return s
}

func (s *GetTableKnowledgeInfoRequest) SetTableSchemaName(v string) *GetTableKnowledgeInfoRequest {
	s.TableSchemaName = &v
	return s
}

func (s *GetTableKnowledgeInfoRequest) Validate() error {
	return dara.Validate(s)
}
