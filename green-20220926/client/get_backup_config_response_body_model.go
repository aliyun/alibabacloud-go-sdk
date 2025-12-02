// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v int32) *GetBackupConfigResponseBody
	GetBackupMode() *int32
	SetBucket(v string) *GetBackupConfigResponseBody
	GetBucket() *string
	SetEnable(v bool) *GetBackupConfigResponseBody
	GetEnable() *bool
	SetEnableBackup(v bool) *GetBackupConfigResponseBody
	GetEnableBackup() *bool
	SetEnableBackupVoice(v bool) *GetBackupConfigResponseBody
	GetEnableBackupVoice() *bool
	SetExpireSeconds(v int32) *GetBackupConfigResponseBody
	GetExpireSeconds() *int32
	SetGmtModified(v string) *GetBackupConfigResponseBody
	GetGmtModified() *string
	SetPath(v string) *GetBackupConfigResponseBody
	GetPath() *string
	SetPathVoice(v string) *GetBackupConfigResponseBody
	GetPathVoice() *string
	SetRegion(v string) *GetBackupConfigResponseBody
	GetRegion() *string
	SetRequestId(v string) *GetBackupConfigResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetBackupConfigResponseBody
	GetResourceType() *string
	SetServiceCode(v string) *GetBackupConfigResponseBody
	GetServiceCode() *string
	SetUid(v string) *GetBackupConfigResponseBody
	GetUid() *string
}

type GetBackupConfigResponseBody struct {
	// Backup scope.
	//
	// example:
	//
	// 0
	BackupMode *int32 `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// File server OSS Bucket.
	//
	// example:
	//
	// buckect_test
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Whether it is enabled. Values:
	//
	// - **true**: Enabled
	//
	// - **false**: Disabled
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Whether to enable backup.
	//
	// example:
	//
	// True
	EnableBackup *bool `json:"EnableBackup,omitempty" xml:"EnableBackup,omitempty"`
	// Whether to enable audio backup.
	//
	// example:
	//
	// True
	EnableBackupVoice *bool `json:"EnableBackupVoice,omitempty" xml:"EnableBackupVoice,omitempty"`
	// Expiration time in seconds.
	//
	// example:
	//
	// 300
	ExpireSeconds *int32 `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2023-01-17 12:29:56
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Path.
	//
	// example:
	//
	// aliyun/template/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Audio backup path.
	//
	// example:
	//
	// /back
	PathVoice *string `json:"PathVoice,omitempty" xml:"PathVoice,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// UID.
	//
	// example:
	//
	// 1772612608370735
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetBackupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupConfigResponseBody) GetBackupMode() *int32 {
	return s.BackupMode
}

func (s *GetBackupConfigResponseBody) GetBucket() *string {
	return s.Bucket
}

func (s *GetBackupConfigResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetBackupConfigResponseBody) GetEnableBackup() *bool {
	return s.EnableBackup
}

func (s *GetBackupConfigResponseBody) GetEnableBackupVoice() *bool {
	return s.EnableBackupVoice
}

func (s *GetBackupConfigResponseBody) GetExpireSeconds() *int32 {
	return s.ExpireSeconds
}

func (s *GetBackupConfigResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBackupConfigResponseBody) GetPath() *string {
	return s.Path
}

func (s *GetBackupConfigResponseBody) GetPathVoice() *string {
	return s.PathVoice
}

func (s *GetBackupConfigResponseBody) GetRegion() *string {
	return s.Region
}

func (s *GetBackupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupConfigResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetBackupConfigResponseBody) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetBackupConfigResponseBody) GetUid() *string {
	return s.Uid
}

func (s *GetBackupConfigResponseBody) SetBackupMode(v int32) *GetBackupConfigResponseBody {
	s.BackupMode = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetBucket(v string) *GetBackupConfigResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnable(v bool) *GetBackupConfigResponseBody {
	s.Enable = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnableBackup(v bool) *GetBackupConfigResponseBody {
	s.EnableBackup = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetEnableBackupVoice(v bool) *GetBackupConfigResponseBody {
	s.EnableBackupVoice = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetExpireSeconds(v int32) *GetBackupConfigResponseBody {
	s.ExpireSeconds = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetGmtModified(v string) *GetBackupConfigResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetPath(v string) *GetBackupConfigResponseBody {
	s.Path = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetPathVoice(v string) *GetBackupConfigResponseBody {
	s.PathVoice = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetRegion(v string) *GetBackupConfigResponseBody {
	s.Region = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetRequestId(v string) *GetBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetResourceType(v string) *GetBackupConfigResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetServiceCode(v string) *GetBackupConfigResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *GetBackupConfigResponseBody) SetUid(v string) *GetBackupConfigResponseBody {
	s.Uid = &v
	return s
}

func (s *GetBackupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
