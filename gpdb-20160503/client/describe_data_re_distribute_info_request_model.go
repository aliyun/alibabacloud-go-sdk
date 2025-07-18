// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataReDistributeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDataReDistributeInfoRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeDataReDistributeInfoRequest
	GetOwnerId() *int64
}

type DescribeDataReDistributeInfoRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeDataReDistributeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataReDistributeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDataReDistributeInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDataReDistributeInfoRequest) SetDBInstanceId(v string) *DescribeDataReDistributeInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataReDistributeInfoRequest) SetOwnerId(v int64) *DescribeDataReDistributeInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDataReDistributeInfoRequest) Validate() error {
	return dara.Validate(s)
}
