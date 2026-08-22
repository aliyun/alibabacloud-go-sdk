// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchNodesRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchNodesRequest
	GetRegionId() *string
}

type DescribeOpenSearchNodesRequest struct {
	// The name of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchNodesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchNodesRequest) SetDBInstanceName(v string) *DescribeOpenSearchNodesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchNodesRequest) SetRegionId(v string) *DescribeOpenSearchNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchNodesRequest) Validate() error {
	return dara.Validate(s)
}
