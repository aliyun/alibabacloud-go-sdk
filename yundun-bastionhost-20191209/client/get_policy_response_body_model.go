// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody
	GetPolicy() *GetPolicyResponseBodyPolicy
	SetRequestId(v string) *GetPolicyResponseBody
	GetRequestId() *string
}

type GetPolicyResponseBody struct {
	// The details of the control policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0D29F2C0-8B4B-5861-9474-F3F23D25594B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) GetPolicy() *GetPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicy struct {
	// The details of the logon period restrictions.
	AccessTimeRangeConfig *GetPolicyResponseBodyPolicyAccessTimeRangeConfig `json:"AccessTimeRangeConfig,omitempty" xml:"AccessTimeRangeConfig,omitempty" type:"Struct"`
	// The O\\&M approval setting.
	ApprovalConfig *GetPolicyResponseBodyPolicyApprovalConfig `json:"ApprovalConfig,omitempty" xml:"ApprovalConfig,omitempty" type:"Struct"`
	// The details of the command policy.
	CommandConfig *GetPolicyResponseBodyPolicyCommandConfig `json:"CommandConfig,omitempty" xml:"CommandConfig,omitempty" type:"Struct"`
	// The description of the control policy.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The access control settings on source IP addresses.
	IPAclConfig *GetPolicyResponseBodyPolicyIPAclConfig `json:"IPAclConfig,omitempty" xml:"IPAclConfig,omitempty" type:"Struct"`
	// The ID of the control policy.
	//
	// example:
	//
	// 3
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the control policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The priority of the control policy. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The details of protocol control.
	ProtocolConfig *GetPolicyResponseBodyPolicyProtocolConfig `json:"ProtocolConfig,omitempty" xml:"ProtocolConfig,omitempty" type:"Struct"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) GetAccessTimeRangeConfig() *GetPolicyResponseBodyPolicyAccessTimeRangeConfig {
	return s.AccessTimeRangeConfig
}

func (s *GetPolicyResponseBodyPolicy) GetApprovalConfig() *GetPolicyResponseBodyPolicyApprovalConfig {
	return s.ApprovalConfig
}

func (s *GetPolicyResponseBodyPolicy) GetCommandConfig() *GetPolicyResponseBodyPolicyCommandConfig {
	return s.CommandConfig
}

func (s *GetPolicyResponseBodyPolicy) GetComment() *string {
	return s.Comment
}

func (s *GetPolicyResponseBodyPolicy) GetIPAclConfig() *GetPolicyResponseBodyPolicyIPAclConfig {
	return s.IPAclConfig
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyResponseBodyPolicy) GetPriority() *int64 {
	return s.Priority
}

func (s *GetPolicyResponseBodyPolicy) GetProtocolConfig() *GetPolicyResponseBodyPolicyProtocolConfig {
	return s.ProtocolConfig
}

func (s *GetPolicyResponseBodyPolicy) SetAccessTimeRangeConfig(v *GetPolicyResponseBodyPolicyAccessTimeRangeConfig) *GetPolicyResponseBodyPolicy {
	s.AccessTimeRangeConfig = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetApprovalConfig(v *GetPolicyResponseBodyPolicyApprovalConfig) *GetPolicyResponseBodyPolicy {
	s.ApprovalConfig = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetCommandConfig(v *GetPolicyResponseBodyPolicyCommandConfig) *GetPolicyResponseBodyPolicy {
	s.CommandConfig = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetComment(v string) *GetPolicyResponseBodyPolicy {
	s.Comment = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetIPAclConfig(v *GetPolicyResponseBodyPolicyIPAclConfig) *GetPolicyResponseBodyPolicy {
	s.IPAclConfig = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyId(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPriority(v int64) *GetPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetProtocolConfig(v *GetPolicyResponseBodyPolicyProtocolConfig) *GetPolicyResponseBodyPolicy {
	s.ProtocolConfig = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyAccessTimeRangeConfig struct {
	// The details of the periods during which logons are allowed.
	EffectiveTime []*GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty" type:"Repeated"`
}

