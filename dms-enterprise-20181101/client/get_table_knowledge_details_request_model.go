// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableKnowledgeDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *GetTableKnowledgeDetailsRequest
	GetDbId() *int32
	SetTableName(v string) *GetTableKnowledgeDetailsRequest
	GetTableName() *string
}

type GetTableKnowledgeDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order_info
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableKnowledgeDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableKnowledgeDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetTableKnowledgeDetailsRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *GetTableKnowledgeDetailsRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTableKnowledgeDetailsRequest) SetDbId(v int32) *GetTableKnowledgeDetailsRequest {
	s.DbId = &v
	return s
}

func (s *GetTableKnowledgeDetailsRequest) SetTableName(v string) *GetTableKnowledgeDetailsRequest {
	s.TableName = &v
	return s
}

func (s *GetTableKnowledgeDetailsRequest) Validate() error {
	return dara.Validate(s)
}
