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
	// The ID of the backup file for the instance. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html) operation to query the ID.
	//
	// example:
	//
	// 78241****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The keys to restore. Regular expressions are supported. Separate multiple values with commas (,).
	//
	// - If you do not specify this parameter, the entire instance is restored.
	//
	// - If you specify this parameter, only the specified keys are restored. This feature is supported only by classic architecture instances.
	//
	// > For example, in a regular expression, an asterisk (`*`) matches the preceding subexpression zero or more times. If you set this parameter to `h.*llo`, the system matches strings such as `hllo` and `heeeello`.
	//
	// example:
	//
	// key:00000007198*
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The instance ID.
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
	// The point in time to which you want to restore data. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format. The time must be in UTC.
	//
	// > The point in time cannot be earlier than the time when the data flashback feature is enabled.
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration type. Valid values:
	//
	// 	- **0*	- (default): This value is deprecated.
	//
	// 	- **1**: Restores data to a specified point in time. You can specify this value only when the [data flashback](https://help.aliyun.com/document_detail/148479.html) feature is enabled for the instance. If you specify this value, you must also specify the **RestoreTime*	- parameter.
	//
	// example:
	//
	// 1
	RestoreType   *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// When you restore a classic architecture instance, you can apply an offset to the expiration time of keys regardless of whether you choose to restore full data or specified keys. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The instance calculates the remaining expiration duration of the keys at the specified flashback point in time, starting from the specified expiration offset point in time.
	//
	// > 	- This feature supports offsetting the expiration time only at the key level. It does not support offsetting the internal expiration time of keys in Tair-developed data structures, such as the fields in exHash and the skeys in TS.
	//
	// > 	- The point in time cannot be earlier than the specified flashback point in time or later than the time when the restoration task is submitted.
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
