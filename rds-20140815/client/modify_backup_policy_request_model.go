// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest
	GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies
	SetAdvancedLogPolicies(v []*ModifyBackupPolicyRequestAdvancedLogPolicies) *ModifyBackupPolicyRequest
	GetAdvancedLogPolicies() []*ModifyBackupPolicyRequestAdvancedLogPolicies
	SetArchiveBackupKeepCount(v int32) *ModifyBackupPolicyRequest
	GetArchiveBackupKeepCount() *int32
	SetArchiveBackupKeepPolicy(v string) *ModifyBackupPolicyRequest
	GetArchiveBackupKeepPolicy() *string
	SetArchiveBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetArchiveBackupRetentionPeriod() *string
	SetBackupInterval(v string) *ModifyBackupPolicyRequest
	GetBackupInterval() *string
	SetBackupLog(v string) *ModifyBackupPolicyRequest
	GetBackupLog() *string
	SetBackupMethod(v string) *ModifyBackupPolicyRequest
	GetBackupMethod() *string
	SetBackupPolicyMode(v string) *ModifyBackupPolicyRequest
	GetBackupPolicyMode() *string
	SetBackupPriority(v int32) *ModifyBackupPolicyRequest
	GetBackupPriority() *int32
	SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetBackupRetentionPeriod() *string
	SetCategory(v string) *ModifyBackupPolicyRequest
	GetCategory() *string
	SetCompressType(v string) *ModifyBackupPolicyRequest
	GetCompressType() *string
	SetDBInstanceId(v string) *ModifyBackupPolicyRequest
	GetDBInstanceId() *string
	SetEnableAdvancedBackupPolicy(v int32) *ModifyBackupPolicyRequest
	GetEnableAdvancedBackupPolicy() *int32
	SetEnableBackupLog(v string) *ModifyBackupPolicyRequest
	GetEnableBackupLog() *string
	SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyRequest
	GetEnableIncrementDataBackup() *bool
	SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest
	GetHighSpaceUsageProtection() *string
	SetLocalLogRetentionHours(v string) *ModifyBackupPolicyRequest
	GetLocalLogRetentionHours() *string
	SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyRequest
	GetLocalLogRetentionSpace() *string
	SetLogBackupFrequency(v string) *ModifyBackupPolicyRequest
	GetLogBackupFrequency() *string
	SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyRequest
	GetLogBackupLocalRetentionNumber() *int32
	SetLogBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest
	GetLogBackupRetentionPeriod() *string
	SetOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupPolicyRequest
	GetOwnerId() *int64
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetReleasedKeepPolicy(v string) *ModifyBackupPolicyRequest
	GetReleasedKeepPolicy() *string
	SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupPolicyRequest struct {
	AdvancedDataPolicies         []*ModifyBackupPolicyRequestAdvancedDataPolicies `json:"AdvancedDataPolicies,omitempty" xml:"AdvancedDataPolicies,omitempty" type:"Repeated"`
	AdvancedLogPolicies          []*ModifyBackupPolicyRequestAdvancedLogPolicies  `json:"AdvancedLogPolicies,omitempty" xml:"AdvancedLogPolicies,omitempty" type:"Repeated"`
	ArchiveBackupKeepCount       *int32                                           `json:"ArchiveBackupKeepCount,omitempty" xml:"ArchiveBackupKeepCount,omitempty"`
	ArchiveBackupKeepPolicy      *string                                          `json:"ArchiveBackupKeepPolicy,omitempty" xml:"ArchiveBackupKeepPolicy,omitempty"`
	ArchiveBackupRetentionPeriod *string                                          `json:"ArchiveBackupRetentionPeriod,omitempty" xml:"ArchiveBackupRetentionPeriod,omitempty"`
	BackupInterval               *string                                          `json:"BackupInterval,omitempty" xml:"BackupInterval,omitempty"`
	BackupLog                    *string                                          `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	BackupMethod                 *string                                          `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupPolicyMode             *string                                          `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	BackupPriority               *int32                                           `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	BackupRetentionPeriod        *string                                          `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	Category                     *string                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	CompressType                 *string                                          `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// This parameter is required.
	DBInstanceId                  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EnableAdvancedBackupPolicy    *int32  `json:"EnableAdvancedBackupPolicy,omitempty" xml:"EnableAdvancedBackupPolicy,omitempty"`
	EnableBackupLog               *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	EnableIncrementDataBackup     *bool   `json:"EnableIncrementDataBackup,omitempty" xml:"EnableIncrementDataBackup,omitempty"`
	HighSpaceUsageProtection      *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	LocalLogRetentionHours        *string `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace        *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	LogBackupFrequency            *string `json:"LogBackupFrequency,omitempty" xml:"LogBackupFrequency,omitempty"`
	LogBackupLocalRetentionNumber *int32  `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	LogBackupRetentionPeriod      *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	OwnerAccount                  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PreferredBackupPeriod         *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	PreferredBackupTime           *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	ReleasedKeepPolicy            *string `json:"ReleasedKeepPolicy,omitempty" xml:"ReleasedKeepPolicy,omitempty"`
	ResourceOwnerAccount          *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetAdvancedDataPolicies() []*ModifyBackupPolicyRequestAdvancedDataPolicies {
	return s.AdvancedDataPolicies
}

func (s *ModifyBackupPolicyRequest) GetAdvancedLogPolicies() []*ModifyBackupPolicyRequestAdvancedLogPolicies {
	return s.AdvancedLogPolicies
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupKeepCount() *int32 {
	return s.ArchiveBackupKeepCount
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupKeepPolicy() *string {
	return s.ArchiveBackupKeepPolicy
}

func (s *ModifyBackupPolicyRequest) GetArchiveBackupRetentionPeriod() *string {
	return s.ArchiveBackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetBackupInterval() *string {
	return s.BackupInterval
}

func (s *ModifyBackupPolicyRequest) GetBackupLog() *string {
	return s.BackupLog
}

func (s *ModifyBackupPolicyRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ModifyBackupPolicyRequest) GetBackupPolicyMode() *string {
	return s.BackupPolicyMode
}

func (s *ModifyBackupPolicyRequest) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyBackupPolicyRequest) GetCompressType() *string {
	return s.CompressType
}

func (s *ModifyBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupPolicyRequest) GetEnableAdvancedBackupPolicy() *int32 {
	return s.EnableAdvancedBackupPolicy
}

func (s *ModifyBackupPolicyRequest) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyRequest) GetEnableIncrementDataBackup() *bool {
	return s.EnableIncrementDataBackup
}

func (s *ModifyBackupPolicyRequest) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *ModifyBackupPolicyRequest) GetLocalLogRetentionHours() *string {
	return s.LocalLogRetentionHours
}

func (s *ModifyBackupPolicyRequest) GetLocalLogRetentionSpace() *string {
	return s.LocalLogRetentionSpace
}

func (s *ModifyBackupPolicyRequest) GetLogBackupFrequency() *string {
	return s.LogBackupFrequency
}

func (s *ModifyBackupPolicyRequest) GetLogBackupLocalRetentionNumber() *int32 {
	return s.LogBackupLocalRetentionNumber
}

func (s *ModifyBackupPolicyRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetReleasedKeepPolicy() *string {
	return s.ReleasedKeepPolicy
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupPolicyRequest) SetAdvancedDataPolicies(v []*ModifyBackupPolicyRequestAdvancedDataPolicies) *ModifyBackupPolicyRequest {
	s.AdvancedDataPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetAdvancedLogPolicies(v []*ModifyBackupPolicyRequestAdvancedLogPolicies) *ModifyBackupPolicyRequest {
	s.AdvancedLogPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupKeepCount(v int32) *ModifyBackupPolicyRequest {
	s.ArchiveBackupKeepCount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupKeepPolicy(v string) *ModifyBackupPolicyRequest {
	s.ArchiveBackupKeepPolicy = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetArchiveBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.ArchiveBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupInterval(v string) *ModifyBackupPolicyRequest {
	s.BackupInterval = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupLog(v string) *ModifyBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupMethod(v string) *ModifyBackupPolicyRequest {
	s.BackupMethod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPolicyMode(v string) *ModifyBackupPolicyRequest {
	s.BackupPolicyMode = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPriority(v int32) *ModifyBackupPolicyRequest {
	s.BackupPriority = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCategory(v string) *ModifyBackupPolicyRequest {
	s.Category = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCompressType(v string) *ModifyBackupPolicyRequest {
	s.CompressType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableAdvancedBackupPolicy(v int32) *ModifyBackupPolicyRequest {
	s.EnableAdvancedBackupPolicy = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableBackupLog(v string) *ModifyBackupPolicyRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyRequest {
	s.EnableIncrementDataBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyRequest {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLocalLogRetentionHours(v string) *ModifyBackupPolicyRequest {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyRequest {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupFrequency(v string) *ModifyBackupPolicyRequest {
	s.LogBackupFrequency = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyRequest {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetReleasedKeepPolicy(v string) *ModifyBackupPolicyRequest {
	s.ReleasedKeepPolicy = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	if s.AdvancedDataPolicies != nil {
		for _, item := range s.AdvancedDataPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AdvancedLogPolicies != nil {
		for _, item := range s.AdvancedLogPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyBackupPolicyRequestAdvancedDataPolicies struct {
	ActionType              *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	BakType                 *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion              *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType                *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	FilterKey               *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType              *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue             *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	OnlyPreserveOneEachDay  *bool   `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	OnlyPreserveOneEachHour *bool   `json:"OnlyPreserveOneEachHour,omitempty" xml:"OnlyPreserveOneEachHour,omitempty"`
	RetentionType           *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue          *int32  `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion               *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType                 *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StrategyId              *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvancedDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetBakType() *string {
	return s.BakType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetOnlyPreserveOneEachHour() *bool {
	return s.OnlyPreserveOneEachHour
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetRetentionValue() *int32 {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetBakType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.BakType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetOnlyPreserveOneEachHour(v bool) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.OnlyPreserveOneEachHour = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetRetentionValue(v int32) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) SetStrategyId(v string) *ModifyBackupPolicyRequestAdvancedDataPolicies {
	s.StrategyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedDataPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyRequestAdvancedLogPolicies struct {
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DestRegion        *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType          *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	EnableLogBackup   *int32  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	FilterKey         *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterValue       *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	LogRetentionType  *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	LogRetentionValue *int32  `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	SrcRegion         *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType           *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvancedLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvancedLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetLogRetentionValue() *int32 {
	return s.LogRetentionValue
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetEnableLogBackup(v int32) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetLogRetentionValue(v int32) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) SetStrategyId(v string) *ModifyBackupPolicyRequestAdvancedLogPolicies {
	s.StrategyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvancedLogPolicies) Validate() error {
	return dara.Validate(s)
}
