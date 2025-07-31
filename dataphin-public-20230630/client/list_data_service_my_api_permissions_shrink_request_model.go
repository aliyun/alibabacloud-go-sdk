// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyApiPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceMyApiPermissionsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceMyApiPermissionsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceMyApiPermissionsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceMyApiPermissionsShrinkRequest struct {
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceMyApiPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyApiPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceMyApiPermissionsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceMyApiPermissionsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) SetProjectId(v int32) *ListDataServiceMyApiPermissionsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyApiPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
