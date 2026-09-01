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
	SetProcessActionEnd(v int64) *DescribeProcessStatisticsRequest
	GetProcessActionEnd() *int64
	SetProcessActionStart(v int64) *DescribeProcessStatisticsRequest
	GetProcessActionStart() *int64
	SetRoleFor(v string) *DescribeProcessStatisticsRequest
	GetRoleFor() *string
	SetRoleType(v string) *DescribeProcessStatisticsRequest
	GetRoleType() *string
}

type DescribeProcessStatisticsRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The end time of the query for response tasks. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 17561XXX77435
	ProcessActionEnd *int64 `json:"ProcessActionEnd,omitempty" xml:"ProcessActionEnd,omitempty"`
	// The start time of the query for response tasks. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 17000XXX83572
	ProcessActionStart *int64 `json:"ProcessActionStart,omitempty" xml:"ProcessActionStart,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 1709821xxxxx3093
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0 (default): the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
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

func (s *DescribeProcessStatisticsRequest) GetProcessActionEnd() *int64 {
	return s.ProcessActionEnd
}

func (s *DescribeProcessStatisticsRequest) GetProcessActionStart() *int64 {
	return s.ProcessActionStart
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

func (s *DescribeProcessStatisticsRequest) SetProcessActionEnd(v int64) *DescribeProcessStatisticsRequest {
	s.ProcessActionEnd = &v
	return s
}

func (s *DescribeProcessStatisticsRequest) SetProcessActionStart(v int64) *DescribeProcessStatisticsRequest {
	s.ProcessActionStart = &v
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
