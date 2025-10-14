// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarClassListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeColumnarClassListRequest
	GetDBInstanceName() *string
	SetInstanceName(v string) *DescribeColumnarClassListRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeColumnarClassListRequest
	GetRegionId() *string
}

type DescribeColumnarClassListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// pxc-shrvdc****wtf5uk-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeColumnarClassListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarClassListRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnarClassListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeColumnarClassListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeColumnarClassListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColumnarClassListRequest) SetDBInstanceName(v string) *DescribeColumnarClassListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColumnarClassListRequest) SetInstanceName(v string) *DescribeColumnarClassListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeColumnarClassListRequest) SetRegionId(v string) *DescribeColumnarClassListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnarClassListRequest) Validate() error {
	return dara.Validate(s)
}
