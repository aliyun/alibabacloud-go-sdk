// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeUsersPasswordRequest
	GetDesktopId() *string
	SetRegionId(v string) *DescribeUsersPasswordRequest
	GetRegionId() *string
}

type DescribeUsersPasswordRequest struct {
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gq8u6whi9f6k8****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUsersPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersPasswordRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersPasswordRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeUsersPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUsersPasswordRequest) SetDesktopId(v string) *DescribeUsersPasswordRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeUsersPasswordRequest) SetRegionId(v string) *DescribeUsersPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUsersPasswordRequest) Validate() error {
	return dara.Validate(s)
}
