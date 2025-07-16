// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateSearchKeywordShrinkRequest
	GetContent() *string
	SetEndTime(v int64) *CreateSearchKeywordShrinkRequest
	GetEndTime() *int64
	SetResId(v string) *CreateSearchKeywordShrinkRequest
	GetResId() *string
	SetStartTime(v int64) *CreateSearchKeywordShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *CreateSearchKeywordShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdListShrink(v string) *CreateSearchKeywordShrinkRequest
	GetUserIdListShrink() *string
}

type CreateSearchKeywordShrinkRequest struct {
	// example:
	//
	// {}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1699265024987
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1028
	ResId *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	// example:
	//
	// 1699265024987
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	UserIdListShrink    *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s CreateSearchKeywordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSearchKeywordShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSearchKeywordShrinkRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateSearchKeywordShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSearchKeywordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateSearchKeywordShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *CreateSearchKeywordShrinkRequest) SetContent(v string) *CreateSearchKeywordShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) SetEndTime(v int64) *CreateSearchKeywordShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) SetResId(v string) *CreateSearchKeywordShrinkRequest {
	s.ResId = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) SetStartTime(v int64) *CreateSearchKeywordShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) SetTenantContextShrink(v string) *CreateSearchKeywordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) SetUserIdListShrink(v string) *CreateSearchKeywordShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *CreateSearchKeywordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
