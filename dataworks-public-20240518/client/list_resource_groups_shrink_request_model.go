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
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The list of Alibaba Cloud tags.
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// The name of the resource group. Fuzzy search is supported.
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
	// The page size.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the resource group. Valid values include:
	//
	// - `PrePaid`: subscription.
	//
	// - `PostPaid`: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The types of the resource groups to query. **If this parameter is not specified, general-purpose resource groups are queried by default.**
	ResourceGroupTypesShrink *string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty"`
	// The sorting criterion for the results. The format is `FieldName SortOrder`. `SortOrder` can be `Asc` (ascending) or `Desc` (descending). If you do not specify `SortOrder`, the default is `Asc`. The following fields are supported:
	//
	// - `Id`: Resource group ID
	//
	// - `Name`: Resource group name
	//
	// - `Remark`: Resource group remarks
	//
	// - `Type`: Resource group type
	//
	// - `Status`: Resource group status
	//
	// - `Spec`: Resource group specifications
	//
	// - `CreateUser`: The user who created the resource group
	//
	// - `CreateTime`: The time when the resource group was created
	//
	// Default value: `CreateTime Asc`
	//
	// example:
	//
	// CreateTime Asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The statuses of the resource groups to query.
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
