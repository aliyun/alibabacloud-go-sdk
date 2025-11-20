// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetCardTemplateRequest
	GetTemplateId() *string
	SetTenantContext(v *GetCardTemplateRequestTenantContext) *GetCardTemplateRequest
	GetTenantContext() *GetCardTemplateRequestTenantContext
}

type GetCardTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	TemplateId    *string                              `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContext *GetCardTemplateRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetCardTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetCardTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetCardTemplateRequest) GetTenantContext() *GetCardTemplateRequestTenantContext {
	return s.TenantContext
}

func (s *GetCardTemplateRequest) SetTemplateId(v string) *GetCardTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCardTemplateRequest) SetTenantContext(v *GetCardTemplateRequestTenantContext) *GetCardTemplateRequest {
	s.TenantContext = v
	return s
}

func (s *GetCardTemplateRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCardTemplateRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetCardTemplateRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetCardTemplateRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetCardTemplateRequestTenantContext) SetTenantId(v string) *GetCardTemplateRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetCardTemplateRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
