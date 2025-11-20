// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *WearOrgHonorRequestTenantContext) *WearOrgHonorRequest
	GetTenantContext() *WearOrgHonorRequestTenantContext
	SetHonorId(v string) *WearOrgHonorRequest
	GetHonorId() *string
	SetOrgId(v int64) *WearOrgHonorRequest
	GetOrgId() *int64
	SetUserId(v string) *WearOrgHonorRequest
	GetUserId() *string
	SetWear(v bool) *WearOrgHonorRequest
	GetWear() *bool
}

type WearOrgHonorRequest struct {
	TenantContext *WearOrgHonorRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 21659595
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Wear *bool `json:"wear,omitempty" xml:"wear,omitempty"`
}

func (s WearOrgHonorRequest) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorRequest) GoString() string {
	return s.String()
}

func (s *WearOrgHonorRequest) GetTenantContext() *WearOrgHonorRequestTenantContext {
	return s.TenantContext
}

func (s *WearOrgHonorRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *WearOrgHonorRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *WearOrgHonorRequest) GetUserId() *string {
	return s.UserId
}

func (s *WearOrgHonorRequest) GetWear() *bool {
	return s.Wear
}

func (s *WearOrgHonorRequest) SetTenantContext(v *WearOrgHonorRequestTenantContext) *WearOrgHonorRequest {
	s.TenantContext = v
	return s
}

func (s *WearOrgHonorRequest) SetHonorId(v string) *WearOrgHonorRequest {
	s.HonorId = &v
	return s
}

func (s *WearOrgHonorRequest) SetOrgId(v int64) *WearOrgHonorRequest {
	s.OrgId = &v
	return s
}

func (s *WearOrgHonorRequest) SetUserId(v string) *WearOrgHonorRequest {
	s.UserId = &v
	return s
}

func (s *WearOrgHonorRequest) SetWear(v bool) *WearOrgHonorRequest {
	s.Wear = &v
	return s
}

func (s *WearOrgHonorRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WearOrgHonorRequestTenantContext struct {
	// example:
	//
	// 306752103647458
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s WearOrgHonorRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorRequestTenantContext) GoString() string {
	return s.String()
}

func (s *WearOrgHonorRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *WearOrgHonorRequestTenantContext) SetTenantId(v string) *WearOrgHonorRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *WearOrgHonorRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
