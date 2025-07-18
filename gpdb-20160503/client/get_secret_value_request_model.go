// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *GetSecretValueRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *GetSecretValueRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetSecretValueRequest
	GetRegionId() *string
	SetSecretArn(v string) *GetSecretValueRequest
	GetSecretArn() *string
	SetSecretName(v string) *GetSecretValueRequest
	GetSecretName() *string
	SetWorkspaceId(v string) *GetSecretValueRequest
	GetWorkspaceId() *string
}

type GetSecretValueRequest struct {
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

func (s GetSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetSecretValueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSecretValueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSecretValueRequest) GetSecretArn() *string {
	return s.SecretArn
}

func (s *GetSecretValueRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *GetSecretValueRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSecretValueRequest) SetDBInstanceId(v string) *GetSecretValueRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GetSecretValueRequest) SetOwnerId(v int64) *GetSecretValueRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSecretValueRequest) SetRegionId(v string) *GetSecretValueRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretValueRequest) SetSecretArn(v string) *GetSecretValueRequest {
	s.SecretArn = &v
	return s
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetWorkspaceId(v string) *GetSecretValueRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSecretValueRequest) Validate() error {
	return dara.Validate(s)
}
