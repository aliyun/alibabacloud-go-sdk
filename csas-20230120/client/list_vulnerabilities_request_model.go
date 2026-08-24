// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulnerabilitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVulnerabilitiesRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListVulnerabilitiesRequest
	GetDepartment() *string
	SetDevTag(v string) *ListVulnerabilitiesRequest
	GetDevTag() *string
	SetDevType(v string) *ListVulnerabilitiesRequest
	GetDevType() *string
	SetHostname(v string) *ListVulnerabilitiesRequest
	GetHostname() *string
	SetPageSize(v int64) *ListVulnerabilitiesRequest
	GetPageSize() *int64
	SetSaseUserId(v string) *ListVulnerabilitiesRequest
	GetSaseUserId() *string
	SetScanTaskId(v string) *ListVulnerabilitiesRequest
	GetScanTaskId() *string
	SetTitle(v string) *ListVulnerabilitiesRequest
	GetTitle() *string
	SetUpdateIds(v []*string) *ListVulnerabilitiesRequest
	GetUpdateIds() []*string
	SetUsername(v string) *ListVulnerabilitiesRequest
	GetUsername() *string
	SetVulLevel(v string) *ListVulnerabilitiesRequest
	GetVulLevel() *string
	SetVulType(v string) *ListVulnerabilitiesRequest
	GetVulType() *string
}

type ListVulnerabilitiesRequest struct {
	// The page number of the current page in a paged query with paging. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The department name. Matches any level of department in the organizational structure to which the user belongs. Specify the department name itself without the full path of the organizational structure.
	//
	// example:
	//
	// R&D Department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The unique identifier of the user endpoint device. Exact match. The value can be up to 64 characters in length. Valid values are obtained from:
	//
	// - [ListUserDevices](~~ListUserDevices~~): lists user endpoint devices.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The operating system type of the user endpoint device. Valid values:
	//
	// - **windows**: Windows. Currently, vulnerability scanning supports only Windows.
	//
	// example:
	//
	// windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// The hostname of the user endpoint device. Fuzzy match is supported. The value can be up to 64 characters in length.
	//
	// example:
	//
	// DESKTOP-8A3F
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The number of entries per page. Settings for paged query with paging. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user ID. Exact match. Valid values are obtained from:
	//
	// - [ListUserDevices](~~ListUserDevices~~): lists user endpoint devices.
	//
	// - [GetUserDevice](~~GetUserDevice~~): queries the details of a user endpoint device.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The ID of the vulnerability scanning node that detected the vulnerability. Used to filter detection results of a specified node. Valid values are obtained from:
	//
	// - [ListVulScanTasks](~~ListVulScanTasks~~): lists vulnerability scanning nodes.
	//
	// - [CreateVulScanTask](~~CreateVulScanTask~~): creates a vulnerability scanning node.
	//
	// example:
	//
	// vul-scan-task-4d7b1e9a6c38****
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The vulnerability title. Fuzzy match is supported. Matches both Chinese and English titles.
	//
	// example:
	//
	// Cumulative Update
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The patch IDs used for filtering. A maximum of 100 IDs can be specified. Duplicate values are not allowed.
	UpdateIds []*string `json:"UpdateIds,omitempty" xml:"UpdateIds,omitempty" type:"Repeated"`
	// The username. Fuzzy match is supported. The value can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), asterisks (*), hyphens (-), at signs (@), spaces, middle dots (·), and parentheses.
	//
	// example:
	//
	// John Smith
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The vulnerability risk level used for filtering. Valid values:
	//
	// - **High**: high risk.
	//
	// - **Mid**: medium risk.
	//
	// - **Low**: low risk.
	//
	// example:
	//
	// High
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
	// The vulnerability type used for filtering. Valid values:
	//
	// - **windows**: Windows system vulnerability.
	//
	// - **ai_agent**: AI Agent vulnerability.
	//
	// example:
	//
	// windows
	VulType *string `json:"VulType,omitempty" xml:"VulType,omitempty"`
}

func (s ListVulnerabilitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVulnerabilitiesRequest) GoString() string {
	return s.String()
}

func (s *ListVulnerabilitiesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVulnerabilitiesRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListVulnerabilitiesRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *ListVulnerabilitiesRequest) GetDevType() *string {
	return s.DevType
}

func (s *ListVulnerabilitiesRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListVulnerabilitiesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVulnerabilitiesRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListVulnerabilitiesRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListVulnerabilitiesRequest) GetTitle() *string {
	return s.Title
}

func (s *ListVulnerabilitiesRequest) GetUpdateIds() []*string {
	return s.UpdateIds
}

func (s *ListVulnerabilitiesRequest) GetUsername() *string {
	return s.Username
}

func (s *ListVulnerabilitiesRequest) GetVulLevel() *string {
	return s.VulLevel
}

func (s *ListVulnerabilitiesRequest) GetVulType() *string {
	return s.VulType
}

func (s *ListVulnerabilitiesRequest) SetCurrentPage(v int64) *ListVulnerabilitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetDepartment(v string) *ListVulnerabilitiesRequest {
	s.Department = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetDevTag(v string) *ListVulnerabilitiesRequest {
	s.DevTag = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetDevType(v string) *ListVulnerabilitiesRequest {
	s.DevType = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetHostname(v string) *ListVulnerabilitiesRequest {
	s.Hostname = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetPageSize(v int64) *ListVulnerabilitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetSaseUserId(v string) *ListVulnerabilitiesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetScanTaskId(v string) *ListVulnerabilitiesRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetTitle(v string) *ListVulnerabilitiesRequest {
	s.Title = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetUpdateIds(v []*string) *ListVulnerabilitiesRequest {
	s.UpdateIds = v
	return s
}

func (s *ListVulnerabilitiesRequest) SetUsername(v string) *ListVulnerabilitiesRequest {
	s.Username = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetVulLevel(v string) *ListVulnerabilitiesRequest {
	s.VulLevel = &v
	return s
}

func (s *ListVulnerabilitiesRequest) SetVulType(v string) *ListVulnerabilitiesRequest {
	s.VulType = &v
	return s
}

func (s *ListVulnerabilitiesRequest) Validate() error {
	return dara.Validate(s)
}
