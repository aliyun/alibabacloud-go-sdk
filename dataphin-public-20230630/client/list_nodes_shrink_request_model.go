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
}

type ListNodesShrinkRequest struct {
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
