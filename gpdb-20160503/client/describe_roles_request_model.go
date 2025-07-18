// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeRolesRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeRolesRequest
	GetOwnerId() *int64
}

type DescribeRolesRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRolesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRolesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRolesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRolesRequest) SetDBInstanceId(v string) *DescribeRolesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRolesRequest) SetOwnerId(v int64) *DescribeRolesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRolesRequest) Validate() error {
	return dara.Validate(s)
}
