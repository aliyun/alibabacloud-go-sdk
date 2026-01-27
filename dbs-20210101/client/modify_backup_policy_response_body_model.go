// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyBackupPolicyResponseBody
	GetCode() *string
	SetData(v *ModifyBackupPolicyResponseBodyData) *ModifyBackupPolicyResponseBody
	GetData() *ModifyBackupPolicyResponseBodyData
	SetErrCode(v string) *ModifyBackupPolicyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyBackupPolicyResponseBody
	GetErrMessage() *string
	SetMessage(v string) *ModifyBackupPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyBackupPolicyResponseBody
	GetSuccess() *string
}

type ModifyBackupPolicyResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the backup policy.
	Data *ModifyBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
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
	// D570F209-A166-50C6-98A3-155A20B218B7
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

func (s ModifyBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBackupPolicyResponseBody) GetData() *ModifyBackupPolicyResponseBodyData {
	return s.Data
}

func (s *ModifyBackupPolicyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyBackupPolicyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyBackupPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPolicyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyBackupPolicyResponseBody) SetCode(v string) *ModifyBackupPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetData(v *ModifyBackupPolicyResponseBodyData) *ModifyBackupPolicyResponseBody {
	s.Data = v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetErrCode(v string) *ModifyBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetErrMessage(v string) *ModifyBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetMessage(v string) *ModifyBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetSuccess(v string) *ModifyBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBackupPolicyResponseBodyData struct {
	// The details of data backup policies.
	AdvanceDataPolicies                    []*ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies `json:"AdvanceDataPolicies,omitempty" xml:"AdvanceDataPolicies,omitempty" type:"Repeated"`
	AdvanceIncPolicies                     []*ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies  `json:"AdvanceIncPolicies,omitempty" xml:"AdvanceIncPolicies,omitempty" type:"Repeated"`
	AdvanceLogPolicies                     []*ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies  `json:"AdvanceLogPolicies,omitempty" xml:"AdvanceLogPolicies,omitempty" type:"Repeated"`
	BackupMethod                           *string                                                  `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupPriority                         *int32                                                   `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string                                                  `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	Category                               *string                                                  `json:"Category,omitempty" xml:"Category,omitempty"`
	EnableIncBackup                        *bool                                                    `json:"EnableIncBackup,omitempty" xml:"EnableIncBackup,omitempty"`
	// The time period during which a basic backup is performed.
	//
	// example:
	//
	// 17:00Z-18:00Z
	PreferredBackupWindow *string `json:"PreferredBackupWindow,omitempty" xml:"PreferredBackupWindow,omitempty"`
	// The start time of a basic backup.
	//
	// example:
	//
	// 17:00Z
	PreferredBackupWindowBegin *string `json:"PreferredBackupWindowBegin,omitempty" xml:"PreferredBackupWindowBegin,omitempty"`
}

func (s ModifyBackupPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBodyData) GetAdvanceDataPolicies() []*ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	return s.AdvanceDataPolicies
}

func (s *ModifyBackupPolicyResponseBodyData) GetAdvanceIncPolicies() []*ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	return s.AdvanceIncPolicies
}

func (s *ModifyBackupPolicyResponseBodyData) GetAdvanceLogPolicies() []*ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	return s.AdvanceLogPolicies
}

func (s *ModifyBackupPolicyResponseBodyData) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ModifyBackupPolicyResponseBodyData) GetBackupPriority() *int32 {
	return s.BackupPriority
}

func (s *ModifyBackupPolicyResponseBodyData) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *ModifyBackupPolicyResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ModifyBackupPolicyResponseBodyData) GetEnableIncBackup() *bool {
	return s.EnableIncBackup
}

func (s *ModifyBackupPolicyResponseBodyData) GetPreferredBackupWindow() *string {
	return s.PreferredBackupWindow
}

func (s *ModifyBackupPolicyResponseBodyData) GetPreferredBackupWindowBegin() *string {
	return s.PreferredBackupWindowBegin
}

func (s *ModifyBackupPolicyResponseBodyData) SetAdvanceDataPolicies(v []*ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) *ModifyBackupPolicyResponseBodyData {
	s.AdvanceDataPolicies = v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetAdvanceIncPolicies(v []*ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) *ModifyBackupPolicyResponseBodyData {
	s.AdvanceIncPolicies = v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetAdvanceLogPolicies(v []*ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) *ModifyBackupPolicyResponseBodyData {
	s.AdvanceLogPolicies = v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetBackupMethod(v string) *ModifyBackupPolicyResponseBodyData {
	s.BackupMethod = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetBackupPriority(v int32) *ModifyBackupPolicyResponseBodyData {
	s.BackupPriority = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyResponseBodyData {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetCategory(v string) *ModifyBackupPolicyResponseBodyData {
	s.Category = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetEnableIncBackup(v bool) *ModifyBackupPolicyResponseBodyData {
	s.EnableIncBackup = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetPreferredBackupWindow(v string) *ModifyBackupPolicyResponseBodyData {
	s.PreferredBackupWindow = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) SetPreferredBackupWindowBegin(v string) *ModifyBackupPolicyResponseBodyData {
	s.PreferredBackupWindowBegin = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyData) Validate() error {
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

type ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies struct {
	// Indicates whether the backup policy is generated by the system. Valid values:
	//
	// 	- **true**: The backup policy is generated by the system.
	//
	// 	- **false**: The backup policy is a custom one.
	//
	// example:
	//
	// false
	AutoCreated *bool `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **F**: full backup.
	//
	// 	- **L**: log backup.
	//
	// example:
	//
	// F
	BakType *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
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
	// >  This parameter is returned only when FilterType is set to **crontab**.
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
	// The ID of the advanced backup policy.
	//
	// example:
	//
	// dc13b153acc91141789122c23835****
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

func (s ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetBakType() *string {
	return s.BakType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetDumpAction() *string {
	return s.DumpAction
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetAutoCreated(v bool) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.AutoCreated = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetBakType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.BakType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDumpAction(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DumpAction = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterKey(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetPolicyId(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) SetStorageClass(v string) *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.StorageClass = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceDataPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies struct {
	AutoCreated            *bool   `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	BakType                *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	DestRegion             *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	DestType               *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	DumpAction             *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
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

func (s ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetAutoCreated() *bool {
	return s.AutoCreated
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetBakType() *string {
	return s.BakType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetDumpAction() *string {
	return s.DumpAction
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetOnlyPreserveOneEachDay() *bool {
	return s.OnlyPreserveOneEachDay
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetRetentionType() *string {
	return s.RetentionType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetRetentionValue() *string {
	return s.RetentionValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetAutoCreated(v bool) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.AutoCreated = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetBakType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.BakType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDestRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDestType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetDumpAction(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.DumpAction = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterKey(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetFilterValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetOnlyPreserveOneEachDay(v bool) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.OnlyPreserveOneEachDay = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetPolicyId(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetRetentionType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.RetentionType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetRetentionValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.RetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetSrcRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) SetSrcType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceIncPolicies) Validate() error {
	return dara.Validate(s)
}

type ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies struct {
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

func (s ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetDestRegion() *string {
	return s.DestRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetDestType() *string {
	return s.DestType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetEnableLogBackup() *int64 {
	return s.EnableLogBackup
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterKey() *string {
	return s.FilterKey
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterType() *string {
	return s.FilterType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetFilterValue() *string {
	return s.FilterValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetLogRetentionType() *string {
	return s.LogRetentionType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetLogRetentionValue() *string {
	return s.LogRetentionValue
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) GetSrcType() *string {
	return s.SrcType
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetEnableLogBackup(v int64) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterKey(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterKey = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetFilterValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.FilterValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionValue(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetPolicyId(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.PolicyId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcRegion(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcType(v string) *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBodyDataAdvanceLogPolicies) Validate() error {
	return dara.Validate(s)
}
