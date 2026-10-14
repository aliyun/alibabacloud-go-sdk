// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ConfigWebCCTemplateRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ConfigWebCCTemplateRequest
	GetResourceGroupId() *string
	SetTemplate(v string) *ConfigWebCCTemplateRequest
	GetTemplate() *string
}

type ConfigWebCCTemplateRequest struct {
	// The domain name of the website service.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. By default, this parameter is empty, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The defense mode for frequency control (HTTP flood mitigation). Valid values:
	//
	// - **default**: Normal.
	//
	// - **gf_under_attack**: Emergency.
	//
	// - **gf_sos_verify**: Strict.
	//
	// - **gf_sos_enhance**: Super Strict.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ConfigWebCCTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigWebCCTemplateRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigWebCCTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigWebCCTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *ConfigWebCCTemplateRequest) SetDomain(v string) *ConfigWebCCTemplateRequest {
	s.Domain = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) SetResourceGroupId(v string) *ConfigWebCCTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) SetTemplate(v string) *ConfigWebCCTemplateRequest {
	s.Template = &v
	return s
}

func (s *ConfigWebCCTemplateRequest) Validate() error {
	return dara.Validate(s)
}
