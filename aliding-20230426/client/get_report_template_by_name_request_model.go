// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateName(v string) *GetReportTemplateByNameRequest
	GetTemplateName() *string
	SetTenantContext(v *GetReportTemplateByNameRequestTenantContext) *GetReportTemplateByNameRequest
	GetTenantContext() *GetReportTemplateByNameRequestTenantContext
}

type GetReportTemplateByNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 我管理的模版
	TemplateName  *string                                      `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContext *GetReportTemplateByNameRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetReportTemplateByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameRequest) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetReportTemplateByNameRequest) GetTenantContext() *GetReportTemplateByNameRequestTenantContext {
	return s.TenantContext
}

func (s *GetReportTemplateByNameRequest) SetTemplateName(v string) *GetReportTemplateByNameRequest {
	s.TemplateName = &v
	return s
}

func (s *GetReportTemplateByNameRequest) SetTenantContext(v *GetReportTemplateByNameRequestTenantContext) *GetReportTemplateByNameRequest {
	s.TenantContext = v
	return s
}

func (s *GetReportTemplateByNameRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReportTemplateByNameRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetReportTemplateByNameRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetReportTemplateByNameRequestTenantContext) SetTenantId(v string) *GetReportTemplateByNameRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetReportTemplateByNameRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
