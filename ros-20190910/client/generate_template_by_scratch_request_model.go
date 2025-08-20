// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplateByScratchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionRegionId(v string) *GenerateTemplateByScratchRequest
	GetProvisionRegionId() *string
	SetRegionId(v string) *GenerateTemplateByScratchRequest
	GetRegionId() *string
	SetTemplateScratchId(v string) *GenerateTemplateByScratchRequest
	GetTemplateScratchId() *string
	SetTemplateType(v string) *GenerateTemplateByScratchRequest
	GetTemplateType() *string
}

type GenerateTemplateByScratchRequest struct {
	// The region ID of the new node.
	//
	// example:
	//
	// cn-hangzhou
	ProvisionRegionId *string `json:"ProvisionRegionId,omitempty" xml:"ProvisionRegionId,omitempty"`
	// The region ID of the resource scenario.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource scenario.
	//
	// For more information about how to query the IDs of resource scenarios, see [ListTemplateScratches](https://help.aliyun.com/document_detail/363050.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ts-aa9c62feab844a6b****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the template that Resource Orchestration Service (ROS) generates. ROS can generate templates of the ROS and Terraform types. Default value: ROS.
	//
	// example:
	//
	// ROS
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GenerateTemplateByScratchRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplateByScratchRequest) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchRequest) GetProvisionRegionId() *string {
	return s.ProvisionRegionId
}

func (s *GenerateTemplateByScratchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateTemplateByScratchRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *GenerateTemplateByScratchRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GenerateTemplateByScratchRequest) SetProvisionRegionId(v string) *GenerateTemplateByScratchRequest {
	s.ProvisionRegionId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetRegionId(v string) *GenerateTemplateByScratchRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetTemplateScratchId(v string) *GenerateTemplateByScratchRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetTemplateType(v string) *GenerateTemplateByScratchRequest {
	s.TemplateType = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) Validate() error {
	return dara.Validate(s)
}
