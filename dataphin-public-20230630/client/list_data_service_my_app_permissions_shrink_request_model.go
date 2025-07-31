// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceMyAppPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceMyAppPermissionsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceMyAppPermissionsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceMyAppPermissionsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceMyAppPermissionsShrinkRequest struct {
	// This parameter is required.
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

func (s ListDataServiceMyAppPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceMyAppPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceMyAppPermissionsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceMyAppPermissionsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) SetProjectId(v int32) *ListDataServiceMyAppPermissionsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceMyAppPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
