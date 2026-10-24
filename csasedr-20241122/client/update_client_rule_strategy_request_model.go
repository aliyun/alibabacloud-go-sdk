// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientRuleStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *UpdateClientRuleStrategyRequest
	GetActionType() *int32
	SetAllGroup(v int32) *UpdateClientRuleStrategyRequest
	GetAllGroup() *int32
	SetApprovalId(v string) *UpdateClientRuleStrategyRequest
	GetApprovalId() *string
	SetCmdline(v string) *UpdateClientRuleStrategyRequest
	GetCmdline() *string
	SetDescription(v string) *UpdateClientRuleStrategyRequest
	GetDescription() *string
	SetExcludeUserNameList(v []*string) *UpdateClientRuleStrategyRequest
	GetExcludeUserNameList() []*string
	SetFilePath(v string) *UpdateClientRuleStrategyRequest
	GetFilePath() *string
	SetId(v int64) *UpdateClientRuleStrategyRequest
	GetId() *int64
	SetInSwitch(v bool) *UpdateClientRuleStrategyRequest
	GetInSwitch() *bool
	SetIncludeGroupIdList(v []*string) *UpdateClientRuleStrategyRequest
	GetIncludeGroupIdList() []*string
	SetIp(v string) *UpdateClientRuleStrategyRequest
	GetIp() *string
	SetMainButtonTextCh(v string) *UpdateClientRuleStrategyRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *UpdateClientRuleStrategyRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *UpdateClientRuleStrategyRequest
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *UpdateClientRuleStrategyRequest
	GetMatchTargetIds() []*string
	SetMd5List(v string) *UpdateClientRuleStrategyRequest
	GetMd5List() *string
	SetMinorButtonTextCh(v string) *UpdateClientRuleStrategyRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *UpdateClientRuleStrategyRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *UpdateClientRuleStrategyRequest
	GetName() *string
	SetNewFilePath(v string) *UpdateClientRuleStrategyRequest
	GetNewFilePath() *string
	SetOutSwitch(v bool) *UpdateClientRuleStrategyRequest
	GetOutSwitch() *bool
	SetParentCmdline(v string) *UpdateClientRuleStrategyRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *UpdateClientRuleStrategyRequest
	GetParentProcPath() *string
	SetPlatform(v string) *UpdateClientRuleStrategyRequest
	GetPlatform() *string
	SetPolicyType(v int32) *UpdateClientRuleStrategyRequest
	GetPolicyType() *int32
	SetPort(v string) *UpdateClientRuleStrategyRequest
	GetPort() *string
	SetPortOnList(v []*string) *UpdateClientRuleStrategyRequest
	GetPortOnList() []*string
	SetPriority(v int32) *UpdateClientRuleStrategyRequest
	GetPriority() *int32
	SetProcPath(v string) *UpdateClientRuleStrategyRequest
	GetProcPath() *string
	SetRegistryContent(v string) *UpdateClientRuleStrategyRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *UpdateClientRuleStrategyRequest
	GetRegistryKey() *string
	SetRuleType(v int32) *UpdateClientRuleStrategyRequest
	GetRuleType() *int32
	SetSoftWareOnList(v []*string) *UpdateClientRuleStrategyRequest
	GetSoftWareOnList() []*string
	SetStatus(v int32) *UpdateClientRuleStrategyRequest
	GetStatus() *int32
	SetStrategySourceType(v string) *UpdateClientRuleStrategyRequest
	GetStrategySourceType() *string
	SetSupportsApproval(v int32) *UpdateClientRuleStrategyRequest
	GetSupportsApproval() *int32
	SetUsContent(v string) *UpdateClientRuleStrategyRequest
	GetUsContent() *string
	SetUsTitle(v string) *UpdateClientRuleStrategyRequest
	GetUsTitle() *string
	SetZhContent(v string) *UpdateClientRuleStrategyRequest
	GetZhContent() *string
	SetZhTitle(v string) *UpdateClientRuleStrategyRequest
	GetZhTitle() *string
}

