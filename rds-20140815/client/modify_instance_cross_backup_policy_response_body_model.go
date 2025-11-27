// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceCrossBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetBackupEnabled() *string
	SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetCrossBackupRegion() *string
	SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetCrossBackupType() *string
	SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetDBInstanceId() *string
	SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetLogBackupEnabled() *string
	SetRegionId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRequestId() *string
	SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRetentType() *int32
	SetRetention(v int32) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRetention() *int32
}

type ModifyInstanceCrossBackupPolicyResponseBody struct {
	// The status of the cross-region backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	BackupEnabled *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	// The ID of the region in which the cross-region backup files of the instance are stored.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The policy that is used to save the cross-region backup files of the instance. Default value: **1**. The value 1 indicates that all cross-region backup files are saved.
	//
	// example:
	//
	// 1
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	// The region ID of the source instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50A6059D-6DBB-46C6-A851-1EE93C9013CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The policy that is used to retain the cross-region backup files of the instance. Default value: **1**. The value 1 indicates that the cross-region backup files of the instance are retained based on the specified retention period.
	//
	// example:
	//
	// 1
	RetentType *int32 `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: **7 to 1825**.
	//
	// example:
	//
	// 15
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifyInstanceCrossBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceCrossBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetBackupEnabled() *string {
	return s.BackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetLogBackupEnabled() *string {
	return s.LogBackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRetentType() *int32 {
	return s.RetentType
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.BackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupRegion = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.LogBackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRegionId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRequestId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RetentType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRetention(v int32) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.Retention = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
