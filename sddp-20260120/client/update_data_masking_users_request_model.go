// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataMaskingUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthRole(v string) *UpdateDataMaskingUsersRequest
	GetAuthRole() *string
	SetExpireTime(v int64) *UpdateDataMaskingUsersRequest
	GetExpireTime() *int64
	SetExpireTimeOperation(v string) *UpdateDataMaskingUsersRequest
	GetExpireTimeOperation() *string
	SetLang(v string) *UpdateDataMaskingUsersRequest
	GetLang() *string
	SetProductCode(v string) *UpdateDataMaskingUsersRequest
	GetProductCode() *string
	SetProductId(v int64) *UpdateDataMaskingUsersRequest
	GetProductId() *int64
	SetUserList(v []*UpdateDataMaskingUsersRequestUserList) *UpdateDataMaskingUsersRequest
	GetUserList() []*UpdateDataMaskingUsersRequestUserList
}

type UpdateDataMaskingUsersRequest struct {
	// The column encryption access permissions. Valid values: fullAccess: plaintext access. restrictedAccess: restricted access. noneAccess: prohibits decryption. noneAccess applies only to RDS MySQL instances whose current effective key pattern is client_key.
	//
	// example:
	//
	// fullAccess
	AuthRole *string `json:"AuthRole,omitempty" xml:"AuthRole,omitempty"`
	// The expiration time of the plaintext access permissions. The value is a UNIX timestamp in milliseconds since 00:00:00 UTC on January 1, 1970. Specify this parameter only when ExpireTimeOperation is set to SET. The value must be later than the current time and no greater than 2145953410000. The server saves the value with second-level precision. Do not specify this parameter when ExpireTimeOperation is set to PRESERVE or CLEAR.
	//
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The processing method for the permission expiration time. Valid values: PRESERVE: retains the original expiration time. CLEAR: clears the expiration time. SET: sets the expiration time to the value of ExpireTime. Default value: PRESERVE. Only fullAccess supports setting an expiration time.
	//
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// The language of the response description. Valid values: zh: Chinese. en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also provided, the two parameters must match.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type identifier. Valid values: 5: RDS. 8: POLARDB. 27: PolarDBX2. If specified together with ProductCode, the values must match.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of database accounts for which you want to set permissions.
	UserList []*UpdateDataMaskingUsersRequestUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s UpdateDataMaskingUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersRequest) GetAuthRole() *string {
	return s.AuthRole
}

func (s *UpdateDataMaskingUsersRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *UpdateDataMaskingUsersRequest) GetExpireTimeOperation() *string {
	return s.ExpireTimeOperation
}

func (s *UpdateDataMaskingUsersRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataMaskingUsersRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpdateDataMaskingUsersRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *UpdateDataMaskingUsersRequest) GetUserList() []*UpdateDataMaskingUsersRequestUserList {
	return s.UserList
}

func (s *UpdateDataMaskingUsersRequest) SetAuthRole(v string) *UpdateDataMaskingUsersRequest {
	s.AuthRole = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetExpireTime(v int64) *UpdateDataMaskingUsersRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetExpireTimeOperation(v string) *UpdateDataMaskingUsersRequest {
	s.ExpireTimeOperation = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetLang(v string) *UpdateDataMaskingUsersRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetProductCode(v string) *UpdateDataMaskingUsersRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetProductId(v int64) *UpdateDataMaskingUsersRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequest) SetUserList(v []*UpdateDataMaskingUsersRequestUserList) *UpdateDataMaskingUsersRequest {
	s.UserList = v
	return s
}

func (s *UpdateDataMaskingUsersRequest) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDataMaskingUsersRequestUserList struct {
	// The name of the database account. Obtain the value from Items[].AccountName in the records returned by ListDataAssetAccounts and pass it as a string exactly as it is. Do not pass Items[].Id. Each AccountId must correspond to the InstanceId in the same item. Accounts with the same name in different instances are considered different targets.
	//
	// example:
	//
	// app_reader
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The authentication database of MongoDB.
	//
	// example:
	//
	// database
	AuthDb *string `json:"AuthDb,omitempty" xml:"AuthDb,omitempty"`
	// The database instance identity. You can obtain the value from Items[].InstanceId returned by ListDataMaskingInstances.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDataMaskingUsersRequestUserList) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataMaskingUsersRequestUserList) GoString() string {
	return s.String()
}

func (s *UpdateDataMaskingUsersRequestUserList) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateDataMaskingUsersRequestUserList) GetAuthDb() *string {
	return s.AuthDb
}

func (s *UpdateDataMaskingUsersRequestUserList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDataMaskingUsersRequestUserList) SetAccountId(v string) *UpdateDataMaskingUsersRequestUserList {
	s.AccountId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequestUserList) SetAuthDb(v string) *UpdateDataMaskingUsersRequestUserList {
	s.AuthDb = &v
	return s
}

func (s *UpdateDataMaskingUsersRequestUserList) SetInstanceId(v string) *UpdateDataMaskingUsersRequestUserList {
	s.InstanceId = &v
	return s
}

func (s *UpdateDataMaskingUsersRequestUserList) Validate() error {
	return dara.Validate(s)
}
