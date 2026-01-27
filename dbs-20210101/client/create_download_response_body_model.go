// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDownloadResponseBody
	GetCode() *string
	SetData(v *CreateDownloadResponseBodyData) *CreateDownloadResponseBody
	GetData() *CreateDownloadResponseBodyData
	SetErrCode(v string) *CreateDownloadResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateDownloadResponseBody
	GetErrMessage() *string
	SetMessage(v string) *CreateDownloadResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDownloadResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDownloadResponseBody
	GetSuccess() *string
}

type CreateDownloadResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the download task.
	Data *CreateDownloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// formatType can not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// formatType can not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A08F908D-2C35-583F-93C1-ED80753F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDownloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDownloadResponseBody) GetData() *CreateDownloadResponseBodyData {
	return s.Data
}

func (s *CreateDownloadResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDownloadResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateDownloadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDownloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDownloadResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDownloadResponseBody) SetCode(v string) *CreateDownloadResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDownloadResponseBody) SetData(v *CreateDownloadResponseBodyData) *CreateDownloadResponseBody {
	s.Data = v
	return s
}

func (s *CreateDownloadResponseBody) SetErrCode(v string) *CreateDownloadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDownloadResponseBody) SetErrMessage(v string) *CreateDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDownloadResponseBody) SetMessage(v string) *CreateDownloadResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDownloadResponseBody) SetRequestId(v string) *CreateDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadResponseBody) SetSuccess(v string) *CreateDownloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDownloadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDownloadResponseBodyData struct {
	// The point in time of the backup set if the task is used to download a backup set at a specific point in time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661373070000
	BackupSetTime *int64 `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	// The ID of the full backup set.
	//
	// example:
	//
	// 146005****
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The database and table information that is returned if databases and tables are filtered by the download task.
	//
	// example:
	//
	// testdb
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The state of the download task. Valid values:
	//
	// 	- initializing: The download task was being initialized.
	//
	// 	- queuing: The download task was queuing.
	//
	// 	- running: The download task was running.
	//
	// 	- failed: The download task failed.
	//
	// 	- finished: The download task was complete.
	//
	// 	- expired: The download task expired.
	//
	// > If the TargetType parameter is set to URL, the download task expires in three days after the task is complete.
	//
	// example:
	//
	// initializing
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The size of the downloaded data. Unit: bytes.
	//
	// example:
	//
	// 0
	ExportDataSize *int64 `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	// The format to which the downloaded data is converted.
	//
	// example:
	//
	// CSV
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the download task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661940917570
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The size of the processed data. Unit: bytes.
	//
	// example:
	//
	// 0
	ImportDataSize *int64 `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	// The number of tables that have been downloaded and the total number of tables to be downloaded.
	//
	// > If the task is in the preparation stage, 0/0 is returned.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// >  This parameter is returned if the value of **TargetType is OSS**.
	//
	// example:
	//
	// test_db/path
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded.
	//
	// example:
	//
	// URL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the download task.
	//
	// example:
	//
	// dt-qxnsfq5s****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDownloadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponseBodyData) GetBackupSetTime() *int64 {
	return s.BackupSetTime
}

func (s *CreateDownloadResponseBodyData) GetBakSetId() *string {
	return s.BakSetId
}

func (s *CreateDownloadResponseBodyData) GetDbList() *string {
	return s.DbList
}

func (s *CreateDownloadResponseBodyData) GetDownloadStatus() *string {
	return s.DownloadStatus
}

func (s *CreateDownloadResponseBodyData) GetExportDataSize() *int64 {
	return s.ExportDataSize
}

func (s *CreateDownloadResponseBodyData) GetFormat() *string {
	return s.Format
}

func (s *CreateDownloadResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateDownloadResponseBodyData) GetImportDataSize() *int64 {
	return s.ImportDataSize
}

func (s *CreateDownloadResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *CreateDownloadResponseBodyData) GetRegionCode() *string {
	return s.RegionCode
}

func (s *CreateDownloadResponseBodyData) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateDownloadResponseBodyData) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDownloadResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDownloadResponseBodyData) SetBackupSetTime(v int64) *CreateDownloadResponseBodyData {
	s.BackupSetTime = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetBakSetId(v string) *CreateDownloadResponseBodyData {
	s.BakSetId = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetDbList(v string) *CreateDownloadResponseBodyData {
	s.DbList = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetDownloadStatus(v string) *CreateDownloadResponseBodyData {
	s.DownloadStatus = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetExportDataSize(v int64) *CreateDownloadResponseBodyData {
	s.ExportDataSize = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetFormat(v string) *CreateDownloadResponseBodyData {
	s.Format = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetGmtCreate(v int64) *CreateDownloadResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetImportDataSize(v int64) *CreateDownloadResponseBodyData {
	s.ImportDataSize = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetProgress(v string) *CreateDownloadResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetRegionCode(v string) *CreateDownloadResponseBodyData {
	s.RegionCode = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTargetPath(v string) *CreateDownloadResponseBodyData {
	s.TargetPath = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTargetType(v string) *CreateDownloadResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTaskId(v string) *CreateDownloadResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateDownloadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
