// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientRuleStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetActionType() *int32
	SetAllGroup(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetAllGroup() *int32
	SetApprovalId(v string) *UpdateClientRuleStrategyShrinkRequest
	GetApprovalId() *string
	SetCmdline(v string) *UpdateClientRuleStrategyShrinkRequest
	GetCmdline() *string
	SetDescription(v string) *UpdateClientRuleStrategyShrinkRequest
	GetDescription() *string
	SetExcludeUserNameListShrink(v string) *UpdateClientRuleStrategyShrinkRequest
	GetExcludeUserNameListShrink() *string
	SetFilePath(v string) *UpdateClientRuleStrategyShrinkRequest
	GetFilePath() *string
	SetId(v int64) *UpdateClientRuleStrategyShrinkRequest
	GetId() *int64
	SetInSwitch(v bool) *UpdateClientRuleStrategyShrinkRequest
	GetInSwitch() *bool
	SetIncludeGroupIdListShrink(v string) *UpdateClientRuleStrategyShrinkRequest
	GetIncludeGroupIdListShrink() *string
	SetIp(v string) *UpdateClientRuleStrategyShrinkRequest
	GetIp() *string
	SetMainButtonTextCh(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMainButtonTextCh() *string
	SetMainButtonTextEn(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMainButtonTextEn() *string
	SetMatchMode(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *UpdateClientRuleStrategyShrinkRequest
	GetMatchTargetIds() []*string
	SetMd5List(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMd5List() *string
	SetMinorButtonTextCh(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMinorButtonTextCh() *string
	SetMinorButtonTextEn(v string) *UpdateClientRuleStrategyShrinkRequest
	GetMinorButtonTextEn() *string
	SetName(v string) *UpdateClientRuleStrategyShrinkRequest
	GetName() *string
	SetNewFilePath(v string) *UpdateClientRuleStrategyShrinkRequest
	GetNewFilePath() *string
	SetOutSwitch(v bool) *UpdateClientRuleStrategyShrinkRequest
	GetOutSwitch() *bool
	SetParentCmdline(v string) *UpdateClientRuleStrategyShrinkRequest
	GetParentCmdline() *string
	SetParentProcPath(v string) *UpdateClientRuleStrategyShrinkRequest
	GetParentProcPath() *string
	SetPlatform(v string) *UpdateClientRuleStrategyShrinkRequest
	GetPlatform() *string
	SetPolicyType(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetPolicyType() *int32
	SetPort(v string) *UpdateClientRuleStrategyShrinkRequest
	GetPort() *string
	SetPortOnListShrink(v string) *UpdateClientRuleStrategyShrinkRequest
	GetPortOnListShrink() *string
	SetPriority(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetPriority() *int32
	SetProcPath(v string) *UpdateClientRuleStrategyShrinkRequest
	GetProcPath() *string
	SetRegistryContent(v string) *UpdateClientRuleStrategyShrinkRequest
	GetRegistryContent() *string
	SetRegistryKey(v string) *UpdateClientRuleStrategyShrinkRequest
	GetRegistryKey() *string
	SetRuleType(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetRuleType() *int32
	SetSoftWareOnListShrink(v string) *UpdateClientRuleStrategyShrinkRequest
	GetSoftWareOnListShrink() *string
	SetStatus(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetStatus() *int32
	SetStrategySourceType(v string) *UpdateClientRuleStrategyShrinkRequest
	GetStrategySourceType() *string
	SetSupportsApproval(v int32) *UpdateClientRuleStrategyShrinkRequest
	GetSupportsApproval() *int32
	SetUsContent(v string) *UpdateClientRuleStrategyShrinkRequest
	GetUsContent() *string
	SetUsTitle(v string) *UpdateClientRuleStrategyShrinkRequest
	GetUsTitle() *string
	SetZhContent(v string) *UpdateClientRuleStrategyShrinkRequest
	GetZhContent() *string
	SetZhTitle(v string) *UpdateClientRuleStrategyShrinkRequest
	GetZhTitle() *string
}

type UpdateClientRuleStrategyShrinkRequest struct {
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
	ExcludeUserNameListShrink *string `json:"ExcludeUserNameList,omitempty" xml:"ExcludeUserNameList,omitempty"`
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
	PortOnListShrink *string `json:"PortOnList,omitempty" xml:"PortOnList,omitempty"`
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
	SoftWareOnListShrink *string `json:"SoftWareOnList,omitempty" xml:"SoftWareOnList,omitempty"`
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

func (s UpdateClientRuleStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientRuleStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetAllGroup() *int32 {
	return s.AllGroup
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetExcludeUserNameListShrink() *string {
	return s.ExcludeUserNameListShrink
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetInSwitch() *bool {
	return s.InSwitch
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetIncludeGroupIdListShrink() *string {
	return s.IncludeGroupIdListShrink
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetIp() *string {
	return s.Ip
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMainButtonTextCh() *string {
	return s.MainButtonTextCh
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMainButtonTextEn() *string {
	return s.MainButtonTextEn
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMd5List() *string {
	return s.Md5List
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMinorButtonTextCh() *string {
	return s.MinorButtonTextCh
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetMinorButtonTextEn() *string {
	return s.MinorButtonTextEn
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetNewFilePath() *string {
	return s.NewFilePath
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetOutSwitch() *bool {
	return s.OutSwitch
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetParentProcPath() *string {
	return s.ParentProcPath
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetPortOnListShrink() *string {
	return s.PortOnListShrink
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetProcPath() *string {
	return s.ProcPath
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetRegistryContent() *string {
	return s.RegistryContent
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetRegistryKey() *string {
	return s.RegistryKey
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetSoftWareOnListShrink() *string {
	return s.SoftWareOnListShrink
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetStrategySourceType() *string {
	return s.StrategySourceType
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetSupportsApproval() *int32 {
	return s.SupportsApproval
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetUsContent() *string {
	return s.UsContent
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetUsTitle() *string {
	return s.UsTitle
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetZhContent() *string {
	return s.ZhContent
}

func (s *UpdateClientRuleStrategyShrinkRequest) GetZhTitle() *string {
	return s.ZhTitle
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetActionType(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetAllGroup(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.AllGroup = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetApprovalId(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ApprovalId = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetCmdline(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Cmdline = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetDescription(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetExcludeUserNameListShrink(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ExcludeUserNameListShrink = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetFilePath(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetId(v int64) *UpdateClientRuleStrategyShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetInSwitch(v bool) *UpdateClientRuleStrategyShrinkRequest {
	s.InSwitch = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetIncludeGroupIdListShrink(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.IncludeGroupIdListShrink = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetIp(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Ip = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMainButtonTextCh(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.MainButtonTextCh = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMainButtonTextEn(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.MainButtonTextEn = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMatchMode(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMatchTargetIds(v []*string) *UpdateClientRuleStrategyShrinkRequest {
	s.MatchTargetIds = v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMd5List(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Md5List = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMinorButtonTextCh(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.MinorButtonTextCh = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetMinorButtonTextEn(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.MinorButtonTextEn = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetName(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetNewFilePath(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.NewFilePath = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetOutSwitch(v bool) *UpdateClientRuleStrategyShrinkRequest {
	s.OutSwitch = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetParentCmdline(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ParentCmdline = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetParentProcPath(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ParentProcPath = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetPlatform(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Platform = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetPolicyType(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetPort(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.Port = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetPortOnListShrink(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.PortOnListShrink = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetPriority(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetProcPath(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ProcPath = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetRegistryContent(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.RegistryContent = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetRegistryKey(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.RegistryKey = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetRuleType(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetSoftWareOnListShrink(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.SoftWareOnListShrink = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetStatus(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetStrategySourceType(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.StrategySourceType = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetSupportsApproval(v int32) *UpdateClientRuleStrategyShrinkRequest {
	s.SupportsApproval = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetUsContent(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.UsContent = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetUsTitle(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.UsTitle = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetZhContent(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ZhContent = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) SetZhTitle(v string) *UpdateClientRuleStrategyShrinkRequest {
	s.ZhTitle = &v
	return s
}

func (s *UpdateClientRuleStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
