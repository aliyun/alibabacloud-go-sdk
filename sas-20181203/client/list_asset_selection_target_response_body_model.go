// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetSelectionTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAssetSelectionTargetResponseBodyData) *ListAssetSelectionTargetResponseBody
	GetData() []*ListAssetSelectionTargetResponseBodyData
	SetPageInfo(v *ListAssetSelectionTargetResponseBodyPageInfo) *ListAssetSelectionTargetResponseBody
	GetPageInfo() *ListAssetSelectionTargetResponseBodyPageInfo
	SetRequestId(v string) *ListAssetSelectionTargetResponseBody
	GetRequestId() *string
}

type ListAssetSelectionTargetResponseBody struct {
	// The data returned.
	Data []*ListAssetSelectionTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListAssetSelectionTargetResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 50A75355-F58F-5D65-8377-98C88DED9C51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssetSelectionTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionTargetResponseBody) GetData() []*ListAssetSelectionTargetResponseBodyData {
	return s.Data
}

func (s *ListAssetSelectionTargetResponseBody) GetPageInfo() *ListAssetSelectionTargetResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListAssetSelectionTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetSelectionTargetResponseBody) SetData(v []*ListAssetSelectionTargetResponseBodyData) *ListAssetSelectionTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListAssetSelectionTargetResponseBody) SetPageInfo(v *ListAssetSelectionTargetResponseBodyPageInfo) *ListAssetSelectionTargetResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListAssetSelectionTargetResponseBody) SetRequestId(v string) *ListAssetSelectionTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAssetSelectionTargetResponseBodyData struct {
	// The ID of the asset.
	//
	// example:
	//
	// 30****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// test****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
}

func (s ListAssetSelectionTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionTargetResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAssetSelectionTargetResponseBodyData) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAssetSelectionTargetResponseBodyData) SetTargetId(v string) *ListAssetSelectionTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBodyData) SetTargetName(v string) *ListAssetSelectionTargetResponseBodyData {
	s.TargetName = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAssetSelectionTargetResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 639
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAssetSelectionTargetResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAssetSelectionTargetResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) SetCurrentPage(v int32) *ListAssetSelectionTargetResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) SetPageSize(v int32) *ListAssetSelectionTargetResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) SetTotalCount(v int32) *ListAssetSelectionTargetResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListAssetSelectionTargetResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
