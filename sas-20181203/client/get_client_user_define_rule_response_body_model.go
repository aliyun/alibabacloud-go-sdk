// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserDefineRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetClientUserDefineRuleResponseBody
	GetRequestId() *string
	SetUserDefineRuleDetail(v *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) *GetClientUserDefineRuleResponseBody
	GetUserDefineRuleDetail() *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail
}

type GetClientUserDefineRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the custom defense rule.
	UserDefineRuleDetail *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail `json:"UserDefineRuleDetail,omitempty" xml:"UserDefineRuleDetail,omitempty" type:"Struct"`
}

func (s GetClientUserDefineRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserDefineRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientUserDefineRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientUserDefineRuleResponseBody) GetUserDefineRuleDetail() *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	return s.UserDefineRuleDetail
}

func (s *GetClientUserDefineRuleResponseBody) SetRequestId(v string) *GetClientUserDefineRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBody) SetUserDefineRuleDetail(v *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) *GetClientUserDefineRuleResponseBody {
	s.UserDefineRuleDetail = v
	return s
}

func (s *GetClientUserDefineRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClientUserDefineRuleResponseBodyUserDefineRuleDetail struct {
	// The action of the custom defense rule. Valid values:
	//
	// 	- **0**: allow
	//
	// 	- **1**: block
	//
	// example:
	//
	// 0
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
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
	// The time when the custom defense rule was created.
	//
	// example:
	//
	// 167118088****
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the custom defense rule was last modified.
	//
	// example:
	//
	// 167118088****
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.240.XX.XX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the custom defense rule.
	//
	// example:
	//
	// 200****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The hash values of processes.
	//
	// example:
	//
	// 0c9045b5bec90f9825f1f3f64dd4****
	Md5List *string `json:"Md5List,omitempty" xml:"Md5List,omitempty"`
	// The name of the custom defense rule.
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
	// c:/windows/system32/i****
	ParentProcPath *string `json:"ParentProcPath,omitempty" xml:"ParentProcPath,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// 	- **all**
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The port number.
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
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GoString() string {
	return s.String()
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetActionType() *string {
	return s.ActionType
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetCmdline() *string {
	return s.Cmdline
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetDomain() *string {
	return s.Domain
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetFilePath() *string {
	return s.FilePath
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetIP() *string {
	return s.IP
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetId() *int64 {
	return s.Id
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetMd5List() *string {
	return s.Md5List
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetName() *string {
	return s.Name
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetPlatform() *string {
	return s.Platform
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetPort() *int32 {
	return s.Port
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetPortStr() *string {
	return s.PortStr
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetProcPath() *string {
	return s.ProcPath
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) GetType() *int32 {
	return s.Type
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetActionType(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.ActionType = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetCmdline(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Cmdline = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetDomain(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Domain = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetFilePath(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.FilePath = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetGmtCreate(v int64) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.GmtCreate = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetGmtModified(v int64) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.GmtModified = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetIP(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.IP = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetId(v int64) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Id = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetMd5List(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Md5List = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetName(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Name = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetNewFilePath(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.NewFilePath = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetParentCmdline(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.ParentCmdline = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetParentProcPath(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.ParentProcPath = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetPlatform(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Platform = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetPort(v int32) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Port = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetPortStr(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.PortStr = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetProcPath(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.ProcPath = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetRegistryContent(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.RegistryContent = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetRegistryKey(v string) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.RegistryKey = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) SetType(v int32) *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail {
	s.Type = &v
	return s
}

func (s *GetClientUserDefineRuleResponseBodyUserDefineRuleDetail) Validate() error {
	return dara.Validate(s)
}
