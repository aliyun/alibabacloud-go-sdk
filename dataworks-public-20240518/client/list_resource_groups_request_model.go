// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *ListResourceGroupsRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTags(v []*ListResourceGroupsRequestAliyunResourceTags) *ListResourceGroupsRequest
	GetAliyunResourceTags() []*ListResourceGroupsRequestAliyunResourceTags
	SetName(v string) *ListResourceGroupsRequest
	GetName() *string
	SetPageNumber(v int32) *ListResourceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListResourceGroupsRequest
	GetPaymentType() *string
	SetProjectId(v int64) *ListResourceGroupsRequest
	GetProjectId() *int64
	SetResourceGroupTypes(v []*string) *ListResourceGroupsRequest
	GetResourceGroupTypes() []*string
	SetSortBy(v string) *ListResourceGroupsRequest
	GetSortBy() *string
	SetStatuses(v []*string) *ListResourceGroupsRequest
	GetStatuses() []*string
}

type ListResourceGroupsRequest struct {
	// Alibaba Cloud Resource Group ID
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// Alibaba Cloud tag list
	AliyunResourceTags []*ListResourceGroupsRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The name of a resource group, which is used for fuzzy match.
	//
	// example:
	//
	// Resource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of resource groups. Valid values:
	//
	// 	- PrePaid
	//
	// 	- PostPaid
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The types of resource groups to query. If you do not configure this parameter, only serverless resource groups are returned by default.
	ResourceGroupTypes []*string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty" type:"Repeated"`
	// The fields used for sorting. Fields such as TriggerTime and StartedTime are supported. The value of this parameter is in the Sort field + Sort by (Desc/Asc) format. By default, results are sorted in ascending order. Valid values:
	//
	// 	- Id (Desc/Asc): the resource group ID
	//
	// 	- Name (Desc/Asc): the name of the resource group
	//
	// 	- Remark (Desc/Asc): the remarks of the resource group
	//
	// 	- Type (Desc/Asc): the type of the resource group
	//
	// 	- Status (Desc/Asc): the status of the resource group
	//
	// 	- Spec (Desc/Asc): the specifications of the resource group
	//
	// 	- CreateUser (Desc/Asc): the creator of the resource group
	//
	// 	- CreateTime (Desc/Asc): the time when the resource group is created
	//
	// Default value: CreateTime Asc
	//
	// example:
	//
	// CreateTime Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The statuses of resource groups.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListResourceGroupsRequest) GetAliyunResourceTags() []*ListResourceGroupsRequestAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *ListResourceGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListResourceGroupsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListResourceGroupsRequest) GetResourceGroupTypes() []*string {
	return s.ResourceGroupTypes
}

func (s *ListResourceGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceGroupsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListResourceGroupsRequest) SetAliyunResourceGroupId(v string) *ListResourceGroupsRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetAliyunResourceTags(v []*ListResourceGroupsRequestAliyunResourceTags) *ListResourceGroupsRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageNumber(v int32) *ListResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPageSize(v int32) *ListResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPaymentType(v string) *ListResourceGroupsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetProjectId(v int64) *ListResourceGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupTypes(v []*string) *ListResourceGroupsRequest {
	s.ResourceGroupTypes = v
	return s
}

func (s *ListResourceGroupsRequest) SetSortBy(v string) *ListResourceGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupsRequest) SetStatuses(v []*string) *ListResourceGroupsRequest {
	s.Statuses = v
	return s
}

func (s *ListResourceGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsRequestAliyunResourceTags struct {
	// Tag Key
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag Value
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsRequestAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequestAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsRequestAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsRequestAliyunResourceTags) SetKey(v string) *ListResourceGroupsRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsRequestAliyunResourceTags) SetValue(v string) *ListResourceGroupsRequestAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsRequestAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
