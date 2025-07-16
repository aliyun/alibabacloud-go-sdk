// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPeriod(v string) *UpdateBackupPolicyRequest
	GetBackupPeriod() *string
	SetBackupPlanBegin(v string) *UpdateBackupPolicyRequest
	GetBackupPlanBegin() *string
	SetBackupSetRetention(v int32) *UpdateBackupPolicyRequest
	GetBackupSetRetention() *int32
	SetBackupType(v string) *UpdateBackupPolicyRequest
	GetBackupType() *string
	SetBackupWay(v string) *UpdateBackupPolicyRequest
	GetBackupWay() *string
	SetColdDataBackupInterval(v int32) *UpdateBackupPolicyRequest
	GetColdDataBackupInterval() *int32
	SetColdDataBackupRetention(v int32) *UpdateBackupPolicyRequest
	GetColdDataBackupRetention() *int32
	SetCrossRegionDataBackupRetention(v int32) *UpdateBackupPolicyRequest
	GetCrossRegionDataBackupRetention() *int32
	SetCrossRegionLogBackupRetention(v int32) *UpdateBackupPolicyRequest
	GetCrossRegionLogBackupRetention() *int32
	SetDBInstanceName(v string) *UpdateBackupPolicyRequest
	GetDBInstanceName() *string
	SetDestCrossRegion(v string) *UpdateBackupPolicyRequest
	GetDestCrossRegion() *string
	SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyRequest
	GetForceCleanOnHighSpaceUsage() *int32
	SetIsCrossRegionDataBackupEnabled(v bool) *UpdateBackupPolicyRequest
	GetIsCrossRegionDataBackupEnabled() *bool
	SetIsCrossRegionLogBackupEnabled(v bool) *UpdateBackupPolicyRequest
	GetIsCrossRegionLogBackupEnabled() *bool
	SetIsEnabled(v int32) *UpdateBackupPolicyRequest
	GetIsEnabled() *int32
	SetLocalLogRetention(v int32) *UpdateBackupPolicyRequest
	GetLocalLogRetention() *int32
	SetLocalLogRetentionNumber(v int32) *UpdateBackupPolicyRequest
	GetLocalLogRetentionNumber() *int32
	SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyRequest
	GetLogLocalRetentionSpace() *int32
	SetRegionId(v string) *UpdateBackupPolicyRequest
	GetRegionId() *string
	SetRemoveLogRetention(v int32) *UpdateBackupPolicyRequest
	GetRemoveLogRetention() *int32
}

type UpdateBackupPolicyRequest struct {
	// example:
	//
	// 1001000
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// example:
	//
	// 03:00Z
	BackupPlanBegin *string `json:"BackupPlanBegin,omitempty" xml:"BackupPlanBegin,omitempty"`
	// example:
	//
	// 30
	BackupSetRetention *int32 `json:"BackupSetRetention,omitempty" xml:"BackupSetRetention,omitempty"`
	// example:
	//
	// 0
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// P
	BackupWay *string `json:"BackupWay,omitempty" xml:"BackupWay,omitempty"`
	// example:
	//
	// 30
	ColdDataBackupInterval *int32 `json:"ColdDataBackupInterval,omitempty" xml:"ColdDataBackupInterval,omitempty"`
	// example:
	//
	// 30
	ColdDataBackupRetention        *int32 `json:"ColdDataBackupRetention,omitempty" xml:"ColdDataBackupRetention,omitempty"`
	CrossRegionDataBackupRetention *int32 `json:"CrossRegionDataBackupRetention,omitempty" xml:"CrossRegionDataBackupRetention,omitempty"`
	CrossRegionLogBackupRetention  *int32 `json:"CrossRegionLogBackupRetention,omitempty" xml:"CrossRegionLogBackupRetention,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasdyuoo
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DestCrossRegion *string `json:"DestCrossRegion,omitempty" xml:"DestCrossRegion,omitempty"`
	// example:
	//
	// 1
	ForceCleanOnHighSpaceUsage     *int32 `json:"ForceCleanOnHighSpaceUsage,omitempty" xml:"ForceCleanOnHighSpaceUsage,omitempty"`
	IsCrossRegionDataBackupEnabled *bool  `json:"IsCrossRegionDataBackupEnabled,omitempty" xml:"IsCrossRegionDataBackupEnabled,omitempty"`
	IsCrossRegionLogBackupEnabled  *bool  `json:"IsCrossRegionLogBackupEnabled,omitempty" xml:"IsCrossRegionLogBackupEnabled,omitempty"`
	// example:
	//
	// 1
	IsEnabled *int32 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	// example:
	//
	// 7
	LocalLogRetention       *int32 `json:"LocalLogRetention,omitempty" xml:"LocalLogRetention,omitempty"`
	LocalLogRetentionNumber *int32 `json:"LocalLogRetentionNumber,omitempty" xml:"LocalLogRetentionNumber,omitempty"`
	// example:
	//
	// 30
	LogLocalRetentionSpace *int32 `json:"LogLocalRetentionSpace,omitempty" xml:"LogLocalRetentionSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 7
	RemoveLogRetention *int32 `json:"RemoveLogRetention,omitempty" xml:"RemoveLogRetention,omitempty"`
}