type UpdateClientRuleStrategyRequest struct {
	// The policy action. Valid values:
	//
	// - **0**: monitor mode
	//
	// - **1**: prevent startup
	//
	// - **2**: block mode
	//
	// Valid values for custom rules of proactive defense (PolicyType=2):
	//
	// - **0**: whitelist and allow
	//
	// - **1**: blacklist and block
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Specifies whether the policy takes effect for all user groups. Valid values:
	//
	// - **0**: no
	//
	// - **1**: yes
	//
	// example:
	//
	// 1
	AllGroup *int32 `json:"AllGroup,omitempty" xml:"AllGroup,omitempty"`
	// The ID of the approval process.
	//
	// example:
	//
	// approval-165e6738ad9d5c46
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The matched process command line.
	//
	// example:
	//
	// C:\\Windows\\System32\\RuntimeBroker.exe -Embedding
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The policy description.
	//
	// > The maximum length is 128 characters.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of usernames for which the policy does not take effect.
	ExcludeUserNameList []*string `json:"ExcludeUserNameList,omitempty" xml:"ExcludeUserNameList,omitempty" type:"Repeated"`
	// The matched file path.
	//
	// example:
	//
	// C:\\Program Files\\Example\\example.exe
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4072
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The inbound network switch.
	//
	// example:
	//
	// true
	InSwitch *bool `json:"InSwitch,omitempty" xml:"InSwitch,omitempty"`
	// The list of IDs of the user groups for which the policy takes effect.
	IncludeGroupIdList []*string `json:"IncludeGroupIdList,omitempty" xml:"IncludeGroupIdList,omitempty" type:"Repeated"`
	// The matched IP address.
	//
	// example:
	//
	// 192.0.2.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The text of the primary button in the Chinese button configuration.
	//
	// example:
	//
	// test
	MainButtonTextCh *string `json:"MainButtonTextCh,omitempty" xml:"MainButtonTextCh,omitempty"`
	// The text of the primary button in the English button configuration.
	//
	// example:
	//
	// test
	MainButtonTextEn *string `json:"MainButtonTextEn,omitempty" xml:"MainButtonTextEn,omitempty"`
	// The effective scope mode. Valid values:
	//
	// - **UserGroupAll**: all user groups
	//
	// - **UserGroupNormal**: specified user groups
	//
	// - **DeviceGroupNormal**: specified device groups
	//
	// - **DevTagNormal**: specified device tags
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The list of targets in the effective scope. The meaning of the targets is determined by MatchMode.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The MD5 match string for the file.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	Md5List *string `json:"Md5List,omitempty" xml:"Md5List,omitempty"`
	// The text of the secondary button in the Chinese button configuration.
	//
	// example:
	//
	// test
	MinorButtonTextCh *string `json:"MinorButtonTextCh,omitempty" xml:"MinorButtonTextCh,omitempty"`
	// The text of the secondary button in the English button configuration.
	//
	// example:
	//
	// test
	MinorButtonTextEn *string `json:"MinorButtonTextEn,omitempty" xml:"MinorButtonTextEn,omitempty"`
	// The policy name.
	//
	// > The maximum length is 128 characters.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The matched new file path.
	//
	// example:
	//
	// C:\\Program Files\\Example\\renamed.exe
	NewFilePath *string `json:"NewFilePath,omitempty" xml:"NewFilePath,omitempty"`
	// The outbound network switch.
	//
	// example:
	//
	// true
	OutSwitch *bool `json:"OutSwitch,omitempty" xml:"OutSwitch,omitempty"`
	// The matched parent process command line.
	//
	// example:
	//
	// C:\\Windows\\System32\\svchost.exe -k DcomLaunch -p
	ParentCmdline *string `json:"ParentCmdline,omitempty" xml:"ParentCmdline,omitempty"`
	// The matched parent process path.
	//
	// example:
	//
	// C:\\Windows\\System32\\svchost.exe
	ParentProcPath *string `json:"ParentProcPath,omitempty" xml:"ParentProcPath,omitempty"`
	// The platform to which the rule applies. Valid values:
	//
	// - **linux**
	//
	// - **windows**
	//
	// - **mac**
	//
	// - **all**: all platforms
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The policy type. Valid values:
	//
	// - **0**: software hardening
	//
	// - **1**: firewall
	//
	// - **2**: proactive defense
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PolicyType *int32 `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The matched port.
	//
	// example:
	//
	// 443
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The list of enabled ports.
	//
	// > This parameter is valid only when the policy type is firewall.
	PortOnList []*string `json:"PortOnList,omitempty" xml:"PortOnList,omitempty" type:"Repeated"`
	// The policy priority. A smaller value indicates a higher priority. Valid values: greater than 0.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The matched process path.
	//
	// example:
	//
	// C:\\Windows\\System32\\RuntimeBroker.exe
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// The matched registry content.
	//
	// example:
	//
	// ExampleValue
	RegistryContent *string `json:"RegistryContent,omitempty" xml:"RegistryContent,omitempty"`
	// The matched registry key.
	//
	// example:
	//
	// HKEY_LOCAL_MACHINE\\Software\\Example
	RegistryKey *string `json:"RegistryKey,omitempty" xml:"RegistryKey,omitempty"`
	// The defense rule type. Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The list of main process names of the software for which integrity check is enabled. The elements correspond to the Process returned by GetSwitchInfoListByType.
	SoftWareOnList []*string `json:"SoftWareOnList,omitempty" xml:"SoftWareOnList,omitempty" type:"Repeated"`
	// The policy status. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule source for proactive defense (PolicyType=2). Valid values:
	//
	// - **SYSTEM**: system rules
	//
	// - **CUSTOM**: custom rules
	//
	// example:
	//
	// SYSTEM
	StrategySourceType *string `json:"StrategySourceType,omitempty" xml:"StrategySourceType,omitempty"`
	// Specifies whether approval requests are supported. Valid values:
	//
	// - **0**: not supported
	//
	// - **1**: supported
	//
	// example:
	//
	// 0
	SupportsApproval *int32 `json:"SupportsApproval,omitempty" xml:"SupportsApproval,omitempty"`
	// The English content in the pop-up prompt configuration.
	//
	// example:
	//
	// test
	UsContent *string `json:"UsContent,omitempty" xml:"UsContent,omitempty"`
	// The English title in the pop-up prompt configuration.
	//
	// example:
	//
	// test
	UsTitle *string `json:"UsTitle,omitempty" xml:"UsTitle,omitempty"`
	// The Chinese content in the pop-up prompt configuration.
	//
	// example:
	//
	// test
	ZhContent *string `json:"ZhContent,omitempty" xml:"ZhContent,omitempty"`
	// The Chinese title in the pop-up prompt configuration.
	//
	// example:
	//
	// test
	ZhTitle *string `json:"ZhTitle,omitempty" xml:"ZhTitle,omitempty"`
}

func (s UpdateClientRuleStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientRuleStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientRuleStrategyRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *UpdateClientRuleStrategyRequest) GetAllGroup() *int32 {
	return s.AllGroup
}

func (s *UpdateClientRuleStrategyRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *UpdateClientRuleStrategyRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateClientRuleStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateClientRuleStrategyRequest) GetExcludeUserNameList() []*string {
	return s.ExcludeUserNameList
}

func (s *UpdateClientRuleStrategyRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateClientRuleStrategyRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateClientRuleStrategyRequest) GetInSwitch() *bool {
	return s.InSwitch
}

func (s *UpdateClientRuleStrategyRequest) GetIncludeGroupIdList() []*string {
	return s.IncludeGroupIdList
}

func (s *UpdateClientRuleStrategyRequest) GetIp() *string {
	return s.Ip
}

func (s *UpdateClientRuleStrategyRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *UpdateClientRuleStrategyRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *UpdateClientRuleStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateClientRuleStrategyRequest) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *UpdateClientRuleStrategyRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *UpdateClientRuleStrategyRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *UpdateClientRuleStrategyRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *UpdateClientRuleStrategyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateClientRuleStrategyRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *UpdateClientRuleStrategyRequest) GetOutSwitch() *bool {
	return s.OutSwitch
}

func (s *UpdateClientRuleStrategyRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *UpdateClientRuleStrategyRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *UpdateClientRuleStrategyRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateClientRuleStrategyRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *UpdateClientRuleStrategyRequest) GetPort() *string {
	return s.Port
}

func (s *UpdateClientRuleStrategyRequest) GetPortOnList() []*string {
	return s.PortOnList
}

func (s *UpdateClientRuleStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateClientRuleStrategyRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *UpdateClientRuleStrategyRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *UpdateClientRuleStrategyRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *UpdateClientRuleStrategyRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *UpdateClientRuleStrategyRequest) GetSoftWareOnList() []*string {
	return s.SoftWareOnList
}

func (s *UpdateClientRuleStrategyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateClientRuleStrategyRequest) GetStrategySourceType() *string {
	return s.StrategySourceType
}

func (s *UpdateClientRuleStrategyRequest) GetSupportsApproval() *int32 {
	return s.SupportsApproval
}

func (s *UpdateClientRuleStrategyRequest) GetUsContent() *string {
	return s.UsContent
}

func (s *UpdateClientRuleStrategyRequest) GetUsTitle() *string {
	return s.UsTitle
}

func (s *UpdateClientRuleStrategyRequest) GetZhContent() *string {
	return s.ZhContent
}

func (s *UpdateClientRuleStrategyRequest) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *UpdateClientRuleStrategyRequest) SetActionType(v int32) *UpdateClientRuleStrategyRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetAllGroup(v int32) *UpdateClientRuleStrategyRequest {
	s.AllGroup = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetApprovalId(v string) *UpdateClientRuleStrategyRequest {
	s.ApprovalId = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetCmdline(v string) *UpdateClientRuleStrategyRequest {
	s.Cmdline = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetDescription(v string) *UpdateClientRuleStrategyRequest {
	s.Description = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetExcludeUserNameList(v []*string) *UpdateClientRuleStrategyRequest {
	s.ExcludeUserNameList = v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetFilePath(v string) *UpdateClientRuleStrategyRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetId(v int64) *UpdateClientRuleStrategyRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetInSwitch(v bool) *UpdateClientRuleStrategyRequest {
	s.InSwitch = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetIncludeGroupIdList(v []*string) *UpdateClientRuleStrategyRequest {
	s.IncludeGroupIdList = v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetIp(v string) *UpdateClientRuleStrategyRequest {
	s.Ip = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMainButtonTextCh(v string) *UpdateClientRuleStrategyRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMainButtonTextEn(v string) *UpdateClientRuleStrategyRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMatchMode(v string) *UpdateClientRuleStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMatchTargetIds(v []*string) *UpdateClientRuleStrategyRequest {
	s.MatchTargetIds = v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMd5List(v string) *UpdateClientRuleStrategyRequest {
	s.Md5List = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMinorButtonTextCh(v string) *UpdateClientRuleStrategyRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetMinorButtonTextEn(v string) *UpdateClientRuleStrategyRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetName(v string) *UpdateClientRuleStrategyRequest {
	s.Name = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetNewFilePath(v string) *UpdateClientRuleStrategyRequest {
	s.NewFilePath = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetOutSwitch(v bool) *UpdateClientRuleStrategyRequest {
	s.OutSwitch = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetParentCmdline(v string) *UpdateClientRuleStrategyRequest {
	s.ParentCmdline = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetParentProcPath(v string) *UpdateClientRuleStrategyRequest {
	s.ParentProcPath = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetPlatform(v string) *UpdateClientRuleStrategyRequest {
	s.Platform = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetPolicyType(v int32) *UpdateClientRuleStrategyRequest {
	s.PolicyType = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetPort(v string) *UpdateClientRuleStrategyRequest {
	s.Port = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetPortOnList(v []*string) *UpdateClientRuleStrategyRequest {
	s.PortOnList = v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetPriority(v int32) *UpdateClientRuleStrategyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetProcPath(v string) *UpdateClientRuleStrategyRequest {
	s.ProcPath = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetRegistryContent(v string) *UpdateClientRuleStrategyRequest {
	s.RegistryContent = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetRegistryKey(v string) *UpdateClientRuleStrategyRequest {
	s.RegistryKey = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetRuleType(v int32) *UpdateClientRuleStrategyRequest {
	s.RuleType = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetSoftWareOnList(v []*string) *UpdateClientRuleStrategyRequest {
	s.SoftWareOnList = v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetStatus(v int32) *UpdateClientRuleStrategyRequest {
	s.Status = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetStrategySourceType(v string) *UpdateClientRuleStrategyRequest {
	s.StrategySourceType = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetSupportsApproval(v int32) *UpdateClientRuleStrategyRequest {
	s.SupportsApproval = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetUsContent(v string) *UpdateClientRuleStrategyRequest {
	s.UsContent = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetUsTitle(v string) *UpdateClientRuleStrategyRequest {
	s.UsTitle = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetZhContent(v string) *UpdateClientRuleStrategyRequest {
	s.ZhContent = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) SetZhTitle(v string) *UpdateClientRuleStrategyRequest {
	s.ZhTitle = &v
	return s
}

func (s *UpdateClientRuleStrategyRequest) Validate() error {
	return dara.Validate(s)
}
