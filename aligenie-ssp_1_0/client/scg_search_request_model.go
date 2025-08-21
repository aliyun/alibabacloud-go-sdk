// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScgSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScgFilter(v *ScgSearchRequestScgFilter) *ScgSearchRequest
	GetScgFilter() *ScgSearchRequestScgFilter
	SetTopicId(v string) *ScgSearchRequest
	GetTopicId() *string
}

type ScgSearchRequest struct {
	// This parameter is required.
	ScgFilter *ScgSearchRequestScgFilter `json:"ScgFilter,omitempty" xml:"ScgFilter,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// MC201132
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s ScgSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchRequest) GoString() string {
	return s.String()
}

func (s *ScgSearchRequest) GetScgFilter() *ScgSearchRequestScgFilter {
	return s.ScgFilter
}

func (s *ScgSearchRequest) GetTopicId() *string {
	return s.TopicId
}

func (s *ScgSearchRequest) SetScgFilter(v *ScgSearchRequestScgFilter) *ScgSearchRequest {
	s.ScgFilter = v
	return s
}

func (s *ScgSearchRequest) SetTopicId(v string) *ScgSearchRequest {
	s.TopicId = &v
	return s
}

func (s *ScgSearchRequest) Validate() error {
	return dara.Validate(s)
}

type ScgSearchRequestScgFilter struct {
	OffSetParam *ScgSearchRequestScgFilterOffSetParam `json:"OffSetParam,omitempty" xml:"OffSetParam,omitempty" type:"Struct"`
	PageParam   *ScgSearchRequestScgFilterPageParam   `json:"PageParam,omitempty" xml:"PageParam,omitempty" type:"Struct"`
	// This parameter is required.
	SortParam *ScgSearchRequestScgFilterSortParam `json:"SortParam,omitempty" xml:"SortParam,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// true
	UseOffSet *bool `json:"UseOffSet,omitempty" xml:"UseOffSet,omitempty"`
}

func (s ScgSearchRequestScgFilter) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchRequestScgFilter) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilter) GetOffSetParam() *ScgSearchRequestScgFilterOffSetParam {
	return s.OffSetParam
}

func (s *ScgSearchRequestScgFilter) GetPageParam() *ScgSearchRequestScgFilterPageParam {
	return s.PageParam
}

func (s *ScgSearchRequestScgFilter) GetSortParam() *ScgSearchRequestScgFilterSortParam {
	return s.SortParam
}

func (s *ScgSearchRequestScgFilter) GetUseOffSet() *bool {
	return s.UseOffSet
}

func (s *ScgSearchRequestScgFilter) SetOffSetParam(v *ScgSearchRequestScgFilterOffSetParam) *ScgSearchRequestScgFilter {
	s.OffSetParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetPageParam(v *ScgSearchRequestScgFilterPageParam) *ScgSearchRequestScgFilter {
	s.PageParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetSortParam(v *ScgSearchRequestScgFilterSortParam) *ScgSearchRequestScgFilter {
	s.SortParam = v
	return s
}

func (s *ScgSearchRequestScgFilter) SetUseOffSet(v bool) *ScgSearchRequestScgFilter {
	s.UseOffSet = &v
	return s
}

func (s *ScgSearchRequestScgFilter) Validate() error {
	return dara.Validate(s)
}

type ScgSearchRequestScgFilterOffSetParam struct {
	// example:
	//
	// 20
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 10
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s ScgSearchRequestScgFilterOffSetParam) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchRequestScgFilterOffSetParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterOffSetParam) GetLimit() *int32 {
	return s.Limit
}

func (s *ScgSearchRequestScgFilterOffSetParam) GetOffset() *int32 {
	return s.Offset
}

func (s *ScgSearchRequestScgFilterOffSetParam) SetLimit(v int32) *ScgSearchRequestScgFilterOffSetParam {
	s.Limit = &v
	return s
}

func (s *ScgSearchRequestScgFilterOffSetParam) SetOffset(v int32) *ScgSearchRequestScgFilterOffSetParam {
	s.Offset = &v
	return s
}

func (s *ScgSearchRequestScgFilterOffSetParam) Validate() error {
	return dara.Validate(s)
}

type ScgSearchRequestScgFilterPageParam struct {
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ScgSearchRequestScgFilterPageParam) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchRequestScgFilterPageParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterPageParam) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ScgSearchRequestScgFilterPageParam) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ScgSearchRequestScgFilterPageParam) SetPageNum(v int32) *ScgSearchRequestScgFilterPageParam {
	s.PageNum = &v
	return s
}

func (s *ScgSearchRequestScgFilterPageParam) SetPageSize(v int32) *ScgSearchRequestScgFilterPageParam {
	s.PageSize = &v
	return s
}

func (s *ScgSearchRequestScgFilterPageParam) Validate() error {
	return dara.Validate(s)
}

type ScgSearchRequestScgFilterSortParam struct {
	// example:
	//
	// internal_id
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// example:
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// “”
	SortText *string `json:"SortText,omitempty" xml:"SortText,omitempty"`
}

func (s ScgSearchRequestScgFilterSortParam) String() string {
	return dara.Prettify(s)
}

func (s ScgSearchRequestScgFilterSortParam) GoString() string {
	return s.String()
}

func (s *ScgSearchRequestScgFilterSortParam) GetSortKey() *string {
	return s.SortKey
}

func (s *ScgSearchRequestScgFilterSortParam) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ScgSearchRequestScgFilterSortParam) GetSortText() *string {
	return s.SortText
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortKey(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortKey = &v
	return s
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortOrder(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortOrder = &v
	return s
}

func (s *ScgSearchRequestScgFilterSortParam) SetSortText(v string) *ScgSearchRequestScgFilterSortParam {
	s.SortText = &v
	return s
}

func (s *ScgSearchRequestScgFilterSortParam) Validate() error {
	return dara.Validate(s)
}
