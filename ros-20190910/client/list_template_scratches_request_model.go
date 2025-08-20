// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateScratchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplateScratchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplateScratchesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTemplateScratchesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTemplateScratchesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListTemplateScratchesRequest
	GetStatus() *string
	SetTags(v []*ListTemplateScratchesRequestTags) *ListTemplateScratchesRequest
	GetTags() []*ListTemplateScratchesRequestTags
	SetTemplateScratchId(v string) *ListTemplateScratchesRequest
	GetTemplateScratchId() *string
	SetTemplateScratchType(v string) *ListTemplateScratchesRequest
	GetTemplateScratchType() *string
}

type ListTemplateScratchesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the scenario. Valid values:
	//
	// 	- GENERATE_IN_PROGRESS: The scenario is being created.
	//
	// 	- GENERATE_COMPLETE: The scenario is created.
	//
	// 	- GENERATE_FAILED: The scenario fails to be created.
	//
	// example:
	//
	// ["GENERATE_COMPLETE"]
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the scenario.
	Tags []*ListTemplateScratchesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the resource scenario. Valid values:
	//
	// 	- ArchitectureReplication: resource replication
	//
	// 	- ArchitectureDetection: resource detection
	//
	// 	- ResourceImport: resource management
	//
	// 	- ResourceMigration: resource migration
	//
	// example:
	//
	// ArchitectureReplication
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
}

func (s ListTemplateScratchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplateScratchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplateScratchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTemplateScratchesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateScratchesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTemplateScratchesRequest) GetTags() []*ListTemplateScratchesRequestTags {
	return s.Tags
}

func (s *ListTemplateScratchesRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *ListTemplateScratchesRequest) GetTemplateScratchType() *string {
	return s.TemplateScratchType
}

func (s *ListTemplateScratchesRequest) SetPageNumber(v int32) *ListTemplateScratchesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetPageSize(v int32) *ListTemplateScratchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetRegionId(v string) *ListTemplateScratchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetResourceGroupId(v string) *ListTemplateScratchesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetStatus(v string) *ListTemplateScratchesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetTags(v []*ListTemplateScratchesRequestTags) *ListTemplateScratchesRequest {
	s.Tags = v
	return s
}

func (s *ListTemplateScratchesRequest) SetTemplateScratchId(v string) *ListTemplateScratchesRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetTemplateScratchType(v string) *ListTemplateScratchesRequest {
	s.TemplateScratchType = &v
	return s
}

func (s *ListTemplateScratchesRequest) Validate() error {
	return dara.Validate(s)
}

type ListTemplateScratchesRequestTags struct {
	// The tag key of the scenario.
	//
	// > Tags is optional. If you want to specify Tags, you must specify Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scenario.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateScratchesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateScratchesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTemplateScratchesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTemplateScratchesRequestTags) SetKey(v string) *ListTemplateScratchesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTemplateScratchesRequestTags) SetValue(v string) *ListTemplateScratchesRequestTags {
	s.Value = &v
	return s
}

func (s *ListTemplateScratchesRequestTags) Validate() error {
	return dara.Validate(s)
}
