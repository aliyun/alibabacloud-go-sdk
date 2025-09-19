// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListAssetSelectionTargetRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListAssetSelectionTargetRequest
	GetPageSize() *int32
	SetSelectionKey(v string) *ListAssetSelectionTargetRequest
	GetSelectionKey() *string
}

type ListAssetSelectionTargetRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The globally unique identifier (GUID) of the asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8ccf9b01-2c64-4cba-8122-10115f29****
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
}

func (s ListAssetSelectionTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionTargetRequest) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionTargetRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAssetSelectionTargetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAssetSelectionTargetRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *ListAssetSelectionTargetRequest) SetCurrentPage(v int32) *ListAssetSelectionTargetRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAssetSelectionTargetRequest) SetPageSize(v int32) *ListAssetSelectionTargetRequest {
	s.PageSize = &v
	return s
}

func (s *ListAssetSelectionTargetRequest) SetSelectionKey(v string) *ListAssetSelectionTargetRequest {
	s.SelectionKey = &v
	return s
}

func (s *ListAssetSelectionTargetRequest) Validate() error {
	return dara.Validate(s)
}
