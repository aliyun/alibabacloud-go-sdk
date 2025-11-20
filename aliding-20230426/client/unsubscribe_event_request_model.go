// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScope(v string) *UnsubscribeEventRequest
	GetScope() *string
	SetScopeId(v string) *UnsubscribeEventRequest
	GetScopeId() *string
	SetTenantContext(v *UnsubscribeEventRequestTenantContext) *UnsubscribeEventRequest
	GetTenantContext() *UnsubscribeEventRequestTenantContext
}

type UnsubscribeEventRequest struct {
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
	ScopeId       *string                               `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	TenantContext *UnsubscribeEventRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UnsubscribeEventRequest) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventRequest) GetScope() *string {
	return s.Scope
}

func (s *UnsubscribeEventRequest) GetScopeId() *string {
	return s.ScopeId
}

func (s *UnsubscribeEventRequest) GetTenantContext() *UnsubscribeEventRequestTenantContext {
	return s.TenantContext
}

func (s *UnsubscribeEventRequest) SetScope(v string) *UnsubscribeEventRequest {
	s.Scope = &v
	return s
}

func (s *UnsubscribeEventRequest) SetScopeId(v string) *UnsubscribeEventRequest {
	s.ScopeId = &v
	return s
}

func (s *UnsubscribeEventRequest) SetTenantContext(v *UnsubscribeEventRequestTenantContext) *UnsubscribeEventRequest {
	s.TenantContext = v
	return s
}

func (s *UnsubscribeEventRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnsubscribeEventRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UnsubscribeEventRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UnsubscribeEventRequestTenantContext) SetTenantId(v string) *UnsubscribeEventRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UnsubscribeEventRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
