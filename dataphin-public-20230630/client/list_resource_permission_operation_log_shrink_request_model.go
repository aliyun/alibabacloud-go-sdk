// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcePermissionOperationLogShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListResourcePermissionOperationLogShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListResourcePermissionOperationLogShrinkRequest
	GetOpTenantId() *int64
}

type ListResourcePermissionOperationLogShrinkRequest struct {
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListResourcePermissionOperationLogShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcePermissionOperationLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourcePermissionOperationLogShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListResourcePermissionOperationLogShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListResourcePermissionOperationLogShrinkRequest) SetListQueryShrink(v string) *ListResourcePermissionOperationLogShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListResourcePermissionOperationLogShrinkRequest) SetOpTenantId(v int64) *ListResourcePermissionOperationLogShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListResourcePermissionOperationLogShrinkRequest) Validate() error {
	return dara.Validate(s)
}
