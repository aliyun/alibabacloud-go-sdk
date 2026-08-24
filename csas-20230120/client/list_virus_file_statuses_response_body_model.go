// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusFileStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileStatuses(v []*ListVirusFileStatusesResponseBodyFileStatuses) *ListVirusFileStatusesResponseBody
	GetFileStatuses() []*ListVirusFileStatusesResponseBodyFileStatuses
	SetRequestId(v string) *ListVirusFileStatusesResponseBody
	GetRequestId() *string
	SetTotalNum(v string) *ListVirusFileStatusesResponseBody
	GetTotalNum() *string
}

type ListVirusFileStatusesResponseBody struct {
	// The list of virus files.
	FileStatuses []*ListVirusFileStatusesResponseBodyFileStatuses `json:"FileStatuses,omitempty" xml:"FileStatuses,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of virus files that match the query conditions.
	//
	// example:
	//
	// 37
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListVirusFileStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusFileStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusFileStatusesResponseBody) GetFileStatuses() []*ListVirusFileStatusesResponseBodyFileStatuses {
	return s.FileStatuses
}

func (s *ListVirusFileStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusFileStatusesResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListVirusFileStatusesResponseBody) SetFileStatuses(v []*ListVirusFileStatusesResponseBodyFileStatuses) *ListVirusFileStatusesResponseBody {
	s.FileStatuses = v
	return s
}

func (s *ListVirusFileStatusesResponseBody) SetRequestId(v string) *ListVirusFileStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusFileStatusesResponseBody) SetTotalNum(v string) *ListVirusFileStatusesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVirusFileStatusesResponseBody) Validate() error {
	if s.FileStatuses != nil {
		for _, item := range s.FileStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusFileStatusesResponseBodyFileStatuses struct {
	// The time when the administrator initiated the disposition, in the format yyyy-MM-dd HH:mm:ss (UTC+8). An empty string is returned when the disposition was not initiated by an administrator.
	//
	// example:
	//
	// 2026-08-21 09:30:12
	ConsoleOperationTime *string `json:"ConsoleOperationTime,omitempty" xml:"ConsoleOperationTime,omitempty"`
	// The name of the department to which the user belongs. Multiple departments are separated by commas (,). The nearest department name in the organizational structure is returned, not the full path.
	//
	// example:
	//
	// R&D Department,Security Team
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The unique identifier of the user\\"s endpoint device that detected this virus file.
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
	// The time when the virus file was discovered, in the format yyyy-MM-dd HH:mm:ss (UTC+8). A hyphen (-) is returned when no record exists.
	//
	// example:
	//
	// 2026-08-21 03:12:07
	DiscoveryTime *string `json:"DiscoveryTime,omitempty" xml:"DiscoveryTime,omitempty"`
	// The MD5 hash of the virus file.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The absolute path of the virus file on the user\\"s endpoint device.
	//
	// example:
	//
	// C:\\Users\\Public\\Downloads\\setup.exe
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The disposition status. Valid values:
	//
	// - **Pending**: Pending disposition.
	//
	// - **Processed**: Disposed.
	//
	// example:
	//
	// Pending
	FileProcessStatus *string `json:"FileProcessStatus,omitempty" xml:"FileProcessStatus,omitempty"`
	// The size of the virus file, in bytes.
	//
	// example:
	//
	// 20480
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The hostname of the user\\"s endpoint device.
	//
	// example:
	//
	// DESKTOP-8A3F
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The disposition action that has been performed. An empty string is returned when no disposition has been performed. Valid values:
	//
	// - **AdminQuarantine**: Quarantined by administrator.
	//
	// - **AdminTrust**: Trusted by administrator.
	//
	// - **UserQuarantine**: Quarantined by endpoint user.
	//
	// - **UserTrust**: Trusted by endpoint user.
	//
	// - **AutoQuarantine**: Automatically quarantined based on policy.
	//
	// - **Fail**: Disposition failed.
	//
	// example:
	//
	// AdminQuarantine
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The effective period of the disposition, in the format yyyy-MM-dd HH:mm:ss (UTC+8). The later of the actual disposition time on the user\\"s endpoint device and the time when the administrator initiated the disposition is used. A hyphen (-) is returned when no disposition has been performed.
	//
	// example:
	//
	// 2026-08-21 09:31:45
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// The risk level. Valid values:
	//
	// - **High**: High risk.
	//
	// - **Mid**: Medium risk.
	//
	// - **Low**: Low risk.
	//
	// example:
	//
	// High
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The user ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The ID of the virus scan task that detected this virus file. An empty string is returned when the file is detected by real-time protection.
	//
	// example:
	//
	// v1:1024772
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The execution result description of the disposition or scan, reported by the user\\"s endpoint device. If a disposition record exists, the execution result of the disposition task is returned. Otherwise, the execution result of the scan task is returned.
	//
	// example:
	//
	// quarantine success
	TaskExecutionInfo *string `json:"TaskExecutionInfo,omitempty" xml:"TaskExecutionInfo,omitempty"`
	// The username.
	//
	// example:
	//
	// John Smith
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The detection source of the virus file. Valid values:
	//
	// - **Task**: Detected by a virus scan task.
	//
	// - **Download**: Detected by real-time protection during file download.
	//
	// - **Process**: Detected by real-time protection during process execution.
	//
	// example:
	//
	// Task
	VirusFileSource *string `json:"VirusFileSource,omitempty" xml:"VirusFileSource,omitempty"`
	// The virus type. Valid values:
	//
	// - **Backdoor**: Backdoor program.
	//
	// - **DDoS**: DDoS Trojan.
	//
	// - **Downloader**: Downloader Trojan.
	//
	// - **Engtest**: DPI engine test program.
	//
	// - **Hacktool**: Hacking tool.
	//
	// - **Trojan**: Self-mutating Trojan.
	//
	// - **Malbaseware**: Contaminated base software.
	//
	// - **MalScript**: Malicious script.
	//
	// - **Malware**: Malicious program.
	//
	// - **Miner**: Mining programs.
	//
	// - **Proxytool**: Proxy tool.
	//
	// - **RansomWare**: Ransomware.
	//
	// - **RiskWare**: Risky software.
	//
	// - **Rootkit**: Kernel-hidden program.
	//
	// - **Stealer**: Credential-stealing tool.
	//
	// - **Scanner**: Scanner.
	//
	// - **Suspicious**: Suspicious program.
	//
	// - **Virus**: File-infecting virus.
	//
	// - **WebShell**: Web shell.
	//
	// - **Worm**: Worms.
	//
	// - **BlackList**: File that hit the blacklist.
	//
	// - **Exp**: Vulnerability exploits program.
	//
	// - **Patcher**: Cracking program.
	//
	// - **Gametool**: Private server tool.
	//
	// - **AdWare**: Adware.
	//
	// - **Maldoc**: Malicious document.
	//
	// example:
	//
	// Virus
	VirusType *string `json:"VirusType,omitempty" xml:"VirusType,omitempty"`
}

func (s ListVirusFileStatusesResponseBodyFileStatuses) String() string {
	return dara.Prettify(s)
}

func (s ListVirusFileStatusesResponseBodyFileStatuses) GoString() string {
	return s.String()
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetConsoleOperationTime() *string {
	return s.ConsoleOperationTime
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetDepartment() *string {
	return s.Department
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetDevTag() *string {
	return s.DevTag
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetDevType() *string {
	return s.DevType
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetDiscoveryTime() *string {
	return s.DiscoveryTime
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetFileMd5() *string {
	return s.FileMd5
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetFilePath() *string {
	return s.FilePath
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetFileProcessStatus() *string {
	return s.FileProcessStatus
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetFileSize() *int32 {
	return s.FileSize
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetHostname() *string {
	return s.Hostname
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetOperation() *string {
	return s.Operation
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetOperationTime() *string {
	return s.OperationTime
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetTaskExecutionInfo() *string {
	return s.TaskExecutionInfo
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetUsername() *string {
	return s.Username
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetVirusFileSource() *string {
	return s.VirusFileSource
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) GetVirusType() *string {
	return s.VirusType
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetConsoleOperationTime(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.ConsoleOperationTime = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetDepartment(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.Department = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetDevTag(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.DevTag = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetDevType(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.DevType = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetDiscoveryTime(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.DiscoveryTime = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetFileMd5(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.FileMd5 = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetFilePath(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.FilePath = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetFileProcessStatus(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.FileProcessStatus = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetFileSize(v int32) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.FileSize = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetHostname(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.Hostname = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetOperation(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.Operation = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetOperationTime(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.OperationTime = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetRiskLevel(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.RiskLevel = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetSaseUserId(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.SaseUserId = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetScanTaskId(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.ScanTaskId = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetTaskExecutionInfo(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.TaskExecutionInfo = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetUsername(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.Username = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetVirusFileSource(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.VirusFileSource = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) SetVirusType(v string) *ListVirusFileStatusesResponseBodyFileStatuses {
	s.VirusType = &v
	return s
}

func (s *ListVirusFileStatusesResponseBodyFileStatuses) Validate() error {
	return dara.Validate(s)
}
