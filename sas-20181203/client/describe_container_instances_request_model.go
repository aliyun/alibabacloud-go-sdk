// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeContainerInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeContainerInstancesRequest
	GetCurrentPage() *int32
	SetLogicalExp(v string) *DescribeContainerInstancesRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *DescribeContainerInstancesRequest
	GetPageSize() *int32
}

type DescribeContainerInstancesRequest struct {
	// The search conditions that are used to filter containers. The value of this parameter is in the JSON format and is case-sensitive. The value contains the following fields:
	//
	// 	- **name**: the search condition.
	//
	// 	- **name**: the value of the search condition.
	//
	// 	- **logicalExp**: the logical relationship among multiple search conditions. Valid values:
	//
	//     	- **OR**: Search conditions are evaluated by using a logical **OR**.
	//
	//     	- **AND**: Search conditions are evaluated by using a logical **AND**.
	//
	// > You can use search conditions such as the container ID, cluster ID, cluster name, cluster type, risk level, and region. You can call the [DescribeContainerCriteria](~~DescribeContainerCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"flag","value":"0|8","logicalExp":"AND"},{"name":"ecsType","value":"!8","logicalExp":"AND"}][{"name":"clusterType","value":"NotManagedKubernetes","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The logical operator that you want to use to evaluate multiple search conditions. Valid values:
	//
	// 	- **OR**: Search conditions are evaluated by using a logical **OR**.
	//
	// 	- **AND**: Search conditions are evaluated by using a logical **AND**.
	//
	// example:
	//
	// AND
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeContainerInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeContainerInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeContainerInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeContainerInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerInstancesRequest) SetCriteria(v string) *DescribeContainerInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeContainerInstancesRequest) SetCurrentPage(v int32) *DescribeContainerInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeContainerInstancesRequest) SetLogicalExp(v string) *DescribeContainerInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeContainerInstancesRequest) SetPageSize(v int32) *DescribeContainerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerInstancesRequest) Validate() error {
	return dara.Validate(s)
}
