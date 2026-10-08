// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetExpired(v string) *DescribeDBInstanceAttributeRequest
	GetExpired() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceAttributeRequest struct {
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// 	Warning: Do not specify multiple instance IDs for batch queries. Otherwise, the query times out and fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The expiration status of the instance. Valid values:
	//
	// 	- **True**: The instance has expired.
	//
	// 	- **False**: The instance has not expired.
	//
	// example:
	//
	// False
	Expired         *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetExpired(v string) *DescribeDBInstanceAttributeRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
