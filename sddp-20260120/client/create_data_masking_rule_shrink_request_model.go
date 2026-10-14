// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataMaskingRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncAlgorithm(v string) *CreateDataMaskingRuleShrinkRequest
	GetEncAlgorithm() *string
	SetEncryptionKeyId(v string) *CreateDataMaskingRuleShrinkRequest
	GetEncryptionKeyId() *string
	SetEncryptionKeyMode(v string) *CreateDataMaskingRuleShrinkRequest
	GetEncryptionKeyMode() *string
	SetEngineType(v string) *CreateDataMaskingRuleShrinkRequest
	GetEngineType() *string
	SetExpireTime(v int64) *CreateDataMaskingRuleShrinkRequest
	GetExpireTime() *int64
	SetExpireTimeOperation(v string) *CreateDataMaskingRuleShrinkRequest
	GetExpireTimeOperation() *string
	SetInstanceId(v string) *CreateDataMaskingRuleShrinkRequest
	GetInstanceId() *string
	SetLang(v string) *CreateDataMaskingRuleShrinkRequest
	GetLang() *string
	SetProductCode(v string) *CreateDataMaskingRuleShrinkRequest
	GetProductCode() *string
	SetProductId(v int64) *CreateDataMaskingRuleShrinkRequest
	GetProductId() *int64
	SetRiskHandleId(v int64) *CreateDataMaskingRuleShrinkRequest
	GetRiskHandleId() *int64
	SetSubRuleListShrink(v string) *CreateDataMaskingRuleShrinkRequest
	GetSubRuleListShrink() *string
	SetUserListShrink(v string) *CreateDataMaskingRuleShrinkRequest
	GetUserListShrink() *string
}

type CreateDataMaskingRuleShrinkRequest struct {
	// The name of the column encryption algorithm. The availability of AES_128_GCM, AES_256_GCM, and SM4_128_GCM depends on the instance. Select an available algorithm from [ListDataMaskingEncryptionAlgorithms](https://help.aliyun.com/document_detail/2859327.html).
	//
	// example:
	//
	// AES_256_GCM
	EncAlgorithm *string `json:"EncAlgorithm,omitempty" xml:"EncAlgorithm,omitempty"`
	// The KMS key ID. You can obtain this value from Items[].KeyId in the response of [ListKmsKeys](https://help.aliyun.com/document_detail/2859328.html). Specify this parameter when you use a KMS key. Do not specify this parameter for local key scenarios.
	//
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The key mode. Valid values:
	//
	// - client_key: local key.
	//
	// - kms_key: KMS key. If you set this parameter to kms_key, you must specify EncryptionKeyId.
	//
	// example:
	//
	// client_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// The database engine type. This value must match the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the plaintext access permissions. This value is a UNIX timestamp in milliseconds since 1970-01-01 00:00:00 UTC. Specify this parameter only when ExpireTimeOperation is set to SET. The value must be later than the current time and no greater than 2145953410000. The server stores the value with second-level precision. Do not specify this parameter when ExpireTimeOperation is set to PRESERVE or CLEAR.
	//
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The handling method for the permission expiration time. Valid values:
	//
	// - PRESERVE: Retains the original expiration time.
	//
	// - CLEAR: Clears the expiration time.
	//
	// - SET: Sets the expiration time to the value specified by ExpireTime.
	//
	// Default value: PRESERVE. Only fullAccess supports expiration time settings.
	//
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// The database instance identity. You can obtain this value from Items[].InstanceId in the response of [ListDataMaskingInstances](https://help.aliyun.com/document_detail/2859326.html).
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If you specify both ProductCode and ProductId, the values must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type ID. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If you specify both ProductId and ProductCode, the values must correspond to each other.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the associated risk handling record. You do not need to specify this parameter for standard column encryption creation.
	//
	// example:
	//
	// 1001
	RiskHandleId *int64 `json:"RiskHandleId,omitempty" xml:"RiskHandleId,omitempty"`
	// The list of column encryption targets grouped by database and table.
	SubRuleListShrink *string `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty"`
	// The list of database accounts to be granted fullAccess plaintext access permissions. The accounts must belong to the target instance.
	UserListShrink *string `json:"UserList,omitempty" xml:"UserList,omitempty"`
}

func (s CreateDataMaskingRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleShrinkRequest) GetEncAlgorithm() *string {
	return s.EncAlgorithm
}

func (s *CreateDataMaskingRuleShrinkRequest) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *CreateDataMaskingRuleShrinkRequest) GetEncryptionKeyMode() *string {
	return s.EncryptionKeyMode
}

func (s *CreateDataMaskingRuleShrinkRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateDataMaskingRuleShrinkRequest) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *CreateDataMaskingRuleShrinkRequest) GetExpireTimeOperation() *string {
	return s.ExpireTimeOperation
}

func (s *CreateDataMaskingRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDataMaskingRuleShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDataMaskingRuleShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CreateDataMaskingRuleShrinkRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *CreateDataMaskingRuleShrinkRequest) GetRiskHandleId() *int64 {
	return s.RiskHandleId
}

func (s *CreateDataMaskingRuleShrinkRequest) GetSubRuleListShrink() *string {
	return s.SubRuleListShrink
}

func (s *CreateDataMaskingRuleShrinkRequest) GetUserListShrink() *string {
	return s.UserListShrink
}

func (s *CreateDataMaskingRuleShrinkRequest) SetEncAlgorithm(v string) *CreateDataMaskingRuleShrinkRequest {
	s.EncAlgorithm = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetEncryptionKeyId(v string) *CreateDataMaskingRuleShrinkRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetEncryptionKeyMode(v string) *CreateDataMaskingRuleShrinkRequest {
	s.EncryptionKeyMode = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetEngineType(v string) *CreateDataMaskingRuleShrinkRequest {
	s.EngineType = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetExpireTime(v int64) *CreateDataMaskingRuleShrinkRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetExpireTimeOperation(v string) *CreateDataMaskingRuleShrinkRequest {
	s.ExpireTimeOperation = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetInstanceId(v string) *CreateDataMaskingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetLang(v string) *CreateDataMaskingRuleShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetProductCode(v string) *CreateDataMaskingRuleShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetProductId(v int64) *CreateDataMaskingRuleShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetRiskHandleId(v int64) *CreateDataMaskingRuleShrinkRequest {
	s.RiskHandleId = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetSubRuleListShrink(v string) *CreateDataMaskingRuleShrinkRequest {
	s.SubRuleListShrink = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) SetUserListShrink(v string) *CreateDataMaskingRuleShrinkRequest {
	s.UserListShrink = &v
	return s
}

func (s *CreateDataMaskingRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
