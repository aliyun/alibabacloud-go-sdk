// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListSecretsRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ListSecretsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListSecretsRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListSecretsRequest
	GetWorkspaceId() *string
}

type ListSecretsRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListSecretsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListSecretsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSecretsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSecretsRequest) SetDBInstanceId(v string) *ListSecretsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListSecretsRequest) SetOwnerId(v int64) *ListSecretsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSecretsRequest) SetRegionId(v string) *ListSecretsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretsRequest) SetWorkspaceId(v string) *ListSecretsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSecretsRequest) Validate() error {
	return dara.Validate(s)
}
