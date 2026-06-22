// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedContainerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeGroupedContainerInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeGroupedContainerInstancesRequest
	GetCurrentPage() *int32
	SetFieldValue(v string) *DescribeGroupedContainerInstancesRequest
	GetFieldValue() *string
	SetGroupField(v string) *DescribeGroupedContainerInstancesRequest
	GetGroupField() *string
	SetLogicalExp(v string) *DescribeGroupedContainerInstancesRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *DescribeGroupedContainerInstancesRequest
	GetPageSize() *int32
}

type DescribeGroupedContainerInstancesRequest struct {
	// The conditions for searching assets. This parameter is in JSON format. Separate multiple conditions with commas (,). Example: `[{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]`.
	//
	// > You can search for assets by instance ID, instance name, VPC ID, region, public IP address, and other conditions. Call [DescribeCriteria](~~DescribeCriteria~~) to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates that the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The search condition for the specified group type. Set the search condition based on the type specified by GroupField:
	//
	// - If **GroupField*	- is set to **pod**: specify the pod name to query.
	//
	// - If **GroupField*	- is set to **appName**: specify the application name to query.
	//
	// - If **GroupField*	- is set to **namespace**: specify the namespace to query.
	//
	// - If **GroupField*	- is set to **clusterId**: specify the cluster ID to query.
	//
	// - If **GroupField*	- is set to **image**: specify the image name to query.
	//
	// > All the preceding search conditions support fuzzy match.
	//
	// example:
	//
	// cas-adad-qeqwe
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The group type to query. Valid values:
	//
	// - **pod**: pod
	//
	// - **appName**: application name
	//
	// - **namespace**: namespace
	//
	// - **clusterId**: cluster ID
	//
	// - **image**: image.
	//
	// This parameter is required.
	//
	// example:
	//
	// pod
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The logical relationship among multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are evaluated with a logical OR.
	//
	// - **AND**: The search conditions are evaluated with a logical AND.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of container assets to display on each page when paging is used. Default value: **20**, which indicates that 20 container assets are displayed on each page.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGroupedContainerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedContainerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeGroupedContainerInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedContainerInstancesRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeGroupedContainerInstancesRequest) GetGroupField() *string {
	return s.GroupField
}

func (s *DescribeGroupedContainerInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeGroupedContainerInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedContainerInstancesRequest) SetCriteria(v string) *DescribeGroupedContainerInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetCurrentPage(v int32) *DescribeGroupedContainerInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetFieldValue(v string) *DescribeGroupedContainerInstancesRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetGroupField(v string) *DescribeGroupedContainerInstancesRequest {
	s.GroupField = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetLogicalExp(v string) *DescribeGroupedContainerInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetPageSize(v int32) *DescribeGroupedContainerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) Validate() error {
	return dara.Validate(s)
}
