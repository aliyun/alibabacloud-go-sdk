// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDrivePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCloudDrivePermissionsRequest
	GetCdsId() *string
	SetRegionId(v string) *DescribeCloudDrivePermissionsRequest
	GetRegionId() *string
}

type DescribeCloudDrivePermissionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-82414*****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudDrivePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDrivePermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDrivePermissionsRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCloudDrivePermissionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudDrivePermissionsRequest) SetCdsId(v string) *DescribeCloudDrivePermissionsRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCloudDrivePermissionsRequest) SetRegionId(v string) *DescribeCloudDrivePermissionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudDrivePermissionsRequest) Validate() error {
	return dara.Validate(s)
}
