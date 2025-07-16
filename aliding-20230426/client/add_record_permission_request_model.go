// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *AddRecordPermissionRequest
	GetBizType() *string
	SetConferenceId(v string) *AddRecordPermissionRequest
	GetConferenceId() *string
	SetTenantContext(v *AddRecordPermissionRequestTenantContext) *AddRecordPermissionRequest
	GetTenantContext() *AddRecordPermissionRequestTenantContext
	SetUserId(v string) *AddRecordPermissionRequest
	GetUserId() *string
}

type AddRecordPermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// minutes
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1dddwrqrq
	ConferenceId  *string                                  `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	TenantContext *AddRecordPermissionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddRecordPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionRequest) GetBizType() *string {
	return s.BizType
}

func (s *AddRecordPermissionRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *AddRecordPermissionRequest) GetTenantContext() *AddRecordPermissionRequestTenantContext {
	return s.TenantContext
}

func (s *AddRecordPermissionRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddRecordPermissionRequest) SetBizType(v string) *AddRecordPermissionRequest {
	s.BizType = &v
	return s
}

func (s *AddRecordPermissionRequest) SetConferenceId(v string) *AddRecordPermissionRequest {
	s.ConferenceId = &v
	return s
}

func (s *AddRecordPermissionRequest) SetTenantContext(v *AddRecordPermissionRequestTenantContext) *AddRecordPermissionRequest {
	s.TenantContext = v
	return s
}

func (s *AddRecordPermissionRequest) SetUserId(v string) *AddRecordPermissionRequest {
	s.UserId = &v
	return s
}

func (s *AddRecordPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type AddRecordPermissionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddRecordPermissionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddRecordPermissionRequestTenantContext) SetTenantId(v string) *AddRecordPermissionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddRecordPermissionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
