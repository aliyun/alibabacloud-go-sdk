// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEndpointListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDeleteCN(v bool) *DescribeCustomEndpointListRequest
	GetCheckDeleteCN() *bool
	SetCustomEndpointIds(v string) *DescribeCustomEndpointListRequest
	GetCustomEndpointIds() *string
	SetDBInstanceName(v string) *DescribeCustomEndpointListRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeCustomEndpointListRequest
	GetRegionId() *string
}

type DescribeCustomEndpointListRequest struct {
	// example:
	//
	// true
	CheckDeleteCN *bool `json:"CheckDeleteCN,omitempty" xml:"CheckDeleteCN,omitempty"`
	// example:
	//
	// pxe-b6e****no4pfap1s
	CustomEndpointIds *string `json:"CustomEndpointIds,omitempty" xml:"CustomEndpointIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCustomEndpointListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEndpointListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEndpointListRequest) GetCheckDeleteCN() *bool {
	return s.CheckDeleteCN
}

func (s *DescribeCustomEndpointListRequest) GetCustomEndpointIds() *string {
	return s.CustomEndpointIds
}

func (s *DescribeCustomEndpointListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCustomEndpointListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomEndpointListRequest) SetCheckDeleteCN(v bool) *DescribeCustomEndpointListRequest {
	s.CheckDeleteCN = &v
	return s
}

func (s *DescribeCustomEndpointListRequest) SetCustomEndpointIds(v string) *DescribeCustomEndpointListRequest {
	s.CustomEndpointIds = &v
	return s
}

func (s *DescribeCustomEndpointListRequest) SetDBInstanceName(v string) *DescribeCustomEndpointListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCustomEndpointListRequest) SetRegionId(v string) *DescribeCustomEndpointListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomEndpointListRequest) Validate() error {
	return dara.Validate(s)
}
