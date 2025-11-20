// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *UpdateLiveRequest
	GetCoverUrl() *string
	SetIntroduction(v string) *UpdateLiveRequest
	GetIntroduction() *string
	SetLiveId(v string) *UpdateLiveRequest
	GetLiveId() *string
	SetPreEndTime(v int64) *UpdateLiveRequest
	GetPreEndTime() *int64
	SetPreStartTime(v int64) *UpdateLiveRequest
	GetPreStartTime() *int64
	SetTenantContext(v *UpdateLiveRequestTenantContext) *UpdateLiveRequest
	GetTenantContext() *UpdateLiveRequestTenantContext
	SetTitle(v string) *UpdateLiveRequest
	GetTitle() *string
}

type UpdateLiveRequest struct {
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
	PreStartTime  *int64                          `json:"PreStartTime,omitempty" xml:"PreStartTime,omitempty"`
	TenantContext *UpdateLiveRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *UpdateLiveRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *UpdateLiveRequest) GetLiveId() *string {
	return s.LiveId
}

func (s *UpdateLiveRequest) GetPreEndTime() *int64 {
	return s.PreEndTime
}

func (s *UpdateLiveRequest) GetPreStartTime() *int64 {
	return s.PreStartTime
}

func (s *UpdateLiveRequest) GetTenantContext() *UpdateLiveRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateLiveRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateLiveRequest) SetCoverUrl(v string) *UpdateLiveRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateLiveRequest) SetIntroduction(v string) *UpdateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetPreEndTime(v int64) *UpdateLiveRequest {
	s.PreEndTime = &v
	return s
}

func (s *UpdateLiveRequest) SetPreStartTime(v int64) *UpdateLiveRequest {
	s.PreStartTime = &v
	return s
}

func (s *UpdateLiveRequest) SetTenantContext(v *UpdateLiveRequestTenantContext) *UpdateLiveRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateLiveRequest) SetTitle(v string) *UpdateLiveRequest {
	s.Title = &v
	return s
}

func (s *UpdateLiveRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLiveRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateLiveRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateLiveRequestTenantContext) SetTenantId(v string) *UpdateLiveRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateLiveRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
