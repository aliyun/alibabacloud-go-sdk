// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusFileStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTag(v string) *UpdateVirusFileStatusRequest
	GetDevTag() *string
	SetFileMd5(v string) *UpdateVirusFileStatusRequest
	GetFileMd5() *string
	SetFilePath(v string) *UpdateVirusFileStatusRequest
	GetFilePath() *string
	SetOperation(v string) *UpdateVirusFileStatusRequest
	GetOperation() *string
	SetVirusType(v string) *UpdateVirusFileStatusRequest
	GetVirusType() *string
}

type UpdateVirusFileStatusRequest struct {
	// The unique identifier of the user terminal device where the virus file is located. The value can be up to 64 characters in length. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// The MD5 value of the virus file. The value must be a 32-character hexadecimal string. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// The absolute path of the virus file on the user terminal device. You can obtain the value from the following operation:
	//
	// - [ListVirusFileStatuses](~~ListVirusFileStatuses~~): lists virus file statuses.
	//
	// This parameter is required.
	//
	// example:
	//
	// C:\\Users\\Public\\Downloads\\setup.exe
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The disposal action. Valid values:
	//
	// - **AdminQuarantine**: quarantines the virus file. The server creates a disposal task and returns a TaskId. The user terminal device pulls and executes the quarantine.
	//
	// - **AdminTrust**: trusts the virus file. Only the disposal status is updated. No disposal task is created, and TaskId returns an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// AdminQuarantine
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The virus type. This parameter is used for synchronization to update the virus type of the file. Valid values:
	//
	// - **Backdoor**: backdoor program.
	//
	// - **DDoS**: DDoS Trojan.
	//
	// - **Downloader**: downloader Trojan.
	//
	// - **Engtest**: DPI engine test program.
	//
	// - **Hacktool**: hacker tool.
	//
	// - **Trojan**: self-mutating Trojan.
	//
	// - **Malbaseware**: contaminated base software.
	//
	// - **MalScript**: malicious script.
	//
	// - **Malware**: malicious program.
	//
	// - **Miner**: mining programs.
	//
	// - **Proxytool**: proxy tool.
	//
	// - **RansomWare**: ransomware.
	//
	// - **RiskWare**: riskware.
	//
	// - **Rootkit**: kernel-hidden program.
	//
	// - **Stealer**: credential stealer.
	//
	// - **Scanner**: scanner.
	//
	// - **Suspicious**: suspicious program.
	//
	// - **Virus**: file-infecting virus.
	//
	// - **WebShell**: webshell.
	//
	// - **Worm**: worms.
	//
	// - **BlackList**: file that hit a blacklist entry.
	//
	// - **Exp**: vulnerability exploits program.
	//
	// - **Patcher**: cracking program.
	//
	// - **Gametool**: private server tool.
	//
	// - **AdWare**: adware.
	//
	// - **Maldoc**: malicious document.
	//
	// example:
	//
	// Virus
	VirusType *string `json:"VirusType,omitempty" xml:"VirusType,omitempty"`
}

func (s UpdateVirusFileStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusFileStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirusFileStatusRequest) GetDevTag() *string {
	return s.DevTag
}

func (s *UpdateVirusFileStatusRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *UpdateVirusFileStatusRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateVirusFileStatusRequest) GetOperation() *string {
	return s.Operation
}

func (s *UpdateVirusFileStatusRequest) GetVirusType() *string {
	return s.VirusType
}

func (s *UpdateVirusFileStatusRequest) SetDevTag(v string) *UpdateVirusFileStatusRequest {
	s.DevTag = &v
	return s
}

func (s *UpdateVirusFileStatusRequest) SetFileMd5(v string) *UpdateVirusFileStatusRequest {
	s.FileMd5 = &v
	return s
}

func (s *UpdateVirusFileStatusRequest) SetFilePath(v string) *UpdateVirusFileStatusRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateVirusFileStatusRequest) SetOperation(v string) *UpdateVirusFileStatusRequest {
	s.Operation = &v
	return s
}

func (s *UpdateVirusFileStatusRequest) SetVirusType(v string) *UpdateVirusFileStatusRequest {
	s.VirusType = &v
	return s
}

func (s *UpdateVirusFileStatusRequest) Validate() error {
	return dara.Validate(s)
}
