// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationTemplate(v *GetApplicationTemplateResponseBodyApplicationTemplate) *GetApplicationTemplateResponseBody
	GetApplicationTemplate() *GetApplicationTemplateResponseBodyApplicationTemplate
	SetRequestId(v string) *GetApplicationTemplateResponseBody
	GetRequestId() *string
}

type GetApplicationTemplateResponseBody struct {
	// The application template information.
	ApplicationTemplate *GetApplicationTemplateResponseBodyApplicationTemplate `json:"ApplicationTemplate,omitempty" xml:"ApplicationTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBody) GetApplicationTemplate() *GetApplicationTemplateResponseBodyApplicationTemplate {
	return s.ApplicationTemplate
}

func (s *GetApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationTemplateResponseBody) SetApplicationTemplate(v *GetApplicationTemplateResponseBodyApplicationTemplate) *GetApplicationTemplateResponseBody {
	s.ApplicationTemplate = v
	return s
}

func (s *GetApplicationTemplateResponseBody) SetRequestId(v string) *GetApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationTemplateResponseBody) Validate() error {
	if s.ApplicationTemplate != nil {
		if err := s.ApplicationTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationTemplateResponseBodyApplicationTemplate struct {
	// The application template ID.
	//
	// example:
	//
	// apt_ramintlrole_ixxxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The application template name.
	//
	// example:
	//
	// Alibaba Cloud
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// The time when the application template was created.
	//
	// example:
	//
	// 1730341123000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The application template description.
	//
	// example:
	//
	// Alibaba Cloud SSO
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The help document URL of the application template.
	//
	// example:
	//
	// https://example.com/document_detail/409xxx.html
	HelpDocumentUrl *string `json:"HelpDocumentUrl,omitempty" xml:"HelpDocumentUrl,omitempty"`
	// The logo URL of the application template.
	//
	// example:
	//
	// https://example.com/imgextra/i4/O1CN01xTLxLb1WtyKksHW1H_!!6000000002847-2-tps-xxx-xxx.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The service code of the Alibaba Cloud service that manages the application template.
	//
	// example:
	//
	// bastionhost
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// The sale information of the application template.
	SaleInfo *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo `json:"SaleInfo,omitempty" xml:"SaleInfo,omitempty" type:"Struct"`
	// The console URL of the Alibaba Cloud service that manages the application template.
	//
	// example:
	//
	// https://example.com/?p=bastion
	ServiceConsoleUrl *string `json:"ServiceConsoleUrl,omitempty" xml:"ServiceConsoleUrl,omitempty"`
	// Indicates whether the application template is managed by an Alibaba Cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The supported SSO protocols.
	SsoTypes []*string `json:"SsoTypes,omitempty" xml:"SsoTypes,omitempty" type:"Repeated"`
	// The time when the application template was last updated.
	//
	// example:
	//
	// 1730341124000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetApplicationTemplateResponseBodyApplicationTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBodyApplicationTemplate) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetApplicationTemplateName() *string {
	return s.ApplicationTemplateName
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetHelpDocumentUrl() *string {
	return s.HelpDocumentUrl
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetManagedServiceCode() *string {
	return s.ManagedServiceCode
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetSaleInfo() *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo {
	return s.SaleInfo
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetServiceConsoleUrl() *string {
	return s.ServiceConsoleUrl
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetSsoTypes() []*string {
	return s.SsoTypes
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetApplicationTemplateId(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetApplicationTemplateName(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ApplicationTemplateName = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetCreateTime(v int64) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetDescription(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.Description = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetHelpDocumentUrl(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.HelpDocumentUrl = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetLogoUrl(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.LogoUrl = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetManagedServiceCode(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ManagedServiceCode = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetSaleInfo(v *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.SaleInfo = v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetServiceConsoleUrl(v string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ServiceConsoleUrl = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetServiceManaged(v bool) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.ServiceManaged = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetSsoTypes(v []*string) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.SsoTypes = v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) SetUpdateTime(v int64) *GetApplicationTemplateResponseBodyApplicationTemplate {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplate) Validate() error {
	if s.SaleInfo != nil {
		if err := s.SaleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo struct {
	// Indicates whether the application template is always free.
	//
	// example:
	//
	// true
	AlwaysFree *bool `json:"AlwaysFree,omitempty" xml:"AlwaysFree,omitempty"`
}

func (s GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) GoString() string {
	return s.String()
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) GetAlwaysFree() *bool {
	return s.AlwaysFree
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) SetAlwaysFree(v bool) *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo {
	s.AlwaysFree = &v
	return s
}

func (s *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo) Validate() error {
	return dara.Validate(s)
}
