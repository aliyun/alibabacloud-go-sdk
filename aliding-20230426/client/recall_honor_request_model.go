// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *RecallHonorRequestTenantContext) *RecallHonorRequest
	GetTenantContext() *RecallHonorRequestTenantContext
	SetHonorId(v string) *RecallHonorRequest
	GetHonorId() *string
	SetOrgId(v int64) *RecallHonorRequest
	GetOrgId() *int64
	SetUserId(v string) *RecallHonorRequest
	GetUserId() *string
}

type RecallHonorRequest struct {
	TenantContext *RecallHonorRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 21660610
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 345391052
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RecallHonorRequest) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorRequest) GoString() string {
	return s.String()
}

func (s *RecallHonorRequest) GetTenantContext() *RecallHonorRequestTenantContext {
	return s.TenantContext
}

func (s *RecallHonorRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *RecallHonorRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *RecallHonorRequest) GetUserId() *string {
	return s.UserId
}

func (s *RecallHonorRequest) SetTenantContext(v *RecallHonorRequestTenantContext) *RecallHonorRequest {
	s.TenantContext = v
	return s
}

func (s *RecallHonorRequest) SetHonorId(v string) *RecallHonorRequest {
	s.HonorId = &v
	return s
}

func (s *RecallHonorRequest) SetOrgId(v int64) *RecallHonorRequest {
	s.OrgId = &v
	return s
}

func (s *RecallHonorRequest) SetUserId(v string) *RecallHonorRequest {
	s.UserId = &v
	return s
}

func (s *RecallHonorRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecallHonorRequestTenantContext struct {
	// example:
	//
	// 4
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RecallHonorRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorRequestTenantContext) GoString() string {
	return s.String()
}

func (s *RecallHonorRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *RecallHonorRequestTenantContext) SetTenantId(v string) *RecallHonorRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *RecallHonorRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
