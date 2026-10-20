// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDisplayKnowledgeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *QueryDisplayKnowledgeShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryDisplayKnowledgeShrinkRequest
	GetPageSize() *int32
	SetRagId(v string) *QueryDisplayKnowledgeShrinkRequest
	GetRagId() *string
	SetSkuId(v string) *QueryDisplayKnowledgeShrinkRequest
	GetSkuId() *string
	SetSkuName(v string) *QueryDisplayKnowledgeShrinkRequest
	GetSkuName() *string
	SetStatusesShrink(v string) *QueryDisplayKnowledgeShrinkRequest
	GetStatusesShrink() *string
}

type QueryDisplayKnowledgeShrinkRequest struct {
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
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s QueryDisplayKnowledgeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDisplayKnowledgeShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetRagId() *string {
	return s.RagId
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetSkuId() *string {
	return s.SkuId
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetSkuName() *string {
	return s.SkuName
}

func (s *QueryDisplayKnowledgeShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetPageNum(v int32) *QueryDisplayKnowledgeShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetPageSize(v int32) *QueryDisplayKnowledgeShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetRagId(v string) *QueryDisplayKnowledgeShrinkRequest {
	s.RagId = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetSkuId(v string) *QueryDisplayKnowledgeShrinkRequest {
	s.SkuId = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetSkuName(v string) *QueryDisplayKnowledgeShrinkRequest {
	s.SkuName = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) SetStatusesShrink(v string) *QueryDisplayKnowledgeShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *QueryDisplayKnowledgeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
