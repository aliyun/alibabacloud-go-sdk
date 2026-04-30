// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableAssetKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *SearchTableAssetKnowledgeRequest
	GetDbId() *int32
	SetOffset(v int32) *SearchTableAssetKnowledgeRequest
	GetOffset() *int32
	SetSearchKey(v string) *SearchTableAssetKnowledgeRequest
	GetSearchKey() *string
	SetShowType(v string) *SearchTableAssetKnowledgeRequest
	GetShowType() *string
	SetSize(v int32) *SearchTableAssetKnowledgeRequest
	GetSize() *int32
	SetTableName(v string) *SearchTableAssetKnowledgeRequest
	GetTableName() *string
}

type SearchTableAssetKnowledgeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1***
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// 订单
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// example:
	//
	// TABLE
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order_info
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SearchTableAssetKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTableAssetKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *SearchTableAssetKnowledgeRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *SearchTableAssetKnowledgeRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *SearchTableAssetKnowledgeRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchTableAssetKnowledgeRequest) GetShowType() *string {
	return s.ShowType
}

func (s *SearchTableAssetKnowledgeRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchTableAssetKnowledgeRequest) GetTableName() *string {
	return s.TableName
}

func (s *SearchTableAssetKnowledgeRequest) SetDbId(v int32) *SearchTableAssetKnowledgeRequest {
	s.DbId = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) SetOffset(v int32) *SearchTableAssetKnowledgeRequest {
	s.Offset = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) SetSearchKey(v string) *SearchTableAssetKnowledgeRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) SetShowType(v string) *SearchTableAssetKnowledgeRequest {
	s.ShowType = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) SetSize(v int32) *SearchTableAssetKnowledgeRequest {
	s.Size = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) SetTableName(v string) *SearchTableAssetKnowledgeRequest {
	s.TableName = &v
	return s
}

func (s *SearchTableAssetKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
