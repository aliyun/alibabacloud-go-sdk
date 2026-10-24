// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientRuleStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *CreateClientRuleStrategyShrinkRequest
	GetActionType() *int32
	SetAllGroup(v int32) *CreateClientRuleStrategyShrinkRequest
	GetAllGroup() *int32
	SetApprovalId(v string) *CreateClientRuleStrategyShrinkRequest
	GetApprovalId() *string
	SetCmdline(v string) *CreateClientRuleStrategyShrinkRequest
	GetCmdline() *string
	SetDescription(v string) *CreateClientRuleStrategyShrinkRequest
	GetDescription() *string
	SetExcludeUserNameListShrink(v string) *CreateClientRuleStrategyShrinkRequest
	GetExcludeUserNameListShrink() *string
	SetFilePath(v string) *CreateClientRuleStrategyShrinkRequest
	GetFilePath() *string
	SetInSwitch(v bool) *CreateClientRuleStrategyShrinkRequest
	GetInSwitch() *bool
	SetIncludeGroupIdListShrink(v string) *CreateClientRuleStrategyShrinkRequest
	GetIncludeGroupIdListShrink() *string
	SetIp(v string) *CreateClientRuleStrategyShrinkRequest
	GetIp() *string
	SetMainButtonTextCh(v string) *CreateClientRuleStrategyShrinkRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *CreateClientRuleStrategyShrinkRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *CreateClientRuleStrategyShrinkRequest
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *CreateClientRuleStrategyShrinkRequest
	GetMatchTargetIds() []*string
	SetMd5List(v string) *CreateClientRuleStrategyShrinkRequest
	GetMd5List() *string
	SetMinorButtonTextCh(v string) *CreateClientRuleStrategyShrinkRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *CreateClientRuleStrategyShrinkRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *CreateClientRuleStrategyShrinkRequest
	GetName() *string
	SetNewFilePath(v string) *CreateClientRuleStrategyShrinkRequest
	GetNewFilePath() *string
	SetOutSwitch(v bool) *CreateClientRuleStrategyShrinkRequest
	GetOutSwitch() *bool
	SetParentCmdline(v string) *CreateClientRuleStrategyShrinkRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *CreateClientRuleStrategyShrinkRequest
	GetParentProcPath() *string
	SetPlatform(v string) *CreateClientRuleStrategyShrinkRequest
	GetPlatform() *string
	SetPolicyType(v int32) *CreateClientRuleStrategyShrinkRequest
	GetPolicyType() *int32
	SetPort(v string) *CreateClientRuleStrategyShrinkRequest
	GetPort() *string
	SetPortOnListShrink(v string) *CreateClientRuleStrategyShrinkRequest
	GetPortOnListShrink() *string
	SetPriority(v int32) *CreateClientRuleStrategyShrinkRequest
	GetPriority() *int32
	SetProcPath(v string) *CreateClientRuleStrategyShrinkRequest
	GetProcPath() *string
	SetRegistryContent(v string) *CreateClientRuleStrategyShrinkRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *CreateClientRuleStrategyShrinkRequest
	GetRegistryKey() *string
	SetRuleType(v int32) *CreateClientRuleStrategyShrinkRequest
	GetRuleType() *int32
	SetSelectionKey(v string) *CreateClientRuleStrategyShrinkRequest
	GetSelectionKey() *string
	SetSoftWareOnListShrink(v string) *CreateClientRuleStrategyShrinkRequest
	GetSoftWareOnListShrink() *string
	SetStatus(v int32) *CreateClientRuleStrategyShrinkRequest
	GetStatus() *int32
	SetStrategySourceType(v string) *CreateClientRuleStrategyShrinkRequest
	GetStrategySourceType() *string
	SetSupportsApproval(v int32) *CreateClientRuleStrategyShrinkRequest
	GetSupportsApproval() *int32
	SetUsContent(v string) *CreateClientRuleStrategyShrinkRequest
	GetUsContent() *string
	SetUsTitle(v string) *CreateClientRuleStrategyShrinkRequest
	GetUsTitle() *string
	SetZhContent(v string) *CreateClientRuleStrategyShrinkRequest
	GetZhContent() *string
	SetZhTitle(v string) *CreateClientRuleStrategyShrinkRequest
	GetZhTitle() *string
}

