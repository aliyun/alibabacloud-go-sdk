// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceId(v string) *DescribeBackupFilesRequest
	GetAndroidInstanceId() *string
	SetAndroidInstanceName(v string) *DescribeBackupFilesRequest
	GetAndroidInstanceName() *string
	SetBackupAll(v bool) *DescribeBackupFilesRequest
	GetBackupAll() *bool
	SetBackupFileId(v string) *DescribeBackupFilesRequest
	GetBackupFileId() *string
	SetBackupFileName(v string) *DescribeBackupFilesRequest
	GetBackupFileName() *string
	SetDescription(v string) *DescribeBackupFilesRequest
	GetDescription() *string
	SetEndTime(v string) *DescribeBackupFilesRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeBackupFilesRequest
	GetEndUserId() *string
	SetInstanceGroupId(v string) *DescribeBackupFilesRequest
	GetInstanceGroupId() *string
	SetMaxResults(v int64) *DescribeBackupFilesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeBackupFilesRequest
	GetNextToken() *string
	SetSaleMode(v string) *DescribeBackupFilesRequest
	GetSaleMode() *string
	SetStartTime(v string) *DescribeBackupFilesRequest
	GetStartTime() *string
	SetStatusList(v []*string) *DescribeBackupFilesRequest
	GetStatusList() []*string
}

type DescribeBackupFilesRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The name of the instance. Fuzzy match is supported.
	//
	// example:
	//
	// acp-34pqe4r0kd9kn****
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	// Specifies whether the whole instance is backed up.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	BackupAll *bool `json:"BackupAll,omitempty" xml:"BackupAll,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// bf-dxrh5jrv0zpb8****
	BackupFileId *string `json:"BackupFileId,omitempty" xml:"BackupFileId,omitempty"`
	// The name of the backup file. Fuzzy match is supported.
	//
	// example:
	//
	// defaulBackupFile
	BackupFileName *string `json:"BackupFileName,omitempty" xml:"BackupFileName,omitempty"`
	// The description of the backup file. Fuzzy match is supported.
	//
	// example:
	//
	// default description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end of the period for querying generated backup files.
	//
	// example:
	//
	// 2024-05-20 10:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The owner of the backup file.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-fxdx91jsfyiy3****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l5V9uON****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SaleMode  *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The beginning of the period for querying generated backup files.
	//
	// example:
	//
	// 2024-05-23 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup files.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupFilesRequest) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeBackupFilesRequest) GetAndroidInstanceName() *string {
	return s.AndroidInstanceName
}

func (s *DescribeBackupFilesRequest) GetBackupAll() *bool {
	return s.BackupAll
}

func (s *DescribeBackupFilesRequest) GetBackupFileId() *string {
	return s.BackupFileId
}

func (s *DescribeBackupFilesRequest) GetBackupFileName() *string {
	return s.BackupFileName
}

func (s *DescribeBackupFilesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackupFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupFilesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeBackupFilesRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DescribeBackupFilesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeBackupFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeBackupFilesRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeBackupFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupFilesRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *DescribeBackupFilesRequest) SetAndroidInstanceId(v string) *DescribeBackupFilesRequest {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetAndroidInstanceName(v string) *DescribeBackupFilesRequest {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupAll(v bool) *DescribeBackupFilesRequest {
	s.BackupAll = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupFileId(v string) *DescribeBackupFilesRequest {
	s.BackupFileId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetBackupFileName(v string) *DescribeBackupFilesRequest {
	s.BackupFileName = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetDescription(v string) *DescribeBackupFilesRequest {
	s.Description = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetEndTime(v string) *DescribeBackupFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetEndUserId(v string) *DescribeBackupFilesRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetInstanceGroupId(v string) *DescribeBackupFilesRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetMaxResults(v int64) *DescribeBackupFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetNextToken(v string) *DescribeBackupFilesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetSaleMode(v string) *DescribeBackupFilesRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetStartTime(v string) *DescribeBackupFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupFilesRequest) SetStatusList(v []*string) *DescribeBackupFilesRequest {
	s.StatusList = v
	return s
}

func (s *DescribeBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
