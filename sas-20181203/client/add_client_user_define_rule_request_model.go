// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientUserDefineRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *AddClientUserDefineRuleRequest
	GetActionType() *int32
	SetCmdline(v string) *AddClientUserDefineRuleRequest
	GetCmdline() *string
	SetDomain(v string) *AddClientUserDefineRuleRequest
	GetDomain() *string
	SetFilePath(v string) *AddClientUserDefineRuleRequest
	GetFilePath() *string
	SetIP(v string) *AddClientUserDefineRuleRequest
	GetIP() *string
	SetMd5List(v string) *AddClientUserDefineRuleRequest
	GetMd5List() *string
	SetName(v string) *AddClientUserDefineRuleRequest
	GetName() *string
	SetNewFilePath(v string) *AddClientUserDefineRuleRequest
	GetNewFilePath() *string
	SetParentCmdline(v string) *AddClientUserDefineRuleRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *AddClientUserDefineRuleRequest
	GetParentProcPath() *string
	SetPlatform(v string) *AddClientUserDefineRuleRequest
	GetPlatform() *string
	SetPort(v int32) *AddClientUserDefineRuleRequest
	GetPort() *int32
	SetPortStr(v string) *AddClientUserDefineRuleRequest
	GetPortStr() *string
	SetProcPath(v string) *AddClientUserDefineRuleRequest
	GetProcPath() *string
	SetRegistryContent(v string) *AddClientUserDefineRuleRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *AddClientUserDefineRuleRequest
	GetRegistryKey() *string
	SetType(v int32) *AddClientUserDefineRuleRequest
	GetType() *int32
}

type AddClientUserDefineRuleRequest struct {
	// The action of the custom defense rule. Valid values:
	//
	// 	- **0**: allow
	//
	// 	- **1**: block
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The command line.
	//
	// example:
	//
	// /usr/sbin/s****
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The file path.
	//
	// example:
	//
	// /etc/pam****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.240.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The hash values of processes.
	//
	// example:
	//
	// 0c9045b5bec90f9825f1f3f64dd4****
	Md5List *string `json:"Md5List,omitempty" xml:"Md5List,omitempty"`
	// The name of the custom defense rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rule 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new file path after the file is renamed.
	//
	// example:
	//
	// /etc/pam****
	NewFilePath *string `json:"NewFilePath,omitempty" xml:"NewFilePath,omitempty"`
	// The parent command line.
	//
	// example:
	//
	// /usr/sbin/s****
	ParentCmdline *string `json:"ParentCmdline,omitempty" xml:"ParentCmdline,omitempty"`
	// The path to the parent process.
	//
	// example:
	//
	// c:/windows/system32/i****
	ParentProcPath *string `json:"ParentProcPath,omitempty" xml:"ParentProcPath,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// 	- **all**: all types
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The port number. This parameter is deprecated.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The port number. Valid values: **1 to 65535**.
	//
	// example:
	//
	// 80
	PortStr *string `json:"PortStr,omitempty" xml:"PortStr,omitempty"`
	// The path to the process.
	//
	// example:
	//
	// c:/windows/system32/i****
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// The registry value.
	//
	// example:
	//
	// *SECOH-QAD.exe*
	RegistryContent *string `json:"RegistryContent,omitempty" xml:"RegistryContent,omitempty"`
	// The registry key.
	//
	// example:
	//
	// HKEY_DYN_****
	RegistryKey *string `json:"RegistryKey,omitempty" xml:"RegistryKey,omitempty"`
	// The type of the custom defense rule. Valid values:
	//
	// 	- **1**: Process hash
	//
	// 	- **2**: Command line
	//
	// 	- **3**: Process Network
	//
	// 	- **4**: File Read and Write
	//
	// 	- **5**: Operation on Registry
	//
	// 	- **6**: Dynamic-link Library Loading
	//
	// 	- **7**: File Renaming
	//
	// 	- **8**: Network domain name
	//
	// 	- **9**: Network ip
	//
	// 	- **10**: File Path
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddClientUserDefineRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddClientUserDefineRuleRequest) GoString() string {
	return s.String()
}

func (s *AddClientUserDefineRuleRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *AddClientUserDefineRuleRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *AddClientUserDefineRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddClientUserDefineRuleRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *AddClientUserDefineRuleRequest) GetIP() *string {
	return s.IP
}

func (s *AddClientUserDefineRuleRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *AddClientUserDefineRuleRequest) GetName() *string {
	return s.Name
}

func (s *AddClientUserDefineRuleRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *AddClientUserDefineRuleRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *AddClientUserDefineRuleRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *AddClientUserDefineRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *AddClientUserDefineRuleRequest) GetPort() *int32 {
	return s.Port
}

func (s *AddClientUserDefineRuleRequest) GetPortStr() *string {
	return s.PortStr
}

func (s *AddClientUserDefineRuleRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *AddClientUserDefineRuleRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *AddClientUserDefineRuleRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *AddClientUserDefineRuleRequest) GetType() *int32 {
	return s.Type
}

func (s *AddClientUserDefineRuleRequest) SetActionType(v int32) *AddClientUserDefineRuleRequest {
	s.ActionType = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetCmdline(v string) *AddClientUserDefineRuleRequest {
	s.Cmdline = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetDomain(v string) *AddClientUserDefineRuleRequest {
	s.Domain = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetFilePath(v string) *AddClientUserDefineRuleRequest {
	s.FilePath = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetIP(v string) *AddClientUserDefineRuleRequest {
	s.IP = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetMd5List(v string) *AddClientUserDefineRuleRequest {
	s.Md5List = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetName(v string) *AddClientUserDefineRuleRequest {
	s.Name = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetNewFilePath(v string) *AddClientUserDefineRuleRequest {
	s.NewFilePath = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetParentCmdline(v string) *AddClientUserDefineRuleRequest {
	s.ParentCmdline = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetParentProcPath(v string) *AddClientUserDefineRuleRequest {
	s.ParentProcPath = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetPlatform(v string) *AddClientUserDefineRuleRequest {
	s.Platform = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetPort(v int32) *AddClientUserDefineRuleRequest {
	s.Port = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetPortStr(v string) *AddClientUserDefineRuleRequest {
	s.PortStr = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetProcPath(v string) *AddClientUserDefineRuleRequest {
	s.ProcPath = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetRegistryContent(v string) *AddClientUserDefineRuleRequest {
	s.RegistryContent = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetRegistryKey(v string) *AddClientUserDefineRuleRequest {
	s.RegistryKey = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) SetType(v int32) *AddClientUserDefineRuleRequest {
	s.Type = &v
	return s
}

func (s *AddClientUserDefineRuleRequest) Validate() error {
	return dara.Validate(s)
}
