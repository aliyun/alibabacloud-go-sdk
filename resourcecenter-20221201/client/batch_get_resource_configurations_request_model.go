// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetResourceConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v []*BatchGetResourceConfigurationsRequestResources) *BatchGetResourceConfigurationsRequest
	GetResources() []*BatchGetResourceConfigurationsRequestResources
}

type BatchGetResourceConfigurationsRequest struct {
	// The list of resources.
	Resources []*BatchGetResourceConfigurationsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s BatchGetResourceConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsRequest) GetResources() []*BatchGetResourceConfigurationsRequestResources {
	return s.Resources
}

func (s *BatchGetResourceConfigurationsRequest) SetResources(v []*BatchGetResourceConfigurationsRequestResources) *BatchGetResourceConfigurationsRequest {
	s.Resources = v
	return s
}

func (s *BatchGetResourceConfigurationsRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetResourceConfigurationsRequestResources struct {
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eip-bp1kyg72m****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For a list of resource types supported by Resource Center, see [Alibaba Cloud services and resource types that are supported by Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s BatchGetResourceConfigurationsRequestResources) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsRequestResources) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsRequestResources) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchGetResourceConfigurationsRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchGetResourceConfigurationsRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *BatchGetResourceConfigurationsRequestResources) SetRegionId(v string) *BatchGetResourceConfigurationsRequestResources {
	s.RegionId = &v
	return s
}

func (s *BatchGetResourceConfigurationsRequestResources) SetResourceId(v string) *BatchGetResourceConfigurationsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *BatchGetResourceConfigurationsRequestResources) SetResourceType(v string) *BatchGetResourceConfigurationsRequestResources {
	s.ResourceType = &v
	return s
}

func (s *BatchGetResourceConfigurationsRequestResources) Validate() error {
	return dara.Validate(s)
}
