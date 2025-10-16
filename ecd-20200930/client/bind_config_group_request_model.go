// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *BindConfigGroupRequest
	GetGroupId() *string
	SetRegionId(v string) *BindConfigGroupRequest
	GetRegionId() *string
	SetResourceInfos(v []*BindConfigGroupRequestResourceInfos) *BindConfigGroupRequest
	GetResourceInfos() []*BindConfigGroupRequestResourceInfos
}

type BindConfigGroupRequest struct {
	// The ID of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccg-0chlk9b65lj8z****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources to which you want to bind the configuration group.
	//
	// This parameter is required.
	ResourceInfos []*BindConfigGroupRequestResourceInfos `json:"ResourceInfos,omitempty" xml:"ResourceInfos,omitempty" type:"Repeated"`
}

func (s BindConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s BindConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *BindConfigGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *BindConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindConfigGroupRequest) GetResourceInfos() []*BindConfigGroupRequestResourceInfos {
	return s.ResourceInfos
}

func (s *BindConfigGroupRequest) SetGroupId(v string) *BindConfigGroupRequest {
	s.GroupId = &v
	return s
}

func (s *BindConfigGroupRequest) SetRegionId(v string) *BindConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *BindConfigGroupRequest) SetResourceInfos(v []*BindConfigGroupRequestResourceInfos) *BindConfigGroupRequest {
	s.ResourceInfos = v
	return s
}

func (s *BindConfigGroupRequest) Validate() error {
	if s.ResourceInfos != nil {
		for _, item := range s.ResourceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BindConfigGroupRequestResourceInfos struct {
	// The service type of the resource.
	//
	// Valid value:
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// ecd-1bo4xotjvwyon****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- RESOURCE_GROUP: the resource group
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s BindConfigGroupRequestResourceInfos) String() string {
	return dara.Prettify(s)
}

func (s BindConfigGroupRequestResourceInfos) GoString() string {
	return s.String()
}

func (s *BindConfigGroupRequestResourceInfos) GetProductType() *string {
	return s.ProductType
}

func (s *BindConfigGroupRequestResourceInfos) GetResourceId() *string {
	return s.ResourceId
}

func (s *BindConfigGroupRequestResourceInfos) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *BindConfigGroupRequestResourceInfos) GetResourceType() *string {
	return s.ResourceType
}

func (s *BindConfigGroupRequestResourceInfos) SetProductType(v string) *BindConfigGroupRequestResourceInfos {
	s.ProductType = &v
	return s
}

func (s *BindConfigGroupRequestResourceInfos) SetResourceId(v string) *BindConfigGroupRequestResourceInfos {
	s.ResourceId = &v
	return s
}

func (s *BindConfigGroupRequestResourceInfos) SetResourceRegionId(v string) *BindConfigGroupRequestResourceInfos {
	s.ResourceRegionId = &v
	return s
}

func (s *BindConfigGroupRequestResourceInfos) SetResourceType(v string) *BindConfigGroupRequestResourceInfos {
	s.ResourceType = &v
	return s
}

func (s *BindConfigGroupRequestResourceInfos) Validate() error {
	return dara.Validate(s)
}
