// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScope(v string) *UnsubscribeEventShrinkRequest
	GetScope() *string
	SetScopeId(v string) *UnsubscribeEventShrinkRequest
	GetScopeId() *string
	SetTenantContextShrink(v string) *UnsubscribeEventShrinkRequest
	GetTenantContextShrink() *string
}

type UnsubscribeEventShrinkRequest struct {
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
	// 23456
	ScopeId             *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UnsubscribeEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *UnsubscribeEventShrinkRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *UnsubscribeEventShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UnsubscribeEventShrinkRequest) SetScope(v string) *UnsubscribeEventShrinkRequest {
	s.Scope = &v
	return s
}

func (s *UnsubscribeEventShrinkRequest) SetScopeId(v string) *UnsubscribeEventShrinkRequest {
	s.ScopeId = &v
	return s
}

func (s *UnsubscribeEventShrinkRequest) SetTenantContextShrink(v string) *UnsubscribeEventShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UnsubscribeEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
