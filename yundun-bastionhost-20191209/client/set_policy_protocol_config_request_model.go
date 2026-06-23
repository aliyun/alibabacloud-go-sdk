// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyProtocolConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPolicyProtocolConfigRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyProtocolConfigRequest
	GetPolicyId() *string
	SetProtocolConfig(v *SetPolicyProtocolConfigRequestProtocolConfig) *SetPolicyProtocolConfigRequest
	GetProtocolConfig() *SetPolicyProtocolConfigRequestProtocolConfig
	SetRegionId(v string) *SetPolicyProtocolConfigRequest
	GetRegionId() *string
}

type SetPolicyProtocolConfigRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > Call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to obtain the policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The protocol control configuration.
	//
	// This parameter is required.
	ProtocolConfig *SetPolicyProtocolConfigRequestProtocolConfig `json:"ProtocolConfig,omitempty" xml:"ProtocolConfig,omitempty" type:"Struct"`
	// The ID of the region where the Bastionhost instance resides.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyProtocolConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyProtocolConfigRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyProtocolConfigRequest) GetProtocolConfig() *SetPolicyProtocolConfigRequestProtocolConfig {
	return s.ProtocolConfig
}

func (s *SetPolicyProtocolConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyProtocolConfigRequest) SetInstanceId(v string) *SetPolicyProtocolConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyProtocolConfigRequest) SetPolicyId(v string) *SetPolicyProtocolConfigRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyProtocolConfigRequest) SetProtocolConfig(v *SetPolicyProtocolConfigRequestProtocolConfig) *SetPolicyProtocolConfigRequest {
	s.ProtocolConfig = v
	return s
}

func (s *SetPolicyProtocolConfigRequest) SetRegionId(v string) *SetPolicyProtocolConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyProtocolConfigRequest) Validate() error {
	if s.ProtocolConfig != nil {
		if err := s.ProtocolConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyProtocolConfigRequestProtocolConfig struct {
	// The RDP options.
	RDP *SetPolicyProtocolConfigRequestProtocolConfigRDP `json:"RDP,omitempty" xml:"RDP,omitempty" type:"Struct"`
	// The SSH and SFTP options.
	SSH *SetPolicyProtocolConfigRequestProtocolConfigSSH `json:"SSH,omitempty" xml:"SSH,omitempty" type:"Struct"`
}

func (s SetPolicyProtocolConfigRequestProtocolConfig) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigRequestProtocolConfig) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigRequestProtocolConfig) GetRDP() *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	return s.RDP
}

func (s *SetPolicyProtocolConfigRequestProtocolConfig) GetSSH() *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	return s.SSH
}

