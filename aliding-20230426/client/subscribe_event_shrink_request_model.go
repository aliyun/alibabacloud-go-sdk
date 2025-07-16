// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScope(v string) *SubscribeEventShrinkRequest
	GetScope() *string
	SetScopeId(v string) *SubscribeEventShrinkRequest
	GetScopeId() *string
	SetTenantContextShrink(v string) *SubscribeEventShrinkRequest
	GetTenantContextShrink() *string
}

type SubscribeEventShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SPACE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 233456
	ScopeId             *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SubscribeEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubscribeEventShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *SubscribeEventShrinkRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *SubscribeEventShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SubscribeEventShrinkRequest) SetScope(v string) *SubscribeEventShrinkRequest {
	s.Scope = &v
	return s
}

func (s *SubscribeEventShrinkRequest) SetScopeId(v string) *SubscribeEventShrinkRequest {
	s.ScopeId = &v
	return s
}

func (s *SubscribeEventShrinkRequest) SetTenantContextShrink(v string) *SubscribeEventShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SubscribeEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
