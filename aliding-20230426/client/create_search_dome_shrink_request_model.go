// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateSearchDomeShrinkRequest
	GetContent() *string
	SetEndTime(v int64) *CreateSearchDomeShrinkRequest
	GetEndTime() *int64
	SetResId(v string) *CreateSearchDomeShrinkRequest
	GetResId() *string
	SetStartTime(v int64) *CreateSearchDomeShrinkRequest
	GetStartTime() *int64
	SetTenantContextShrink(v string) *CreateSearchDomeShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdListShrink(v string) *CreateSearchDomeShrinkRequest
	GetUserIdListShrink() *string
}

type CreateSearchDomeShrinkRequest struct {
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
	// 1030
	ResId *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	// example:
	//
	// 1699265024987
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	UserIdListShrink    *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s CreateSearchDomeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSearchDomeShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSearchDomeShrinkRequest) GetResId() *string {
	return s.ResId
}

func (s *CreateSearchDomeShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSearchDomeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateSearchDomeShrinkRequest) GetUserIdListShrink() *string {
	return s.UserIdListShrink
}

func (s *CreateSearchDomeShrinkRequest) SetContent(v string) *CreateSearchDomeShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) SetEndTime(v int64) *CreateSearchDomeShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) SetResId(v string) *CreateSearchDomeShrinkRequest {
	s.ResId = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) SetStartTime(v int64) *CreateSearchDomeShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) SetTenantContextShrink(v string) *CreateSearchDomeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) SetUserIdListShrink(v string) *CreateSearchDomeShrinkRequest {
	s.UserIdListShrink = &v
	return s
}

func (s *CreateSearchDomeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
