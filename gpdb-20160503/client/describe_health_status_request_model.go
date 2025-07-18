// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeHealthStatusRequest
	GetDBInstanceId() *string
	SetKey(v string) *DescribeHealthStatusRequest
	GetKey() *string
}

type DescribeHealthStatusRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The performance metric that you want to query. Separate multiple values with commas (,). For more information, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// node_master_status,node_master_connection_status,node_segment_connection_status,node_segment_disk_status
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHealthStatusRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeHealthStatusRequest) SetDBInstanceId(v string) *DescribeHealthStatusRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHealthStatusRequest) SetKey(v string) *DescribeHealthStatusRequest {
	s.Key = &v
	return s
}

func (s *DescribeHealthStatusRequest) Validate() error {
	return dara.Validate(s)
}
