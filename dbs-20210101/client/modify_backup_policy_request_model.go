// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvanceDataPolicies(v []*ModifyBackupPolicyRequestAdvanceDataPolicies) *ModifyBackupPolicyRequest
	GetAdvanceDataPolicies() []*ModifyBackupPolicyRequestAdvanceDataPolicies
	SetAdvanceIncPolicies(v []*ModifyBackupPolicyRequestAdvanceIncPolicies) *ModifyBackupPolicyRequest
	GetAdvanceIncPolicies() []*ModifyBackupPolicyRequestAdvanceIncPolicies
	SetAdvanceLogPolicies(v []*ModifyBackupPolicyRequestAdvanceLogPolicies) *ModifyBackupPolicyRequest
	GetAdvanceLogPolicies() []*ModifyBackupPolicyRequestAdvanceLogPolicies
	SetBackupMethod(v string) *ModifyBackupPolicyRequest
	GetBackupMethod() *string
	SetBackupPriority(v int32) *ModifyBackupPolicyRequest
	GetBackupPriority() *int32
	SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetCategory(v string) *ModifyBackupPolicyRequest
	GetCategory() *string
	SetEnableIncBackup(v bool) *ModifyBackupPolicyRequest
	GetEnableIncBackup() *bool
	SetInstanceName(v string) *ModifyBackupPolicyRequest
	GetInstanceName() *string
	SetPreferredBackupWindowBegin(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupWindowBegin() *string
	SetRegionCode(v string) *ModifyBackupPolicyRequest
	GetRegionCode() *string
}

