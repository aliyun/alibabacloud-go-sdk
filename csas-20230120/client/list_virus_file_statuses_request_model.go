// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusFileStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVirusFileStatusesRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListVirusFileStatusesRequest
	GetDepartment() *string
	SetDevTag(v string) *ListVirusFileStatusesRequest
	GetDevTag() *string
	SetDevType(v string) *ListVirusFileStatusesRequest
	GetDevType() *string
	SetEndTime(v int64) *ListVirusFileStatusesRequest
	GetEndTime() *int64
	SetFileMd5(v string) *ListVirusFileStatusesRequest
	GetFileMd5() *string
	SetFileProcessStatus(v string) *ListVirusFileStatusesRequest
	GetFileProcessStatus() *string
	SetHostname(v string) *ListVirusFileStatusesRequest
	GetHostname() *string
	SetOperations(v []*string) *ListVirusFileStatusesRequest
	GetOperations() []*string
	SetPageSize(v int64) *ListVirusFileStatusesRequest
	GetPageSize() *int64
	SetRiskLevels(v []*string) *ListVirusFileStatusesRequest
	GetRiskLevels() []*string
	SetSaseUserId(v string) *ListVirusFileStatusesRequest
	GetSaseUserId() *string
	SetScanTaskId(v string) *ListVirusFileStatusesRequest
	GetScanTaskId() *string
	SetStartTime(v int64) *ListVirusFileStatusesRequest
	GetStartTime() *int64
	SetUsername(v string) *ListVirusFileStatusesRequest
	GetUsername() *string
	SetVirusTypes(v []*string) *ListVirusFileStatusesRequest
	GetVirusTypes() []*string
}

type ListVirusFileStatusesRequest struct {
	// The page number of the current page in paging. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The department name. Matches any level of the organizational structure to which the user belongs. Specify the department name itself without the full path of the organizational structure. The value can contain Chinese characters, uppercase and lowercase letters, digits, spaces, periods (.), commas (,), forward slashes (/), at signs (@), hyphens (-), and underscores (_).
	//
	// example:
	//
	// R&D Department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The unique identifier of the user terminal device. Exact match. The value can be up to 64 characters in length. You can obtain the value from the following operation:
	//
	// - [ListUserDevices](~~ListUserDevices~~): Lists user terminal devices.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The operating system type of the user terminal device. Valid values:
	//
	// - **windows**: Windows.
	//
	// - **macOS**: macOS.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// The end time for filtering by virus file discovery time. The value is a UNIX timestamp in seconds. This parameter must be specified together with StartTime and must be later than StartTime.
	//
	// example:
	//
	// 1786377600
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The MD5 value of the virus file. Fuzzy match is supported. The value can be up to 64 characters in length.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// Filters by disposition status. If this parameter is not specified, no filtering by disposition status is applied. Valid values:
	//
	// - **Pending**: Pending disposition.
	//
	// - **Processed**: Disposed.
	//
	// example:
	//
	// Pending
	FileProcessStatus *string `json:"FileProcessStatus,omitempty" xml:"FileProcessStatus,omitempty"`
	// The hostname of the user terminal device. Fuzzy match is supported. The value can be up to 128 characters in length.
	//
	// example:
	//
	// DESKTOP-8A3F
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Filters by disposition action. Duplicate values are not allowed. If this parameter is not specified, no filtering by disposition action is applied.
	Operations []*string `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Repeated"`
	// The number of entries per page in paging. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by risk level. Duplicate values are not allowed. If this parameter is not specified, no filtering by risk level is applied.
	RiskLevels []*string `json:"RiskLevels,omitempty" xml:"RiskLevels,omitempty" type:"Repeated"`
	// The user ID. Exact match. The value can be up to 128 characters in length. You can obtain the value from the following operations:
	//
	// - [ListUserDevices](~~ListUserDevices~~): Lists user terminal devices.
	//
	// - [GetUserDevice](~~GetUserDevice~~): Queries user terminal device details.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The ID of the virus scan task that detected the virus file. This parameter is used to filter detection results of a specified task. You can obtain the value from the following operations:
	//
	// - [ListVirusScanTasks](~~ListVirusScanTasks~~): Lists virus scan tasks.
	//
	// - [CreateVirusScanTask](~~CreateVirusScanTask~~): Creates a virus scan task.
	//
	// example:
	//
	// v1:1024772
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The start time for filtering by virus file discovery time. The value is a UNIX timestamp in seconds. This parameter must be specified together with EndTime and must be earlier than EndTime.
	//
	// example:
	//
	// 1786291200
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username. Fuzzy match is supported. The value can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), asterisks (*), hyphens (-), at signs (@), spaces, middle dots (·), and parentheses.
	//
	// example:
	//
	// John Smith
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Filters by virus type. Duplicate values are not allowed. If this parameter is not specified, no filtering by virus type is applied.
	VirusTypes []*string `json:"VirusTypes,omitempty" xml:"VirusTypes,omitempty" type:"Repeated"`
}

