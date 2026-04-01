// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDBInstanceNetTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SwitchDBInstanceNetTypeRequest
	GetClientToken() *string
	SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest
	GetConnectionStringPrefix() *string
	SetConnectionStringType(v string) *SwitchDBInstanceNetTypeRequest
	GetConnectionStringType() *string
	SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *SwitchDBInstanceNetTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchDBInstanceNetTypeRequest
	GetOwnerId() *int64
	SetPort(v string) *SwitchDBInstanceNetTypeRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *SwitchDBInstanceNetTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchDBInstanceNetTypeRequest
	GetResourceOwnerId() *int64
}

type SwitchDBInstanceNetTypeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The prefix of the custom endpoint. The prefix must be 8 to 64 characters in length and can contain letters and digits. It must start with a lowercase letter. A valid endpoint is in the following format: Prefix.Database engine.rds.aliyuncs.com. Example: test1234.mysql.rds.aliyuncs.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// new**********
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **ReadWriteSplitting**
	//
	// By default, the system returns both types of endpoints.
	//
	// example:
	//
	// Normal
	ConnectionStringType *string `json:"ConnectionStringType,omitempty" xml:"ConnectionStringType,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1**************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the port that is used to connect to the instance. Valid values: **3001 to 3999**.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SwitchDBInstanceNetTypeRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *SwitchDBInstanceNetTypeRequest) GetConnectionStringType() *string {
	return s.ConnectionStringType
}

func (s *SwitchDBInstanceNetTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *SwitchDBInstanceNetTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchDBInstanceNetTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchDBInstanceNetTypeRequest) GetPort() *string {
	return s.Port
}

func (s *SwitchDBInstanceNetTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchDBInstanceNetTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchDBInstanceNetTypeRequest) SetClientToken(v string) *SwitchDBInstanceNetTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringType(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringType = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetOwnerAccount(v string) *SwitchDBInstanceNetTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetOwnerId(v int64) *SwitchDBInstanceNetTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetResourceOwnerAccount(v string) *SwitchDBInstanceNetTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetResourceOwnerId(v int64) *SwitchDBInstanceNetTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) Validate() error {
	return dara.Validate(s)
}
