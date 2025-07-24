// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *UpdateApiTemplateRequest
	GetApiName() *string
	SetContent(v string) *UpdateApiTemplateRequest
	GetContent() *string
	SetRegionId(v string) *UpdateApiTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateApiTemplateRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *UpdateApiTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateApiTemplateRequest
	GetTemplateName() *string
}

type UpdateApiTemplateRequest struct {
	// The name of the API operation. You can create only a cluster API operation template. Set the value to CreateCluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateCluster
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The content of the cluster API operation template. Set the value to JSON strings of the request parameters of the [CreateCluster](https://help.aliyun.com/document_detail/454393.html) API operation for creating a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateCluster
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzabjyop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID (it is recommended to use the parameter TemplateId).
	//
	// This parameter is required.
	//
	// example:
	//
	// AT-****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// datalakeTest1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateApiTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiTemplateRequest) GetApiName() *string {
	return s.ApiName
}

func (s *UpdateApiTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateApiTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApiTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateApiTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateApiTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateApiTemplateRequest) SetApiName(v string) *UpdateApiTemplateRequest {
	s.ApiName = &v
	return s
}

func (s *UpdateApiTemplateRequest) SetContent(v string) *UpdateApiTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateApiTemplateRequest) SetRegionId(v string) *UpdateApiTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApiTemplateRequest) SetResourceGroupId(v string) *UpdateApiTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateApiTemplateRequest) SetTemplateId(v string) *UpdateApiTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateApiTemplateRequest) SetTemplateName(v string) *UpdateApiTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateApiTemplateRequest) Validate() error {
	return dara.Validate(s)
}
