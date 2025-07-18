// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DisableDBResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DisableDBResourceGroupRequest
	GetOwnerId() *int64
}

type DisableDBResourceGroupRequest struct {
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

func (s DisableDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableDBResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DisableDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableDBResourceGroupRequest) SetDBInstanceId(v string) *DisableDBResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DisableDBResourceGroupRequest) SetOwnerId(v int64) *DisableDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
