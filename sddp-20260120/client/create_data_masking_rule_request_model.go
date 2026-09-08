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
	RiskHandleId *int64                                     `json:"RiskHandleId,omitempty" xml:"RiskHandleId,omitempty"`
	SubRuleList  []*CreateDataMaskingRuleRequestSubRuleList `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty" type:"Repeated"`
	UserList     []*CreateDataMaskingRuleRequestUserList    `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
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
	// example:
	//
	// phone,email
	Columns *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// example:
	//
	// business_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
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
	// example:
	//
	// 1001
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *CreateDataMaskingRuleRequestUserList) SetAccountId(v string) *CreateDataMaskingRuleRequestUserList {
	s.AccountId = &v
	return s
}

func (s *CreateDataMaskingRuleRequestUserList) Validate() error {
	return dara.Validate(s)
}
