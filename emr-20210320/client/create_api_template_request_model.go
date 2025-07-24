// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *CreateApiTemplateRequest
	GetApiName() *string
	SetContent(v string) *CreateApiTemplateRequest
	GetContent() *string
	SetRegionId(v string) *CreateApiTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApiTemplateRequest
	GetResourceGroupId() *string
	SetTemplateName(v string) *CreateApiTemplateRequest
	GetTemplateName() *string
}

type CreateApiTemplateRequest struct {
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
	// content
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
	// Cluster template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATALAKE template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateApiTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateApiTemplateRequest) GetApiName() *string {
	return s.ApiName
}

func (s *CreateApiTemplateRequest) GetContent() *string {
	return s.Content
}

func (s *CreateApiTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApiTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApiTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateApiTemplateRequest) SetApiName(v string) *CreateApiTemplateRequest {
	s.ApiName = &v
	return s
}

func (s *CreateApiTemplateRequest) SetContent(v string) *CreateApiTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateApiTemplateRequest) SetRegionId(v string) *CreateApiTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApiTemplateRequest) SetResourceGroupId(v string) *CreateApiTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApiTemplateRequest) SetTemplateName(v string) *CreateApiTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateApiTemplateRequest) Validate() error {
	return dara.Validate(s)
}
