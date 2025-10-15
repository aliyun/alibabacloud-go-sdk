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
	ApplicationTemplate *GetApplicationTemplateResponseBodyApplicationTemplate `json:"ApplicationTemplate,omitempty" xml:"ApplicationTemplate,omitempty" type:"Struct"`
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
	// 应用模板Id
	//
	// example:
	//
	// apt_ramintlrole_ixxxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// 应用模板名称
	ApplicationTemplateName *string `json:"ApplicationTemplateName,omitempty" xml:"ApplicationTemplateName,omitempty"`
	// 应用模板创建时间
	//
	// example:
	//
	// 1730341123000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用模板描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用模板对应帮助文档地址
	//
	// example:
	//
	// https://example.com/document_detail/409xxx.html
	HelpDocumentUrl *string `json:"HelpDocumentUrl,omitempty" xml:"HelpDocumentUrl,omitempty"`
	// 应用模板Logo地址
	//
	// example:
	//
	// https://example.com/imgextra/i4/O1CN01xTLxLb1WtyKksHW1H_!!6000000002847-2-tps-xxx-xxx.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// 托管应用模板的云产品ServiceCode。当且仅当ServiceManaged为true是返回。
	//
	// example:
	//
	// bastionhost
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// 应用模板售卖信息
	SaleInfo *GetApplicationTemplateResponseBodyApplicationTemplateSaleInfo `json:"SaleInfo,omitempty" xml:"SaleInfo,omitempty" type:"Struct"`
	// 托管应用模板的云产品控制台地址。当且仅当ServiceManaged为true是返回。
	//
	// example:
	//
	// https://example.com/?p=bastion
	ServiceConsoleUrl *string `json:"ServiceConsoleUrl,omitempty" xml:"ServiceConsoleUrl,omitempty"`
	// 应用模板是否被云产品托管。
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// 支持SSO协议
	SsoTypes []*string `json:"SsoTypes,omitempty" xml:"SsoTypes,omitempty" type:"Repeated"`
	// 应用模板更新时间
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
	// 是否永久免费
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
