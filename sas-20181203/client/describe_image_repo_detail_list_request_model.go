// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoDetailListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageRepoDetailListRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *DescribeImageRepoDetailListRequest
	GetCurrentPage() *int32
	SetLogicalExp(v string) *DescribeImageRepoDetailListRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *DescribeImageRepoDetailListRequest
	GetPageSize() *int32
}

type DescribeImageRepoDetailListRequest struct {
	// The search conditions for assets. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **name**: the search condition.
	//
	// 	- **name**: the value of the search condition.
	//
	// 	- **logicalExp**: the logical relation for multiple search conditions. Valid values:
	//
	//     	- **OR**: The search conditions use a logical **OR**.
	//
	//     	- **AND**: The search conditions use a logical **AND**.
	//
	// > You can call the [DescribeImageRepoCriteria](~~DescribeImageRepoCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{"name":"vulStatus","value":"YES","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
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
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeImageRepoDetailListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoDetailListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoDetailListRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageRepoDetailListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageRepoDetailListRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *DescribeImageRepoDetailListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageRepoDetailListRequest) SetCriteria(v string) *DescribeImageRepoDetailListRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageRepoDetailListRequest) SetCurrentPage(v int32) *DescribeImageRepoDetailListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageRepoDetailListRequest) SetLogicalExp(v string) *DescribeImageRepoDetailListRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeImageRepoDetailListRequest) SetPageSize(v int32) *DescribeImageRepoDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageRepoDetailListRequest) Validate() error {
	return dara.Validate(s)
}
