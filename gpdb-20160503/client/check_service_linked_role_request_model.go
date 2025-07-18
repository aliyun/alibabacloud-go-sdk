// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckServiceLinkedRoleRequest
	GetRegionId() *string
}

type CheckServiceLinkedRoleRequest struct {
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
