// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *RestoreInstanceRequest
	GetBackupId() *string
	SetFilterKey(v string) *RestoreInstanceRequest
	GetFilterKey() *string
	SetInstanceId(v string) *RestoreInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RestoreInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RestoreInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RestoreInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RestoreInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *RestoreInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *RestoreInstanceRequest
	GetRestoreType() *string
	SetSecurityToken(v string) *RestoreInstanceRequest
	GetSecurityToken() *string
	SetTimeShift(v string) *RestoreInstanceRequest
	GetTimeShift() *string
}

type RestoreInstanceRequest struct {
	// The ID of the backup file. You can find backup file IDs by calling the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation.
	//
	// example:
	//
	// 78241****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The keys to restore, which can be specified as a regular expression. To specify multiple keys, separate them with commas (,).
	//
	// - If you do not specify this parameter, the entire instance is restored.
	//
	// - If you specify this parameter, only the specified keys are restored. This feature is available only for instances that use the classic architecture.
	//
	// > In a regular expression, the asterisk (`*`) matches the preceding element zero or more times. For example, if you set this parameter to `h.*llo`, strings such as `hllo` and `heeeello` are matched.
	//
	// example:
	//
	// key:00000007198*
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The restore point in time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format (UTC).
	//
	// > This point in time cannot be earlier than the time when the Data Flashback feature was enabled.
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restore method. Valid values:
	//
	// - **0*	- (default): This value is deprecated.
	//
	// - **1**: Restores data to a specific point in time. You can set this parameter to 1 only if the [Data Flashback](https://help.aliyun.com/document_detail/148479.html) feature is enabled for the instance. If you set this parameter to 1, you must also specify the **RestoreTime*	- parameter.
	//
	// example:
	//
	// 1
	RestoreType   *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// For instances that use the classic architecture, you can apply an offset to the expiration time of restored keys. This applies whether you restore the entire instance or only specific keys. The system calculates a key\\"s remaining time-to-live (TTL) at the specified flashback point in time and then adds this TTL to the `TimeShift` value to set the key\\"s new expiration time. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format (UTC).
	//
	// > - This feature adjusts the expiration time for top-level keys only. It does not apply to the expiration times of elements within Tair-specific data structures, such as fields in an exHash or secondary keys (Skeys) in a Time Series (TS) data structure.
	//
	// >
	//
	// > - The specified time must be later than `RestoreTime` and earlier than the task submission time.
	//
	// example:
	//
	// 2021-07-06T08:25:57Z
	TimeShift *string `json:"TimeShift,omitempty" xml:"TimeShift,omitempty"`
}

func (s RestoreInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *RestoreInstanceRequest) GetFilterKey() *string {
	return s.FilterKey
}

func (s *RestoreInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestoreInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RestoreInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RestoreInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RestoreInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RestoreInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *RestoreInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *RestoreInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestoreInstanceRequest) GetTimeShift() *string {
	return s.TimeShift
}

func (s *RestoreInstanceRequest) SetBackupId(v string) *RestoreInstanceRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreInstanceRequest) SetFilterKey(v string) *RestoreInstanceRequest {
	s.FilterKey = &v
	return s
}

func (s *RestoreInstanceRequest) SetInstanceId(v string) *RestoreInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestoreInstanceRequest) SetOwnerAccount(v string) *RestoreInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreInstanceRequest) SetOwnerId(v int64) *RestoreInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreInstanceRequest) SetResourceOwnerAccount(v string) *RestoreInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreInstanceRequest) SetResourceOwnerId(v int64) *RestoreInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreInstanceRequest) SetRestoreTime(v string) *RestoreInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *RestoreInstanceRequest) SetRestoreType(v string) *RestoreInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *RestoreInstanceRequest) SetSecurityToken(v string) *RestoreInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestoreInstanceRequest) SetTimeShift(v string) *RestoreInstanceRequest {
	s.TimeShift = &v
	return s
}

func (s *RestoreInstanceRequest) Validate() error {
	return dara.Validate(s)
}
