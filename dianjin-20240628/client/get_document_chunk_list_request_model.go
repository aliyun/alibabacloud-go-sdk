// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentChunkListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIdList(v []*string) *GetDocumentChunkListRequest
	GetChunkIdList() []*string
	SetDocId(v string) *GetDocumentChunkListRequest
	GetDocId() *string
	SetLibraryId(v string) *GetDocumentChunkListRequest
	GetLibraryId() *string
	SetOrder(v string) *GetDocumentChunkListRequest
	GetOrder() *string
	SetOrderBy(v string) *GetDocumentChunkListRequest
	GetOrderBy() *string
	SetPage(v int32) *GetDocumentChunkListRequest
	GetPage() *int32
	SetPageSize(v int32) *GetDocumentChunkListRequest
	GetPageSize() *int32
	SetSearchQuery(v string) *GetDocumentChunkListRequest
	GetSearchQuery() *string
}

type GetDocumentChunkListRequest struct {
	ChunkIdList []*string `json:"chunkIdList,omitempty" xml:"chunkIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 182364872346
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dsjgfdjgfxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
}

func (s GetDocumentChunkListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListRequest) GetChunkIdList() []*string {
	return s.ChunkIdList
}

func (s *GetDocumentChunkListRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetDocumentChunkListRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetDocumentChunkListRequest) GetOrder() *string {
	return s.Order
}

func (s *GetDocumentChunkListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetDocumentChunkListRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetDocumentChunkListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDocumentChunkListRequest) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *GetDocumentChunkListRequest) SetChunkIdList(v []*string) *GetDocumentChunkListRequest {
	s.ChunkIdList = v
	return s
}

func (s *GetDocumentChunkListRequest) SetDocId(v string) *GetDocumentChunkListRequest {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetLibraryId(v string) *GetDocumentChunkListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrder(v string) *GetDocumentChunkListRequest {
	s.Order = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrderBy(v string) *GetDocumentChunkListRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPage(v int32) *GetDocumentChunkListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPageSize(v int32) *GetDocumentChunkListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetSearchQuery(v string) *GetDocumentChunkListRequest {
	s.SearchQuery = &v
	return s
}

func (s *GetDocumentChunkListRequest) Validate() error {
	return dara.Validate(s)
}
