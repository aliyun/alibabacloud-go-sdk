// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetUserIdShrinkRequest
	GetTenantContextShrink() *string
	SetUnionId(v string) *GetUserIdShrinkRequest
	GetUnionId() *string
}

type GetUserIdShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// unionId
	//
	// example:
	//
	// ****iE
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
}

func (s GetUserIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetUserIdShrinkRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *GetUserIdShrinkRequest) SetTenantContextShrink(v string) *GetUserIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetUserIdShrinkRequest) SetUnionId(v string) *GetUserIdShrinkRequest {
	s.UnionId = &v
	return s
}

func (s *GetUserIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
