// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*ListUserBackupFilesResponseBodyRecords) *ListUserBackupFilesResponseBody
	GetRecords() []*ListUserBackupFilesResponseBodyRecords
	SetRequestId(v string) *ListUserBackupFilesResponseBody
	GetRequestId() *string
}

type ListUserBackupFilesResponseBody struct {
	// The list of user backup file details.
	Records []*ListUserBackupFilesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AD67C22F-64F3-4448-A9A8-D1606D242879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserBackupFilesResponseBody) GetRecords() []*ListUserBackupFilesResponseBodyRecords {
	return s.Records
}

func (s *ListUserBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserBackupFilesResponseBody) SetRecords(v []*ListUserBackupFilesResponseBodyRecords) *ListUserBackupFilesResponseBody {
	s.Records = v
	return s
}

func (s *ListUserBackupFilesResponseBody) SetRequestId(v string) *ListUserBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserBackupFilesResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserBackupFilesResponseBodyRecords struct {
	// The user backup ID.
	//
	// example:
	//
	// b-kwwvr7v8t7of****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The binary log file information in the backup file. This parameter is returned if incremental data exists during the backup process.
	//
	// example:
	//
	// {\\"binlogPosition\\":\\"154\\",\\"binlogFile\\":\\"0.000002\\"}
	BinlogInfo *string `json:"BinlogInfo,omitempty" xml:"BinlogInfo,omitempty"`
	// The comment of the user backup.
	//
	// example:
	//
	// BackupTest
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the user backup import started. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1623231084000
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The database engine.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the user backup was successfully imported. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1623231750000
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The time when the user backup import was completed. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1623231750000
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The name of the OSS bucket in which the user backup file is stored.
	//
	// example:
	//
	// BackupTest
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The metadata of the user backup file. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// {\\"Accept-Ranges\\":\\"bytes\\",\\"Connection\\":\\"keep-alive\\",\\"Content-Length\\":81014337,\\"Content-Type\\":\\"application/octet-stream\\",\\"Date\\":1623309548000,\\"ETag\\":\\"889FE9E5FCEBFE4781829488A352863B-1\\",\\"Last-Modified\\":1622186844000,\\"Server\\":\\"AliyunOSS\\",\\"x-oss-hash-crc64ecma\\":\\"5793608435727323129\\",\\"x-oss-object-type\\":\\"Multipart\\",\\"x-oss-request-id\\":\\"60C1BCEC92572F37318BD499\\",\\"x-oss-server-time\\":\\"166\\",\\"x-oss-storage-class\\":\\"Standard\\"}
	OssFileMetaData *string `json:"OssFileMetaData,omitempty" xml:"OssFileMetaData,omitempty"`
	// The name of the user backup file in OSS.
	//
	// example:
	//
	// backup_qp.xb
	OssFileName *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	// The path of the user backup file in OSS.
	//
	// example:
	//
	// test/backup_qp.xb
	OssFilePath *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	// The size of the user backup file in OSS. Unit: KB.
	//
	// example:
	//
	// 79115
	OssFileSize *int64 `json:"OssFileSize,omitempty" xml:"OssFileSize,omitempty"`
	// The OSS download URL of the user backup file.
	//
	// example:
	//
	// https://****.oss-ap-****.aliyuncs.com/backup_qp.xb
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// The reason why the user backup file failed to be imported.
	//
	// example:
	//
	// success
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The storage space required to restore the user backup. Unit: GB.
	//
	// example:
	//
	// 20
	RestoreSize *string `json:"RestoreSize,omitempty" xml:"RestoreSize,omitempty"`
	// The retention period of the user backup file. Unit: days.
	//
	// example:
	//
	// 3
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The status of the user backup file. Valid values:
	//
	// 	- **Importing**: The backup is being imported.
	//
	// 	- **Failed**: The import failed.
	//
	// 	- **CheckSuccess**: The verification passed.
	//
	// 	- **BackupSuccess**: The import succeeded.
	//
	// 	- **Deleted**: The backup is deleted.
	//
	// example:
	//
	// BackupSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The zone ID of the user backup.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListUserBackupFilesResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListUserBackupFilesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListUserBackupFilesResponseBodyRecords) GetBackupId() *string {
	return s.BackupId
}

func (s *ListUserBackupFilesResponseBodyRecords) GetBinlogInfo() *string {
	return s.BinlogInfo
}

func (s *ListUserBackupFilesResponseBodyRecords) GetComment() *string {
	return s.Comment
}

func (s *ListUserBackupFilesResponseBodyRecords) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListUserBackupFilesResponseBodyRecords) GetEngine() *string {
	return s.Engine
}

func (s *ListUserBackupFilesResponseBodyRecords) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ListUserBackupFilesResponseBodyRecords) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListUserBackupFilesResponseBodyRecords) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssFileMetaData() *string {
	return s.OssFileMetaData
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssFileName() *string {
	return s.OssFileName
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssFilePath() *string {
	return s.OssFilePath
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssFileSize() *int64 {
	return s.OssFileSize
}

func (s *ListUserBackupFilesResponseBodyRecords) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ListUserBackupFilesResponseBodyRecords) GetReason() *string {
	return s.Reason
}

func (s *ListUserBackupFilesResponseBodyRecords) GetRestoreSize() *string {
	return s.RestoreSize
}

func (s *ListUserBackupFilesResponseBodyRecords) GetRetention() *int32 {
	return s.Retention
}

func (s *ListUserBackupFilesResponseBodyRecords) GetStatus() *string {
	return s.Status
}

func (s *ListUserBackupFilesResponseBodyRecords) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListUserBackupFilesResponseBodyRecords) SetBackupId(v string) *ListUserBackupFilesResponseBodyRecords {
	s.BackupId = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetBinlogInfo(v string) *ListUserBackupFilesResponseBodyRecords {
	s.BinlogInfo = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetComment(v string) *ListUserBackupFilesResponseBodyRecords {
	s.Comment = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetCreationTime(v string) *ListUserBackupFilesResponseBodyRecords {
	s.CreationTime = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetEngine(v string) *ListUserBackupFilesResponseBodyRecords {
	s.Engine = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetEngineVersion(v string) *ListUserBackupFilesResponseBodyRecords {
	s.EngineVersion = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetFinishTime(v string) *ListUserBackupFilesResponseBodyRecords {
	s.FinishTime = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetModificationTime(v string) *ListUserBackupFilesResponseBodyRecords {
	s.ModificationTime = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssBucket(v string) *ListUserBackupFilesResponseBodyRecords {
	s.OssBucket = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssFileMetaData(v string) *ListUserBackupFilesResponseBodyRecords {
	s.OssFileMetaData = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssFileName(v string) *ListUserBackupFilesResponseBodyRecords {
	s.OssFileName = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssFilePath(v string) *ListUserBackupFilesResponseBodyRecords {
	s.OssFilePath = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssFileSize(v int64) *ListUserBackupFilesResponseBodyRecords {
	s.OssFileSize = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetOssUrl(v string) *ListUserBackupFilesResponseBodyRecords {
	s.OssUrl = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetReason(v string) *ListUserBackupFilesResponseBodyRecords {
	s.Reason = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetRestoreSize(v string) *ListUserBackupFilesResponseBodyRecords {
	s.RestoreSize = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetRetention(v int32) *ListUserBackupFilesResponseBodyRecords {
	s.Retention = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetStatus(v string) *ListUserBackupFilesResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) SetZoneId(v string) *ListUserBackupFilesResponseBodyRecords {
	s.ZoneId = &v
	return s
}

func (s *ListUserBackupFilesResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
