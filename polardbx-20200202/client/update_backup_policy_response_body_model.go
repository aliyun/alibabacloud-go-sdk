// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateBackupPolicyResponseBodyData) *UpdateBackupPolicyResponseBody
	GetData() *UpdateBackupPolicyResponseBodyData
	SetMessage(v string) *UpdateBackupPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBackupPolicyResponseBody
	GetSuccess() *bool
}

type UpdateBackupPolicyResponseBody struct {
	// The returned policy details.
	Data *UpdateBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. If the call is successful, success is returned. If the call fails, the corresponding error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBody) GetData() *UpdateBackupPolicyResponseBodyData {
	return s.Data
}

func (s *UpdateBackupPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBackupPolicyResponseBody) SetData(v *UpdateBackupPolicyResponseBodyData) *UpdateBackupPolicyResponseBody {
	s.Data = v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetMessage(v string) *UpdateBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetRequestId(v string) *UpdateBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) SetSuccess(v bool) *UpdateBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupPolicyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBackupPolicyResponseBodyData struct {
	// The backup cycle. Specify at least two days. The value is a 7-digit string, where each digit from left to right represents whether backups are enabled for Monday through Sunday. A value of 0 indicates that backups are disabled, and a value of 1 indicates that backups are enabled:
	//
	// - First digit: Monday
	//
	// - Second digit: Tuesday
	//
	// - Third digit: Wednesday
	//
	// - Fourth digit: Thursday
	//
	// - Fifth digit: Friday
	//
	// - Sixth digit: Saturday
	//
	// - Seventh digit: Sunday
	//
	// example:
	//
	// 1001000
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The time when the daily backup starts, in UTC.
	//
	// example:
	//
	// 03:00Z
	BackupPlanBegin *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	// The retention period of the backup set, in days.
	//
	// example:
	//
	// 30
	BackupSetRetention *int32 `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	// The backup type. A value of 0 indicates a quick backup.
	//
	// example:
	//
	// 0
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The backup method. Valid values:
	//
	// - **P**: physical backup
	//
	// - **L**: logical backup
	//
	// example:
	//
	// P
	BackupWay *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	// The interval at which cold data is backed up, in days. Valid values: 1 to 59.
	//
	// example:
	//
	// 30
	ColdDataBackupInterval *int32 `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	// The retention period of cold data backups, in days.
	//
	// example:
	//
	// 30
	ColdDataBackupRetention *int32 `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	// The retention period of cross-region data backups, in days.
	//
	// example:
	//
	// 30
	CrossRegionDataBackupRetention *int32 `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	// The number of days selected for cross-region backups.
	//
	// example:
	//
	// 3,7,1,5,6,2,4
	CrossRegionFilterValue *string `json:"CrossRegionFilterValue,omitempty" xml:"CrossRegionFilterValue,omitempty"`
	// The retention period of cross-region log backups, in days.
	//
	// example:
	//
	// 30
	CrossRegionLogBackupRetention *int32 `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The destination region for cross-region backups.
	//
	// example:
	//
	// cn-shanghai
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// Specifies whether to forcibly clear binary logs when the storage usage of the instance is excessively high. A value of 1 indicates yes, and a value of 0 indicates no.
	//
	// example:
	//
	// 1
	ForceCleanOnHighSpaceUsage *int32 `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	// Specifies whether to enable cross-region data backups. Default value: false.
	IsCrossRegionDataBackupEnabled *bool `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	// Specifies whether to enable cross-region log backups. Default value: false.
	IsCrossRegionLogBackupEnabled *bool `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	// Specifies whether to enable backups. The fixed value is 1, which indicates that backups are enabled.
	//
	// example:
	//
	// 1
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	// The local retention period of log backups, in hours. Valid values: 0 to 168. A value of 0 indicates that log backups are not retained. Default value: 7.
	//
	// example:
	//
	// 7
	LocalLogRetention *int32 `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	// The number of local binary logs to retain. Valid values: 6 to 100. Default value: 60.
	//
	// example:
	//
	// 60
	LocalLogRetentionNumber *int32 `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	// The maximum space usage for binary logs. The parameter value ranges from 0 to 50. This is a loop space. Default value: 30.
	//
	// example:
	//
	// 30
	LogLocalRetentionSpace *int32 `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	// The remote retention period of log backups, in days. Valid values: 7 to 730. Default value: 7.
	//
	// example:
	//
	// 7
	RemoveLogRetention *int32 `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
}

