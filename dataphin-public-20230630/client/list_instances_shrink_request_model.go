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
}

type ListInstancesShrinkRequest struct {
	// example:
	//
	// PROD
	Env             *string `json:"Env,omitempty" xml:"Env,omitempty"`
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
