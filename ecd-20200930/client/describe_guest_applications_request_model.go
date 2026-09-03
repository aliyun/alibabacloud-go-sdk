// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeGuestApplicationsRequest
	GetDesktopId() *string
	SetEndUserId(v string) *DescribeGuestApplicationsRequest
	GetEndUserId() *string
	SetRegionId(v string) *DescribeGuestApplicationsRequest
	GetRegionId() *string
}

type DescribeGuestApplicationsRequest struct {
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-asw3giklqvfqe****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGuestApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestApplicationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGuestApplicationsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeGuestApplicationsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeGuestApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGuestApplicationsRequest) SetDesktopId(v string) *DescribeGuestApplicationsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeGuestApplicationsRequest) SetEndUserId(v string) *DescribeGuestApplicationsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeGuestApplicationsRequest) SetRegionId(v string) *DescribeGuestApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGuestApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
