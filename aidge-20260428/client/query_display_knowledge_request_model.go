// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDisplayKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryDisplayKnowledgeRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryDisplayKnowledgeRequest
	GetPageSize() *int32
	SetRagId(v string) *QueryDisplayKnowledgeRequest
	GetRagId() *string
	SetSkuId(v string) *QueryDisplayKnowledgeRequest
	GetSkuId() *string
	SetSkuName(v string) *QueryDisplayKnowledgeRequest
	GetSkuName() *string
	SetStatuses(v []*string) *QueryDisplayKnowledgeRequest
	GetStatuses() []*string
}

type QueryDisplayKnowledgeRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rag_xxx
	RagId *string `json:"RagId,omitempty" xml:"RagId,omitempty"`
	// The SKU ID for exact match.
	//
	// example:
	//
	// sku-001
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// The keyword in the SKU name.
	//
	// example:
	//
	// Product front image
	SkuName *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	// The list of record statuses. Valid values: UPDATING, EMBEDDED, and FAILED.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s QueryDisplayKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDisplayKnowledgeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDisplayKnowledgeRequest) GetRagId() *string {
	return s.RagId
}

func (s *QueryDisplayKnowledgeRequest) GetSkuId() *string {
	return s.SkuId
}

func (s *QueryDisplayKnowledgeRequest) GetSkuName() *string {
	return s.SkuName
}

func (s *QueryDisplayKnowledgeRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *QueryDisplayKnowledgeRequest) SetPageNum(v int32) *QueryDisplayKnowledgeRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDisplayKnowledgeRequest) SetPageSize(v int32) *QueryDisplayKnowledgeRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDisplayKnowledgeRequest) SetRagId(v string) *QueryDisplayKnowledgeRequest {
	s.RagId = &v
	return s
}

func (s *QueryDisplayKnowledgeRequest) SetSkuId(v string) *QueryDisplayKnowledgeRequest {
	s.SkuId = &v
	return s
}

func (s *QueryDisplayKnowledgeRequest) SetSkuName(v string) *QueryDisplayKnowledgeRequest {
	s.SkuName = &v
	return s
}

func (s *QueryDisplayKnowledgeRequest) SetStatuses(v []*string) *QueryDisplayKnowledgeRequest {
	s.Statuses = v
	return s
}

func (s *QueryDisplayKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
