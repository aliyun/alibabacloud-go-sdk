// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlidingAssistantInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *GetAlidingAssistantInfoRequest
	GetAssistantId() *string
	SetTenantContext(v *GetAlidingAssistantInfoRequestTenantContext) *GetAlidingAssistantInfoRequest
	GetTenantContext() *GetAlidingAssistantInfoRequestTenantContext
}

type GetAlidingAssistantInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AssistantId   *string                                      `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	TenantContext *GetAlidingAssistantInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetAlidingAssistantInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *GetAlidingAssistantInfoRequest) GetTenantContext() *GetAlidingAssistantInfoRequestTenantContext {
	return s.TenantContext
}

func (s *GetAlidingAssistantInfoRequest) SetAssistantId(v string) *GetAlidingAssistantInfoRequest {
	s.AssistantId = &v
	return s
}

func (s *GetAlidingAssistantInfoRequest) SetTenantContext(v *GetAlidingAssistantInfoRequestTenantContext) *GetAlidingAssistantInfoRequest {
	s.TenantContext = v
	return s
}

func (s *GetAlidingAssistantInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlidingAssistantInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetAlidingAssistantInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetAlidingAssistantInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetAlidingAssistantInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAlidingAssistantInfoRequestTenantContext) SetTenantId(v string) *GetAlidingAssistantInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetAlidingAssistantInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
