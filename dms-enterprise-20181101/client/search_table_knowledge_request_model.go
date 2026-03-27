// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v string) *SearchTableKnowledgeRequest
	GetDbId() *string
	SetModel(v string) *SearchTableKnowledgeRequest
	GetModel() *string
	SetQuery(v string) *SearchTableKnowledgeRequest
	GetQuery() *string
}

type SearchTableKnowledgeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 查看当前的灰度情况
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s SearchTableKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTableKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *SearchTableKnowledgeRequest) GetDbId() *string {
	return s.DbId
}

func (s *SearchTableKnowledgeRequest) GetModel() *string {
	return s.Model
}

func (s *SearchTableKnowledgeRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchTableKnowledgeRequest) SetDbId(v string) *SearchTableKnowledgeRequest {
	s.DbId = &v
	return s
}

func (s *SearchTableKnowledgeRequest) SetModel(v string) *SearchTableKnowledgeRequest {
	s.Model = &v
	return s
}

func (s *SearchTableKnowledgeRequest) SetQuery(v string) *SearchTableKnowledgeRequest {
	s.Query = &v
	return s
}

func (s *SearchTableKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