func (s UpdateBackupPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyResponseBodyData) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *UpdateBackupPolicyResponseBodyData) GetBackupPlanBegin() *string {
	return s.BackupPlanBegin
}

func (s *UpdateBackupPolicyResponseBodyData) GetBackupSetRetention() *int32 {
	return s.BackupSetRetention
}

func (s *UpdateBackupPolicyResponseBodyData) GetBackupType() *string {
	return s.BackupType
}

func (s *UpdateBackupPolicyResponseBodyData) GetBackupWay() *string {
	return s.BackupWay
}

func (s *UpdateBackupPolicyResponseBodyData) GetColdDataBackupInterval() *int32 {
	return s.ColdDataBackupInterval
}

func (s *UpdateBackupPolicyResponseBodyData) GetColdDataBackupRetention() *int32 {
	return s.ColdDataBackupRetention
}

func (s *UpdateBackupPolicyResponseBodyData) GetCrossRegionDataBackupRetention() *int32 {
	return s.CrossRegionDataBackupRetention
}

func (s *UpdateBackupPolicyResponseBodyData) GetCrossRegionFilterValue() *string {
	return s.CrossRegionFilterValue
}

func (s *UpdateBackupPolicyResponseBodyData) GetCrossRegionLogBackupRetention() *int32 {
	return s.CrossRegionLogBackupRetention
}

func (s *UpdateBackupPolicyResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdateBackupPolicyResponseBodyData) GetDestCrossRegion() *string {
	return s.DestCrossRegion
}

func (s *UpdateBackupPolicyResponseBodyData) GetForceCleanOnHighSpaceUsage() *int32 {
	return s.ForceCleanOnHighSpaceUsage
}

func (s *UpdateBackupPolicyResponseBodyData) GetIsCrossRegionDataBackupEnabled() *bool {
	return s.IsCrossRegionDataBackupEnabled
}

func (s *UpdateBackupPolicyResponseBodyData) GetIsCrossRegionLogBackupEnabled() *bool {
	return s.IsCrossRegionLogBackupEnabled
}

func (s *UpdateBackupPolicyResponseBodyData) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdateBackupPolicyResponseBodyData) GetLocalLogRetention() *int32 {
	return s.LocalLogRetention
}

func (s *UpdateBackupPolicyResponseBodyData) GetLocalLogRetentionNumber() *int32 {
	return s.LocalLogRetentionNumber
}

func (s *UpdateBackupPolicyResponseBodyData) GetLogLocalRetentionSpace() *int32 {
	return s.LogLocalRetentionSpace
}

func (s *UpdateBackupPolicyResponseBodyData) GetRemoveLogRetention() *int32 {
	return s.RemoveLogRetention
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupPeriod(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupPlanBegin(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupPlanBegin = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupSetRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.BackupSetRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupType(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetBackupWay(v string) *UpdateBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetColdDataBackupInterval(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetColdDataBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetCrossRegionDataBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetCrossRegionFilterValue(v string) *UpdateBackupPolicyResponseBodyData {
	s.CrossRegionFilterValue = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetCrossRegionLogBackupRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDBInstanceName(v string) *UpdateBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetDestCrossRegion(v string) *UpdateBackupPolicyResponseBodyData {
	s.DestCrossRegion = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsCrossRegionDataBackupEnabled(v bool) *UpdateBackupPolicyResponseBodyData {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsCrossRegionLogBackupEnabled(v bool) *UpdateBackupPolicyResponseBodyData {
	s.IsCrossRegionLogBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetIsEnabled(v int32) *UpdateBackupPolicyResponseBodyData {
	s.IsEnabled = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLocalLogRetentionNumber(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LocalLogRetentionNumber = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *UpdateBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
