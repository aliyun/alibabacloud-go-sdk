// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody
	GetResources() []*ListResourcesResponseBodyResources
	SetPageNumber(v int32) *ListResourcesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourcesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourcesResponseBody
	GetTotalCount() *int32
}

type ListResourcesResponseBody struct {
	Resources []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 67585D9B-EFA5-5E51-BAB1-8FF07DA1B36F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 9
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) GetResources() []*ListResourcesResponseBodyResources {
	return s.Resources
}

func (s *ListResourcesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourcesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourcesResponseBody) SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourcesResponseBody) Validate() error {
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

type ListResourcesResponseBodyResources struct {
	// example:
	//
	// 365845
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 2024-05-03T02:22:59Z
	CreateTime           *string   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DependsOnResourceIds []*string `json:"dependsOnResourceIds,omitempty" xml:"dependsOnResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// example:
	//
	// {}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// {}
	PropertyVariables map[string]interface{} `json:"propertyVariables,omitempty" xml:"propertyVariables,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// acs:vpc:cn-hangzhou:12345:test
	ResourceArn *string `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	// example:
	//
	// rg-aekzyqyghofqbxy
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// i-efegsewrttfd
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// test
	ResourceName    *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourcePageUrl *string `json:"resourcePageUrl,omitempty" xml:"resourcePageUrl,omitempty"`
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// enable
	Status *string                                   `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListResourcesResponseBodyResourcesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// terraform arn
	//
	// example:
	//
	// alicloud_vpc
	TerraformArn *string `json:"terraformArn,omitempty" xml:"terraformArn,omitempty"`
	// terraform code
	//
	// example:
	//
	// alicloud_vpc
	TerraformCode *string `json:"terraformCode,omitempty" xml:"terraformCode,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) GetAccountId() *string {
	return s.AccountId
}

func (s *ListResourcesResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourcesResponseBodyResources) GetDependsOnResourceIds() []*string {
	return s.DependsOnResourceIds
}

func (s *ListResourcesResponseBodyResources) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListResourcesResponseBodyResources) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListResourcesResponseBodyResources) GetPropertyVariables() map[string]interface{} {
	return s.PropertyVariables
}

func (s *ListResourcesResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResourcesResponseBodyResources) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListResourcesResponseBodyResources) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListResourcesResponseBodyResources) GetResourcePageUrl() *string {
	return s.ResourcePageUrl
}

func (s *ListResourcesResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourcesResponseBodyResources) GetStatus() *string {
	return s.Status
}

func (s *ListResourcesResponseBodyResources) GetTags() []*ListResourcesResponseBodyResourcesTags {
	return s.Tags
}

func (s *ListResourcesResponseBodyResources) GetTerraformArn() *string {
	return s.TerraformArn
}

func (s *ListResourcesResponseBodyResources) GetTerraformCode() *string {
	return s.TerraformCode
}

func (s *ListResourcesResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListResourcesResponseBodyResources) SetAccountId(v string) *ListResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetCreateTime(v string) *ListResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetDependsOnResourceIds(v []*string) *ListResourcesResponseBodyResources {
	s.DependsOnResourceIds = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProductCode(v string) *ListResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProperties(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.Properties = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetPropertyVariables(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.PropertyVariables = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetRegionId(v string) *ListResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceArn(v string) *ListResourcesResponseBodyResources {
	s.ResourceArn = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceGroupId(v string) *ListResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceId(v string) *ListResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceName(v string) *ListResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourcePageUrl(v string) *ListResourcesResponseBodyResources {
	s.ResourcePageUrl = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceType(v string) *ListResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetStatus(v string) *ListResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTags(v []*ListResourcesResponseBodyResourcesTags) *ListResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTerraformArn(v string) *ListResourcesResponseBodyResources {
	s.TerraformArn = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTerraformCode(v string) *ListResourcesResponseBodyResources {
	s.TerraformCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetZoneId(v string) *ListResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourcesResponseBodyResourcesTags struct {
	// example:
	//
	// test
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// example:
	//
	// test
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListResourcesResponseBodyResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListResourcesResponseBodyResourcesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListResourcesResponseBodyResourcesTags) SetTagKey(v string) *ListResourcesResponseBodyResourcesTags {
	s.TagKey = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesTags) SetTagValue(v string) *ListResourcesResponseBodyResourcesTags {
	s.TagValue = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesTags) Validate() error {
	return dara.Validate(s)
}
