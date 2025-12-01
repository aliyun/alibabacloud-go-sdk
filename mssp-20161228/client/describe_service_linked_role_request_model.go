// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeServiceLinkedRoleRequest
	GetLang() *string
	SetRegionId(v string) *DescribeServiceLinkedRoleRequest
	GetRegionId() *string
}

type DescribeServiceLinkedRoleRequest struct {
	// Language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceLinkedRoleRequest) SetLang(v string) *DescribeServiceLinkedRoleRequest {
	s.Lang = &v
	return s
}

func (s *DescribeServiceLinkedRoleRequest) SetRegionId(v string) *DescribeServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
