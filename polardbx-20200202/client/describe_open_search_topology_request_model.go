// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchTopologyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchTopologyRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchTopologyRequest
	GetRegionId() *string
}

type DescribeOpenSearchTopologyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchTopologyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchTopologyRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchTopologyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchTopologyRequest) SetDBInstanceName(v string) *DescribeOpenSearchTopologyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchTopologyRequest) SetRegionId(v string) *DescribeOpenSearchTopologyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchTopologyRequest) Validate() error {
	return dara.Validate(s)
}
