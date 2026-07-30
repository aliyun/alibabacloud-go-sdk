// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UnbindConfigGroupRequest
	GetRegionId() *string
	SetResourceInfos(v []*UnbindConfigGroupRequestResourceInfos) *UnbindConfigGroupRequest
	GetResourceInfos() []*UnbindConfigGroupRequestResourceInfos
	SetType(v string) *UnbindConfigGroupRequest
	GetType() *string
}

type UnbindConfigGroupRequest struct {
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources from which you want to unbind the configuration group.
	//
	// This parameter is required.
	ResourceInfos []*UnbindConfigGroupRequestResourceInfos `json:"ResourceInfos,omitempty" xml:"ResourceInfos,omitempty" type:"Repeated"`
	// The type of the configuration group.
	//
	// Valid value:
	//
	// 	- Timer: the scheduled task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Timer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UnbindConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *UnbindConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnbindConfigGroupRequest) GetResourceInfos() []*UnbindConfigGroupRequestResourceInfos {
	return s.ResourceInfos
}

func (s *UnbindConfigGroupRequest) GetType() *string {
	return s.Type
}

func (s *UnbindConfigGroupRequest) SetRegionId(v string) *UnbindConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindConfigGroupRequest) SetResourceInfos(v []*UnbindConfigGroupRequestResourceInfos) *UnbindConfigGroupRequest {
	s.ResourceInfos = v
	return s
}

func (s *UnbindConfigGroupRequest) SetType(v string) *UnbindConfigGroupRequest {
	s.Type = &v
	return s
}

func (s *UnbindConfigGroupRequest) Validate() error {
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

type UnbindConfigGroupRequestResourceInfos struct {
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
	// ecd-ctwj0bk3l5nz****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-chengdu
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- RESOURCE_GROUP: the resource group.
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UnbindConfigGroupRequestResourceInfos) String() string {
	return dara.Prettify(s)
}

func (s UnbindConfigGroupRequestResourceInfos) GoString() string {
	return s.String()
}

func (s *UnbindConfigGroupRequestResourceInfos) GetProductType() *string {
	return s.ProductType
}

func (s *UnbindConfigGroupRequestResourceInfos) GetResourceId() *string {
	return s.ResourceId
}

func (s *UnbindConfigGroupRequestResourceInfos) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *UnbindConfigGroupRequestResourceInfos) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnbindConfigGroupRequestResourceInfos) SetProductType(v string) *UnbindConfigGroupRequestResourceInfos {
	s.ProductType = &v
	return s
}

func (s *UnbindConfigGroupRequestResourceInfos) SetResourceId(v string) *UnbindConfigGroupRequestResourceInfos {
	s.ResourceId = &v
	return s
}

func (s *UnbindConfigGroupRequestResourceInfos) SetResourceRegionId(v string) *UnbindConfigGroupRequestResourceInfos {
	s.ResourceRegionId = &v
	return s
}

func (s *UnbindConfigGroupRequestResourceInfos) SetResourceType(v string) *UnbindConfigGroupRequestResourceInfos {
	s.ResourceType = &v
	return s
}

func (s *UnbindConfigGroupRequestResourceInfos) Validate() error {
	return dara.Validate(s)
}