func (s UpdateBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPolicyRequest) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *UpdateBackupPolicyRequest) GetBackupPlanBegin() *string {
	return s.BackupPlanBegin
}

func (s *UpdateBackupPolicyRequest) GetBackupSetRetention() *int32 {
	return s.BackupSetRetention
}

func (s *UpdateBackupPolicyRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *UpdateBackupPolicyRequest) GetBackupWay() *string {
	return s.BackupWay
}

func (s *UpdateBackupPolicyRequest) GetColdDataBackupInterval() *int32 {
	return s.ColdDataBackupInterval
}

func (s *UpdateBackupPolicyRequest) GetColdDataBackupRetention() *int32 {
	return s.ColdDataBackupRetention
}

func (s *UpdateBackupPolicyRequest) GetCrossRegionDataBackupRetention() *int32 {
	return s.CrossRegionDataBackupRetention
}

func (s *UpdateBackupPolicyRequest) GetCrossRegionLogBackupRetention() *int32 {
	return s.CrossRegionLogBackupRetention
}

func (s *UpdateBackupPolicyRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpdateBackupPolicyRequest) GetDestCrossRegion() *string {
	return s.DestCrossRegion
}

func (s *UpdateBackupPolicyRequest) GetForceCleanOnHighSpaceUsage() *int32 {
	return s.ForceCleanOnHighSpaceUsage
}

func (s *UpdateBackupPolicyRequest) GetIsCrossRegionDataBackupEnabled() *bool {
	return s.IsCrossRegionDataBackupEnabled
}

func (s *UpdateBackupPolicyRequest) GetIsCrossRegionLogBackupEnabled() *bool {
	return s.IsCrossRegionLogBackupEnabled
}

func (s *UpdateBackupPolicyRequest) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *UpdateBackupPolicyRequest) GetLocalLogRetention() *int32 {
	return s.LocalLogRetention
}

func (s *UpdateBackupPolicyRequest) GetLocalLogRetentionNumber() *int32 {
	return s.LocalLogRetentionNumber
}

func (s *UpdateBackupPolicyRequest) GetLogLocalRetentionSpace() *int32 {
	return s.LogLocalRetentionSpace
}

func (s *UpdateBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBackupPolicyRequest) GetRemoveLogRetention() *int32 {
	return s.RemoveLogRetention
}

func (s *UpdateBackupPolicyRequest) SetBackupPeriod(v string) *UpdateBackupPolicyRequest {
	s.BackupPeriod = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupPlanBegin(v string) *UpdateBackupPolicyRequest {
	s.BackupPlanBegin = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupSetRetention(v int32) *UpdateBackupPolicyRequest {
	s.BackupSetRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupType(v string) *UpdateBackupPolicyRequest {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetBackupWay(v string) *UpdateBackupPolicyRequest {
	s.BackupWay = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetColdDataBackupInterval(v int32) *UpdateBackupPolicyRequest {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetColdDataBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetCrossRegionDataBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetCrossRegionLogBackupRetention(v int32) *UpdateBackupPolicyRequest {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetDBInstanceName(v string) *UpdateBackupPolicyRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetDestCrossRegion(v string) *UpdateBackupPolicyRequest {
	s.DestCrossRegion = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetForceCleanOnHighSpaceUsage(v int32) *UpdateBackupPolicyRequest {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsCrossRegionDataBackupEnabled(v bool) *UpdateBackupPolicyRequest {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsCrossRegionLogBackupEnabled(v bool) *UpdateBackupPolicyRequest {
	s.IsCrossRegionLogBackupEnabled = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetIsEnabled(v int32) *UpdateBackupPolicyRequest {
	s.IsEnabled = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetLocalLogRetention(v int32) *UpdateBackupPolicyRequest {
	s.LocalLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetLocalLogRetentionNumber(v int32) *UpdateBackupPolicyRequest {
	s.LocalLogRetentionNumber = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetLogLocalRetentionSpace(v int32) *UpdateBackupPolicyRequest {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRegionId(v string) *UpdateBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupPolicyRequest) SetRemoveLogRetention(v int32) *UpdateBackupPolicyRequest {
	s.RemoveLogRetention = &v
	return s
}

func (s *UpdateBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
