// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SyncDepartmentShrinkRequest
	GetOpTenantId() *int64
	SetSyncDepartmentCommandShrink(v string) *SyncDepartmentShrinkRequest
	GetSyncDepartmentCommandShrink() *string
}

type SyncDepartmentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SyncDepartmentCommandShrink *string `json:"SyncDepartmentCommand,omitempty" xml:"SyncDepartmentCommand,omitempty"`
}

func (s SyncDepartmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncDepartmentShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SyncDepartmentShrinkRequest) GetSyncDepartmentCommandShrink() *string {
	return s.SyncDepartmentCommandShrink
}

func (s *SyncDepartmentShrinkRequest) SetOpTenantId(v int64) *SyncDepartmentShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *SyncDepartmentShrinkRequest) SetSyncDepartmentCommandShrink(v string) *SyncDepartmentShrinkRequest {
	s.SyncDepartmentCommandShrink = &v
	return s
}

func (s *SyncDepartmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
