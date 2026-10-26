// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataMaskingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncAlgorithm(v string) *CreateDataMaskingRuleRequest
	GetEncAlgorithm() *string
	SetEncryptionKeyId(v string) *CreateDataMaskingRuleRequest
	GetEncryptionKeyId() *string
	SetEncryptionKeyMode(v string) *CreateDataMaskingRuleRequest
	GetEncryptionKeyMode() *string
	SetEngineType(v string) *CreateDataMaskingRuleRequest
	GetEngineType() *string
	SetExpireTime(v int64) *CreateDataMaskingRuleRequest
	GetExpireTime() *int64
	SetExpireTimeOperation(v string) *CreateDataMaskingRuleRequest
	GetExpireTimeOperation() *string
	SetInstanceId(v string) *CreateDataMaskingRuleRequest
	GetInstanceId() *string
	SetLang(v string) *CreateDataMaskingRuleRequest
	GetLang() *string
	SetProductCode(v string) *CreateDataMaskingRuleRequest
	GetProductCode() *string
	SetProductId(v int64) *CreateDataMaskingRuleRequest
	GetProductId() *int64
	SetRiskHandleId(v int64) *CreateDataMaskingRuleRequest
	GetRiskHandleId() *int64
	SetSubRuleList(v []*CreateDataMaskingRuleRequestSubRuleList) *CreateDataMaskingRuleRequest
	GetSubRuleList() []*CreateDataMaskingRuleRequestSubRuleList
	SetUserList(v []*CreateDataMaskingRuleRequestUserList) *CreateDataMaskingRuleRequest
	GetUserList() []*CreateDataMaskingRuleRequestUserList
}

type CreateDataMaskingRuleRequest struct {
	// The name of the column encryption algorithm. The availability of AES_128_GCM, AES_256_GCM, and SM4_128_GCM depends on the instance. Select an available algorithm from ListDataMaskingEncryptionAlgorithms.
	//
	// example:
	//
	// AES_256_GCM
	EncAlgorithm *string `json:"EncAlgorithm,omitempty" xml:"EncAlgorithm,omitempty"`
	// The identifier of the Key Management Service (KMS) key. You can obtain it from Items[].KeyId in the response of the ListKmsKeys operation. Provide this parameter when you use a KMS key. Do not pass this parameter in local key scenarios.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The key mode. Valid values:
	//
	// - client_key: local key
	//
	// - kms_key: KMS key
	//
	// When you use kms_key, you must specify EncryptionKeyId.
	//
	// example:
	//
	// client_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// The type of the database engine. The value must be consistent with the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the plaintext permission. The value is a UNIX timestamp in milliseconds since 00:00:00 UTC on January 1, 1970. Specify this parameter only when ExpireTimeOperation is set to SET. The value must be later than the current time and no greater than 2145953410000. The server saves the value with second precision. Do not specify this parameter when ExpireTimeOperation is set to PRESERVE or CLEAR.
	//
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The processing method for the permission expiration time. Valid values:
	//
	// - PRESERVE: retains the original expiration time.
	//
	// - CLEAR: clears the expiration time.
	//
	// - SET: sets ExpireTime.
	//
	// Default value: PRESERVE. Only fullAccess supports setting a validity period.
	//
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// The instance identity. You can obtain it from Items[].InstanceId in the response of the ListDataMaskingInstances operation.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The code of the cloud service. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also passed, the two must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The identifier of the cloud service type. Valid values:
	//
	// - 5: RDS
	//
	// - 8: POLARDB
	//
	// - 27: PolarDBX2
	//
	// When specified together with ProductCode, the values must correspond to each other.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The identifier of the associated risk handling record. You do not need to specify this parameter in the standard column encryption creation process.
	//
	// example:
	//
	// 1001
	RiskHandleId *int64 `json:"RiskHandleId,omitempty" xml:"RiskHandleId,omitempty"`
	// The list of column encryption targets grouped by database and table.
	SubRuleList []*CreateDataMaskingRuleRequestSubRuleList `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty" type:"Repeated"`
	// The list of database accounts to which the fullAccess plaintext permission is granted. The accounts must belong to the target instance.
	UserList []*CreateDataMaskingRuleRequestUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s CreateDataMaskingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleRequest) GetEncAlgorithm() *string {
	return s.EncAlgorithm
}

func (s *CreateDataMaskingRuleRequest) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *CreateDataMaskingRuleRequest) GetEncryptionKeyMode() *string {
	return s.EncryptionKeyMode
}

func (s *CreateDataMaskingRuleRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateDataMaskingRuleRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateDataMaskingRuleRequest) GetExpireTimeOperation() *string {
	return s.ExpireTimeOperation
}

