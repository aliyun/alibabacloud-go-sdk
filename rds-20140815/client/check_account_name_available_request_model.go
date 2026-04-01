// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountNameAvailableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CheckAccountNameAvailableRequest
	GetAccountName() *string
	SetClientToken(v string) *CheckAccountNameAvailableRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CheckAccountNameAvailableRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CheckAccountNameAvailableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckAccountNameAvailableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckAccountNameAvailableRequest
	GetResourceOwnerAccount() *string
}

type CheckAccountNameAvailableRequest struct {
	// The username of the account.
	//
	// 	- The value must be unique.
	//
	// 	- The value must start with a lowercase letter, and end with a lowercase letter or a digit.
	//
	// 	- The value can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The length of the value must meet the following requirements:
	//
	//     	- If the instance runs MySQL 5.7 or MySQL 8.0, the value must be 2 to 32 characters in length.
	//
	//     	- If the instance runs MySQL 5.6, the value must be 2 to 16 characters in length.
	//
	//     	- If the instance runs SQL Server, the value must be 2 to 64 characters in length.
	//
	//     	- If the instance runs PostgreSQL with cloud disks, the value must be 2 to 63 characters in length.
	//
	//     	- If the instance runs PostgreSQL with local disks, the value must be 2 to 16 characters in length.
	//
	//     	- If the instance runs MariaDB, the value must be 2 to 16 characters in length.
	//
	// 	- For more information about invalid characters, see [Forbidden keywords table](https://help.aliyun.com/document_detail/26317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// DatabaseTest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
	// rm-uf6wjk5xxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s CheckAccountNameAvailableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountNameAvailableRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountNameAvailableRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CheckAccountNameAvailableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckAccountNameAvailableRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckAccountNameAvailableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckAccountNameAvailableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckAccountNameAvailableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckAccountNameAvailableRequest) SetAccountName(v string) *CheckAccountNameAvailableRequest {
	s.AccountName = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetClientToken(v string) *CheckAccountNameAvailableRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetDBInstanceId(v string) *CheckAccountNameAvailableRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetOwnerAccount(v string) *CheckAccountNameAvailableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetOwnerId(v int64) *CheckAccountNameAvailableRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) SetResourceOwnerAccount(v string) *CheckAccountNameAvailableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckAccountNameAvailableRequest) Validate() error {
	return dara.Validate(s)
}
