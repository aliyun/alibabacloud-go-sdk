// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeature(v string) *GetFeatureDetailsRequest
	GetFeature() *string
	SetRegionId(v string) *GetFeatureDetailsRequest
	GetRegionId() *string
}

type GetFeatureDetailsRequest struct {
	// The one or more features that you want to query. Valid values:
	//
	// 	- Terraform: the Terraform hosting feature.
	//
	// 	- ResourceCleaner: the resource cleaner feature. You can use ALIYUN::ROS::ResourceCleaner to create a resource cleaner.
	//
	// 	- TemplateScratch: the scenario feature.
	//
	// 	- All: all features that are supported by ROS.
	//
	// This parameter is required.
	//
	// example:
	//
	// Terraform
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetFeatureDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsRequest) GetFeature() *string {
	return s.Feature
}

func (s *GetFeatureDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetFeatureDetailsRequest) SetFeature(v string) *GetFeatureDetailsRequest {
	s.Feature = &v
	return s
}

func (s *GetFeatureDetailsRequest) SetRegionId(v string) *GetFeatureDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *GetFeatureDetailsRequest) Validate() error {
	return dara.Validate(s)
}
