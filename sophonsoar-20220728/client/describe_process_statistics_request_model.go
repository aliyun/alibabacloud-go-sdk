// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeProcessStatisticsRequest
	GetLang() *string
	SetRoleFor(v string) *DescribeProcessStatisticsRequest
	GetRoleFor() *string
	SetRoleType(v string) *DescribeProcessStatisticsRequest
	GetRoleType() *string
}

type DescribeProcessStatisticsRequest struct {
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 1709821xxxxx3093
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0 (default): the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeProcessStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeProcessStatisticsRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *DescribeProcessStatisticsRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeProcessStatisticsRequest) SetLang(v string) *DescribeProcessStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) SetRoleFor(v string) *DescribeProcessStatisticsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) SetRoleType(v string) *DescribeProcessStatisticsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
