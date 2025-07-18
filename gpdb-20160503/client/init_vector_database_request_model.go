// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitVectorDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *InitVectorDatabaseRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *InitVectorDatabaseRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *InitVectorDatabaseRequest
	GetManagerAccountPassword() *string
	SetOwnerId(v int64) *InitVectorDatabaseRequest
	GetOwnerId() *int64
	SetRegionId(v string) *InitVectorDatabaseRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *InitVectorDatabaseRequest
	GetWorkspaceId() *string
}

type InitVectorDatabaseRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the manager account that has the rds_superuser permission.
	//
	// >  You can create an account on the Account Management page of the AnalyticDB for PostgreSQL console or by calling the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InitVectorDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s InitVectorDatabaseRequest) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *InitVectorDatabaseRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *InitVectorDatabaseRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *InitVectorDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InitVectorDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitVectorDatabaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InitVectorDatabaseRequest) SetDBInstanceId(v string) *InitVectorDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetManagerAccount(v string) *InitVectorDatabaseRequest {
	s.ManagerAccount = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetManagerAccountPassword(v string) *InitVectorDatabaseRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetOwnerId(v int64) *InitVectorDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetRegionId(v string) *InitVectorDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *InitVectorDatabaseRequest) SetWorkspaceId(v string) *InitVectorDatabaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *InitVectorDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
