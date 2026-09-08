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
	// example:
	//
	// AES_256_GCM
	EncAlgorithm *string `json:"EncAlgorithm,omitempty" xml:"EncAlgorithm,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-12345678****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// example:
	//
	// client_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2145953410000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// PRESERVE
	ExpireTimeOperation *string `json:"ExpireTimeOperation,omitempty" xml:"ExpireTimeOperation,omitempty"`
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1001
	RiskHandleId      *int64  `json:"RiskHandleId,omitempty" xml:"RiskHandleId,omitempty"`
	SubRuleListShrink *string `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty"`
	UserListShrink    *string `json:"UserList,omitempty" xml:"UserList,omitempty"`
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
