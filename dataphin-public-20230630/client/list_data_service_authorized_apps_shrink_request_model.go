// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedAppsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceAuthorizedAppsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceAuthorizedAppsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceAuthorizedAppsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceAuthorizedAppsShrinkRequest struct {
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

func (s ListDataServiceAuthorizedAppsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceAuthorizedAppsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceAuthorizedAppsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) SetProjectId(v int32) *ListDataServiceAuthorizedAppsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
