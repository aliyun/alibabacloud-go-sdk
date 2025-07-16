// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScope(v string) *SubscribeEventRequest
	GetScope() *string
	SetScopeId(v string) *SubscribeEventRequest
	GetScopeId() *string
	SetTenantContext(v *SubscribeEventRequestTenantContext) *SubscribeEventRequest
	GetTenantContext() *SubscribeEventRequestTenantContext
}

type SubscribeEventRequest struct {
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
	ScopeId       *string                             `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	TenantContext *SubscribeEventRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SubscribeEventRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *SubscribeEventRequest) GetScope() *string {
	return s.Scope
}

func (s *SubscribeEventRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *SubscribeEventRequest) GetTenantContext() *SubscribeEventRequestTenantContext {
	return s.TenantContext
}

func (s *SubscribeEventRequest) SetScope(v string) *SubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *SubscribeEventRequest) SetScopeId(v string) *SubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *SubscribeEventRequest) SetTenantContext(v *SubscribeEventRequestTenantContext) *SubscribeEventRequest {
	s.TenantContext = v
	return s
}

func (s *SubscribeEventRequest) Validate() error {
	return dara.Validate(s)
}

type SubscribeEventRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SubscribeEventRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SubscribeEventRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SubscribeEventRequestTenantContext) SetTenantId(v string) *SubscribeEventRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SubscribeEventRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
