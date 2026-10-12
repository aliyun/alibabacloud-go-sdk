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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The Alibaba Cloud tags.
	AliyunResourceTags []*ListResourceGroupsRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The name of the resource group for fuzzy match.
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
	// The billing method of the resource group. Valid values:
	//
	// - PrePaid: Subscription.
	//
	// - PostPaid: Pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the workspace to which the resource group belongs.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The types of the resource groups to query. **If this parameter is left empty, general-purpose resource groups are queried by default.**
	ResourceGroupTypes []*string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty" type:"Repeated"`
	// The fields by which to sort the results. You can sort the results by fields such as scheduled time and start time. The format is \\"Sort field+Sort order (Desc/Asc)\\". Asc can be omitted. Valid values for the sort field:
	//
	// - Id (Desc/Asc): The ID of the resource group.
	//
	// - Name (Desc/Asc): The name of the resource group.
	//
	// - Remark (Desc/Asc): The description of the resource group.
	//
	// - Type (Desc/Asc): The type of the resource group.
	//
	// - Status (Desc/Asc): The status of the resource group.
	//
	// - Spec (Desc/Asc): The specifications of the resource group.
	//
	// - CreateUser (Desc/Asc): The creator of the resource group.
	//
	// - CreateTime (Desc/Asc): The time when the resource group was created.
	//
	// Default value: CreateTime Asc.
	//
	// example:
	//
	// CreateTime Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The statuses of the resource groups to query.
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
	if s.AliyunResourceTags != nil {
		for _, item := range s.AliyunResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsRequestAliyunResourceTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
