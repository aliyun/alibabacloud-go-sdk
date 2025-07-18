// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DeleteDBInstanceRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeleteDBInstanceRequest
	GetResourceGroupId() *string
}

type DeleteDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/327176.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88**********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceGroupId(v string) *DeleteDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
