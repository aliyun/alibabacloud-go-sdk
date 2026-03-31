// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComplianceOption(v int32) *GetDiscoveredResourceRequest
	GetComplianceOption() *int32
	SetRegion(v string) *GetDiscoveredResourceRequest
	GetRegion() *string
	SetResourceId(v string) *GetDiscoveredResourceRequest
	GetResourceId() *string
	SetResourceType(v string) *GetDiscoveredResourceRequest
	GetResourceType() *string
}

type GetDiscoveredResourceRequest struct {
	// Specifies whether to query the compliance results of the resource. Valid values:
	//
	// 	- 0 (default): does not query the compliance results of the resource.
	//
	// 	- 1: queries the compliance results of the resource.
	//
	// example:
	//
	// 0
	ComplianceOption *int32 `json:"ComplianceOption,omitempty" xml:"ComplianceOption,omitempty"`
	// The ID of the region in which the resource resides.
	//
	// For more information about how to query the region ID of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/411702.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource ID.
	//
	// For more information about how to obtain the ID of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/411702.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to obtain the type of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/411702.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceRequest) GetComplianceOption() *int32 {
	return s.ComplianceOption
}

func (s *GetDiscoveredResourceRequest) GetRegion() *string {
	return s.Region
}

func (s *GetDiscoveredResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetDiscoveredResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetDiscoveredResourceRequest) SetComplianceOption(v int32) *GetDiscoveredResourceRequest {
	s.ComplianceOption = &v
	return s
}

func (s *GetDiscoveredResourceRequest) SetRegion(v string) *GetDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceRequest) SetResourceId(v string) *GetDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetDiscoveredResourceRequest) SetResourceType(v string) *GetDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceRequest) Validate() error {
	return dara.Validate(s)
}
