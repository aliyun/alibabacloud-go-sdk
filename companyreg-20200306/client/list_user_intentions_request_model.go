// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *ListUserIntentionsRequest
	GetArea() *string
	SetBizType(v string) *ListUserIntentionsRequest
	GetBizType() *string
	SetBizTypes(v string) *ListUserIntentionsRequest
	GetBizTypes() *string
	SetIntentionBizId(v string) *ListUserIntentionsRequest
	GetIntentionBizId() *string
	SetPageNum(v int32) *ListUserIntentionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserIntentionsRequest
	GetPageSize() *int32
	SetSortFiled(v string) *ListUserIntentionsRequest
	GetSortFiled() *string
	SetSortOrder(v string) *ListUserIntentionsRequest
	GetSortOrder() *string
	SetStatus(v int32) *ListUserIntentionsRequest
	GetStatus() *int32
	SetWithExtInfo(v bool) *ListUserIntentionsRequest
	GetWithExtInfo() *bool
}

type ListUserIntentionsRequest struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// example:
	//
	// esp.companyreg_cloud
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// esp.bookkeeping,esp.bookkeeping_cloud
	BizTypes *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	// example:
	//
	// I20210917170147000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// UpdateTime
	SortFiled *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// 37
	Status      *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	WithExtInfo *bool  `json:"WithExtInfo,omitempty" xml:"WithExtInfo,omitempty"`
}

func (s ListUserIntentionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsRequest) GetArea() *string {
	return s.Area
}

func (s *ListUserIntentionsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserIntentionsRequest) GetBizTypes() *string {
	return s.BizTypes
}

func (s *ListUserIntentionsRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserIntentionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserIntentionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserIntentionsRequest) GetSortFiled() *string {
	return s.SortFiled
}

func (s *ListUserIntentionsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListUserIntentionsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserIntentionsRequest) GetWithExtInfo() *bool {
	return s.WithExtInfo
}

func (s *ListUserIntentionsRequest) SetArea(v string) *ListUserIntentionsRequest {
	s.Area = &v
	return s
}

func (s *ListUserIntentionsRequest) SetBizType(v string) *ListUserIntentionsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionsRequest) SetBizTypes(v string) *ListUserIntentionsRequest {
	s.BizTypes = &v
	return s
}

func (s *ListUserIntentionsRequest) SetIntentionBizId(v string) *ListUserIntentionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserIntentionsRequest) SetPageNum(v int32) *ListUserIntentionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionsRequest) SetPageSize(v int32) *ListUserIntentionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionsRequest) SetSortFiled(v string) *ListUserIntentionsRequest {
	s.SortFiled = &v
	return s
}

func (s *ListUserIntentionsRequest) SetSortOrder(v string) *ListUserIntentionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserIntentionsRequest) SetStatus(v int32) *ListUserIntentionsRequest {
	s.Status = &v
	return s
}

func (s *ListUserIntentionsRequest) SetWithExtInfo(v bool) *ListUserIntentionsRequest {
	s.WithExtInfo = &v
	return s
}

func (s *ListUserIntentionsRequest) Validate() error {
	return dara.Validate(s)
}
