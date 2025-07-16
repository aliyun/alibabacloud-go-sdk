// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *CreateLiveShrinkRequest
	GetCoverUrl() *string
	SetIntroduction(v string) *CreateLiveShrinkRequest
	GetIntroduction() *string
	SetPreEndTime(v int64) *CreateLiveShrinkRequest
	GetPreEndTime() *int64
	SetPreStartTime(v int64) *CreateLiveShrinkRequest
	GetPreStartTime() *int64
	SetPublicType(v int64) *CreateLiveShrinkRequest
	GetPublicType() *int64
	SetTenantContextShrink(v string) *CreateLiveShrinkRequest
	GetTenantContextShrink() *string
	SetTitle(v string) *CreateLiveShrinkRequest
	GetTitle() *string
}

type CreateLiveShrinkRequest struct {
	// example:
	//
	// http://sss/sss
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// example:
	//
	// 这是一个直播
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1698596800000
	PreEndTime *int64 `json:"PreEndTime,omitempty" xml:"PreEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1691596800000
	PreStartTime *int64 `json:"PreStartTime,omitempty" xml:"PreStartTime,omitempty"`
	// example:
	//
	// 0
	PublicType          *int64  `json:"PublicType,omitempty" xml:"PublicType,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateLiveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveShrinkRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *CreateLiveShrinkRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *CreateLiveShrinkRequest) GetPreEndTime() *int64 {
	return s.PreEndTime
}

func (s *CreateLiveShrinkRequest) GetPreStartTime() *int64 {
	return s.PreStartTime
}

func (s *CreateLiveShrinkRequest) GetPublicType() *int64 {
	return s.PublicType
}

func (s *CreateLiveShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateLiveShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateLiveShrinkRequest) SetCoverUrl(v string) *CreateLiveShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetIntroduction(v string) *CreateLiveShrinkRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetPreEndTime(v int64) *CreateLiveShrinkRequest {
	s.PreEndTime = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetPreStartTime(v int64) *CreateLiveShrinkRequest {
	s.PreStartTime = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetPublicType(v int64) *CreateLiveShrinkRequest {
	s.PublicType = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetTenantContextShrink(v string) *CreateLiveShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateLiveShrinkRequest) SetTitle(v string) *CreateLiveShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
