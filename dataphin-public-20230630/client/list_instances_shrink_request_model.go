// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListInstancesShrinkRequest
	GetEnv() *string
	SetListQueryShrink(v string) *ListInstancesShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListInstancesShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListInstancesShrinkRequest
	GetOpUserId() *string
}

type ListInstancesShrinkRequest struct {
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
	// The query request.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator user ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListInstancesShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListInstancesShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListInstancesShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListInstancesShrinkRequest) SetEnv(v string) *ListInstancesShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetListQueryShrink(v string) *ListInstancesShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOpTenantId(v int64) *ListInstancesShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetOpUserId(v string) *ListInstancesShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
