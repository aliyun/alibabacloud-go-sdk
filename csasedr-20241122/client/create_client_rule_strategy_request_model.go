// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientRuleStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *CreateClientRuleStrategyRequest
	GetActionType() *int32
	SetAllGroup(v int32) *CreateClientRuleStrategyRequest
	GetAllGroup() *int32
	SetApprovalId(v string) *CreateClientRuleStrategyRequest
	GetApprovalId() *string
	SetCmdline(v string) *CreateClientRuleStrategyRequest
	GetCmdline() *string
	SetDescription(v string) *CreateClientRuleStrategyRequest
	GetDescription() *string
	SetExcludeUserNameList(v []*string) *CreateClientRuleStrategyRequest
	GetExcludeUserNameList() []*string
	SetFilePath(v string) *CreateClientRuleStrategyRequest
	GetFilePath() *string
	SetInSwitch(v bool) *CreateClientRuleStrategyRequest
	GetInSwitch() *bool
	SetIncludeGroupIdList(v []*string) *CreateClientRuleStrategyRequest
	GetIncludeGroupIdList() []*string
	SetIp(v string) *CreateClientRuleStrategyRequest
	GetIp() *string
	SetMainButtonTextCh(v string) *CreateClientRuleStrategyRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *CreateClientRuleStrategyRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *CreateClientRuleStrategyRequest
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *CreateClientRuleStrategyRequest
	GetMatchTargetIds() []*string
	SetMd5List(v string) *CreateClientRuleStrategyRequest
	GetMd5List() *string
	SetMinorButtonTextCh(v string) *CreateClientRuleStrategyRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *CreateClientRuleStrategyRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *CreateClientRuleStrategyRequest
	GetName() *string
	SetNewFilePath(v string) *CreateClientRuleStrategyRequest
	GetNewFilePath() *string
	SetOutSwitch(v bool) *CreateClientRuleStrategyRequest
	GetOutSwitch() *bool
	SetParentCmdline(v string) *CreateClientRuleStrategyRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *CreateClientRuleStrategyRequest
	GetParentProcPath() *string
	SetPlatform(v string) *CreateClientRuleStrategyRequest
	GetPlatform() *string
	SetPolicyType(v int32) *CreateClientRuleStrategyRequest
	GetPolicyType() *int32
	SetPort(v string) *CreateClientRuleStrategyRequest
	GetPort() *string
	SetPortOnList(v []*string) *CreateClientRuleStrategyRequest
	GetPortOnList() []*string
	SetPriority(v int32) *CreateClientRuleStrategyRequest
	GetPriority() *int32
	SetProcPath(v string) *CreateClientRuleStrategyRequest
	GetProcPath() *string
	SetRegistryContent(v string) *CreateClientRuleStrategyRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *CreateClientRuleStrategyRequest
	GetRegistryKey() *string
	SetRuleType(v int32) *CreateClientRuleStrategyRequest
	GetRuleType() *int32
	SetSelectionKey(v string) *CreateClientRuleStrategyRequest
	GetSelectionKey() *string
	SetSoftWareOnList(v []*string) *CreateClientRuleStrategyRequest
	GetSoftWareOnList() []*string
	SetStatus(v int32) *CreateClientRuleStrategyRequest
	GetStatus() *int32
	SetStrategySourceType(v string) *CreateClientRuleStrategyRequest
	GetStrategySourceType() *string
	SetSupportsApproval(v int32) *CreateClientRuleStrategyRequest
	GetSupportsApproval() *int32
	SetUsContent(v string) *CreateClientRuleStrategyRequest
	GetUsContent() *string
	SetUsTitle(v string) *CreateClientRuleStrategyRequest
	GetUsTitle() *string
	SetZhContent(v string) *CreateClientRuleStrategyRequest
	GetZhContent() *string
	SetZhTitle(v string) *CreateClientRuleStrategyRequest
	GetZhTitle() *string
}

