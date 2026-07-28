// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeProfileRequest
	GetDBInstanceId() *string
	SetQueryId(v string) *DescribeProfileRequest
	GetQueryId() *string
	SetRegionId(v string) *DescribeProfileRequest
	GetRegionId() *string
}

type DescribeProfileRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The query ID.
	//
	// example:
	//
	// sq202510231018sh3b69ad10014154
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProfileRequest) GoString() string {
	return s.String()
}

func (s *DescribeProfileRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeProfileRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeProfileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProfileRequest) SetDBInstanceId(v string) *DescribeProfileRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeProfileRequest) SetQueryId(v string) *DescribeProfileRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeProfileRequest) SetRegionId(v string) *DescribeProfileRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProfileRequest) Validate() error {
	return dara.Validate(s)
}
