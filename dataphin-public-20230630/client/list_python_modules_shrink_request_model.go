// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPythonModulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListPythonModulesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListPythonModulesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListPythonModulesShrinkRequest
	GetOpUserId() *string
}

type ListPythonModulesShrinkRequest struct {
	// The paging conditions for the paged query.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListPythonModulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPythonModulesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListPythonModulesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPythonModulesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListPythonModulesShrinkRequest) SetListQueryShrink(v string) *ListPythonModulesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListPythonModulesShrinkRequest) SetOpTenantId(v int64) *ListPythonModulesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPythonModulesShrinkRequest) SetOpUserId(v string) *ListPythonModulesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListPythonModulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
