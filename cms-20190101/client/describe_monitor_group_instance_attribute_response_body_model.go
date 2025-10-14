// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetRequestId() *string
	SetResources(v *DescribeMonitorGroupInstanceAttributeResponseBodyResources) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetResources() *DescribeMonitorGroupInstanceAttributeResponseBodyResources
	SetSuccess(v bool) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody
	GetTotal() *int32
}

type DescribeMonitorGroupInstanceAttributeResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9FB8EA79-7279-4482-8D6D-3D28EEDD871A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that are associated with the application group.
	Resources *DescribeMonitorGroupInstanceAttributeResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetResources() *DescribeMonitorGroupInstanceAttributeResponseBodyResources {
	return s.Resources
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetCode(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetMessage(v string) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetRequestId(v string) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetResources(v *DescribeMonitorGroupInstanceAttributeResponseBodyResources) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetSuccess(v bool) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetTotal(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) Validate() error {
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResources struct {
	Resource []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResources) GetResource() []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResources) SetResource(v []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) *DescribeMonitorGroupInstanceAttributeResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource struct {
	// The name of the cloud service.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The resource description.
	//
	// example:
	//
	// desc_test
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The dimensions of the resource that is associated with the application group.
	//
	// example:
	//
	// {"instanceId":"i-m5e0k0bexac8tykr****"}
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-m5e0k0bexac8tykr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// hostName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The region.
	Region *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	// The tag of the resource.
	Tags *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The VPC description.
	Vpc *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetDesc() *string {
	return s.Desc
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetRegion() *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion {
	return s.Region
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetTags() *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags {
	return s.Tags
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GetVpc() *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc {
	return s.Vpc
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetCategory(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetDesc(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Desc = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetDimension(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Dimension = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetInstanceName(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.InstanceName = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetNetworkType(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.NetworkType = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetRegion(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Region = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetTags(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Tags = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetVpc(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Vpc = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) Validate() error {
	if s.Region != nil {
		if err := s.Region.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion struct {
	// The zone.
	//
	// example:
	//
	// cn-hangzhou-f
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) GetAvailabilityZone() *string {
	return s.AvailabilityZone
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) SetAvailabilityZone(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) SetRegionId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags struct {
	Tag []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) GetTag() []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag {
	return s.Tag
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) SetTag(v []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags {
	s.Tag = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// instanceNetworkType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// VPC
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) SetKey(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) SetValue(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc struct {
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zew7etgiceg21****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2ze36seq79n992****
	VswitchInstanceId *string `json:"VswitchInstanceId,omitempty" xml:"VswitchInstanceId,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) GetVswitchInstanceId() *string {
	return s.VswitchInstanceId
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) SetVpcInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) SetVswitchInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc {
	s.VswitchInstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) Validate() error {
	return dara.Validate(s)
}
