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
	Data *UpdateBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BackupPeriod                   *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	BackupPlanBegin                *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	BackupSetRetention             *int32  `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	BackupType                     *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupWay                      *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	ColdDataBackupInterval         *int32  `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	ColdDataBackupRetention        *int32  `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	CrossRegionDataBackupRetention *int32  `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	CrossRegionLogBackupRetention  *int32  `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	DBInstanceName                 *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion                *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	ForceCleanOnHighSpaceUsage     *int32  `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsCrossRegionDataBackupEnabled *bool   `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	IsCrossRegionLogBackupEnabled  *bool   `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	IsEnabled                      *int32  `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	LocalLogRetention              *int32  `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LocalLogRetentionNumber        *int32  `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	LogLocalRetentionSpace         *int32  `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	RemoveLogRetention             *int32  `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
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
