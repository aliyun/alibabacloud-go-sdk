// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgePointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetKnowledgePointRequest
	GetId() *string
	SetKnowledgeBaseId(v string) *GetKnowledgePointRequest
	GetKnowledgeBaseId() *string
	SetMaxResults(v int32) *GetKnowledgePointRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetKnowledgePointRequest
	GetNextToken() *string
	SetType(v string) *GetKnowledgePointRequest
	GetType() *string
}

type GetKnowledgePointRequest struct {
	// The location ID of the knowledge point. Copy this value exactly as it appears in the search results and use it together with the corresponding KnowledgeBaseId and Type. Do not parse or convert this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// node-sales
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The maximum number of relationships to return per page. This parameter is valid only when Type is set to GRAPH. Default value: 20. Valid values: 1 to 100. This parameter is not supported for WIKI and KNOWLEDGE_CONTENT.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination cursor. Do not specify this parameter when you request the first page. For subsequent pages, use the NextToken value returned in the previous response. An empty response indicates that no more pages are available. If the current page is empty but the cursor is not empty, continue pagination. This parameter is used only when Type is set to GRAPH. You must keep the same knowledge base, entry point, and MaxResults values.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The read category of the knowledge point. Valid values: GRAPH for local graphs, WIKI for topic pages, and KNOWLEDGE_CONTENT for knowledge content.
	//
	// This parameter is required.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgePointRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointRequest) GetId() *string {
	return s.Id
}

func (s *GetKnowledgePointRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgePointRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetKnowledgePointRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetKnowledgePointRequest) GetType() *string {
	return s.Type
}

func (s *GetKnowledgePointRequest) SetId(v string) *GetKnowledgePointRequest {
	s.Id = &v
	return s
}

func (s *GetKnowledgePointRequest) SetKnowledgeBaseId(v string) *GetKnowledgePointRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgePointRequest) SetMaxResults(v int32) *GetKnowledgePointRequest {
	s.MaxResults = &v
	return s
}

func (s *GetKnowledgePointRequest) SetNextToken(v string) *GetKnowledgePointRequest {
	s.NextToken = &v
	return s
}

func (s *GetKnowledgePointRequest) SetType(v string) *GetKnowledgePointRequest {
	s.Type = &v
	return s
}

func (s *GetKnowledgePointRequest) Validate() error {
	return dara.Validate(s)
}
