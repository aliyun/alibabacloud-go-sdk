// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRCSecurityGroupPermissionRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *DescribeRCSecurityGroupPermissionRequest
	GetSecurityGroupId() *string
}

type DescribeRCSecurityGroupPermissionRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DescribeRCSecurityGroupPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCSecurityGroupPermissionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCSecurityGroupPermissionRequest) SetRegionId(v string) *DescribeRCSecurityGroupPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionRequest) SetSecurityGroupId(v string) *DescribeRCSecurityGroupPermissionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCSecurityGroupPermissionRequest) Validate() error {
	return dara.Validate(s)
}
