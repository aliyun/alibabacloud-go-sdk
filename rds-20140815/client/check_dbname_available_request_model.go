// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameAvailableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CheckDBNameAvailableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CheckDBNameAvailableRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CheckDBNameAvailableRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CheckDBNameAvailableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckDBNameAvailableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckDBNameAvailableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckDBNameAvailableRequest
	GetResourceOwnerId() *int64
}

type CheckDBNameAvailableRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3axxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name that you want to check.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckDBNameAvailableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameAvailableRequest) GoString() string {
	return s.String()
}

func (s *CheckDBNameAvailableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckDBNameAvailableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckDBNameAvailableRequest) GetDBName() *string {
	return s.DBName
}

func (s *CheckDBNameAvailableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckDBNameAvailableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckDBNameAvailableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckDBNameAvailableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckDBNameAvailableRequest) SetClientToken(v string) *CheckDBNameAvailableRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetDBInstanceId(v string) *CheckDBNameAvailableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetDBName(v string) *CheckDBNameAvailableRequest {
	s.DBName = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetOwnerAccount(v string) *CheckDBNameAvailableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetOwnerId(v int64) *CheckDBNameAvailableRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetResourceOwnerAccount(v string) *CheckDBNameAvailableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDBNameAvailableRequest) SetResourceOwnerId(v int64) *CheckDBNameAvailableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDBNameAvailableRequest) Validate() error {
	return dara.Validate(s)
}