func (s GetPolicyResponseBodyPolicyAccessTimeRangeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyAccessTimeRangeConfig) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfig) GetEffectiveTime() []*GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime {
	return s.EffectiveTime
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfig) SetEffectiveTime(v []*GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) *GetPolicyResponseBodyPolicyAccessTimeRangeConfig {
	s.EffectiveTime = v
	return s
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfig) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime struct {
	// The days of a week on which logons are allowed.
	Days []*string `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The time periods during which logons are allowed.
	Hours []*string `json:"Hours,omitempty" xml:"Hours,omitempty" type:"Repeated"`
}

func (s GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) GetDays() []*string {
	return s.Days
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) GetHours() []*string {
	return s.Hours
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) SetDays(v []*string) *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime {
	s.Days = v
	return s
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) SetHours(v []*string) *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime {
	s.Hours = v
	return s
}

func (s *GetPolicyResponseBodyPolicyAccessTimeRangeConfigEffectiveTime) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyApprovalConfig struct {
	// Indicates whether O\\&M approval is enabled in the control policy. Valid values:
	//
	// 	- **On**: O\\&M approval is enabled.
	//
	// 	- **Off**: O\\&M approval is disabled.
	//
	// example:
	//
	// Off
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s GetPolicyResponseBodyPolicyApprovalConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyApprovalConfig) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyApprovalConfig) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *GetPolicyResponseBodyPolicyApprovalConfig) SetSwitchStatus(v string) *GetPolicyResponseBodyPolicyApprovalConfig {
	s.SwitchStatus = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyApprovalConfig) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyCommandConfig struct {
	// The details of the command approval settings.
	Approval *GetPolicyResponseBodyPolicyCommandConfigApproval `json:"Approval,omitempty" xml:"Approval,omitempty" type:"Struct"`
	// The details of the command control setting.
	Deny *GetPolicyResponseBodyPolicyCommandConfigDeny `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Struct"`
}

func (s GetPolicyResponseBodyPolicyCommandConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyCommandConfig) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyCommandConfig) GetApproval() *GetPolicyResponseBodyPolicyCommandConfigApproval {
	return s.Approval
}

func (s *GetPolicyResponseBodyPolicyCommandConfig) GetDeny() *GetPolicyResponseBodyPolicyCommandConfigDeny {
	return s.Deny
}

func (s *GetPolicyResponseBodyPolicyCommandConfig) SetApproval(v *GetPolicyResponseBodyPolicyCommandConfigApproval) *GetPolicyResponseBodyPolicyCommandConfig {
	s.Approval = v
	return s
}

func (s *GetPolicyResponseBodyPolicyCommandConfig) SetDeny(v *GetPolicyResponseBodyPolicyCommandConfigDeny) *GetPolicyResponseBodyPolicyCommandConfig {
	s.Deny = v
	return s
}

func (s *GetPolicyResponseBodyPolicyCommandConfig) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyCommandConfigApproval struct {
	// An array of commands that can be run only after approval.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s GetPolicyResponseBodyPolicyCommandConfigApproval) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyCommandConfigApproval) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyCommandConfigApproval) GetCommands() []*string {
	return s.Commands
}

func (s *GetPolicyResponseBodyPolicyCommandConfigApproval) SetCommands(v []*string) *GetPolicyResponseBodyPolicyCommandConfigApproval {
	s.Commands = v
	return s
}

func (s *GetPolicyResponseBodyPolicyCommandConfigApproval) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyCommandConfigDeny struct {
	// The type of command control. Valid values:
	//
	// 	- white: whitelist mode.
	//
	// 	- black: blacklist mode.
	//
	// example:
	//
	// black
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// An array of controlled commands.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s GetPolicyResponseBodyPolicyCommandConfigDeny) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyCommandConfigDeny) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyCommandConfigDeny) GetAclType() *string {
	return s.AclType
}

