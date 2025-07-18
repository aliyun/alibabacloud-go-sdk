// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteSecretRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DeleteSecretRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSecretRequest
	GetRegionId() *string
	SetSecretArn(v string) *DeleteSecretRequest
	GetSecretArn() *string
	SetSecretName(v string) *DeleteSecretRequest
	GetSecretName() *string
	SetWorkspaceId(v string) *DeleteSecretRequest
	GetWorkspaceId() *string
}

type DeleteSecretRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// >
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the access credential for the created Data API account. Format: `acs:gpdb:{{region}}:{{accountId}}:secret/{{secretName}}-{{32 digits random string}`.
	//
	// >  You must specify one of the SecretArn and SecretName parameters.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the access credential.
	//
	// >  You must specify one of the SecretArn and SecretName parameters.
	//
	// example:
	//
	// testsecret
	SecretName  *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteSecretRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSecretRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSecretRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DeleteSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteSecretRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteSecretRequest) SetDBInstanceId(v string) *DeleteSecretRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteSecretRequest) SetOwnerId(v int64) *DeleteSecretRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSecretRequest) SetRegionId(v string) *DeleteSecretRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretArn(v string) *DeleteSecretRequest {
	s.SecretArn = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretRequest) SetWorkspaceId(v string) *DeleteSecretRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteSecretRequest) Validate() error {
	return dara.Validate(s)
}
