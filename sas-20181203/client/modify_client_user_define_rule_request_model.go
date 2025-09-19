// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDefineRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *ModifyClientUserDefineRuleRequest
	GetActionType() *int32
	SetCmdline(v string) *ModifyClientUserDefineRuleRequest
	GetCmdline() *string
	SetDomain(v string) *ModifyClientUserDefineRuleRequest
	GetDomain() *string
	SetFilePath(v string) *ModifyClientUserDefineRuleRequest
	GetFilePath() *string
	SetIP(v string) *ModifyClientUserDefineRuleRequest
	GetIP() *string
	SetId(v int64) *ModifyClientUserDefineRuleRequest
	GetId() *int64
	SetMd5List(v string) *ModifyClientUserDefineRuleRequest
	GetMd5List() *string
	SetName(v string) *ModifyClientUserDefineRuleRequest
	GetName() *string
	SetNewFilePath(v string) *ModifyClientUserDefineRuleRequest
	GetNewFilePath() *string
	SetParentCmdline(v string) *ModifyClientUserDefineRuleRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *ModifyClientUserDefineRuleRequest
	GetParentProcPath() *string
	SetPlatform(v string) *ModifyClientUserDefineRuleRequest
	GetPlatform() *string
	SetPort(v int32) *ModifyClientUserDefineRuleRequest
	GetPort() *int32
	SetPortStr(v string) *ModifyClientUserDefineRuleRequest
	GetPortStr() *string
	SetProcPath(v string) *ModifyClientUserDefineRuleRequest
	GetProcPath() *string
	SetRegistryContent(v string) *ModifyClientUserDefineRuleRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *ModifyClientUserDefineRuleRequest
	GetRegistryKey() *string
	SetType(v int32) *ModifyClientUserDefineRuleRequest
	GetType() *int32
}

type ModifyClientUserDefineRuleRequest struct {
	// The action of the rule. Valid values:
	//
	// 	- **0**: allow
	//
	// 	- **1**: block
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
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
	// /etc/pam.d/su****
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 1.1.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the rule.
	//
	// >  You can call the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) operation to query the IDs of rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// 210****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The hash values of processes.
	//
	// example:
	//
	// aa5ee3ed4363c9d195a591a70281****,3e522d6f3bf5cf88bb77e9ff3d13****
	Md5List *string `json:"Md5List,omitempty" xml:"Md5List,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule\\*\\*\\*\\*
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
	// C:/Windows/System32/cmd****
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
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	PortStr *string `json:"PortStr,omitempty" xml:"PortStr,omitempty"`
	// The path to the process.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// /root/1111/****
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// The registry value.
	//
	// example:
	//
	// SECOH-QAD****
	RegistryContent *string `json:"RegistryContent,omitempty" xml:"RegistryContent,omitempty"`
	// The registry key.
	//
	// example:
	//
	// HKEY_DYN_****
	RegistryKey *string `json:"RegistryKey,omitempty" xml:"RegistryKey,omitempty"`
	// The type of the rule. Valid values:
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
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyClientUserDefineRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDefineRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDefineRuleRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModifyClientUserDefineRuleRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *ModifyClientUserDefineRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyClientUserDefineRuleRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ModifyClientUserDefineRuleRequest) GetIP() *string {
	return s.IP
}

func (s *ModifyClientUserDefineRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyClientUserDefineRuleRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *ModifyClientUserDefineRuleRequest) GetName() *string {
	return s.Name
}

func (s *ModifyClientUserDefineRuleRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *ModifyClientUserDefineRuleRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *ModifyClientUserDefineRuleRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *ModifyClientUserDefineRuleRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ModifyClientUserDefineRuleRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyClientUserDefineRuleRequest) GetPortStr() *string {
	return s.PortStr
}

func (s *ModifyClientUserDefineRuleRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *ModifyClientUserDefineRuleRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *ModifyClientUserDefineRuleRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *ModifyClientUserDefineRuleRequest) GetType() *int32 {
	return s.Type
}

func (s *ModifyClientUserDefineRuleRequest) SetActionType(v int32) *ModifyClientUserDefineRuleRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetCmdline(v string) *ModifyClientUserDefineRuleRequest {
	s.Cmdline = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetDomain(v string) *ModifyClientUserDefineRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetFilePath(v string) *ModifyClientUserDefineRuleRequest {
	s.FilePath = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetIP(v string) *ModifyClientUserDefineRuleRequest {
	s.IP = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetId(v int64) *ModifyClientUserDefineRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetMd5List(v string) *ModifyClientUserDefineRuleRequest {
	s.Md5List = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetName(v string) *ModifyClientUserDefineRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetNewFilePath(v string) *ModifyClientUserDefineRuleRequest {
	s.NewFilePath = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetParentCmdline(v string) *ModifyClientUserDefineRuleRequest {
	s.ParentCmdline = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetParentProcPath(v string) *ModifyClientUserDefineRuleRequest {
	s.ParentProcPath = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetPlatform(v string) *ModifyClientUserDefineRuleRequest {
	s.Platform = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetPort(v int32) *ModifyClientUserDefineRuleRequest {
	s.Port = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetPortStr(v string) *ModifyClientUserDefineRuleRequest {
	s.PortStr = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetProcPath(v string) *ModifyClientUserDefineRuleRequest {
	s.ProcPath = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetRegistryContent(v string) *ModifyClientUserDefineRuleRequest {
	s.RegistryContent = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetRegistryKey(v string) *ModifyClientUserDefineRuleRequest {
	s.RegistryKey = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) SetType(v int32) *ModifyClientUserDefineRuleRequest {
	s.Type = &v
	return s
}

func (s *ModifyClientUserDefineRuleRequest) Validate() error {
	return dara.Validate(s)
}
