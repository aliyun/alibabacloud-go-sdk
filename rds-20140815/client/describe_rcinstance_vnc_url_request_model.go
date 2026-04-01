// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceVncUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *DescribeRCInstanceVncUrlRequest
	GetDbType() *string
	SetInstanceId(v string) *DescribeRCInstanceVncUrlRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRCInstanceVncUrlRequest
	GetRegionId() *string
}

type DescribeRCInstanceVncUrlRequest struct {
	// The database engine. Valid values:
	//
	// 	- **mssql**: SQL Server
	//
	// 	- **mysql**: MySQL
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-e6e3757b8px27oa5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceVncUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceVncUrlRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeRCInstanceVncUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceVncUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceVncUrlRequest) SetDbType(v string) *DescribeRCInstanceVncUrlRequest {
	s.DbType = &v
	return s
}

func (s *DescribeRCInstanceVncUrlRequest) SetInstanceId(v string) *DescribeRCInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceVncUrlRequest) SetRegionId(v string) *DescribeRCInstanceVncUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceVncUrlRequest) Validate() error {
	return dara.Validate(s)
}
