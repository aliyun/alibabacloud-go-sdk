// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDetailSolutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListUserDetailSolutionsRequest
	GetBizType() *string
	SetIntentionBizId(v string) *ListUserDetailSolutionsRequest
	GetIntentionBizId() *string
	SetPageNum(v int32) *ListUserDetailSolutionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserDetailSolutionsRequest
	GetPageSize() *int32
}

type ListUserDetailSolutionsRequest struct {
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// I20211222161651000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUserDetailSolutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserDetailSolutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserDetailSolutionsRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserDetailSolutionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserDetailSolutionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserDetailSolutionsRequest) SetBizType(v string) *ListUserDetailSolutionsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetIntentionBizId(v string) *ListUserDetailSolutionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetPageNum(v int32) *ListUserDetailSolutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) SetPageSize(v int32) *ListUserDetailSolutionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserDetailSolutionsRequest) Validate() error {
	return dara.Validate(s)
}
