// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBackupPolicyResponseBody
	GetCode() *string
	SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody
	GetData() *DescribeBackupPolicyResponseBodyData
	SetErrCode(v string) *DescribeBackupPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupPolicyResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeBackupPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeBackupPolicyResponseBody
	GetSuccess() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the backup policy.
	Data *DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// instanceName can not be empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54A63B3B-AA10-1CC3-A6BB-6CCE98D19628
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBackupPolicyResponseBody) GetData() *DescribeBackupPolicyResponseBodyData {
	return s.Data
}

func (s *DescribeBackupPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeBackupPolicyResponseBody) SetCode(v string) *DescribeBackupPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetErrCode(v string) *DescribeBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetErrMessage(v string) *DescribeBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetMessage(v string) *DescribeBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v string) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyData struct {
	// The details of data backup policies.
	AdvanceDataPolicies []*DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies `json:"AdvanceDataPolicies,omitempty" xml:"AdvanceDataPolicies,omitempty" type:"Repeated"`
	AdvanceIncPolicies  []*DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies  `json:"AdvanceIncPolicies,omitempty" xml:"AdvanceIncPolicies,omitempty" type:"Repeated"`
	// The details of log backup policies.
	AdvanceLogPolicies []*DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies `json:"AdvanceLogPolicies,omitempty" xml:"AdvanceLogPolicies,omitempty" type:"Repeated"`
	// The method that is used to generate backup files. Valid values:
	//
	// 	- **Physical**: physical backup
	//
	// 	- **Snapshot**: snapshot backup
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// Indicates whether the secondary database is preferentially backed up. Valid values:
	//
	// 	- **1**: The secondary database is preferentially backed up.
	//
	// 	- **2**: Only the primary database is backed up.
	//
	// >  This parameter is returned only for ApsaraDB RDS for SQL Server instances. If the instance runs a different database engine, **0*	- is returned.
	//
	// example:
	//
	// 0
	BackupPriority *int32 `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	// The retention period of basic backup files. If an advanced backup policy is enabled, the value indicates the retention period of the level-1 backup policy.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The policy that is used to retain the archived backup files if the instance is deleted. Valid values:
	//
	// 	- **NONE**: No archived backup file is retained.
	//
	// 	- **LATEST**: Only the latest archived backup file is retained.
	//
	// 	- **ALL**: All archived backup files are retained.
	//
	// example:
	//
	// LATEST
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// Indicates whether the second-granularity backup feature is enabled. Valid values:
	//
	// 	- **Flash**: The second-granularity backup feature is enabled.
	//
	// 	- **Standard**: The standard backup feature is enabled.
	//
	// >  This parameter is returned only for MySQL instances.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether the backup feature is enabled.
	//
	// example:
	//
	// 1
	EnableBackup *int32 `json:"EnableBackup,omitempty" xml:"EnableBackup,omitempty"`
	// Indicates whether the incremental backup feature is enabled.
	//
	// example:
	//
	// 0
	EnableIncBackup *int32 `json:"EnableIncBackup,omitempty" xml:"EnableIncBackup,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **1**: The log backup feature is enabled.
	//
	// 	- **0**: The log backup feature is disabled.
	//
	// example:
	//
	// 1
	EnableLogBackup *int32 `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	// The interval for high-frequency backups. Unit: minutes. For example, a value of 120 indicates that a backup is performed every two hours.
	//
	// example:
	//
	// 120
	HighFrequencyBakInterval *int32 `json:"HighFrequencyBakInterval,omitempty" xml:"HighFrequencyBakInterval,omitempty"`
	// Indicates whether logs are forcibly cleared if the storage usage of the instance is greater than 80% or the available storage space of the instance is smaller than 5 GB. Valid values:
	//
	// 	- **Disable**: Logs are not forcibly cleared.
	//
	// 	- **Enable**: Logs are forcibly cleared.
	//
	// example:
	//
	// Enable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The interval for high-frequency incremental backups.
	//
	// example:
	//
	// -1
	IncBackupInterval *int32 `json:"IncBackupInterval,omitempty" xml:"IncBackupInterval,omitempty"`
	// The maximum storage usage that is allowed for log files on the instance.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *int32 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The number of local binary log files that are retained.
	//
	// example:
	//
	// 10
	LogBackupLocalRetentionNumber *string `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The retention period of log backups.
	//
	// example:
	//
	// 7
	LogBackupRetention *int32 `json:"LogBackupRetention,omitempty" xml:"LogBackupRetention,omitempty"`
	// The backup cycle of a basic backup. The backup cycle is represented by a seven-digit string, which corresponds to the days of the week from Monday to Sunday. A value of 1 indicates that a basic backup is performed on that day, and a value of 0 indicates that no basic backup is performed on that day.
	//
	// example:
	//
	// 1010101
	PreferredBackupDate *string `json:"PreferredBackupDate,omitempty" xml:"PreferredBackupDate,omitempty"`
	// The time period during which a basic backup is performed.
	//
	// example:
	//
	// 23:00Z-24:00Z
	PreferredBackupWindow *string `json:"PreferredBackupWindow,omitempty" xml:"PreferredBackupWindow,omitempty"`
	// The start time of a basic backup.
	//
	// example:
	//
	// 23:00Z
	PreferredBackupWindowBegin *string `json:"PreferredBackupWindowBegin,omitempty" xml:"PreferredBackupWindowBegin,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyData) GetAdvanceDataPolicies() []*DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	return s.AdvanceDataPolicies
}

func (s *DescribeBackupPolicyResponseBodyData) GetAdvanceIncPolicies() []*DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	return s.AdvanceIncPolicies
}

func (s *DescribeBackupPolicyResponseBodyData) GetAdvanceLogPolicies() []*DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	return s.AdvanceLogPolicies
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *DescribeBackupPolicyResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *DescribeBackupPolicyResponseBodyData) GetEnableBackup() *int32 {
	return s.EnableBackup
}

func (s *DescribeBackupPolicyResponseBodyData) GetEnableIncBackup() *int32 {
	return s.EnableIncBackup
}

func (s *DescribeBackupPolicyResponseBodyData) GetEnableLogBackup() *int32 {
	return s.EnableLogBackup
}

func (s *DescribeBackupPolicyResponseBodyData) GetHighFrequencyBakInterval() *int32 {
	return s.HighFrequencyBakInterval
}

func (s *DescribeBackupPolicyResponseBodyData) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *DescribeBackupPolicyResponseBodyData) GetIncBackupInterval() *int32 {
	return s.IncBackupInterval
}

func (s *DescribeBackupPolicyResponseBodyData) GetLocalLogRetentionSpace() *int32 {
	return s.LocalLogRetentionSpace
}

func (s *DescribeBackupPolicyResponseBodyData) GetLogBackupLocalRetentionNumber() *string {
	return s.LogBackupLocalRetentionNumber
}

func (s *DescribeBackupPolicyResponseBodyData) GetLogBackupRetention() *int32 {
	return s.LogBackupRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetPreferredBackupDate() *string {
	return s.PreferredBackupDate
}

func (s *DescribeBackupPolicyResponseBodyData) GetPreferredBackupWindow() *string {
	return s.PreferredBackupWindow
}

func (s *DescribeBackupPolicyResponseBodyData) GetPreferredBackupWindowBegin() *string {
	return s.PreferredBackupWindowBegin
}

func (s *DescribeBackupPolicyResponseBodyData) SetAdvanceDataPolicies(v []*DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) *DescribeBackupPolicyResponseBodyData {
	s.AdvanceDataPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetAdvanceIncPolicies(v []*DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) *DescribeBackupPolicyResponseBodyData {
	s.AdvanceIncPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetAdvanceLogPolicies(v []*DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) *DescribeBackupPolicyResponseBodyData {
	s.AdvanceLogPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupMethod(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPriority(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupPriority = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCategory(v string) *DescribeBackupPolicyResponseBodyData {
	s.Category = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableIncBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableIncBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableLogBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetHighFrequencyBakInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.HighFrequencyBakInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetHighSpaceUsageProtection(v string) *DescribeBackupPolicyResponseBodyData {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIncBackupInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.IncBackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetentionSpace(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogBackupLocalRetentionNumber(v string) *DescribeBackupPolicyResponseBodyData {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LogBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupDate(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupDate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupWindow(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupWindow = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupWindowBegin(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupWindowBegin = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) Validate() error {
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

type DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies struct {
	// Indicates whether the backup policy is generated by the system. Valid values:
	//
	// 	- **true**: The backup policy is generated by the system.
	//
	// 	- **false**: The backup policy is a custom one.
	//
	// example:
	//
	// true
	AutoCreated *bool `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **F**: full backup
	//
	// 	- **L**: log backup
	//
	// example:
	//
	// F
	BakType *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	// The region in which the backup files are stored.
	//
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The storage method of backup files. Valid values:
	//
	// 	- **db**: database
	//
	// 	- **level1**: level-1 backup
	//
	// 	- **level2**: level-2 backup
	//
	// 	- **level2Cross**: level-2 cross-region backup
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// The information about the dump policy. Valid values:
	//
	// 	- **copy**: The backup data is copied from the data source to the destination.
	//
	// 	- **move**: The backup data is moved from the data source to the destination.
	//
	// example:
	//
	// copy
	DumpAction *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
	// The scheduling frequency. Valid values:
	//
	// 	- **dayOfWeek**: scheduling by week
	//
	// 	- **dayOfMonth**: scheduling by month
	//
	// 	- **dayOfYear**: scheduling by year
	//
	// 	- **backupInterval**: scheduling at a specific interval
	//
	// >  This parameter is returned only if the value of FilterType is **crontab**.
	//
	// example:
	//
	// dayOfWeek
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The scheduling mode of the advanced backup policy. Valid values:
	//
	// 	- **crontab**: periodic scheduling
	//
	// 	- **event**: event-based scheduling
	//
	// example:
	//
	// crontab
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// The day of a week on which you want to back up data.
	//
	// example:
	//
	// 1,2,3,4,5,6,7
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The ID of the advanced backup policy.
	//
	// example:
	//
	// 71930ac2e9f15e41615e10627c******
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
	// The validity period, in days.
	//
	// example:
	//
	// 7
	RetentionValue *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	// The region in which the data source of the backup policy resides.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the data source of the backup policy. Valid values:
	//
	// 	- **db**: database
	//
	// 	- **level1**: level-1 backup
	//
	// 	- **level2**: level-2 backup
	//
	// 	- **level2Cross**: level-2 cross-region backup
	//
	// example:
	//
	// db
	SrcType      *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetBakType() *string {
	return s.BakType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDumpAction() *string {
	return s.DumpAction
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetAutoCreated(v bool) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.AutoCreated = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetBakType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.BakType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDumpAction(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DumpAction = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetPolicyId(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetStorageClass(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.StorageClass = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies struct {
	AutoCreated    *bool   `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	BakType        *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion     *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType       *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	DumpAction     *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
	FilterKey      *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType     *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue    *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	PolicyId       *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RetentionType  *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	RetentionValue *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	SrcRegion      *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	SrcType        *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetBakType() *string {
	return s.BakType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDumpAction() *string {
	return s.DumpAction
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetAutoCreated(v bool) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.AutoCreated = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetBakType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.BakType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDestType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDumpAction(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DumpAction = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetPolicyId(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetRetentionType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.RetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetRetentionValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.RetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) SetSrcType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceIncPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies struct {
	// The region in which the backup files are stored.
	//
	// example:
	//
	// cn-shanghai
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The storage method of backup files. Valid values:
	//
	// 	- **db**: database
	//
	// 	- **level1**: level-1 backup
	//
	// 	- **level2**: level-2 backup
	//
	// 	- **level2Cross**: level-2 cross-region backup
	//
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 1
	EnableLogBackup *int64  `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	FilterKey       *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterType      *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	FilterValue     *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The retention type of log backups. Valid values:
	//
	// 	- **never**: Log backups never expire.
	//
	// 	- **delay**: Log backups are retained for a specific number of days.
	//
	// example:
	//
	// delay
	LogRetentionType *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	// The retention period of log backups.
	//
	// example:
	//
	// 3
	LogRetentionValue *string `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	// The backup policy ID.
	//
	// example:
	//
	// 6s67c7i3y8f8p72808p******
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region in which the data source of the backup policy resides.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the data source of the backup policy. Valid values:
	//
	// 	- **db**: database
	//
	// 	- **level1**: level-1 backup
	//
	// 	- **level2**: level-2 backup
	//
	// 	- **level2Cross**: level-2 cross-region backup
	//
	// example:
	//
	// level1
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetEnableLogBackup() *int64 {
	return s.EnableLogBackup
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetLogRetentionValue() *string {
	return s.LogRetentionValue
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetEnableLogBackup(v int64) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetPolicyId(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) Validate() error {
	return dara.Validate(s)
}
