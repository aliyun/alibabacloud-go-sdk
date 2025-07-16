// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *AddRecordPermissionShrinkRequest
	GetBizType() *string
	SetConferenceId(v string) *AddRecordPermissionShrinkRequest
	GetConferenceId() *string
	SetTenantContextShrink(v string) *AddRecordPermissionShrinkRequest
	GetTenantContextShrink() *string
	SetUserId(v string) *AddRecordPermissionShrinkRequest
	GetUserId() *string
}

type AddRecordPermissionShrinkRequest struct {
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
	ConferenceId        *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddRecordPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecordPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddRecordPermissionShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *AddRecordPermissionShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *AddRecordPermissionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddRecordPermissionShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddRecordPermissionShrinkRequest) SetBizType(v string) *AddRecordPermissionShrinkRequest {
	s.BizType = &v
	return s
}

func (s *AddRecordPermissionShrinkRequest) SetConferenceId(v string) *AddRecordPermissionShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *AddRecordPermissionShrinkRequest) SetTenantContextShrink(v string) *AddRecordPermissionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddRecordPermissionShrinkRequest) SetUserId(v string) *AddRecordPermissionShrinkRequest {
	s.UserId = &v
	return s
}

func (s *AddRecordPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
