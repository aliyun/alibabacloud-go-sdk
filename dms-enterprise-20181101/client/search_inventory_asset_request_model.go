// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *SearchInventoryAssetRequest
	GetJobId() *int64
	SetOffset(v int32) *SearchInventoryAssetRequest
	GetOffset() *int32
	SetQuery(v string) *SearchInventoryAssetRequest
	GetQuery() *string
	SetSize(v int32) *SearchInventoryAssetRequest
	GetSize() *int32
	SetSortBy(v string) *SearchInventoryAssetRequest
	GetSortBy() *string
	SetSortOrder(v string) *SearchInventoryAssetRequest
	GetSortOrder() *string
}

type SearchInventoryAssetRequest struct {
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
	// order
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
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

func (s SearchInventoryAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryAssetRequest) GoString() string {
	return s.String()
}

func (s *SearchInventoryAssetRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *SearchInventoryAssetRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *SearchInventoryAssetRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchInventoryAssetRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchInventoryAssetRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchInventoryAssetRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *SearchInventoryAssetRequest) SetJobId(v int64) *SearchInventoryAssetRequest {
	s.JobId = &v
	return s
}

func (s *SearchInventoryAssetRequest) SetOffset(v int32) *SearchInventoryAssetRequest {
	s.Offset = &v
	return s
}

func (s *SearchInventoryAssetRequest) SetQuery(v string) *SearchInventoryAssetRequest {
	s.Query = &v
	return s
}

func (s *SearchInventoryAssetRequest) SetSize(v int32) *SearchInventoryAssetRequest {
	s.Size = &v
	return s
}

func (s *SearchInventoryAssetRequest) SetSortBy(v string) *SearchInventoryAssetRequest {
	s.SortBy = &v
	return s
}

func (s *SearchInventoryAssetRequest) SetSortOrder(v string) *SearchInventoryAssetRequest {
	s.SortOrder = &v
	return s
}

func (s *SearchInventoryAssetRequest) Validate() error {
	return dara.Validate(s)
}