func (s *GetPolicyResponseBodyPolicyCommandConfigDeny) GetCommands() []*string {
	return s.Commands
}

func (s *GetPolicyResponseBodyPolicyCommandConfigDeny) SetAclType(v string) *GetPolicyResponseBodyPolicyCommandConfigDeny {
	s.AclType = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyCommandConfigDeny) SetCommands(v []*string) *GetPolicyResponseBodyPolicyCommandConfigDeny {
	s.Commands = v
	return s
}

func (s *GetPolicyResponseBodyPolicyCommandConfigDeny) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyIPAclConfig struct {
	// The mode of access control on source IP addresses. Valid values:
	//
	// 	- white: whitelist mode.
	//
	// 	- black: blacklist mode.
	//
	// example:
	//
	// black
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The IP addresses from which logons are not allowed.
	IPs []*string `json:"IPs,omitempty" xml:"IPs,omitempty" type:"Repeated"`
}

func (s GetPolicyResponseBodyPolicyIPAclConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyIPAclConfig) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyIPAclConfig) GetAclType() *string {
	return s.AclType
}

func (s *GetPolicyResponseBodyPolicyIPAclConfig) GetIPs() []*string {
	return s.IPs
}

func (s *GetPolicyResponseBodyPolicyIPAclConfig) SetAclType(v string) *GetPolicyResponseBodyPolicyIPAclConfig {
	s.AclType = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyIPAclConfig) SetIPs(v []*string) *GetPolicyResponseBodyPolicyIPAclConfig {
	s.IPs = v
	return s
}

func (s *GetPolicyResponseBodyPolicyIPAclConfig) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyProtocolConfig struct {
	// The configuration details of Remote Desktop Protocol (RDP) options.
	RDP *GetPolicyResponseBodyPolicyProtocolConfigRDP `json:"RDP,omitempty" xml:"RDP,omitempty" type:"Struct"`
	// The configuration details of SSH and SSH File Transfer Protocol (SFTP) options.
	SSH *GetPolicyResponseBodyPolicyProtocolConfigSSH `json:"SSH,omitempty" xml:"SSH,omitempty" type:"Struct"`
}

func (s GetPolicyResponseBodyPolicyProtocolConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyProtocolConfig) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyProtocolConfig) GetRDP() *GetPolicyResponseBodyPolicyProtocolConfigRDP {
	return s.RDP
}

func (s *GetPolicyResponseBodyPolicyProtocolConfig) GetSSH() *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	return s.SSH
}

func (s *GetPolicyResponseBodyPolicyProtocolConfig) SetRDP(v *GetPolicyResponseBodyPolicyProtocolConfigRDP) *GetPolicyResponseBodyPolicyProtocolConfig {
	s.RDP = v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfig) SetSSH(v *GetPolicyResponseBodyPolicyProtocolConfigSSH) *GetPolicyResponseBodyPolicyProtocolConfig {
	s.SSH = v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfig) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyProtocolConfigRDP struct {
	// Indicates whether downloading from the clipboard is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	ClipboardDownload *string `json:"ClipboardDownload,omitempty" xml:"ClipboardDownload,omitempty"`
	// Indicates whether file uploading from the clipboard is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	ClipboardUpload *string `json:"ClipboardUpload,omitempty" xml:"ClipboardUpload,omitempty"`
	// Indicates whether driver mapping is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	DiskRedirection *string `json:"DiskRedirection,omitempty" xml:"DiskRedirection,omitempty"`
	// Indicates whether keyboard recording is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	RecordKeyboard *string `json:"RecordKeyboard,omitempty" xml:"RecordKeyboard,omitempty"`
}

