// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDownStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListNodeDownStreamShrinkRequest
	GetEnv() *string
	SetListQueryShrink(v string) *ListNodeDownStreamShrinkRequest
	GetListQueryShrink() *string
	SetOpTenantId(v int64) *ListNodeDownStreamShrinkRequest
	GetOpTenantId() *int64
}

type ListNodeDownStreamShrinkRequest struct {
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

func (s ListNodeDownStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamShrinkRequest) GetEnv() *string {
	return s.Env
}

func (s *ListNodeDownStreamShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListNodeDownStreamShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListNodeDownStreamShrinkRequest) SetEnv(v string) *ListNodeDownStreamShrinkRequest {
	s.Env = &v
	return s
}

func (s *ListNodeDownStreamShrinkRequest) SetListQueryShrink(v string) *ListNodeDownStreamShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListNodeDownStreamShrinkRequest) SetOpTenantId(v int64) *ListNodeDownStreamShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListNodeDownStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
