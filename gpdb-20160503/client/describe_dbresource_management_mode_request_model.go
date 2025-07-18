// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceManagementModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBResourceManagementModeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDBResourceManagementModeRequest
	GetOwnerId() *int64
}

type DescribeDBResourceManagementModeRequest struct {
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

func (s DescribeDBResourceManagementModeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceManagementModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceManagementModeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBResourceManagementModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBResourceManagementModeRequest) SetDBInstanceId(v string) *DescribeDBResourceManagementModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBResourceManagementModeRequest) SetOwnerId(v int64) *DescribeDBResourceManagementModeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourceManagementModeRequest) Validate() error {
	return dara.Validate(s)
}
