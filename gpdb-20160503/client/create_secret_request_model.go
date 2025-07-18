// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateSecretRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateSecretRequest
	GetDescription() *string
	SetOwnerId(v int64) *CreateSecretRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateSecretRequest
	GetPassword() *string
	SetRegionId(v string) *CreateSecretRequest
	GetRegionId() *string
	SetSecretName(v string) *CreateSecretRequest
	GetSecretName() *string
	SetTestConnection(v bool) *CreateSecretRequest
	GetTestConnection() *bool
	SetUsername(v string) *CreateSecretRequest
	GetUsername() *string
	SetWorkspaceId(v string) *CreateSecretRequest
	GetWorkspaceId() *string
}

type CreateSecretRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the access credential.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the database account that is used to access the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the access credential. The name must be 1 to 16 characters in length and can contain letters, digits, and underscores (_). If you leave this parameter empty, the value of the Username parameter is used.
	//
	// example:
	//
	// testsecret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// Specifies whether to check the connectivity to the instance by using the name and password of the database account.
	//
	// example:
	//
	// true
	TestConnection *bool `json:"TestConnection,omitempty" xml:"TestConnection,omitempty"`
	// The name of the database account that is used to access the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// testacc
	Username    *string `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateSecretRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSecretRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateSecretRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretRequest) GetTestConnection() *bool {
	return s.TestConnection
}

func (s *CreateSecretRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateSecretRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateSecretRequest) SetDBInstanceId(v string) *CreateSecretRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetOwnerId(v int64) *CreateSecretRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSecretRequest) SetPassword(v string) *CreateSecretRequest {
	s.Password = &v
	return s
}

func (s *CreateSecretRequest) SetRegionId(v string) *CreateSecretRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetTestConnection(v bool) *CreateSecretRequest {
	s.TestConnection = &v
	return s
}

func (s *CreateSecretRequest) SetUsername(v string) *CreateSecretRequest {
	s.Username = &v
	return s
}

func (s *CreateSecretRequest) SetWorkspaceId(v string) *CreateSecretRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateSecretRequest) Validate() error {
	return dara.Validate(s)
}
