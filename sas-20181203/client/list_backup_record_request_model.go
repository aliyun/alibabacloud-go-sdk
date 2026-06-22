// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupEndTime(v int64) *ListBackupRecordRequest
	GetBackupEndTime() *int64
	SetBackupStartTime(v int64) *ListBackupRecordRequest
	GetBackupStartTime() *int64
	SetCurrentPage(v int32) *ListBackupRecordRequest
	GetCurrentPage() *int32
	SetMachineRemark(v string) *ListBackupRecordRequest
	GetMachineRemark() *string
	SetPageSize(v int32) *ListBackupRecordRequest
	GetPageSize() *int32
	SetStatusList(v []*string) *ListBackupRecordRequest
	GetStatusList() []*string
}

type ListBackupRecordRequest struct {
	// The backup end time. The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1699600611000
	BackupEndTime *int64 `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The backup start time. The value is a timestamp in milliseconds.
	//
	// example:
	//
	// 1699514211000
	BackupStartTime *int64 `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The identification information of the server protected by the anti-ransomware policy that you want to query. You can enter the IP address or instance ID of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	MachineRemark *string `json:"MachineRemark,omitempty" xml:"MachineRemark,omitempty"`
	// The maximum number of entries per page in a paged query. Default value: 20. If you leave this parameter empty, 20 entries are returned.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of backup task statuses. Valid values:
	//
	// - **BACKUP_COMPLETE**: backup succeeded
	//
	// - **BACKUP_FAILED**: backup failed
	//
	// - **PARTIAL_COMPLETE**: partial backup succeeded.
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListBackupRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBackupRecordRequest) GoString() string {
	return s.String()
}

func (s *ListBackupRecordRequest) GetBackupEndTime() *int64 {
	return s.BackupEndTime
}

func (s *ListBackupRecordRequest) GetBackupStartTime() *int64 {
	return s.BackupStartTime
}

func (s *ListBackupRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBackupRecordRequest) GetMachineRemark() *string {
	return s.MachineRemark
}

func (s *ListBackupRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBackupRecordRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListBackupRecordRequest) SetBackupEndTime(v int64) *ListBackupRecordRequest {
	s.BackupEndTime = &v
	return s
}

func (s *ListBackupRecordRequest) SetBackupStartTime(v int64) *ListBackupRecordRequest {
	s.BackupStartTime = &v
	return s
}

func (s *ListBackupRecordRequest) SetCurrentPage(v int32) *ListBackupRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupRecordRequest) SetMachineRemark(v string) *ListBackupRecordRequest {
	s.MachineRemark = &v
	return s
}

func (s *ListBackupRecordRequest) SetPageSize(v int32) *ListBackupRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListBackupRecordRequest) SetStatusList(v []*string) *ListBackupRecordRequest {
	s.StatusList = v
	return s
}

func (s *ListBackupRecordRequest) Validate() error {
	return dara.Validate(s)
}
