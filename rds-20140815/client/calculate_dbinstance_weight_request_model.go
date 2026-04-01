// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateDBInstanceWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CalculateDBInstanceWeightRequest
	GetDBInstanceId() *string
	SetResourceOwnerId(v int64) *CalculateDBInstanceWeightRequest
	GetResourceOwnerId() *int64
}

type CalculateDBInstanceWeightRequest struct {
	// The primary instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CalculateDBInstanceWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s CalculateDBInstanceWeightRequest) GoString() string {
	return s.String()
}

func (s *CalculateDBInstanceWeightRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CalculateDBInstanceWeightRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CalculateDBInstanceWeightRequest) SetDBInstanceId(v string) *CalculateDBInstanceWeightRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CalculateDBInstanceWeightRequest) SetResourceOwnerId(v int64) *CalculateDBInstanceWeightRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CalculateDBInstanceWeightRequest) Validate() error {
	return dara.Validate(s)
}
