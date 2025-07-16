// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody
	GetData() *DescribeBackupPolicyResponseBodyData
	SetMessage(v string) *DescribeBackupPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPolicyResponseBody
	GetSuccess() *bool
}

type DescribeBackupPolicyResponseBody struct {
	Data *DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetData() *DescribeBackupPolicyResponseBodyData {
	return s.Data
}

func (s *DescribeBackupPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPolicyResponseBody) SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
	s.Data = v
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

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyData struct {
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

func (s DescribeBackupPolicyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupPlanBegin() *string {
	return s.BackupPlanBegin
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupSetRetention() *int32 {
	return s.BackupSetRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupPolicyResponseBodyData) GetBackupWay() *string {
	return s.BackupWay
}

func (s *DescribeBackupPolicyResponseBodyData) GetColdDataBackupInterval() *int32 {
	return s.ColdDataBackupInterval
}

func (s *DescribeBackupPolicyResponseBodyData) GetColdDataBackupRetention() *int32 {
	return s.ColdDataBackupRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetCrossRegionDataBackupRetention() *int32 {
	return s.CrossRegionDataBackupRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetCrossRegionLogBackupRetention() *int32 {
	return s.CrossRegionLogBackupRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeBackupPolicyResponseBodyData) GetDestCrossRegion() *string {
	return s.DestCrossRegion
}

func (s *DescribeBackupPolicyResponseBodyData) GetForceCleanOnHighSpaceUsage() *int32 {
	return s.ForceCleanOnHighSpaceUsage
}

func (s *DescribeBackupPolicyResponseBodyData) GetIsCrossRegionDataBackupEnabled() *bool {
	return s.IsCrossRegionDataBackupEnabled
}

func (s *DescribeBackupPolicyResponseBodyData) GetIsCrossRegionLogBackupEnabled() *bool {
	return s.IsCrossRegionLogBackupEnabled
}

func (s *DescribeBackupPolicyResponseBodyData) GetIsEnabled() *int32 {
	return s.IsEnabled
}

func (s *DescribeBackupPolicyResponseBodyData) GetLocalLogRetention() *int32 {
	return s.LocalLogRetention
}

func (s *DescribeBackupPolicyResponseBodyData) GetLocalLogRetentionNumber() *int32 {
	return s.LocalLogRetentionNumber
}

func (s *DescribeBackupPolicyResponseBodyData) GetLogLocalRetentionSpace() *int32 {
	return s.LogLocalRetentionSpace
}

func (s *DescribeBackupPolicyResponseBodyData) GetRemoveLogRetention() *int32 {
	return s.RemoveLogRetention
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPeriod(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPlanBegin(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupPlanBegin = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupSetRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupSetRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupType(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupWay(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupWay = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetColdDataBackupInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ColdDataBackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetColdDataBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ColdDataBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCrossRegionDataBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.CrossRegionDataBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCrossRegionLogBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.CrossRegionLogBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDBInstanceName(v string) *DescribeBackupPolicyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetDestCrossRegion(v string) *DescribeBackupPolicyResponseBodyData {
	s.DestCrossRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetForceCleanOnHighSpaceUsage(v int32) *DescribeBackupPolicyResponseBodyData {
	s.ForceCleanOnHighSpaceUsage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsCrossRegionDataBackupEnabled(v bool) *DescribeBackupPolicyResponseBodyData {
	s.IsCrossRegionDataBackupEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsCrossRegionLogBackupEnabled(v bool) *DescribeBackupPolicyResponseBodyData {
	s.IsCrossRegionLogBackupEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIsEnabled(v int32) *DescribeBackupPolicyResponseBodyData {
	s.IsEnabled = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetentionNumber(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetentionNumber = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogLocalRetentionSpace(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LogLocalRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetRemoveLogRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.RemoveLogRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
