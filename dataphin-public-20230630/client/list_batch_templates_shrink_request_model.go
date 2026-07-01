// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListBatchTemplatesShrinkRequest
	GetEnv() *string
	SetListQueryShrink(v string) *ListBatchTemplatesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListBatchTemplatesShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *ListBatchTemplatesShrinkRequest
	GetProjectId() *int64
}

type ListBatchTemplatesShrinkRequest struct {
	// The runtime environment. Default value: PROD.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The paged query conditions.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListBatchTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListBatchTemplatesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListBatchTemplatesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBatchTemplatesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBatchTemplatesShrinkRequest) SetEnv(v string) *ListBatchTemplatesShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListBatchTemplatesShrinkRequest) SetListQueryShrink(v string) *ListBatchTemplatesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListBatchTemplatesShrinkRequest) SetOpTenantId(v int64) *ListBatchTemplatesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBatchTemplatesShrinkRequest) SetProjectId(v int64) *ListBatchTemplatesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListBatchTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
