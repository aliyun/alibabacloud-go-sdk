// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSolutionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListUserSolutionsShrinkRequest
	GetBizType() *string
	SetExistStatusShrink(v string) *ListUserSolutionsShrinkRequest
	GetExistStatusShrink() *string
	SetIntentionBizId(v string) *ListUserSolutionsShrinkRequest
	GetIntentionBizId() *string
	SetPageNum(v int32) *ListUserSolutionsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserSolutionsShrinkRequest
	GetPageSize() *int32
}

type ListUserSolutionsShrinkRequest struct {
	BizType           *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExistStatusShrink *string `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty"`
	// example:
	//
	// I20210924151843000001
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

func (s ListUserSolutionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserSolutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserSolutionsShrinkRequest) GetExistStatusShrink() *string {
	return s.ExistStatusShrink
}

func (s *ListUserSolutionsShrinkRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserSolutionsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserSolutionsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSolutionsShrinkRequest) SetBizType(v string) *ListUserSolutionsShrinkRequest {
	s.BizType = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetExistStatusShrink(v string) *ListUserSolutionsShrinkRequest {
	s.ExistStatusShrink = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetIntentionBizId(v string) *ListUserSolutionsShrinkRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetPageNum(v int32) *ListUserSolutionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) SetPageSize(v int32) *ListUserSolutionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserSolutionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