func (s *SetPolicyProtocolConfigRequestProtocolConfig) SetRDP(v *SetPolicyProtocolConfigRequestProtocolConfigRDP) *SetPolicyProtocolConfigRequestProtocolConfig {
	s.RDP = v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfig) SetSSH(v *SetPolicyProtocolConfigRequestProtocolConfigSSH) *SetPolicyProtocolConfigRequestProtocolConfig {
	s.SSH = v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfig) Validate() error {
	if s.RDP != nil {
		if err := s.RDP.Validate(); err != nil {
			return err
		}
	}
	if s.SSH != nil {
		if err := s.SSH.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyProtocolConfigRequestProtocolConfigRDP struct {
	// Specifies whether to allow clipboard content to be downloaded. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	ClipboardDownload *string `json:"ClipboardDownload,omitempty" xml:"ClipboardDownload,omitempty"`
	// Specifies whether to allow clipboard content to be uploaded. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	ClipboardUpload *string `json:"ClipboardUpload,omitempty" xml:"ClipboardUpload,omitempty"`
	// Specifies whether to enable drive and printer mapping. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	DiskRedirection *string `json:"DiskRedirection,omitempty" xml:"DiskRedirection,omitempty"`
	// example:
	//
	// Enable
	DiskRedirectionDownload *string `json:"DiskRedirectionDownload,omitempty" xml:"DiskRedirectionDownload,omitempty"`
	// example:
	//
	// Enable
	DiskRedirectionUpload *string `json:"DiskRedirectionUpload,omitempty" xml:"DiskRedirectionUpload,omitempty"`
	// Specifies whether to record keyboard input. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	RecordKeyboard *string `json:"RecordKeyboard,omitempty" xml:"RecordKeyboard,omitempty"`
}

func (s SetPolicyProtocolConfigRequestProtocolConfigRDP) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigRequestProtocolConfigRDP) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetClipboardDownload() *string {
	return s.ClipboardDownload
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetClipboardUpload() *string {
	return s.ClipboardUpload
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetDiskRedirection() *string {
	return s.DiskRedirection
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetDiskRedirectionDownload() *string {
	return s.DiskRedirectionDownload
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetDiskRedirectionUpload() *string {
	return s.DiskRedirectionUpload
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) GetRecordKeyboard() *string {
	return s.RecordKeyboard
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetClipboardDownload(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.ClipboardDownload = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetClipboardUpload(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.ClipboardUpload = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetDiskRedirection(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.DiskRedirection = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetDiskRedirectionDownload(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.DiskRedirectionDownload = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetDiskRedirectionUpload(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.DiskRedirectionUpload = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetRecordKeyboard(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.RecordKeyboard = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) Validate() error {
	return dara.Validate(s)
}

type SetPolicyProtocolConfigRequestProtocolConfigSSH struct {
	// example:
	//
	// Enable
	AllowDirectTcp *string `json:"AllowDirectTcp,omitempty" xml:"AllowDirectTcp,omitempty"`
	// example:
	//
	// Enable
	AllowTcpForwarding *string `json:"AllowTcpForwarding,omitempty" xml:"AllowTcpForwarding,omitempty"`
	// Specifies whether to allow remote command execution. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	ExecCommand *string `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty"`
	// Specifies whether to enable the SFTP channel. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > 	- The default value is Disable.
	//
	// >
	//
	// > 	- At least one of the SSH channel and the SFTP channel must be enabled.
	//
	// >
	//
	// > 	- If you grant only SFTP permissions to a host account, do not disable the SSH and SFTP channels for that account in the control policy. Otherwise, you cannot use the host account to access the target server through Bastionhost.
	//
	// example:
	//
	// Enable
	SFTPChannel *string `json:"SFTPChannel,omitempty" xml:"SFTPChannel,omitempty"`
	// Specifies whether to allow file downloads over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPDownloadFile *string `json:"SFTPDownloadFile,omitempty" xml:"SFTPDownloadFile,omitempty"`
	// Specifies whether to allow folder creation over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPMkdir *string `json:"SFTPMkdir,omitempty" xml:"SFTPMkdir,omitempty"`
	// Specifies whether to allow file deletions over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPRemoveFile *string `json:"SFTPRemoveFile,omitempty" xml:"SFTPRemoveFile,omitempty"`
	// Specifies whether to allow file renames over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPRenameFile *string `json:"SFTPRenameFile,omitempty" xml:"SFTPRenameFile,omitempty"`
	// Specifies whether to allow folder deletion over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPRmdir *string `json:"SFTPRmdir,omitempty" xml:"SFTPRmdir,omitempty"`
	// Specifies whether to allow file uploads over SFTP. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	SFTPUploadFile *string `json:"SFTPUploadFile,omitempty" xml:"SFTPUploadFile,omitempty"`
	// Specifies whether to enable the SSH channel. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > 	- The default value is Disable.
	//
	// >
	//
	// > 	- At least one of the SSH channel and the SFTP channel must be enabled. If you disable the SSH channel, you cannot use SSH permissions to log on to the asset account. Configure this parameter with caution.
	//
	// >
	//
	// > 	- If you grant only SFTP permissions to a host account, do not disable the SSH and SFTP channels for that account in the control policy. Otherwise, you cannot use the host account to access the target server through Bastionhost.
	//
	// example:
	//
	// Enable
	SSHChannel *string `json:"SSHChannel,omitempty" xml:"SSHChannel,omitempty"`
	// example:
	//
	// Enable
	TcpForwarding *string `json:"TcpForwarding,omitempty" xml:"TcpForwarding,omitempty"`
	// Specifies whether to enable X11 forwarding. Valid values:
	//
	// - Enable
	//
	// - Disable
	//
	// > The default value is Disable.
	//
	// example:
	//
	// Enable
	X11Forwarding *string `json:"X11Forwarding,omitempty" xml:"X11Forwarding,omitempty"`
}

func (s SetPolicyProtocolConfigRequestProtocolConfigSSH) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyProtocolConfigRequestProtocolConfigSSH) GoString() string {
	return s.String()
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetAllowDirectTcp() *string {
	return s.AllowDirectTcp
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetAllowTcpForwarding() *string {
	return s.AllowTcpForwarding
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetExecCommand() *string {
	return s.ExecCommand
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPChannel() *string {
	return s.SFTPChannel
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPDownloadFile() *string {
	return s.SFTPDownloadFile
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPMkdir() *string {
	return s.SFTPMkdir
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPRemoveFile() *string {
	return s.SFTPRemoveFile
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPRenameFile() *string {
	return s.SFTPRenameFile
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPRmdir() *string {
	return s.SFTPRmdir
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSFTPUploadFile() *string {
	return s.SFTPUploadFile
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetSSHChannel() *string {
	return s.SSHChannel
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetTcpForwarding() *string {
	return s.TcpForwarding
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetX11Forwarding() *string {
	return s.X11Forwarding
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetAllowDirectTcp(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.AllowDirectTcp = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetAllowTcpForwarding(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.AllowTcpForwarding = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetExecCommand(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.ExecCommand = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPChannel(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPChannel = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPDownloadFile(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPDownloadFile = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPMkdir(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPMkdir = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPRemoveFile(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPRemoveFile = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPRenameFile(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPRenameFile = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPRmdir(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPRmdir = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSFTPUploadFile(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SFTPUploadFile = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetSSHChannel(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.SSHChannel = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetTcpForwarding(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.TcpForwarding = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetX11Forwarding(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.X11Forwarding = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) Validate() error {
	return dara.Validate(s)
}
