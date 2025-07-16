// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReportTemplateByNameShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateName(v string) *GetReportTemplateByNameShrinkRequest
	GetTemplateName() *string
	SetTenantContextShrink(v string) *GetReportTemplateByNameShrinkRequest
	GetTenantContextShrink() *string
}

type GetReportTemplateByNameShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 我管理的模版
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetReportTemplateByNameShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReportTemplateByNameShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetReportTemplateByNameShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetReportTemplateByNameShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetReportTemplateByNameShrinkRequest) SetTemplateName(v string) *GetReportTemplateByNameShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetReportTemplateByNameShrinkRequest) SetTenantContextShrink(v string) *GetReportTemplateByNameShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetReportTemplateByNameShrinkRequest) Validate() error {
	return dara.Validate(s)
}
