// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeBackupPolicyResponseBodyAccessDeniedDetail) *DescribeBackupPolicyResponseBody
	GetAccessDeniedDetail() *DescribeBackupPolicyResponseBodyAccessDeniedDetail
	SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody
	GetBackupRetentionPeriod() *string
	SetDbsInstance(v string) *DescribeBackupPolicyResponseBody
	GetDbsInstance() *string
	SetEnableBackupLog(v int32) *DescribeBackupPolicyResponseBody
	GetEnableBackupLog() *int32
	SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredBackupTime() *string
	SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody
	GetPreferredNextBackupTime() *string
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The following parameters are no longer used. Ignore the parameters.
	AccessDeniedDetail *DescribeBackupPolicyResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The retention period of the backup data. Unit: days.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *string `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Indicates whether the backup-as-a-service feature is enabled for the instance. Valid values:
	//
	// 	- **1**: The backup-as-a-service feature is enabled for the instance.
	//
	// 	- **0**: The backup-as-a-service feature is disabled for the instance.
	//
	// example:
	//
	// 0
	DbsInstance *string `json:"DbsInstance,omitempty" xml:"DbsInstance,omitempty"`
	// Indicates whether incremental data backup is enabled. Valid values:
	//
	// 	- **1**: Incremental data backup is enabled.
	//
	// 	- **0**: Incremental data backup is disabled.
	//
	// example:
	//
	// 1
	EnableBackupLog *int32 `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// example:
	//
	// Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time range during which the backup was created. The time follows the ISO 8601 standard in the *HH:mm*Z-*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 05:00Z-06:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The next backup time. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-03-14T05:28Z
	PreferredNextBackupTime *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90B82DB7-FB28-4CC2-ADBF-1F8659F3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetAccessDeniedDetail() *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeBackupPolicyResponseBody) GetBackupRetentionPeriod() *string {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetDbsInstance() *string {
	return s.DbsInstance
}

func (s *DescribeBackupPolicyResponseBody) GetEnableBackupLog() *int32 {
	return s.EnableBackupLog
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetPreferredNextBackupTime() *string {
	return s.PreferredNextBackupTime
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) SetAccessDeniedDetail(v *DescribeBackupPolicyResponseBodyAccessDeniedDetail) *DescribeBackupPolicyResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDbsInstance(v string) *DescribeBackupPolicyResponseBody {
	s.DbsInstance = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyAccessDeniedDetail struct {
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeBackupPolicyResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
