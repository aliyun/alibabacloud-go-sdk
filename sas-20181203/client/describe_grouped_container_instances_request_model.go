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
	// The search conditions for assets. Specify the value in the JSON format. Separate multiple search conditions with commas (,). Example: `[{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]`.
	//
	// >  Supported search conditions include the instance ID, instance name, virtual private cloud (VPC) ID, region, and public IP address. You can call the [DescribeCriteria](~~DescribeCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"riskStatus","value":"YES"},{"name":"riskLevel","value":"2"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The keyword that you want to use to query containers. This parameter depends on the value of the GroupField parameter.
	//
	// 	- If the **GroupField*	- parameter is set to **pod**, set this parameter to the name of the pod that you want to query.
	//
	// 	- If the **GroupField*	- parameter is set to **appName**, set this parameter to the name of the application that you want to query.
	//
	// 	- If the **GroupField*	- parameter is set to **namespace**, set this parameter to the namespace that you want to query.
	//
	// 	- If the **GroupField*	- parameter is set to **clusterId**, set this parameter to the ID of the cluster that you want to query.
	//
	// 	- If the **GroupField*	- parameter is set to **image**, set this parameter to the name of the image that you want to query.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// cas-adad-qeqwe
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The group type that you want to use to query containers. Valid values:
	//
	// 	- **pod**
	//
	// 	- **appName**
	//
	// 	- **namespace**
	//
	// 	- **clusterId**
	//
	// 	- **image**
	//
	// This parameter is required.
	//
	// example:
	//
	// pod
	GroupField *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	// The logical relationship that you want to use to evaluate multiple search conditions. Valid values:
	//
	// 	- **OR**: Search conditions are evaluated by using a logical **OR**.
	//
	// 	- **AND**: Search conditions are evaluated by using a logical **AND**.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// >  We recommend that you do not leave this parameter empty.
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
