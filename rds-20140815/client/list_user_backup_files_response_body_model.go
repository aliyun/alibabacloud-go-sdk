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
	Records   []*ListUserBackupFilesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BackupId         *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BinlogInfo       *string `json:"BinlogInfo,omitempty" xml:"BinlogInfo,omitempty"`
	Comment          *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	OssBucket        *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFileMetaData  *string `json:"OssFileMetaData,omitempty" xml:"OssFileMetaData,omitempty"`
	OssFileName      *string `json:"OssFileName,omitempty" xml:"OssFileName,omitempty"`
	OssFilePath      *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	OssFileSize      *int64  `json:"OssFileSize,omitempty" xml:"OssFileSize,omitempty"`
	OssUrl           *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	Reason           *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RestoreSize      *string `json:"RestoreSize,omitempty" xml:"RestoreSize,omitempty"`
	Retention        *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