type CreateClientRuleStrategyShrinkRequest struct {
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
	ExcludeUserNameListShrink *string `json:"ExcludeUserNameList,omitempty" xml:"ExcludeUserNameList,omitempty"`
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
	IncludeGroupIdListShrink *string `json:"IncludeGroupIdList,omitempty" xml:"IncludeGroupIdList,omitempty"`
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
	PortOnListShrink *string `json:"PortOnList,omitempty" xml:"PortOnList,omitempty"`
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
	SoftWareOnListShrink *string `json:"SoftWareOnList,omitempty" xml:"SoftWareOnList,omitempty"`
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

func (s CreateClientRuleStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientRuleStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClientRuleStrategyShrinkRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *CreateClientRuleStrategyShrinkRequest) GetAllGroup() *int32 {
	return s.AllGroup
}

func (s *CreateClientRuleStrategyShrinkRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *CreateClientRuleStrategyShrinkRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *CreateClientRuleStrategyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateClientRuleStrategyShrinkRequest) GetExcludeUserNameListShrink() *string {
	return s.ExcludeUserNameListShrink
}

func (s *CreateClientRuleStrategyShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateClientRuleStrategyShrinkRequest) GetInSwitch() *bool {
	return s.InSwitch
}

func (s *CreateClientRuleStrategyShrinkRequest) GetIncludeGroupIdListShrink() *string {
	return s.IncludeGroupIdListShrink
}

func (s *CreateClientRuleStrategyShrinkRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *CreateClientRuleStrategyShrinkRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *CreateClientRuleStrategyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateClientRuleStrategyShrinkRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *CreateClientRuleStrategyShrinkRequest) GetOutSwitch() *bool {
	return s.OutSwitch
}

func (s *CreateClientRuleStrategyShrinkRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *CreateClientRuleStrategyShrinkRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *CreateClientRuleStrategyShrinkRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateClientRuleStrategyShrinkRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *CreateClientRuleStrategyShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateClientRuleStrategyShrinkRequest) GetPortOnListShrink() *string {
	return s.PortOnListShrink
}

func (s *CreateClientRuleStrategyShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateClientRuleStrategyShrinkRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *CreateClientRuleStrategyShrinkRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *CreateClientRuleStrategyShrinkRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *CreateClientRuleStrategyShrinkRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *CreateClientRuleStrategyShrinkRequest) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *CreateClientRuleStrategyShrinkRequest) GetSoftWareOnListShrink() *string {
	return s.SoftWareOnListShrink
}

func (s *CreateClientRuleStrategyShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateClientRuleStrategyShrinkRequest) GetStrategySourceType() *string {
	return s.StrategySourceType
}

func (s *CreateClientRuleStrategyShrinkRequest) GetSupportsApproval() *int32 {
	return s.SupportsApproval
}

func (s *CreateClientRuleStrategyShrinkRequest) GetUsContent() *string {
	return s.UsContent
}

func (s *CreateClientRuleStrategyShrinkRequest) GetUsTitle() *string {
	return s.UsTitle
}

func (s *CreateClientRuleStrategyShrinkRequest) GetZhContent() *string {
	return s.ZhContent
}

func (s *CreateClientRuleStrategyShrinkRequest) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *CreateClientRuleStrategyShrinkRequest) SetActionType(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.ActionType = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetAllGroup(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.AllGroup = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetApprovalId(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ApprovalId = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetCmdline(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Cmdline = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetDescription(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetExcludeUserNameListShrink(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ExcludeUserNameListShrink = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetFilePath(v string) *CreateClientRuleStrategyShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetInSwitch(v bool) *CreateClientRuleStrategyShrinkRequest {
	s.InSwitch = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetIncludeGroupIdListShrink(v string) *CreateClientRuleStrategyShrinkRequest {
	s.IncludeGroupIdListShrink = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetIp(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Ip = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMainButtonTextCh(v string) *CreateClientRuleStrategyShrinkRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMainButtonTextEn(v string) *CreateClientRuleStrategyShrinkRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMatchMode(v string) *CreateClientRuleStrategyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMatchTargetIds(v []*string) *CreateClientRuleStrategyShrinkRequest {
	s.MatchTargetIds = v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMd5List(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Md5List = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMinorButtonTextCh(v string) *CreateClientRuleStrategyShrinkRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetMinorButtonTextEn(v string) *CreateClientRuleStrategyShrinkRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetName(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetNewFilePath(v string) *CreateClientRuleStrategyShrinkRequest {
	s.NewFilePath = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetOutSwitch(v bool) *CreateClientRuleStrategyShrinkRequest {
	s.OutSwitch = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetParentCmdline(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ParentCmdline = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetParentProcPath(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ParentProcPath = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetPlatform(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Platform = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetPolicyType(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetPort(v string) *CreateClientRuleStrategyShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetPortOnListShrink(v string) *CreateClientRuleStrategyShrinkRequest {
	s.PortOnListShrink = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetPriority(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetProcPath(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ProcPath = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetRegistryContent(v string) *CreateClientRuleStrategyShrinkRequest {
	s.RegistryContent = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetRegistryKey(v string) *CreateClientRuleStrategyShrinkRequest {
	s.RegistryKey = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetRuleType(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetSelectionKey(v string) *CreateClientRuleStrategyShrinkRequest {
	s.SelectionKey = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetSoftWareOnListShrink(v string) *CreateClientRuleStrategyShrinkRequest {
	s.SoftWareOnListShrink = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetStatus(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetStrategySourceType(v string) *CreateClientRuleStrategyShrinkRequest {
	s.StrategySourceType = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetSupportsApproval(v int32) *CreateClientRuleStrategyShrinkRequest {
	s.SupportsApproval = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetUsContent(v string) *CreateClientRuleStrategyShrinkRequest {
	s.UsContent = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetUsTitle(v string) *CreateClientRuleStrategyShrinkRequest {
	s.UsTitle = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetZhContent(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ZhContent = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) SetZhTitle(v string) *CreateClientRuleStrategyShrinkRequest {
	s.ZhTitle = &v
	return s
}

func (s *CreateClientRuleStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
