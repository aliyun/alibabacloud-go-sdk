// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServiceApiCallsShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServiceApiCallsShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceApiCallsShrinkRequest
	GetProjectId() *int32
}

type ListDataServiceApiCallsShrinkRequest struct {
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

func (s ListDataServiceApiCallsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallsShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServiceApiCallsShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiCallsShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallsShrinkRequest) SetListQueryShrink(v string) *ListDataServiceApiCallsShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServiceApiCallsShrinkRequest) SetOpTenantId(v int64) *ListDataServiceApiCallsShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiCallsShrinkRequest) SetProjectId(v int32) *ListDataServiceApiCallsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
