// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateScratchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplateScratchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplateScratchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTemplateScratchesResponseBody
	GetRequestId() *string
	SetTemplateScratches(v []*ListTemplateScratchesResponseBodyTemplateScratches) *ListTemplateScratchesResponseBody
	GetTemplateScratches() []*ListTemplateScratchesResponseBodyTemplateScratches
	SetTotalCount(v int32) *ListTemplateScratchesResponseBody
	GetTotalCount() *int32
}

type ListTemplateScratchesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D1C09606-C58B-558F-9B4E-5BF263D17D09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource scenarios.
	TemplateScratches []*ListTemplateScratchesResponseBodyTemplateScratches `json:"TemplateScratches,omitempty" xml:"TemplateScratches,omitempty" type:"Repeated"`
	// The total number of scenarios.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplateScratchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplateScratchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplateScratchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateScratchesResponseBody) GetTemplateScratches() []*ListTemplateScratchesResponseBodyTemplateScratches {
	return s.TemplateScratches
}

func (s *ListTemplateScratchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplateScratchesResponseBody) SetPageNumber(v int32) *ListTemplateScratchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetPageSize(v int32) *ListTemplateScratchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetRequestId(v string) *ListTemplateScratchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetTemplateScratches(v []*ListTemplateScratchesResponseBodyTemplateScratches) *ListTemplateScratchesResponseBody {
	s.TemplateScratches = v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetTotalCount(v int32) *ListTemplateScratchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) Validate() error {
	if s.TemplateScratches != nil {
		for _, item := range s.TemplateScratches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTemplateScratchesResponseBodyTemplateScratches struct {
	// The time when the resource scenario was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-07T08:06:44
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource scenario.
	//
	// example:
	//
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status code of the resource scenario that failed to be generated.
	//
	// >  This parameter is returned only if the value of Status is GENERATE_FAILED.
	//
	// example:
	//
	// InvalidZoneId
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// 	- LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	//
	// 	- LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	//
	// 	- ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	//
	// example:
	//
	// LongTypePrefixAndIndexSuffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the resource scenario.
	PreferenceParameters []*ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4nxcvht4pmi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	SourceResources []*ListTemplateScratchesResponseBodyTemplateScratchesSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The state of the resource scenario.
	//
	// example:
	//
	// GENERATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource scenario failed to be generated.
	//
	// >  This parameter is returned only if the value of Status is GENERATE_FAILED.
	//
	// example:
	//
	// Resource ALIYUN::ECS::VPC vpc-m5eauuq80anx59v28***	- could not be found for template scratch.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the resource scenario.
	Tags []*ListTemplateScratchesResponseBodyTemplateScratchesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the resource scenario.
	//
	// example:
	//
	// ts-48ad85d66cca4620****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the resource scenario. Valid values:
	//
	// 	- ResourceImport: resource management
	//
	// 	- ArchitectureReplication: resource replication
	//
	// example:
	//
	// ResourceImport
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
	// The time when the resource scenario was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-07T08:06:44
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratches) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratches) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetDescription() *string {
	return s.Description
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetFailedCode() *string {
	return s.FailedCode
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetPreferenceParameters() []*ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters {
	return s.PreferenceParameters
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetSourceResourceGroup() *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup {
	return s.SourceResourceGroup
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetSourceResources() []*ListTemplateScratchesResponseBodyTemplateScratchesSourceResources {
	return s.SourceResources
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetSourceTag() *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag {
	return s.SourceTag
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetStatus() *string {
	return s.Status
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetTags() []*ListTemplateScratchesResponseBodyTemplateScratchesTags {
	return s.Tags
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetTemplateScratchType() *string {
	return s.TemplateScratchType
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetCreateTime(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetDescription(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Description = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetFailedCode(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.FailedCode = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetLogicalIdStrategy(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.LogicalIdStrategy = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetPreferenceParameters(v []*ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.PreferenceParameters = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetResourceGroupId(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceResourceGroup(v *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceResourceGroup = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceResources(v []*ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceResources = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceTag(v *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceTag = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetStatus(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Status = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetStatusReason(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.StatusReason = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTags(v []*ListTemplateScratchesResponseBodyTemplateScratchesTags) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Tags = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTemplateScratchId(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.TemplateScratchId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTemplateScratchType(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.TemplateScratchType = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetUpdateTime(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.UpdateTime = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) Validate() error {
	if s.PreferenceParameters != nil {
		for _, item := range s.PreferenceParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceResourceGroup != nil {
		if err := s.SourceResourceGroup.Validate(); err != nil {
			return err
		}
	}
	if s.SourceResources != nil {
		for _, item := range s.SourceResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceTag != nil {
		if err := s.SourceTag.Validate(); err != nil {
			return err
		}
	}
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

type ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters struct {
	// The parameter name.
	//
	// example:
	//
	// DeletionPolicy
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// Retain
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) SetParameterKey(v string) *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) SetParameterValue(v string) *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters {
	s.ParameterValue = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) Validate() error {
	return dara.Validate(s)
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup struct {
	// The ID of the source resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource types for filtering resources.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) SetResourceGroupId(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) SetResourceTypeFilter(v []*string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) Validate() error {
	return dara.Validate(s)
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceResources struct {
	// The resource ID.
	//
	// example:
	//
	// vpc-m5eauuq80anx59v28****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) SetResourceId(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources {
	s.ResourceId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) SetResourceType(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources {
	s.ResourceType = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) Validate() error {
	return dara.Validate(s)
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceTag struct {
	// The source tags.
	//
	// example:
	//
	// {"a": "b"}
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The resource types for filtering resources.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) GetResourceTags() map[string]interface{} {
	return s.ResourceTags
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) SetResourceTags(v map[string]interface{}) *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag {
	s.ResourceTags = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) SetResourceTypeFilter(v []*string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag {
	s.ResourceTypeFilter = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) Validate() error {
	return dara.Validate(s)
}

type ListTemplateScratchesResponseBodyTemplateScratchesTags struct {
	// The tag key of the resource scenario.
	//
	// example:
	//
	// usage1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource scenario.
	//
	// example:
	//
	// test1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesTags) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesTags) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) GetKey() *string {
	return s.Key
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) GetValue() *string {
	return s.Value
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) SetKey(v string) *ListTemplateScratchesResponseBodyTemplateScratchesTags {
	s.Key = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) SetValue(v string) *ListTemplateScratchesResponseBodyTemplateScratchesTags {
	s.Value = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) Validate() error {
	return dara.Validate(s)
}
