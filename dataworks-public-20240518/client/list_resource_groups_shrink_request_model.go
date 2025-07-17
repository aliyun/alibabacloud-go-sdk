// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *ListResourceGroupsShrinkRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTagsShrink(v string) *ListResourceGroupsShrinkRequest
	GetAliyunResourceTagsShrink() *string
	SetName(v string) *ListResourceGroupsShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListResourceGroupsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceGroupsShrinkRequest
	GetPageSize() *int32
	SetPaymentType(v string) *ListResourceGroupsShrinkRequest
	GetPaymentType() *string
	SetProjectId(v int64) *ListResourceGroupsShrinkRequest
	GetProjectId() *int64
	SetResourceGroupTypesShrink(v string) *ListResourceGroupsShrinkRequest
	GetResourceGroupTypesShrink() *string
	SetSortBy(v string) *ListResourceGroupsShrinkRequest
	GetSortBy() *string
	SetStatusesShrink(v string) *ListResourceGroupsShrinkRequest
	GetStatusesShrink() *string
}

type ListResourceGroupsShrinkRequest struct {
	// Alibaba Cloud Resource Group ID
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// Alibaba Cloud tag list
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
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
	ResourceGroupTypesShrink *string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty"`
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
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListResourceGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsShrinkRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListResourceGroupsShrinkRequest) GetAliyunResourceTagsShrink() *string {
	return s.AliyunResourceTagsShrink
}

func (s *ListResourceGroupsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListResourceGroupsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListResourceGroupsShrinkRequest) GetResourceGroupTypesShrink() *string {
	return s.ResourceGroupTypesShrink
}

func (s *ListResourceGroupsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListResourceGroupsShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListResourceGroupsShrinkRequest) SetAliyunResourceGroupId(v string) *ListResourceGroupsShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetAliyunResourceTagsShrink(v string) *ListResourceGroupsShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetName(v string) *ListResourceGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetPageNumber(v int32) *ListResourceGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetPageSize(v int32) *ListResourceGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetPaymentType(v string) *ListResourceGroupsShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetProjectId(v int64) *ListResourceGroupsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetResourceGroupTypesShrink(v string) *ListResourceGroupsShrinkRequest {
	s.ResourceGroupTypesShrink = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetSortBy(v string) *ListResourceGroupsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetStatusesShrink(v string) *ListResourceGroupsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
