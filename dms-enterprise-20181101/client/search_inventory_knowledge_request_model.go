// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *SearchInventoryKnowledgeRequest
	GetJobId() *int64
	SetOffset(v int32) *SearchInventoryKnowledgeRequest
	GetOffset() *int32
	SetQuery(v string) *SearchInventoryKnowledgeRequest
	GetQuery() *string
	SetShowType(v string) *SearchInventoryKnowledgeRequest
	GetShowType() *string
	SetSize(v int32) *SearchInventoryKnowledgeRequest
	GetSize() *int32
	SetSortBy(v string) *SearchInventoryKnowledgeRequest
	GetSortBy() *string
	SetSortOrder(v string) *SearchInventoryKnowledgeRequest
	GetSortOrder() *string
}

type SearchInventoryKnowledgeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// 订单
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// TABLE
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// confidence
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s SearchInventoryKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *SearchInventoryKnowledgeRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *SearchInventoryKnowledgeRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *SearchInventoryKnowledgeRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchInventoryKnowledgeRequest) GetShowType() *string {
	return s.ShowType
}

func (s *SearchInventoryKnowledgeRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchInventoryKnowledgeRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchInventoryKnowledgeRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *SearchInventoryKnowledgeRequest) SetJobId(v int64) *SearchInventoryKnowledgeRequest {
	s.JobId = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetOffset(v int32) *SearchInventoryKnowledgeRequest {
	s.Offset = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetQuery(v string) *SearchInventoryKnowledgeRequest {
	s.Query = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetShowType(v string) *SearchInventoryKnowledgeRequest {
	s.ShowType = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetSize(v int32) *SearchInventoryKnowledgeRequest {
	s.Size = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetSortBy(v string) *SearchInventoryKnowledgeRequest {
	s.SortBy = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) SetSortOrder(v string) *SearchInventoryKnowledgeRequest {
	s.SortOrder = &v
	return s
}

func (s *SearchInventoryKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