type CreateClientRuleStrategyRequest struct {
	// The action of the policy. Valid values:
	//
	// - **0**: monitor mode.
	//
	// - **1**: block startup.
	//
	// - **2**: Block Mode.
	//
	// Valid values for custom rules of proactive defense (PolicyType=2):
	//
	// - **0**: add to allowlist and allow.
	//
	// - **1**: add to blocklist and block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// Specifies whether the policy takes effect for all user groups. Valid values:
	//
	// - **0**: no.
	//
	// - **1**: yes.
	//
	// example:
	//
	// 1
	AllGroup *int32 `json:"AllGroup,omitempty" xml:"AllGroup,omitempty"`
	// The ID of the reporting and approval process.
	//
	// example:
	//
	// approval-b567d830cd17674c
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The matched process command line.
	//
	// example:
	//
	// C:\\Windows\\System32\\RuntimeBroker.exe -Embedding
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The description of the policy.
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
	// - **UserGroupAll**: all user groups.
	//
	// - **UserGroupNormal**: specified user groups.
	//
	// - **DeviceGroupNormal**: specified device groups.
	//
	// - **DevTagNormal**: specified device tags.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The list of effective scope targets. The meaning of the targets is determined by MatchMode.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
	// The MD5 match string of the file.
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
	// The Policy Name.
	//
	// > The maximum length is 128 characters.
	//
	// This parameter is required.
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
	// - **all**: all platforms.
	//
	// example:
	//
	// windows
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The type of the policy. Valid values:
	//
	// - **0**: software hardening.
	//
	// - **1**: firewall.
	//
	// - **2**: proactive defense.
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
	// This parameter is required.
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
	// The type of the defense rule. Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The key of the switch configuration for the rule associated with the policy. This key can be used in ListSystemConfigRules and OperateTargetConfig.
	//
	// example:
	//
	// a8d64565-ab8b-41a6-bd3d-14d0fc639285
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The list of main process names of the software for which integrity check is enabled. The elements correspond to the Process returned by GetSwitchInfoListByType.
	SoftWareOnList []*string `json:"SoftWareOnList,omitempty" xml:"SoftWareOnList,omitempty" type:"Repeated"`
	// The status of the policy. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rule source for proactive defense (PolicyType=2). Valid values:
	//
	// - **SYSTEM**: system rules.
	//
	// - **CUSTOM**: custom rules.
	//
	// example:
	//
	// SYSTEM
	StrategySourceType *string `json:"StrategySourceType,omitempty" xml:"StrategySourceType,omitempty"`
	// Specifies whether reporting and approval are supported. Valid values:
	//
	// - **0**: not supported.
	//
	// - **1**: supported.
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

func (s CreateClientRuleStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientRuleStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateClientRuleStrategyRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *CreateClientRuleStrategyRequest) GetAllGroup() *int32 {
	return s.AllGroup
}

func (s *CreateClientRuleStrategyRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *CreateClientRuleStrategyRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateClientRuleStrategyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClientRuleStrategyRequest) GetExcludeUserNameList() []*string {
	return s.ExcludeUserNameList
}

func (s *CreateClientRuleStrategyRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateClientRuleStrategyRequest) GetInSwitch() *bool {
	return s.InSwitch
}

func (s *CreateClientRuleStrategyRequest) GetIncludeGroupIdList() []*string {
	return s.IncludeGroupIdList
}

func (s *CreateClientRuleStrategyRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateClientRuleStrategyRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *CreateClientRuleStrategyRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *CreateClientRuleStrategyRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateClientRuleStrategyRequest) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *CreateClientRuleStrategyRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *CreateClientRuleStrategyRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *CreateClientRuleStrategyRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *CreateClientRuleStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateClientRuleStrategyRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *CreateClientRuleStrategyRequest) GetOutSwitch() *bool {
	return s.OutSwitch
}

func (s *CreateClientRuleStrategyRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *CreateClientRuleStrategyRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *CreateClientRuleStrategyRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateClientRuleStrategyRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *CreateClientRuleStrategyRequest) GetPort() *string {
	return s.Port
}

func (s *CreateClientRuleStrategyRequest) GetPortOnList() []*string {
	return s.PortOnList
}

func (s *CreateClientRuleStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateClientRuleStrategyRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *CreateClientRuleStrategyRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *CreateClientRuleStrategyRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *CreateClientRuleStrategyRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *CreateClientRuleStrategyRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *CreateClientRuleStrategyRequest) GetSoftWareOnList() []*string {
	return s.SoftWareOnList
}

func (s *CreateClientRuleStrategyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateClientRuleStrategyRequest) GetStrategySourceType() *string {
	return s.StrategySourceType
}