func (s *CreateDataMaskingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDataMaskingRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataMaskingRuleRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateDataMaskingRuleRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *CreateDataMaskingRuleRequest) GetRiskHandleId() *int64 {
	return s.RiskHandleId
}

func (s *CreateDataMaskingRuleRequest) GetSubRuleList() []*CreateDataMaskingRuleRequestSubRuleList {
	return s.SubRuleList
}

func (s *CreateDataMaskingRuleRequest) GetUserList() []*CreateDataMaskingRuleRequestUserList {
	return s.UserList
}

func (s *CreateDataMaskingRuleRequest) SetEncAlgorithm(v string) *CreateDataMaskingRuleRequest {
	s.EncAlgorithm = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetEncryptionKeyId(v string) *CreateDataMaskingRuleRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetEncryptionKeyMode(v string) *CreateDataMaskingRuleRequest {
	s.EncryptionKeyMode = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetEngineType(v string) *CreateDataMaskingRuleRequest {
	s.EngineType = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetExpireTime(v int64) *CreateDataMaskingRuleRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetExpireTimeOperation(v string) *CreateDataMaskingRuleRequest {
	s.ExpireTimeOperation = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetInstanceId(v string) *CreateDataMaskingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetLang(v string) *CreateDataMaskingRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetProductCode(v string) *CreateDataMaskingRuleRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetProductId(v int64) *CreateDataMaskingRuleRequest {
	s.ProductId = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetRiskHandleId(v int64) *CreateDataMaskingRuleRequest {
	s.RiskHandleId = &v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetSubRuleList(v []*CreateDataMaskingRuleRequestSubRuleList) *CreateDataMaskingRuleRequest {
	s.SubRuleList = v
	return s
}

func (s *CreateDataMaskingRuleRequest) SetUserList(v []*CreateDataMaskingRuleRequestUserList) *CreateDataMaskingRuleRequest {
	s.UserList = v
	return s
}

func (s *CreateDataMaskingRuleRequest) Validate() error {
	if s.SubRuleList != nil {
		for _, item := range s.SubRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateDataMaskingRuleRequestSubRuleList struct {
	// The list of target column names. Multiple column names are separated by commas (,).
	//
	// example:
	//
	// phone,email
	Columns *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The name of the target database.
	//
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The name of the target table.
	//
	// example:
	//
	// customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreateDataMaskingRuleRequestSubRuleList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleRequestSubRuleList) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleRequestSubRuleList) GetColumns() *string {
	return s.Columns
}

func (s *CreateDataMaskingRuleRequestSubRuleList) GetDbName() *string {
	return s.DbName
}

func (s *CreateDataMaskingRuleRequestSubRuleList) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataMaskingRuleRequestSubRuleList) SetColumns(v string) *CreateDataMaskingRuleRequestSubRuleList {
	s.Columns = &v
	return s
}

func (s *CreateDataMaskingRuleRequestSubRuleList) SetDbName(v string) *CreateDataMaskingRuleRequestSubRuleList {
	s.DbName = &v
	return s
}

func (s *CreateDataMaskingRuleRequestSubRuleList) SetTableName(v string) *CreateDataMaskingRuleRequestSubRuleList {
	s.TableName = &v
	return s
}

func (s *CreateDataMaskingRuleRequestSubRuleList) Validate() error {
	return dara.Validate(s)
}

type CreateDataMaskingRuleRequestUserList struct {
	// The identifier of the database account. You can obtain it from Items[].Id in the response of the ListDataAssetAccounts operation. Pass the value as a string.
	//
	// example:
	//
	// 1001
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The authentication database for MongoDB accounts. This parameter is required when you select a MongoDB account. You do not need to pass this parameter for other services.
	//
	// example:
	//
	// admin
	AuthDb *string `json:"AuthDb,omitempty" xml:"AuthDb,omitempty"`
}

func (s CreateDataMaskingRuleRequestUserList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleRequestUserList) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleRequestUserList) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateDataMaskingRuleRequestUserList) GetAuthDb() *string {
	return s.AuthDb
}

func (s *CreateDataMaskingRuleRequestUserList) SetAccountId(v string) *CreateDataMaskingRuleRequestUserList {
	s.AccountId = &v
	return s
}

func (s *CreateDataMaskingRuleRequestUserList) SetAuthDb(v string) *CreateDataMaskingRuleRequestUserList {
	s.AuthDb = &v
	return s
}

func (s *CreateDataMaskingRuleRequestUserList) Validate() error {
	return dara.Validate(s)
}
