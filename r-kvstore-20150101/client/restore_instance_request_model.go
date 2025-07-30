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
	// The ID of the backup file. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/473823.html)operation to query the IDs of backup files.
	//
	// example:
	//
	// 78241****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The key that you want to restore. You can specify multiple keys. Separate multiple keys with commas (,). Regular expressions are supported.
	//
	// 	- If you do not specify this parameter, the entire instance is restored.
	//
	// 	- If you specify this parameter, only the involved keys are restored. Only classic instances support this feature.
	//
	// >  In a regular expression, an asterisk (`*`) matches zero or more occurrences of a subexpression that occurs before. For example, if you set this parameter to `h.*llo`, strings such as `hllo` and `heeeello` are matched.
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
	// The point in time to which you want to restore data. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// >  The point in time cannot be earlier than the point in time when the data flashback feature is enabled.
	//
	// example:
	//
	// 2021-07-06T07:25:57Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration mode. Valid values:
	//
	// 	- **0*	- (default): The parameter is invalid.
	//
	// 	- **1**: restores data to a specified point in time. You can specify this value only if the [data flashback](https://help.aliyun.com/document_detail/148479.html) feature is enabled for the instance. If you specify this value, you also need to set the **RestoreTime*	- parameter.
	//
	// example:
	//
	// 1
	RestoreType   *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// When you restore a classic instance, regardless of whether you choose to restore all data or specific keys, you can apply an offset to the expiration time of the keys. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. A key expires after the remaining validity period of the key elapses based on the expiration offset time point.
	//
	// >
	//
	// 	- This feature applies only to keys and does not work on elements in the self-developed data structures of Tair, such as fields in exHash and skeys in TairTS.
	//
	// 	- This time point must be between the specified flashback time point and the submission time of the data restoration task.
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
