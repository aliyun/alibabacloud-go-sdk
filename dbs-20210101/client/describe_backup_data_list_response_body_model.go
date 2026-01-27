// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDataListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeBackupDataListResponseBody
	GetCode() *string
	SetData(v *DescribeBackupDataListResponseBodyData) *DescribeBackupDataListResponseBody
	GetData() *DescribeBackupDataListResponseBodyData
	SetErrCode(v string) *DescribeBackupDataListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupDataListResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeBackupDataListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBackupDataListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeBackupDataListResponseBody
	GetSuccess() *string
}

type DescribeBackupDataListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeBackupDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 95A5FFD0-7F46-5A7D-9DFE-6A376B4E2A28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupDataListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBackupDataListResponseBody) GetData() *DescribeBackupDataListResponseBodyData {
	return s.Data
}

func (s *DescribeBackupDataListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupDataListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupDataListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupDataListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupDataListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeBackupDataListResponseBody) SetCode(v string) *DescribeBackupDataListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetData(v *DescribeBackupDataListResponseBodyData) *DescribeBackupDataListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetErrCode(v string) *DescribeBackupDataListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetErrMessage(v string) *DescribeBackupDataListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetMessage(v string) *DescribeBackupDataListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetRequestId(v string) *DescribeBackupDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetSuccess(v string) *DescribeBackupDataListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupDataListResponseBodyData struct {
	// The information about the task.
	Content []*DescribeBackupDataListResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The additional information.
	//
	// example:
	//
	// dbtest
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of backup sets.
	//
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeBackupDataListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyData) GetContent() []*DescribeBackupDataListResponseBodyDataContent {
	return s.Content
}

func (s *DescribeBackupDataListResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeBackupDataListResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeBackupDataListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeBackupDataListResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *DescribeBackupDataListResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *DescribeBackupDataListResponseBodyData) SetContent(v []*DescribeBackupDataListResponseBodyDataContent) *DescribeBackupDataListResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetExtra(v string) *DescribeBackupDataListResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetPageNumber(v int64) *DescribeBackupDataListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetPageSize(v int64) *DescribeBackupDataListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetTotalElements(v int64) *DescribeBackupDataListResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetTotalPages(v int64) *DescribeBackupDataListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupDataListResponseBodyDataContent struct {
	// The URL that is used to download the backup set over the Internet.
	//
	// >  This parameter is returned only when the value of BackupMethod is **Physical*	- or **Logical**.
	//
	// example:
	//
	// http://oss.com/****
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2024-04-17T17:00:32Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 213088****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The URL that is used to download the backup set over the internal network.
	//
	// >  This parameter is returned only when the value of BackupMethod is **Physical*	- or **Logical**.
	//
	// example:
	//
	// http://oss.com/****
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The storage path of backup files.
	//
	// example:
	//
	// logic
	BackupLocation *string `json:"BackupLocation,omitempty" xml:"BackupLocation,omitempty"`
	// The backup method. Valid values:
	//
	// 	- **Physical**
	//
	// 	- **Logical**
	//
	// 	- **Snapshot**
	//
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The name of the backup set.
	//
	// example:
	//
	// logic_backup
	BackupName *string `json:"BackupName,omitempty" xml:"BackupName,omitempty"`
	// The backup scale. Valid values:
	//
	// 	- **DBInstance**
	//
	// 	- **DBTable**
	//
	// example:
	//
	// DBInstance
	BackupScale *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	// The size of the backup set. Unit: byte.
	//
	// example:
	//
	// 25669140480
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2024-04-17T17:00:16Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The status of the backup set. Valid values:
	//
	// 	- **OK**: The backup succeeded.
	//
	// 	- **ERROR**: The backup failed.
	//
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The checksum value.
	//
	// example:
	//
	// 84a4c16431f969712e6895992719****
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The snapshot checkpoint time. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1713373221
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The information about the encryption.
	//
	// example:
	//
	// psk2
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// polardb_mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the backup set expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-19T05:00:49Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// Indicates whether the backup set expires. Valid values:
	//
	// 	- NEVER: The backup set does not expire.
	//
	// 	- EXPIRED: The backup set expired.
	//
	// 	- DELAY: The backup set expires after a specific period of time.
	//
	// example:
	//
	// DELAY
	ExpectExpireType *string `json:"ExpectExpireType,omitempty" xml:"ExpectExpireType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **1**: The backup set is available.
	//
	// 	- **0**: The backup set is unavailable.
	//
	// example:
	//
	// 1
	IsAvail *int32 `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// The information about the PolarDB level-2 dump.
	//
	// >  This parameter is returned only if the level-2 dump feature is enabled for the PolarDB for MySQL cluster and the level-1 backup is dumped.
	PolarSnapshot *DescribeBackupDataListResponseBodyDataContentPolarSnapshot `json:"PolarSnapshot,omitempty" xml:"PolarSnapshot,omitempty" type:"Struct"`
	// Indicates whether the backup set can be deleted. Valid values:
	//
	// 	- **0**: The backup set can be deleted.
	//
	// 	- **1**: The backup set cannot be deleted.
	//
	// example:
	//
	// 0
	SupportDeletion *int32 `json:"SupportDeletion,omitempty" xml:"SupportDeletion,omitempty"`
}

