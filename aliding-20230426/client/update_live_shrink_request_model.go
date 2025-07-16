// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *UpdateLiveShrinkRequest
	GetCoverUrl() *string
	SetIntroduction(v string) *UpdateLiveShrinkRequest
	GetIntroduction() *string
	SetLiveId(v string) *UpdateLiveShrinkRequest
	GetLiveId() *string
	SetPreEndTime(v int64) *UpdateLiveShrinkRequest
	GetPreEndTime() *int64
	SetPreStartTime(v int64) *UpdateLiveShrinkRequest
	GetPreStartTime() *int64
	SetTenantContextShrink(v string) *UpdateLiveShrinkRequest
	GetTenantContextShrink() *string
	SetTitle(v string) *UpdateLiveShrinkRequest
	GetTitle() *string
}

type UpdateLiveShrinkRequest struct {
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
	// 4d38xxxxx
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
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
	PreStartTime        *int64  `json:"PreStartTime,omitempty" xml:"PreStartTime,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLiveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveShrinkRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *UpdateLiveShrinkRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *UpdateLiveShrinkRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *UpdateLiveShrinkRequest) GetPreEndTime() *int64 {
	return s.PreEndTime
}

func (s *UpdateLiveShrinkRequest) GetPreStartTime() *int64 {
	return s.PreStartTime
}

func (s *UpdateLiveShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateLiveShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateLiveShrinkRequest) SetCoverUrl(v string) *UpdateLiveShrinkRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetIntroduction(v string) *UpdateLiveShrinkRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetLiveId(v string) *UpdateLiveShrinkRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetPreEndTime(v int64) *UpdateLiveShrinkRequest {
	s.PreEndTime = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetPreStartTime(v int64) *UpdateLiveShrinkRequest {
	s.PreStartTime = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetTenantContextShrink(v string) *UpdateLiveShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateLiveShrinkRequest) SetTitle(v string) *UpdateLiveShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
