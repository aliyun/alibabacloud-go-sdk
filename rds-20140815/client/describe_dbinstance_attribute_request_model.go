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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// 	Notice: Do not query the details of multiple instances at a time by using multiple instance IDs. Otherwise, the query times out and fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether the instance expires. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
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
