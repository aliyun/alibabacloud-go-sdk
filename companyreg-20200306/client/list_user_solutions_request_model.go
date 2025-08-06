// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSolutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListUserSolutionsRequest
	GetBizType() *string
	SetExistStatus(v []*int64) *ListUserSolutionsRequest
	GetExistStatus() []*int64
	SetIntentionBizId(v string) *ListUserSolutionsRequest
	GetIntentionBizId() *string
	SetPageNum(v int32) *ListUserSolutionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserSolutionsRequest
	GetPageSize() *int32
}

type ListUserSolutionsRequest struct {
	BizType     *string  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExistStatus []*int64 `json:"ExistStatus,omitempty" xml:"ExistStatus,omitempty" type:"Repeated"`
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

func (s ListUserSolutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserSolutionsRequest) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListUserSolutionsRequest) GetExistStatus() []*int64 {
	return s.ExistStatus
}

func (s *ListUserSolutionsRequest) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserSolutionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserSolutionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSolutionsRequest) SetBizType(v string) *ListUserSolutionsRequest {
	s.BizType = &v
	return s
}

func (s *ListUserSolutionsRequest) SetExistStatus(v []*int64) *ListUserSolutionsRequest {
	s.ExistStatus = v
	return s
}

func (s *ListUserSolutionsRequest) SetIntentionBizId(v string) *ListUserSolutionsRequest {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsRequest) SetPageNum(v int32) *ListUserSolutionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListUserSolutionsRequest) SetPageSize(v int32) *ListUserSolutionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserSolutionsRequest) Validate() error {
	return dara.Validate(s)
}
