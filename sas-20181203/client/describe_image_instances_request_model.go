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
	// Sets the conditions for searching assets. This parameter is in JSON format and contains the following fields:
	//
	// - **name**: The search item.
	//
	// - **value**: The value of the search item.
	//
	// - **logicalExp**: The logical relationship between multiple search item values. Valid values:
	//
	//     - **OR**: The multiple search item values are in an **OR*	- relationship.
	//
	//     - **AND**: The multiple search item values are in an **AND*	- relationship.
	//
	// > You can call the [DescribeImageRepoCriteria](~~DescribeImageRepoCriteria~~) operation to query supported search conditions.
	//
	// example:
	//
	// [{"name":"instanceId","value":"390100182","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the returned results to start displaying. The starting value is **1**. Default value: **1**, which indicates that page 1 is displayed.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Sets the logical relationship between multiple search conditions. Valid values:
	//
	// - **OR**: The multiple search conditions are in an **OR*	- relationship.
	//
	// - **AND**: The multiple search conditions are in an **AND*	- relationship.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The maximum number of entries to return on each page during a paged query. The default number of entries per page is 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > We recommend that you do not leave the PageSize parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the image has been scanned. Valid values:
	//
	// - **true**: processed.
	//
	// - **false**: not processed.
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
