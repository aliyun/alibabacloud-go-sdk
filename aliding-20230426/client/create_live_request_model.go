// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverUrl(v string) *CreateLiveRequest
	GetCoverUrl() *string
	SetIntroduction(v string) *CreateLiveRequest
	GetIntroduction() *string
	SetPreEndTime(v int64) *CreateLiveRequest
	GetPreEndTime() *int64
	SetPreStartTime(v int64) *CreateLiveRequest
	GetPreStartTime() *int64
	SetPublicType(v int64) *CreateLiveRequest
	GetPublicType() *int64
	SetTenantContext(v *CreateLiveRequestTenantContext) *CreateLiveRequest
	GetTenantContext() *CreateLiveRequestTenantContext
	SetTitle(v string) *CreateLiveRequest
	GetTitle() *string
}

type CreateLiveRequest struct {
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
	PublicType    *int64                          `json:"PublicType,omitempty" xml:"PublicType,omitempty"`
	TenantContext *CreateLiveRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *CreateLiveRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *CreateLiveRequest) GetPreEndTime() *int64 {
	return s.PreEndTime
}

func (s *CreateLiveRequest) GetPreStartTime() *int64 {
	return s.PreStartTime
}

func (s *CreateLiveRequest) GetPublicType() *int64 {
	return s.PublicType
}

func (s *CreateLiveRequest) GetTenantContext() *CreateLiveRequestTenantContext {
	return s.TenantContext
}

func (s *CreateLiveRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateLiveRequest) SetCoverUrl(v string) *CreateLiveRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateLiveRequest) SetIntroduction(v string) *CreateLiveRequest {
	s.Introduction = &v
	return s
}

func (s *CreateLiveRequest) SetPreEndTime(v int64) *CreateLiveRequest {
	s.PreEndTime = &v
	return s
}

func (s *CreateLiveRequest) SetPreStartTime(v int64) *CreateLiveRequest {
	s.PreStartTime = &v
	return s
}

func (s *CreateLiveRequest) SetPublicType(v int64) *CreateLiveRequest {
	s.PublicType = &v
	return s
}

func (s *CreateLiveRequest) SetTenantContext(v *CreateLiveRequestTenantContext) *CreateLiveRequest {
	s.TenantContext = v
	return s
}

func (s *CreateLiveRequest) SetTitle(v string) *CreateLiveRequest {
	s.Title = &v
	return s
}

func (s *CreateLiveRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLiveRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateLiveRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateLiveRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateLiveRequestTenantContext) SetTenantId(v string) *CreateLiveRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateLiveRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
