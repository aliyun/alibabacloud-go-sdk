// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAccountRoleRequest
	GetRegionId() *string
}

type DescribeAccountRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountRoleRequest) SetRegionId(v string) *DescribeAccountRoleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
