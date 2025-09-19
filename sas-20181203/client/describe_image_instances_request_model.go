// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeImageInstancesRequest
	GetCurrentPage() *int32
	SetLogicalExp(v string) *DescribeImageInstancesRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *DescribeImageInstancesRequest
	GetPageSize() *int32
	SetScanned(v bool) *DescribeImageInstancesRequest
	GetScanned() *bool
}

type DescribeImageInstancesRequest struct {
	// The search condition that is used to filter the server. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **name**: the search condition
	//
	// 	- **name**: the value of the search condition
	//
	// 	- **logicalExp**: the logical relation for multiple search conditions Valid values:
	//
	//     	- **OR**: The search conditions use a logical **OR**.
	//
	//     	- **AND**: The search conditions use a logical **AND**.
	//
	// > You can call the [DescribeImageCriteria](https://help.aliyun.com/document_detail/471822.html) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"instanceId","value":"390100182","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The logical relationship that you want to use to evaluate multiple search conditions. Valid values:
	//
	// 	- **OR**: The search conditions are evaluated by using a logical **OR**.
	//
	// 	- **AND**: The search conditions are evaluated by using a logical **AND**.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > : We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the image is scanned. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Scanned *bool `json:"Scanned,omitempty" xml:"Scanned,omitempty"`
}

func (s DescribeImageInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeImageInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageInstancesRequest) GetScanned() *bool {
	return s.Scanned
}

func (s *DescribeImageInstancesRequest) SetCriteria(v string) *DescribeImageInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageInstancesRequest) SetCurrentPage(v int32) *DescribeImageInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageInstancesRequest) SetLogicalExp(v string) *DescribeImageInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeImageInstancesRequest) SetPageSize(v int32) *DescribeImageInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageInstancesRequest) SetScanned(v bool) *DescribeImageInstancesRequest {
	s.Scanned = &v
	return s
}

func (s *DescribeImageInstancesRequest) Validate() error {
	return dara.Validate(s)
}