func (s GetPolicyResponseBodyPolicyProtocolConfigRDP) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyProtocolConfigRDP) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) GetClipboardDownload() *string {
	return s.ClipboardDownload
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) GetClipboardUpload() *string {
	return s.ClipboardUpload
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) GetDiskRedirection() *string {
	return s.DiskRedirection
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) GetRecordKeyboard() *string {
	return s.RecordKeyboard
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) SetClipboardDownload(v string) *GetPolicyResponseBodyPolicyProtocolConfigRDP {
	s.ClipboardDownload = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) SetClipboardUpload(v string) *GetPolicyResponseBodyPolicyProtocolConfigRDP {
	s.ClipboardUpload = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) SetDiskRedirection(v string) *GetPolicyResponseBodyPolicyProtocolConfigRDP {
	s.DiskRedirection = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) SetRecordKeyboard(v string) *GetPolicyResponseBodyPolicyProtocolConfigRDP {
	s.RecordKeyboard = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigRDP) Validate() error {
	return dara.Validate(s)
}

type GetPolicyResponseBodyPolicyProtocolConfigSSH struct {
	// Indicates whether remote command execution is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	ExecCommand *string `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty"`
	// Indicates whether the SFTP channel option is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPChannel *string `json:"SFTPChannel,omitempty" xml:"SFTPChannel,omitempty"`
	// Indicates whether file downloading is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPDownloadFile *string `json:"SFTPDownloadFile,omitempty" xml:"SFTPDownloadFile,omitempty"`
	// Indicates whether folder creation is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPMkdir *string `json:"SFTPMkdir,omitempty" xml:"SFTPMkdir,omitempty"`
	// Indicates whether file deletion is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPRemoveFile *string `json:"SFTPRemoveFile,omitempty" xml:"SFTPRemoveFile,omitempty"`
	// Indicates whether file renaming is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPRenameFile *string `json:"SFTPRenameFile,omitempty" xml:"SFTPRenameFile,omitempty"`
	// Indicates whether folder deletion is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPRmdir *string `json:"SFTPRmdir,omitempty" xml:"SFTPRmdir,omitempty"`
	// Indicates whether file uploading is enabled in SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SFTPUploadFile *string `json:"SFTPUploadFile,omitempty" xml:"SFTPUploadFile,omitempty"`
	// Indicates whether the SSH channel option is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	SSHChannel *string `json:"SSHChannel,omitempty" xml:"SSHChannel,omitempty"`
	// Indicates whether X11 forwarding is enabled. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// example:
	//
	// Enable
	X11Forwarding *string `json:"X11Forwarding,omitempty" xml:"X11Forwarding,omitempty"`
}

func (s GetPolicyResponseBodyPolicyProtocolConfigSSH) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicyProtocolConfigSSH) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetExecCommand() *string {
	return s.ExecCommand
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPChannel() *string {
	return s.SFTPChannel
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPDownloadFile() *string {
	return s.SFTPDownloadFile
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPMkdir() *string {
	return s.SFTPMkdir
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPRemoveFile() *string {
	return s.SFTPRemoveFile
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPRenameFile() *string {
	return s.SFTPRenameFile
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPRmdir() *string {
	return s.SFTPRmdir
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSFTPUploadFile() *string {
	return s.SFTPUploadFile
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetSSHChannel() *string {
	return s.SSHChannel
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) GetX11Forwarding() *string {
	return s.X11Forwarding
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetExecCommand(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.ExecCommand = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPChannel(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPChannel = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPDownloadFile(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPDownloadFile = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPMkdir(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPMkdir = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPRemoveFile(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPRemoveFile = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPRenameFile(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPRenameFile = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPRmdir(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPRmdir = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSFTPUploadFile(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SFTPUploadFile = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetSSHChannel(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.SSHChannel = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) SetX11Forwarding(v string) *GetPolicyResponseBodyPolicyProtocolConfigSSH {
	s.X11Forwarding = &v
	return s
}

func (s *GetPolicyResponseBodyPolicyProtocolConfigSSH) Validate() error {
	return dara.Validate(s)
}