func (s ListVirusFileStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusFileStatusesRequest) GoString() string {
	return s.String()
}

func (s *ListVirusFileStatusesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVirusFileStatusesRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListVirusFileStatusesRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *ListVirusFileStatusesRequest) GetDevType() *string {
	return s.DevType
}

func (s *ListVirusFileStatusesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVirusFileStatusesRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *ListVirusFileStatusesRequest) GetFileProcessStatus() *string {
	return s.FileProcessStatus
}

func (s *ListVirusFileStatusesRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListVirusFileStatusesRequest) GetOperations() []*string {
	return s.Operations
}

func (s *ListVirusFileStatusesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVirusFileStatusesRequest) GetRiskLevels() []*string {
	return s.RiskLevels
}

func (s *ListVirusFileStatusesRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListVirusFileStatusesRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListVirusFileStatusesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVirusFileStatusesRequest) GetUsername() *string {
	return s.Username
}

func (s *ListVirusFileStatusesRequest) GetVirusTypes() []*string {
	return s.VirusTypes
}

func (s *ListVirusFileStatusesRequest) SetCurrentPage(v int64) *ListVirusFileStatusesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetDepartment(v string) *ListVirusFileStatusesRequest {
	s.Department = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetDevTag(v string) *ListVirusFileStatusesRequest {
	s.DevTag = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetDevType(v string) *ListVirusFileStatusesRequest {
	s.DevType = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetEndTime(v int64) *ListVirusFileStatusesRequest {
	s.EndTime = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetFileMd5(v string) *ListVirusFileStatusesRequest {
	s.FileMd5 = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetFileProcessStatus(v string) *ListVirusFileStatusesRequest {
	s.FileProcessStatus = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetHostname(v string) *ListVirusFileStatusesRequest {
	s.Hostname = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetOperations(v []*string) *ListVirusFileStatusesRequest {
	s.Operations = v
	return s
}

func (s *ListVirusFileStatusesRequest) SetPageSize(v int64) *ListVirusFileStatusesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetRiskLevels(v []*string) *ListVirusFileStatusesRequest {
	s.RiskLevels = v
	return s
}

func (s *ListVirusFileStatusesRequest) SetSaseUserId(v string) *ListVirusFileStatusesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetScanTaskId(v string) *ListVirusFileStatusesRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetStartTime(v int64) *ListVirusFileStatusesRequest {
	s.StartTime = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetUsername(v string) *ListVirusFileStatusesRequest {
	s.Username = &v
	return s
}

func (s *ListVirusFileStatusesRequest) SetVirusTypes(v []*string) *ListVirusFileStatusesRequest {
	s.VirusTypes = v
	return s
}

func (s *ListVirusFileStatusesRequest) Validate() error {
	return dara.Validate(s)
}
