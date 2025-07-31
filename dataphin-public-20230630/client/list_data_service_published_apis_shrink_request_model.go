// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListDataServicePublishedApisShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListDataServicePublishedApisShrinkRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServicePublishedApisShrinkRequest
	GetProjectId() *int32
}

type ListDataServicePublishedApisShrinkRequest struct {
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

func (s ListDataServicePublishedApisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListDataServicePublishedApisShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServicePublishedApisShrinkRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServicePublishedApisShrinkRequest) SetListQueryShrink(v string) *ListDataServicePublishedApisShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListDataServicePublishedApisShrinkRequest) SetOpTenantId(v int64) *ListDataServicePublishedApisShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServicePublishedApisShrinkRequest) SetProjectId(v int32) *ListDataServicePublishedApisShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServicePublishedApisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
