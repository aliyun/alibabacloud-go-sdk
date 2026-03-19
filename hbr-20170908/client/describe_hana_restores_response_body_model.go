// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaRestoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaRestoresResponseBody
	GetCode() *string
	SetHanaRestore(v *DescribeHanaRestoresResponseBodyHanaRestore) *DescribeHanaRestoresResponseBody
	GetHanaRestore() *DescribeHanaRestoresResponseBodyHanaRestore
	SetMessage(v string) *DescribeHanaRestoresResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHanaRestoresResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaRestoresResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHanaRestoresResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaRestoresResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeHanaRestoresResponseBody
	GetTotalCount() *int32
}

type DescribeHanaRestoresResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HanaRestore *DescribeHanaRestoresResponseBodyHanaRestore `json:"HanaRestore,omitempty" xml:"HanaRestore,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7DEFC897-8F05-5C05-912C-C9A9510FBFF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaRestoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRestoresResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaRestoresResponseBody) GetHanaRestore() *DescribeHanaRestoresResponseBodyHanaRestore {
	return s.HanaRestore
}

func (s *DescribeHanaRestoresResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaRestoresResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaRestoresResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaRestoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaRestoresResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaRestoresResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHanaRestoresResponseBody) SetCode(v string) *DescribeHanaRestoresResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetHanaRestore(v *DescribeHanaRestoresResponseBodyHanaRestore) *DescribeHanaRestoresResponseBody {
	s.HanaRestore = v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetMessage(v string) *DescribeHanaRestoresResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetPageNumber(v int32) *DescribeHanaRestoresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetPageSize(v int32) *DescribeHanaRestoresResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetRequestId(v string) *DescribeHanaRestoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetSuccess(v bool) *DescribeHanaRestoresResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) SetTotalCount(v int32) *DescribeHanaRestoresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHanaRestoresResponseBody) Validate() error {
	if s.HanaRestore != nil {
		if err := s.HanaRestore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHanaRestoresResponseBodyHanaRestore struct {
	HanaRestores []*DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores `json:"HanaRestores,omitempty" xml:"HanaRestores,omitempty" type:"Repeated"`
}

func (s DescribeHanaRestoresResponseBodyHanaRestore) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRestoresResponseBodyHanaRestore) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBodyHanaRestore) GetHanaRestores() []*DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	return s.HanaRestores
}

func (s *DescribeHanaRestoresResponseBodyHanaRestore) SetHanaRestores(v []*DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) *DescribeHanaRestoresResponseBodyHanaRestore {
	s.HanaRestores = v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestore) Validate() error {
	if s.HanaRestores != nil {
		for _, item := range s.HanaRestores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores struct {
	BackupID            *int64  `json:"BackupID,omitempty" xml:"BackupID,omitempty"`
	BackupPrefix        *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	CheckAccess         *bool   `json:"CheckAccess,omitempty" xml:"CheckAccess,omitempty"`
	ClearLog            *bool   `json:"ClearLog,omitempty" xml:"ClearLog,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CurrentPhase        *int32  `json:"CurrentPhase,omitempty" xml:"CurrentPhase,omitempty"`
	CurrentProgress     *int64  `json:"CurrentProgress,omitempty" xml:"CurrentProgress,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DatabaseRestoreId   *int64  `json:"DatabaseRestoreId,omitempty" xml:"DatabaseRestoreId,omitempty"`
	EndTime             *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogPosition         *int64  `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	MaxPhase            *int32  `json:"MaxPhase,omitempty" xml:"MaxPhase,omitempty"`
	MaxProgress         *int64  `json:"MaxProgress,omitempty" xml:"MaxProgress,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Mode                *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Phase               *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	ReachedTime         *int64  `json:"ReachedTime,omitempty" xml:"ReachedTime,omitempty"`
	RecoveryPointInTime *int64  `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	RestoreId           *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceClusterId     *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemCopy          *bool   `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	UseCatalog          *bool   `json:"UseCatalog,omitempty" xml:"UseCatalog,omitempty"`
	UseDelta            *bool   `json:"UseDelta,omitempty" xml:"UseDelta,omitempty"`
	VaultId             *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VolumeId            *int32  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GoString() string {
	return s.String()
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetBackupID() *int64 {
	return s.BackupID
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetCheckAccess() *bool {
	return s.CheckAccess
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetClearLog() *bool {
	return s.ClearLog
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetCurrentPhase() *int32 {
	return s.CurrentPhase
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetCurrentProgress() *int64 {
	return s.CurrentProgress
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetDatabaseRestoreId() *int64 {
	return s.DatabaseRestoreId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetLogPosition() *int64 {
	return s.LogPosition
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetMaxPhase() *int32 {
	return s.MaxPhase
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetMaxProgress() *int64 {
	return s.MaxProgress
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetMode() *string {
	return s.Mode
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetPhase() *string {
	return s.Phase
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetReachedTime() *int64 {
	return s.ReachedTime
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetRecoveryPointInTime() *int64 {
	return s.RecoveryPointInTime
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetRestoreId() *string {
	return s.RestoreId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetSource() *string {
	return s.Source
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetSourceClusterId() *string {
	return s.SourceClusterId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetState() *string {
	return s.State
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetStatus() *string {
	return s.Status
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetSystemCopy() *bool {
	return s.SystemCopy
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetUseCatalog() *bool {
	return s.UseCatalog
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetUseDelta() *bool {
	return s.UseDelta
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) GetVolumeId() *int32 {
	return s.VolumeId
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetBackupID(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.BackupID = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetBackupPrefix(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.BackupPrefix = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCheckAccess(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CheckAccess = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetClearLog(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ClearLog = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetClusterId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCurrentPhase(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CurrentPhase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetCurrentProgress(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.CurrentProgress = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetDatabaseName(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetDatabaseRestoreId(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.DatabaseRestoreId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetEndTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.EndTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetLogPosition(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.LogPosition = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMaxPhase(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.MaxPhase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMaxProgress(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.MaxProgress = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMessage(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Message = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetMode(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Mode = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetPhase(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Phase = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetReachedTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.ReachedTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetRecoveryPointInTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.RecoveryPointInTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetRestoreId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.RestoreId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSource(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Source = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSourceClusterId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.SourceClusterId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetStartTime(v int64) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.StartTime = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetState(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.State = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetStatus(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.Status = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetSystemCopy(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.SystemCopy = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetUseCatalog(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.UseCatalog = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetUseDelta(v bool) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.UseDelta = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetVaultId(v string) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) SetVolumeId(v int32) *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores {
	s.VolumeId = &v
	return s
}

func (s *DescribeHanaRestoresResponseBodyHanaRestoreHanaRestores) Validate() error {
	return dara.Validate(s)
}
