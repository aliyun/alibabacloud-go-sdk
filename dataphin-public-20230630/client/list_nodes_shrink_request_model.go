// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListNodesShrinkRequest
	GetEnv() *string
	SetListQueryShrink(v string) *ListNodesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListNodesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListNodesShrinkRequest
	GetOpUserId() *string
}

type ListNodesShrinkRequest struct {
	// The environment identifier. Valid values:
	//
	// - DEV: Development environment.
	//
	// - PROD (default): Production environment.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The query conditions.
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

func (s ListNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListNodesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListNodesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListNodesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListNodesShrinkRequest) SetEnv(v string) *ListNodesShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListNodesShrinkRequest) SetListQueryShrink(v string) *ListNodesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOpTenantId(v int64) *ListNodesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetOpUserId(v string) *ListNodesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
