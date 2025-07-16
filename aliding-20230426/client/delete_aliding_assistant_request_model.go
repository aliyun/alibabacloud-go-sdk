// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlidingAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *DeleteAlidingAssistantRequest
	GetAssistantId() *string
	SetTenantContext(v *DeleteAlidingAssistantRequestTenantContext) *DeleteAlidingAssistantRequest
	GetTenantContext() *DeleteAlidingAssistantRequestTenantContext
}

type DeleteAlidingAssistantRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AssistantId   *string                                     `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	TenantContext *DeleteAlidingAssistantRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeleteAlidingAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *DeleteAlidingAssistantRequest) GetTenantContext() *DeleteAlidingAssistantRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteAlidingAssistantRequest) SetAssistantId(v string) *DeleteAlidingAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *DeleteAlidingAssistantRequest) SetTenantContext(v *DeleteAlidingAssistantRequestTenantContext) *DeleteAlidingAssistantRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteAlidingAssistantRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteAlidingAssistantRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteAlidingAssistantRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlidingAssistantRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteAlidingAssistantRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteAlidingAssistantRequestTenantContext) SetTenantId(v string) *DeleteAlidingAssistantRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteAlidingAssistantRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
