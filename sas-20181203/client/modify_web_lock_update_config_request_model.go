// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUpdateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenceMode(v string) *ModifyWebLockUpdateConfigRequest
	GetDefenceMode() *string
	SetDir(v string) *ModifyWebLockUpdateConfigRequest
	GetDir() *string
	SetExclusiveDir(v string) *ModifyWebLockUpdateConfigRequest
	GetExclusiveDir() *string
	SetExclusiveFile(v string) *ModifyWebLockUpdateConfigRequest
	GetExclusiveFile() *string
	SetExclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest
	GetExclusiveFileType() *string
	SetId(v int32) *ModifyWebLockUpdateConfigRequest
	GetId() *int32
	SetInclusiveFile(v string) *ModifyWebLockUpdateConfigRequest
	GetInclusiveFile() *string
	SetInclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest
	GetInclusiveFileType() *string
	SetLang(v string) *ModifyWebLockUpdateConfigRequest
	GetLang() *string
	SetLocalBackupDir(v string) *ModifyWebLockUpdateConfigRequest
	GetLocalBackupDir() *string
	SetMode(v string) *ModifyWebLockUpdateConfigRequest
	GetMode() *string
	SetSourceIp(v string) *ModifyWebLockUpdateConfigRequest
	GetSourceIp() *string
	SetUuid(v string) *ModifyWebLockUpdateConfigRequest
	GetUuid() *string
}

type ModifyWebLockUpdateConfigRequest struct {
	// The defense mode. Valid values:
	//
	// - **block**: Block.
	//
	// - **audit**: Alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	DefenceMode *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	// The full path of the directory that you want to protect.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/admin/tomcat
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The directory that does not require web tamper-proofing protection (excluded directory).
	//
	// > This parameter is required when the protection pattern **Mode*	- is set to **blacklist**.
	//
	// example:
	//
	// /home/admin/test
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The file that does not require web tamper-proofing protection (excluded file).
	//
	// > This parameter is required when the protection pattern **Mode*	- is set to **blacklist**.
	//
	// example:
	//
	// /home/admin/apache.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The file types that do not require web tamper-proofing protection (excluded file types). Separate multiple file types with semicolons (;). Valid values:
	//
	// - php
	//
	// - jsp
	//
	// - asp
	//
	// - aspx
	//
	// - js
	//
	// - cgi
	//
	// - html
	//
	// - htm
	//
	// - xml
	//
	// - shtml
	//
	// - shtm
	//
	// - jpg
	//
	// - gif
	//
	// - png
	//
	// > This parameter is required when the protection pattern **Mode*	- is set to **blacklist**.
	//
	// example:
	//
	// jpg
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	// The ID of the protected directory configuration that you want to modify.
	//
	// > You can call the [DescribeWebLockConfigList](~~DescribeWebLockConfigList~~) operation to obtain the ID of the protected directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 312077
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The file that requires protection.
	//
	// > This parameter is required when the protection pattern **Mode*	- is set to **whitelist**.
	//
	// example:
	//
	// /home/admin/test.log
	InclusiveFile *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	// The file types that require web tamper-proofing protection. Separate multiple file types with semicolons (;). Valid values:
	//
	// - php
	//
	// - jsp
	//
	// - asp
	//
	// - aspx
	//
	// - js
	//
	// - cgi
	//
	// - html
	//
	// - htm
	//
	// - xml
	//
	// - shtml
	//
	// - shtm
	//
	// - jpg
	//
	// - gif
	//
	// - png
	//
	// > This parameter is required when the protection pattern **Mode*	- is set to **whitelist**.
	//
	// example:
	//
	// jpg
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The local backup path for securely backing up the protected directory.
	//
	// The format of the directory path may differ between Linux servers and Windows servers. Make sure that you enter the correct format. Refer to the following directory formats:
	//
	//  - Linux server: /usr/local/aegis/bak
	//
	//  - Windows server: C:\\Program Files (x86)\\Alibaba\\Aegis\\bak.
	//
	// This parameter is required.
	//
	// example:
	//
	// /usr/local/backup
	LocalBackupDir *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	// The protection pattern. Valid values:
	//
	// - **whitelist**: whitelist mode. Protects only the specified directories and file types.
	//
	// - **blacklist**: blacklist mode. Protects all subdirectories, file types, and files under the specified directory except those that are excluded.
	//
	// example:
	//
	// blacklist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 36.112.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server whose protected directory you want to modify.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the UUID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockUpdateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigRequest) GetDefenceMode() *string {
	return s.DefenceMode
}

func (s *ModifyWebLockUpdateConfigRequest) GetDir() *string {
	return s.Dir
}

func (s *ModifyWebLockUpdateConfigRequest) GetExclusiveDir() *string {
	return s.ExclusiveDir
}

func (s *ModifyWebLockUpdateConfigRequest) GetExclusiveFile() *string {
	return s.ExclusiveFile
}

func (s *ModifyWebLockUpdateConfigRequest) GetExclusiveFileType() *string {
	return s.ExclusiveFileType
}

func (s *ModifyWebLockUpdateConfigRequest) GetId() *int32 {
	return s.Id
}

func (s *ModifyWebLockUpdateConfigRequest) GetInclusiveFile() *string {
	return s.InclusiveFile
}

func (s *ModifyWebLockUpdateConfigRequest) GetInclusiveFileType() *string {
	return s.InclusiveFileType
}

func (s *ModifyWebLockUpdateConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyWebLockUpdateConfigRequest) GetLocalBackupDir() *string {
	return s.LocalBackupDir
}

func (s *ModifyWebLockUpdateConfigRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyWebLockUpdateConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyWebLockUpdateConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockUpdateConfigRequest) SetDefenceMode(v string) *ModifyWebLockUpdateConfigRequest {
	s.DefenceMode = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveFile(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetId(v int32) *ModifyWebLockUpdateConfigRequest {
	s.Id = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetInclusiveFile(v string) *ModifyWebLockUpdateConfigRequest {
	s.InclusiveFile = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetInclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetLang(v string) *ModifyWebLockUpdateConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetLocalBackupDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetMode(v string) *ModifyWebLockUpdateConfigRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetSourceIp(v string) *ModifyWebLockUpdateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetUuid(v string) *ModifyWebLockUpdateConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) Validate() error {
	return dara.Validate(s)
}
