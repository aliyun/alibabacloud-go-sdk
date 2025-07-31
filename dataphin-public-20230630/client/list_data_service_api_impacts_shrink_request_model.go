// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiImpactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceApiImpactsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceApiImpactsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceApiImpactsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceApiImpactsShrinkRequest struct {
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

func (s ListDataServiceApiImpactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceApiImpactsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiImpactsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiImpactsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceApiImpactsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceApiImpactsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceApiImpactsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiImpactsShrinkRequest) SetProjectId(v int32) *ListDataServiceApiImpactsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiImpactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
