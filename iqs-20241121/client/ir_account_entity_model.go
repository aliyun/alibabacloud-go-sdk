// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIrAccountEntity interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *IrAccountEntity
	GetAccountId() *string
	SetAccountName(v string) *IrAccountEntity
	GetAccountName() *string
	SetAccountStatus(v string) *IrAccountEntity
	GetAccountStatus() *string
	SetAdjustedNormalQps(v int32) *IrAccountEntity
	GetAdjustedNormalQps() *int32
	SetConfiguration(v string) *IrAccountEntity
	GetConfiguration() *string
	SetCreateTime(v string) *IrAccountEntity
	GetCreateTime() *string
	SetId(v int64) *IrAccountEntity
	GetId() *int64
	SetIsDeleted(v int32) *IrAccountEntity
	GetIsDeleted() *int32
	SetModifiedTime(v string) *IrAccountEntity
	GetModifiedTime() *string
	SetQuarkKey(v string) *IrAccountEntity
	GetQuarkKey() *string
	SetQuarkOsr(v string) *IrAccountEntity
	GetQuarkOsr() *string
	SetQuarkUsername(v string) *IrAccountEntity
	GetQuarkUsername() *string
	SetSearchType(v string) *IrAccountEntity
	GetSearchType() *string
	SetTestQps(v int32) *IrAccountEntity
	GetTestQps() *int32
	SetTestQueryPerDay(v int32) *IrAccountEntity
	GetTestQueryPerDay() *int32
	SetTestStartTime(v string) *IrAccountEntity
	GetTestStartTime() *string
}

type IrAccountEntity struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName       *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	AccountStatus     *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	AdjustedNormalQps *int32  `json:"adjustedNormalQps,omitempty" xml:"adjustedNormalQps,omitempty"`
	Configuration     *string `json:"configuration,omitempty" xml:"configuration,omitempty"`
	CreateTime        *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted         *int32  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	ModifiedTime      *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	QuarkKey          *string `json:"quarkKey,omitempty" xml:"quarkKey,omitempty"`
	QuarkOsr          *string `json:"quarkOsr,omitempty" xml:"quarkOsr,omitempty"`
	QuarkUsername     *string `json:"quarkUsername,omitempty" xml:"quarkUsername,omitempty"`
	SearchType        *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	TestQps           *int32  `json:"testQps,omitempty" xml:"testQps,omitempty"`
	TestQueryPerDay   *int32  `json:"testQueryPerDay,omitempty" xml:"testQueryPerDay,omitempty"`
	TestStartTime     *string `json:"testStartTime,omitempty" xml:"testStartTime,omitempty"`
}

func (s IrAccountEntity) String() string {
	return dara.Prettify(s)
}

func (s IrAccountEntity) GoString() string {
	return s.String()
}

func (s *IrAccountEntity) GetAccountId() *string {
	return s.AccountId
}

func (s *IrAccountEntity) GetAccountName() *string {
	return s.AccountName
}

func (s *IrAccountEntity) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *IrAccountEntity) GetAdjustedNormalQps() *int32 {
	return s.AdjustedNormalQps
}

func (s *IrAccountEntity) GetConfiguration() *string {
	return s.Configuration
}

func (s *IrAccountEntity) GetCreateTime() *string {
	return s.CreateTime
}

func (s *IrAccountEntity) GetId() *int64 {
	return s.Id
}

func (s *IrAccountEntity) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *IrAccountEntity) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *IrAccountEntity) GetQuarkKey() *string {
	return s.QuarkKey
}

func (s *IrAccountEntity) GetQuarkOsr() *string {
	return s.QuarkOsr
}

func (s *IrAccountEntity) GetQuarkUsername() *string {
	return s.QuarkUsername
}

func (s *IrAccountEntity) GetSearchType() *string {
	return s.SearchType
}

func (s *IrAccountEntity) GetTestQps() *int32 {
	return s.TestQps
}

func (s *IrAccountEntity) GetTestQueryPerDay() *int32 {
	return s.TestQueryPerDay
}

func (s *IrAccountEntity) GetTestStartTime() *string {
	return s.TestStartTime
}

func (s *IrAccountEntity) SetAccountId(v string) *IrAccountEntity {
	s.AccountId = &v
	return s
}

func (s *IrAccountEntity) SetAccountName(v string) *IrAccountEntity {
	s.AccountName = &v
	return s
}

func (s *IrAccountEntity) SetAccountStatus(v string) *IrAccountEntity {
	s.AccountStatus = &v
	return s
}

func (s *IrAccountEntity) SetAdjustedNormalQps(v int32) *IrAccountEntity {
	s.AdjustedNormalQps = &v
	return s
}

func (s *IrAccountEntity) SetConfiguration(v string) *IrAccountEntity {
	s.Configuration = &v
	return s
}

func (s *IrAccountEntity) SetCreateTime(v string) *IrAccountEntity {
	s.CreateTime = &v
	return s
}

func (s *IrAccountEntity) SetId(v int64) *IrAccountEntity {
	s.Id = &v
	return s
}

func (s *IrAccountEntity) SetIsDeleted(v int32) *IrAccountEntity {
	s.IsDeleted = &v
	return s
}

func (s *IrAccountEntity) SetModifiedTime(v string) *IrAccountEntity {
	s.ModifiedTime = &v
	return s
}

func (s *IrAccountEntity) SetQuarkKey(v string) *IrAccountEntity {
	s.QuarkKey = &v
	return s
}

func (s *IrAccountEntity) SetQuarkOsr(v string) *IrAccountEntity {
	s.QuarkOsr = &v
	return s
}

func (s *IrAccountEntity) SetQuarkUsername(v string) *IrAccountEntity {
	s.QuarkUsername = &v
	return s
}

func (s *IrAccountEntity) SetSearchType(v string) *IrAccountEntity {
	s.SearchType = &v
	return s
}

func (s *IrAccountEntity) SetTestQps(v int32) *IrAccountEntity {
	s.TestQps = &v
	return s
}

func (s *IrAccountEntity) SetTestQueryPerDay(v int32) *IrAccountEntity {
	s.TestQueryPerDay = &v
	return s
}

func (s *IrAccountEntity) SetTestStartTime(v string) *IrAccountEntity {
	s.TestStartTime = &v
	return s
}

func (s *IrAccountEntity) Validate() error {
	return dara.Validate(s)
}
