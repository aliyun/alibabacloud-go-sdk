// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *ApiTemplate
	GetApiName() *string
	SetContent(v string) *ApiTemplate
	GetContent() *string
	SetRegionId(v string) *ApiTemplate
	GetRegionId() *string
	SetResourceGroupId(v string) *ApiTemplate
	GetResourceGroupId() *string
	SetStatus(v string) *ApiTemplate
	GetStatus() *string
	SetTemplateId(v string) *ApiTemplate
	GetTemplateId() *string
	SetTemplateName(v string) *ApiTemplate
	GetTemplateName() *string
}

type ApiTemplate struct {
	// 接口名。
	//
	// example:
	//
	// CreateCluster
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 模版接口参数。
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 区域ID。
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源组ID。
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 模板状态。
	//
	// example:
	//
	// READY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板ID。
	//
	// example:
	//
	// AT-Af***
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板ID。
	//
	// example:
	//
	// AT-Af***
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ApiTemplate) String() string {
	return dara.Prettify(s)
}

func (s ApiTemplate) GoString() string {
	return s.String()
}

func (s *ApiTemplate) GetApiName() *string {
	return s.ApiName
}

func (s *ApiTemplate) GetContent() *string {
	return s.Content
}

func (s *ApiTemplate) GetRegionId() *string {
	return s.RegionId
}

func (s *ApiTemplate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ApiTemplate) GetStatus() *string {
	return s.Status
}

func (s *ApiTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ApiTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ApiTemplate) SetApiName(v string) *ApiTemplate {
	s.ApiName = &v
	return s
}

func (s *ApiTemplate) SetContent(v string) *ApiTemplate {
	s.Content = &v
	return s
}

func (s *ApiTemplate) SetRegionId(v string) *ApiTemplate {
	s.RegionId = &v
	return s
}

func (s *ApiTemplate) SetResourceGroupId(v string) *ApiTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *ApiTemplate) SetStatus(v string) *ApiTemplate {
	s.Status = &v
	return s
}

func (s *ApiTemplate) SetTemplateId(v string) *ApiTemplate {
	s.TemplateId = &v
	return s
}

func (s *ApiTemplate) SetTemplateName(v string) *ApiTemplate {
	s.TemplateName = &v
	return s
}

func (s *ApiTemplate) Validate() error {
	return dara.Validate(s)
}
