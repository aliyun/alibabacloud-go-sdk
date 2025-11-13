// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUniBackupRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRegionId(v string) *ListUniBackupRecordRequest
	GetBackupRegionId() *string
	SetCurrentPage(v int32) *ListUniBackupRecordRequest
	GetCurrentPage() *int32
	SetMachineRemark(v string) *ListUniBackupRecordRequest
	GetMachineRemark() *string
	SetPageSize(v int32) *ListUniBackupRecordRequest
	GetPageSize() *int32
	SetState(v string) *ListUniBackupRecordRequest
	GetState() *string
}

type ListUniBackupRecordRequest struct {
	// The region where the anti-ransomware backup service is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	BackupRegionId *string `json:"BackupRegionId,omitempty" xml:"BackupRegionId,omitempty"`
	// When performing a paginated query, set the page number for the current page. The default value is **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The identification information of the server protected by the anti-ransomware policy. You can enter the IP address or instance ID of the server.
	//
	// example:
	//
	// 1.1.XX.XX
	MachineRemark *string `json:"MachineRemark,omitempty" xml:"MachineRemark,omitempty"`
	// The maximum number of data entries to display per page in a paginated query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Backup status. Values:
	//
	// - **completed**: Success
	//
	// - **error**: Failure
	//
	// - **canceled**: Closed
	//
	// example:
	//
	// completed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListUniBackupRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUniBackupRecordRequest) GoString() string {
	return s.String()
}

func (s *ListUniBackupRecordRequest) GetBackupRegionId() *string {
	return s.BackupRegionId
}

func (s *ListUniBackupRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUniBackupRecordRequest) GetMachineRemark() *string {
	return s.MachineRemark
}

func (s *ListUniBackupRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUniBackupRecordRequest) GetState() *string {
	return s.State
}

func (s *ListUniBackupRecordRequest) SetBackupRegionId(v string) *ListUniBackupRecordRequest {
	s.BackupRegionId = &v
	return s
}

func (s *ListUniBackupRecordRequest) SetCurrentPage(v int32) *ListUniBackupRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUniBackupRecordRequest) SetMachineRemark(v string) *ListUniBackupRecordRequest {
	s.MachineRemark = &v
	return s
}

func (s *ListUniBackupRecordRequest) SetPageSize(v int32) *ListUniBackupRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListUniBackupRecordRequest) SetState(v string) *ListUniBackupRecordRequest {
	s.State = &v
	return s
}

func (s *ListUniBackupRecordRequest) Validate() error {
	return dara.Validate(s)
}
