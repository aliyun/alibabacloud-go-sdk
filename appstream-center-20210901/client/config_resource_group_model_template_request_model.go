// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigResourceGroupModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *ConfigResourceGroupModelTemplateRequest
	GetModelTemplateId() *string
	SetResourceGroupId(v string) *ConfigResourceGroupModelTemplateRequest
	GetResourceGroupId() *string
}

type ConfigResourceGroupModelTemplateRequest struct {
	// The model group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The WUYING resource group ID, which starts with rg-. This resource group is managed by the WUYING Enterprise Edition management and control service, and its automatic creation is performed by the system after the deployment of agent runtimes such as JVS Computer, OpenClaw, and Hermes Agent. You can obtain it from the ResourceGroup[].ResourceGroupId field in the response of the DescribeResourceGroups operation of WUYING Workspace eds-user (API version 2021-03-08), or from the Data.ResourceGroupId field in the response of the GetRuntimeModelConfig operation of this service.
	//
	// > Other IDs such as the account-level resource group ID and node pool ID in Alibaba Cloud Resource Manager cannot be passed as this parameter. Only WUYING resource group IDs are valid.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ConfigResourceGroupModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigResourceGroupModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigResourceGroupModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ConfigResourceGroupModelTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigResourceGroupModelTemplateRequest) SetModelTemplateId(v string) *ConfigResourceGroupModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateRequest) SetResourceGroupId(v string) *ConfigResourceGroupModelTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigResourceGroupModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
