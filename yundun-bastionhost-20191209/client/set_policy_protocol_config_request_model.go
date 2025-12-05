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
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The protocol control settings.
	//
	// This parameter is required.
	ProtocolConfig *SetPolicyProtocolConfigRequestProtocolConfig `json:"ProtocolConfig,omitempty" xml:"ProtocolConfig,omitempty" type:"Struct"`
	// The region ID of the bastion host.
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
	// The settings of the Remote Desktop Protocol (RDP) options.
	RDP *SetPolicyProtocolConfigRequestProtocolConfigRDP `json:"RDP,omitempty" xml:"RDP,omitempty" type:"Struct"`
	// The settings of the SSH and SSH Fine Transfer Protocol (SFTP) options.
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
	// Specifies whether to enable downloading from the clipboard. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	ClipboardDownload *string `json:"ClipboardDownload,omitempty" xml:"ClipboardDownload,omitempty"`
	// Specifies whether to enable uploading from the clipboard. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	ClipboardUpload *string `json:"ClipboardUpload,omitempty" xml:"ClipboardUpload,omitempty"`
	// Specifies whether to enable driver mapping. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	DiskRedirection *string `json:"DiskRedirection,omitempty" xml:"DiskRedirection,omitempty"`
	// Specifies whether to enable keyboard operation recording. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
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

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) SetRecordKeyboard(v string) *SetPolicyProtocolConfigRequestProtocolConfigRDP {
	s.RecordKeyboard = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigRDP) Validate() error {
	return dara.Validate(s)
}

type SetPolicyProtocolConfigRequestProtocolConfigSSH struct {
	// Specifies whether to enable remote command execution. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	ExecCommand *string `json:"ExecCommand,omitempty" xml:"ExecCommand,omitempty"`
	// Specifies whether to enable SFTP channels. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// >
	//
	// 	- If you do not specify this parameter, the default value Disable is used.
	//
	// 	- You must set at least one of the following parameters to Enable: SSHChannel and SFTPChannel.
	//
	// 	- If you select Enable Only SFTP Permission for a host account, do not set SSHChannel and SFTPChannel to Disable for the account. Otherwise, users of the bastion host cannot use the account to access the host.
	//
	// example:
	//
	// Enable
	SFTPChannel *string `json:"SFTPChannel,omitempty" xml:"SFTPChannel,omitempty"`
	// Specifies whether to enable file downloading during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPDownloadFile *string `json:"SFTPDownloadFile,omitempty" xml:"SFTPDownloadFile,omitempty"`
	// Specifies whether to enable folder creation during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPMkdir *string `json:"SFTPMkdir,omitempty" xml:"SFTPMkdir,omitempty"`
	// Specifies whether to enable file deletion during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPRemoveFile *string `json:"SFTPRemoveFile,omitempty" xml:"SFTPRemoveFile,omitempty"`
	// Specifies whether to enable file renaming during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPRenameFile *string `json:"SFTPRenameFile,omitempty" xml:"SFTPRenameFile,omitempty"`
	// Specifies whether to enable folder deletion during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPRmdir *string `json:"SFTPRmdir,omitempty" xml:"SFTPRmdir,omitempty"`
	// Specifies whether to enable file uploading during SFTP-based O\\&M. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
	//
	// example:
	//
	// Enable
	SFTPUploadFile *string `json:"SFTPUploadFile,omitempty" xml:"SFTPUploadFile,omitempty"`
	// Specifies whether to enable SSH channels. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// >
	//
	// 	- If you do not specify this parameter, the default value Disable is used.
	//
	// 	- You must set at least one of the following parameters to Enable: SSHChannel and SFTPChannel. If you set SSHChannel to Disable, SSH-based logon is disabled for the asset account. Proceed with caution.
	//
	// 	- If you select Enable Only SFTP Permission for a host account, do not set SSHChannel and SFTPChannel to Disable for the account. Otherwise, users of the bastion host cannot use the account to access the host.
	//
	// example:
	//
	// Enable
	SSHChannel *string `json:"SSHChannel,omitempty" xml:"SSHChannel,omitempty"`
	// Specifies whether to enable X11 forwarding. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// > If you do not specify this parameter, the default value Disable is used.
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

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) GetX11Forwarding() *string {
	return s.X11Forwarding
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

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) SetX11Forwarding(v string) *SetPolicyProtocolConfigRequestProtocolConfigSSH {
	s.X11Forwarding = &v
	return s
}

func (s *SetPolicyProtocolConfigRequestProtocolConfigSSH) Validate() error {
	return dara.Validate(s)
}