type ModifyBackupPolicyRequest struct {
	// The details of data backup policies.
	AdvanceDataPolicies                    []*ModifyBackupPolicyRequestAdvanceDataPolicies `json:"AdvanceDataPolicies,omitempty" xml:"AdvanceDataPolicies,omitempty" type:"Repeated"`
	AdvanceIncPolicies                     []*ModifyBackupPolicyRequestAdvanceIncPolicies  `json:"AdvanceIncPolicies,omitempty" xml:"AdvanceIncPolicies,omitempty" type:"Repeated"`
	AdvanceLogPolicies                     []*ModifyBackupPolicyRequestAdvanceLogPolicies  `json:"AdvanceLogPolicies,omitempty" xml:"AdvanceLogPolicies,omitempty" type:"Repeated"`
	BackupMethod                           *string                                         `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupPriority                         *int32                                          `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string                                         `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	Category                               *string                                         `json:"Category,omitempty" xml:"Category,omitempty"`
	EnableIncBackup                        *bool                                           `json:"EnableIncBackup,omitempty" xml:"EnableIncBackup,omitempty"`
	// The ID of the PolarDB instance.
	//
	// example:
	//
	// pc-2ze3nrr64c5****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The start time of a basic backup.
	//
	// example:
	//
	// 17:00Z
	PreferredBackupWindowBegin *string `json:"PreferredBackupWindowBegin,omitempty" xml:"PreferredBackupWindowBegin,omitempty"`
	// The region in which backup sets reside.
	//
	// example:
	//
	// cn-shanghai
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetAdvanceDataPolicies() []*ModifyBackupPolicyRequestAdvanceDataPolicies {
	return s.AdvanceDataPolicies
}

func (s *ModifyBackupPolicyRequest) GetAdvanceIncPolicies() []*ModifyBackupPolicyRequestAdvanceIncPolicies {
	return s.AdvanceIncPolicies
}

func (s *ModifyBackupPolicyRequest) GetAdvanceLogPolicies() []*ModifyBackupPolicyRequestAdvanceLogPolicies {
	return s.AdvanceLogPolicies
}

func (s *ModifyBackupPolicyRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ModifyBackupPolicyRequest) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyBackupPolicyRequest) GetEnableIncBackup() *bool {
	return s.EnableIncBackup
}

func (s *ModifyBackupPolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupWindowBegin() *string {
	return s.PreferredBackupWindowBegin
}

func (s *ModifyBackupPolicyRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ModifyBackupPolicyRequest) SetAdvanceDataPolicies(v []*ModifyBackupPolicyRequestAdvanceDataPolicies) *ModifyBackupPolicyRequest {
	s.AdvanceDataPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetAdvanceIncPolicies(v []*ModifyBackupPolicyRequestAdvanceIncPolicies) *ModifyBackupPolicyRequest {
	s.AdvanceIncPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetAdvanceLogPolicies(v []*ModifyBackupPolicyRequestAdvanceLogPolicies) *ModifyBackupPolicyRequest {
	s.AdvanceLogPolicies = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupMethod(v string) *ModifyBackupPolicyRequest {
	s.BackupMethod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupPriority(v int32) *ModifyBackupPolicyRequest {
	s.BackupPriority = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetCategory(v string) *ModifyBackupPolicyRequest {
	s.Category = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableIncBackup(v bool) *ModifyBackupPolicyRequest {
	s.EnableIncBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetInstanceName(v string) *ModifyBackupPolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupWindowBegin(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupWindowBegin = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRegionCode(v string) *ModifyBackupPolicyRequest {
	s.RegionCode = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	if s.AdvanceDataPolicies != nil {
		for _, item := range s.AdvanceDataPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AdvanceIncPolicies != nil {
		for _, item := range s.AdvanceIncPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AdvanceLogPolicies != nil {
		for _, item := range s.AdvanceLogPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyBackupPolicyRequestAdvanceDataPolicies struct {
	// The type of the operation. Valid values:
	//
	// 	- **CREATE**
	//
	// 	- **UPDATE**
	//
	// 	- **DELETE**
	//
	// example:
	//
	// UPDATE
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The region in which backup files are stored.
	//
	// example:
	//
	// cn-shanghai
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The storage method of backup files. Valid values:
	//
	// 	- **db**: database.
	//
	// 	- **level1**: level-1 backup.
	//
	// 	- **level2**: level-2 backup.
	//
	// 	- **level2Cross**: level-2 cross-region backup.
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// The scheduling cycle. Valid values:
	//
	// 	- **dayOfWeek**: scheduling by week.
	//
	// 	- **dayOfMonth**: scheduling by month.
	//
	// 	- **dayOfYear**: scheduling by year.
	//
	// 	- **backupInterval**: scheduling at a specific interval.
	//
	// >  This parameter is required only when FilterType is set to **crontab**.
	//
	// example:
	//
	// backupInterval
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The scheduling mode of the advanced backup policy. Valid values:
	//
	// 	- **crontab**: periodic scheduling.
	//
	// 	- **event**: event-based scheduling.
	//
	// example:
	//
	// crontab
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The backup cycle.
	//
	// example:
	//
	// 180
	FilterValue            *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	OnlyPreserveOneEachDay *bool   `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	// The ID of the advanced backup policy. You can call the [DescribeBackupPolicy](https://help.aliyun.com/document_detail/2869783.html) operation to query the ID.
	//
	// example:
	//
	// 6s67c7i3y8f8p72808p******
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The retention type of backup sets. Valid values:
	//
	// 	- **never**: Backup sets never expire.
	//
	// 	- **delay**: Backup sets are retained for a specific number of days.
	//
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// The retention period. Unit: day.
	//
	// example:
	//
	// 4
	RetentionValue *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	// The region in which the data source of the backup policy resides.
	//
	// example:
	//
	// cn-shanghai
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the data source of the backup policy. Valid values:
	//
	// 	- **db**: database.
	//
	// 	- **level1**: level-1 backup.
	//
	// 	- **level2**: level-2 backup.
	//
	// 	- **level2Cross**: level-2 cross-region backup.
	//
	// example:
	//
	// db
	SrcType      *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvanceDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvanceDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetPolicyId(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetRetentionType(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetRetentionValue(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) SetStorageClass(v string) *ModifyBackupPolicyRequestAdvanceDataPolicies {
	s.StorageClass = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceDataPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyRequestAdvanceIncPolicies struct {
	ActionType             *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DestRegion             *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType               *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	FilterKey              *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType             *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue            *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	OnlyPreserveOneEachDay *bool   `json:"OnlyPreserveOneEachDay,omitempty" xml:"OnlyPreserveOneEachDay,omitempty"`
	PolicyId               *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RetentionType          *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue         *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion              *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType                *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvanceIncPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvanceIncPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetPolicyId(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetRetentionType(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetRetentionValue(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvanceIncPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceIncPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyRequestAdvanceLogPolicies struct {
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	DestRegion        *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType          *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	EnableLogBackup   *int64  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	FilterKey         *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType        *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue       *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	LogRetentionType  *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	LogRetentionValue *string `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	PolicyId          *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SrcRegion         *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType           *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s ModifyBackupPolicyRequestAdvanceLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequestAdvanceLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetActionType() *string {
	return s.ActionType
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetEnableLogBackup() *int64 {
	return s.EnableLogBackup
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetLogRetentionValue() *string {
	return s.LogRetentionValue
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetActionType(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.ActionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetDestRegion(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetDestType(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetEnableLogBackup(v int64) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetFilterKey(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetFilterType(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetFilterValue(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetLogRetentionType(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetLogRetentionValue(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetPolicyId(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetSrcRegion(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) SetSrcType(v string) *ModifyBackupPolicyRequestAdvanceLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyRequestAdvanceLogPolicies) Validate() error {
	return dara.Validate(s)
}