func (s *CreateClientRuleStrategyRequest) GetSupportsApproval() *int32 {
	return s.SupportsApproval
}

func (s *CreateClientRuleStrategyRequest) GetUsContent() *string {
	return s.UsContent
}

func (s *CreateClientRuleStrategyRequest) GetUsTitle() *string {
	return s.UsTitle
}

func (s *CreateClientRuleStrategyRequest) GetZhContent() *string {
	return s.ZhContent
}

func (s *CreateClientRuleStrategyRequest) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *CreateClientRuleStrategyRequest) SetActionType(v int32) *CreateClientRuleStrategyRequest {
	s.ActionType = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetAllGroup(v int32) *CreateClientRuleStrategyRequest {
	s.AllGroup = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetApprovalId(v string) *CreateClientRuleStrategyRequest {
	s.ApprovalId = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetCmdline(v string) *CreateClientRuleStrategyRequest {
	s.Cmdline = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetDescription(v string) *CreateClientRuleStrategyRequest {
	s.Description = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetExcludeUserNameList(v []*string) *CreateClientRuleStrategyRequest {
	s.ExcludeUserNameList = v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetFilePath(v string) *CreateClientRuleStrategyRequest {
	s.FilePath = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetInSwitch(v bool) *CreateClientRuleStrategyRequest {
	s.InSwitch = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetIncludeGroupIdList(v []*string) *CreateClientRuleStrategyRequest {
	s.IncludeGroupIdList = v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetIp(v string) *CreateClientRuleStrategyRequest {
	s.Ip = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMainButtonTextCh(v string) *CreateClientRuleStrategyRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMainButtonTextEn(v string) *CreateClientRuleStrategyRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMatchMode(v string) *CreateClientRuleStrategyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMatchTargetIds(v []*string) *CreateClientRuleStrategyRequest {
	s.MatchTargetIds = v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMd5List(v string) *CreateClientRuleStrategyRequest {
	s.Md5List = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMinorButtonTextCh(v string) *CreateClientRuleStrategyRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetMinorButtonTextEn(v string) *CreateClientRuleStrategyRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetName(v string) *CreateClientRuleStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetNewFilePath(v string) *CreateClientRuleStrategyRequest {
	s.NewFilePath = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetOutSwitch(v bool) *CreateClientRuleStrategyRequest {
	s.OutSwitch = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetParentCmdline(v string) *CreateClientRuleStrategyRequest {
	s.ParentCmdline = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetParentProcPath(v string) *CreateClientRuleStrategyRequest {
	s.ParentProcPath = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetPlatform(v string) *CreateClientRuleStrategyRequest {
	s.Platform = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetPolicyType(v int32) *CreateClientRuleStrategyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetPort(v string) *CreateClientRuleStrategyRequest {
	s.Port = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetPortOnList(v []*string) *CreateClientRuleStrategyRequest {
	s.PortOnList = v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetPriority(v int32) *CreateClientRuleStrategyRequest {
	s.Priority = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetProcPath(v string) *CreateClientRuleStrategyRequest {
	s.ProcPath = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetRegistryContent(v string) *CreateClientRuleStrategyRequest {
	s.RegistryContent = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetRegistryKey(v string) *CreateClientRuleStrategyRequest {
	s.RegistryKey = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetRuleType(v int32) *CreateClientRuleStrategyRequest {
	s.RuleType = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetSelectionKey(v string) *CreateClientRuleStrategyRequest {
	s.SelectionKey = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetSoftWareOnList(v []*string) *CreateClientRuleStrategyRequest {
	s.SoftWareOnList = v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetStatus(v int32) *CreateClientRuleStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetStrategySourceType(v string) *CreateClientRuleStrategyRequest {
	s.StrategySourceType = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetSupportsApproval(v int32) *CreateClientRuleStrategyRequest {
	s.SupportsApproval = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetUsContent(v string) *CreateClientRuleStrategyRequest {
	s.UsContent = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetUsTitle(v string) *CreateClientRuleStrategyRequest {
	s.UsTitle = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetZhContent(v string) *CreateClientRuleStrategyRequest {
	s.ZhContent = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) SetZhTitle(v string) *CreateClientRuleStrategyRequest {
	s.ZhTitle = &v
	return s
}

func (s *CreateClientRuleStrategyRequest) Validate() error {
	return dara.Validate(s)
}