func (s DescribeBackupDataListResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupLocation() *string {
	return s.BackupLocation
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupName() *string {
	return s.BackupName
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupScale() *string {
	return s.BackupScale
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetChecksum() *string {
	return s.Checksum
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetConsistentTime() *int64 {
	return s.ConsistentTime
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetEncryption() *string {
	return s.Encryption
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetEngine() *string {
	return s.Engine
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetExpectExpireType() *string {
	return s.ExpectExpireType
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetIsAvail() *int32 {
	return s.IsAvail
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetPolarSnapshot() *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	return s.PolarSnapshot
}

func (s *DescribeBackupDataListResponseBodyDataContent) GetSupportDeletion() *int32 {
	return s.SupportDeletion
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupDownloadURL(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupEndTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupId(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupIntranetDownloadURL(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupLocation(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupLocation = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupMethod(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupMode(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupName(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupName = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupScale(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupScale = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupSize(v int64) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupStartTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupStatus(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupType(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetChecksum(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Checksum = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetConsistentTime(v int64) *DescribeBackupDataListResponseBodyDataContent {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEncryption(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Encryption = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEngine(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Engine = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEngineVersion(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.EngineVersion = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetExpectExpireTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetExpectExpireType(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.ExpectExpireType = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetInstanceName(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetIsAvail(v int32) *DescribeBackupDataListResponseBodyDataContent {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetPolarSnapshot(v *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) *DescribeBackupDataListResponseBodyDataContent {
	s.PolarSnapshot = v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetSupportDeletion(v int32) *DescribeBackupDataListResponseBodyDataContent {
	s.SupportDeletion = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) Validate() error {
	if s.PolarSnapshot != nil {
		if err := s.PolarSnapshot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupDataListResponseBodyDataContentPolarSnapshot struct {
	// The dump backup ID.
	//
	// example:
	//
	// abc****
	DumpId *int64 `json:"DumpId,omitempty" xml:"DumpId,omitempty"`
	// The size of the dump backup. Unit: byte.
	//
	// example:
	//
	// 25669140589
	DumpSize *int64 `json:"DumpSize,omitempty" xml:"DumpSize,omitempty"`
	// The time when the backup set expires. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-19T05:00:49Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// Indicates whether the backup set expires. Valid values:
	//
	// 	- NEVER: The backup set does not expire.
	//
	// 	- EXPIRED: The backup set expired.
	//
	// 	- DELAY: The backup set expires after a specific period of time.
	//
	// example:
	//
	// DELAY
	ExpectExpireType *string `json:"expectExpireType,omitempty" xml:"expectExpireType,omitempty"`
}

func (s DescribeBackupDataListResponseBodyDataContentPolarSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GetDumpId() *int64 {
	return s.DumpId
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GetDumpSize() *int64 {
	return s.DumpSize
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GetExpectExpireType() *string {
	return s.ExpectExpireType
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetDumpId(v int64) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.DumpId = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetDumpSize(v int64) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.DumpSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetExpectExpireTime(v string) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetExpectExpireType(v string) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.ExpectExpireType = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) Validate() error {
	return dara.Validate(s)
}
