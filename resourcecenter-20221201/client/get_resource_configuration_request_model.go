// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *GetResourceConfigurationRequest
	GetResourceId() *string
	SetResourceRegionId(v string) *GetResourceConfigurationRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *GetResourceConfigurationRequest
	GetResourceType() *string
}

type GetResourceConfigurationRequest struct {
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-bp1kyg72m****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// For more information about the resource types supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceConfigurationRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *GetResourceConfigurationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationRequest) SetResourceId(v string) *GetResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceRegionId(v string) *GetResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceType(v string) *GetResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
