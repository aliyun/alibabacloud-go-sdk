// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *DisableSceneGroupTemplateRequest
	GetOpenConversationId() *string
	SetTemplateId(v string) *DisableSceneGroupTemplateRequest
	GetTemplateId() *string
	SetTenantContext(v *DisableSceneGroupTemplateRequestTenantContext) *DisableSceneGroupTemplateRequest
	GetTenantContext() *DisableSceneGroupTemplateRequestTenantContext
}

type DisableSceneGroupTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2efdt*****fswe==
	TemplateId    *string                                        `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContext *DisableSceneGroupTemplateRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DisableSceneGroupTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateRequest) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *DisableSceneGroupTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DisableSceneGroupTemplateRequest) GetTenantContext() *DisableSceneGroupTemplateRequestTenantContext {
	return s.TenantContext
}

func (s *DisableSceneGroupTemplateRequest) SetOpenConversationId(v string) *DisableSceneGroupTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DisableSceneGroupTemplateRequest) SetTemplateId(v string) *DisableSceneGroupTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DisableSceneGroupTemplateRequest) SetTenantContext(v *DisableSceneGroupTemplateRequestTenantContext) *DisableSceneGroupTemplateRequest {
	s.TenantContext = v
	return s
}

func (s *DisableSceneGroupTemplateRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableSceneGroupTemplateRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DisableSceneGroupTemplateRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DisableSceneGroupTemplateRequestTenantContext) SetTenantId(v string) *DisableSceneGroupTemplateRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DisableSceneGroupTemplateRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
