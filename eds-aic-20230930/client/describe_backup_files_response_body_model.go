// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBackupFilesResponseBodyData) *DescribeBackupFilesResponseBody
	GetData() []*DescribeBackupFilesResponseBodyData
	SetMaxResults(v string) *DescribeBackupFilesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *DescribeBackupFilesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeBackupFilesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeBackupFilesResponseBody
	GetTotalCount() *int64
}

type DescribeBackupFilesResponseBody struct {
	// The list of returned result objects.
	Data []*DescribeBackupFilesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned in this call.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that indicates the position up to which data has been read in the current call. An empty value indicates that all data has been read.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID. If you encounter an issue, provide this request ID for troubleshooting.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 91
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBody) GetData() []*DescribeBackupFilesResponseBodyData {
	return s.Data
}

func (s *DescribeBackupFilesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeBackupFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupFilesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBackupFilesResponseBody) SetData(v []*DescribeBackupFilesResponseBodyData) *DescribeBackupFilesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetMaxResults(v string) *DescribeBackupFilesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetNextToken(v string) *DescribeBackupFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetRequestId(v string) *DescribeBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) SetTotalCount(v int64) *DescribeBackupFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupFilesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupFilesResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// Test cloud phone
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// Specifies whether to perform a full device backup.
	//
	// example:
	//
	// false
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The backup file ID.
	//
	// example:
	//
	// bf-b0qbg3pbpjkn7****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The backup file name.
	//
	// example:
	//
	// a-58ftsoo90p0qa****.ab
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The directory where the backup file is stored.
	//
	// example:
	//
	// oss://cloudphone-saved-bucket-cn-shanghai/backup/aic-58ftsoo90p0qa****.ab
	BackupFilePath *string `json:"BackupFilePath,omitempty" xml:"BackupFilePath,omitempty"`
	Channel        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// Backup file for /data/media directory.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The owner of the backup file.
	//
	// example:
	//
	// test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The total size of the source backup files.
	//
	// example:
	//
	// 10227168
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-05-15 17:33:59
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-05-15 17:33:59
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ag-58ftsoo90p0qi****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of backed-up application package names.
	SourceAppInfoList []*string `json:"SourceAppInfoList,omitempty" xml:"SourceAppInfoList,omitempty" type:"Repeated"`
	// The list of source backup file directories.
	SourceFilePathList []*string `json:"SourceFilePathList,omitempty" xml:"SourceFilePathList,omitempty" type:"Repeated"`
	// The backup file status.
	//
	// example:
	//
	// AVAILABLE
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemVersion *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The domain name of the upload endpoint.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	UploadEndpoint *string `json:"UploadEndpoint,omitempty" xml:"UploadEndpoint,omitempty"`
	// The backup type.
	//
	// example:
	//
	// OSS
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s DescribeBackupFilesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesResponseBodyData) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeBackupFilesResponseBodyData) GetAndroidInstanceName() *string {
	return s.AndroidInstanceName
}

func (s *DescribeBackupFilesResponseBodyData) GetBackupAll() *bool {
	return s.BackupAll
}

func (s *DescribeBackupFilesResponseBodyData) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *DescribeBackupFilesResponseBodyData) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *DescribeBackupFilesResponseBodyData) GetBackupFilePath() *string {
	return s.BackupFilePath
}

func (s *DescribeBackupFilesResponseBodyData) GetChannel() *string {
	return s.Channel
}

func (s *DescribeBackupFilesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackupFilesResponseBodyData) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeBackupFilesResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeBackupFilesResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeBackupFilesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeBackupFilesResponseBodyData) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DescribeBackupFilesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupFilesResponseBodyData) GetSourceAppInfoList() []*string {
	return s.SourceAppInfoList
}

func (s *DescribeBackupFilesResponseBodyData) GetSourceFilePathList() []*string {
	return s.SourceFilePathList
}

func (s *DescribeBackupFilesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupFilesResponseBodyData) GetSystemVersion() *string {
	return s.SystemVersion
}

func (s *DescribeBackupFilesResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeBackupFilesResponseBodyData) GetUploadEndpoint() *string {
	return s.UploadEndpoint
}

func (s *DescribeBackupFilesResponseBodyData) GetUploadType() *string {
	return s.UploadType
}

func (s *DescribeBackupFilesResponseBodyData) SetAndroidInstanceId(v string) *DescribeBackupFilesResponseBodyData {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetAndroidInstanceName(v string) *DescribeBackupFilesResponseBodyData {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupAll(v bool) *DescribeBackupFilesResponseBodyData {
	s.BackupAll = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFileId(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFileId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFileName(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFileName = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetBackupFilePath(v string) *DescribeBackupFilesResponseBodyData {
	s.BackupFilePath = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetChannel(v string) *DescribeBackupFilesResponseBodyData {
	s.Channel = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetDescription(v string) *DescribeBackupFilesResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetEndUserId(v string) *DescribeBackupFilesResponseBodyData {
	s.EndUserId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetFileSize(v int64) *DescribeBackupFilesResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetGmtCreated(v string) *DescribeBackupFilesResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetGmtModified(v string) *DescribeBackupFilesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetInstanceGroupId(v string) *DescribeBackupFilesResponseBodyData {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetRegionId(v string) *DescribeBackupFilesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetSourceAppInfoList(v []*string) *DescribeBackupFilesResponseBodyData {
	s.SourceAppInfoList = v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetSourceFilePathList(v []*string) *DescribeBackupFilesResponseBodyData {
	s.SourceFilePathList = v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetStatus(v string) *DescribeBackupFilesResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetSystemVersion(v string) *DescribeBackupFilesResponseBodyData {
	s.SystemVersion = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetTaskId(v string) *DescribeBackupFilesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetUploadEndpoint(v string) *DescribeBackupFilesResponseBodyData {
	s.UploadEndpoint = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) SetUploadType(v string) *DescribeBackupFilesResponseBodyData {
	s.UploadType = &v
	return s
}

func (s *DescribeBackupFilesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
