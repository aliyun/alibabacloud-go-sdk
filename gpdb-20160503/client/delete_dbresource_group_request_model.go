// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DeleteDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *DeleteDBResourceGroupRequest
	GetResourceGroupName() *string
}

type DeleteDBResourceGroupRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DeleteDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBResourceGroupRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DeleteDBResourceGroupRequest) SetDBInstanceId(v string) *DeleteDBResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetOwnerId(v int64) *DeleteDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetResourceGroupName(v string) *DeleteDBResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
