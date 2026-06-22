// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockCreateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenceMode(v string) *ModifyWebLockCreateConfigRequest
	GetDefenceMode() *string
	SetDir(v string) *ModifyWebLockCreateConfigRequest
	GetDir() *string
	SetExclusiveDir(v string) *ModifyWebLockCreateConfigRequest
	GetExclusiveDir() *string
	SetExclusiveFile(v string) *ModifyWebLockCreateConfigRequest
	GetExclusiveFile() *string
	SetExclusiveFileType(v string) *ModifyWebLockCreateConfigRequest
	GetExclusiveFileType() *string
	SetInclusiveFile(v string) *ModifyWebLockCreateConfigRequest
	GetInclusiveFile() *string
	SetInclusiveFileType(v string) *ModifyWebLockCreateConfigRequest
	GetInclusiveFileType() *string
	SetLang(v string) *ModifyWebLockCreateConfigRequest
	GetLang() *string
	SetLocalBackupDir(v string) *ModifyWebLockCreateConfigRequest
	GetLocalBackupDir() *string
	SetMode(v string) *ModifyWebLockCreateConfigRequest
	GetMode() *string
	SetSourceIp(v string) *ModifyWebLockCreateConfigRequest
	GetSourceIp() *string
	SetUuid(v string) *ModifyWebLockCreateConfigRequest
	GetUuid() *string
}

type ModifyWebLockCreateConfigRequest struct {
	// The defense mode. Valid values:
	//
	// - **block**: Block mode.
	//
	// - **audit**: Alert mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	DefenceMode *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	// The protected directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/admin/tomcat
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The folder to exclude from web tamper proofing protection.
	//
	// > This parameter is required when the Defense mode **Mode*	- is set to **blacklist*	- pattern.
	//
	// example:
	//
	// /home/admin/test
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The file to exclude from web tamper proofing protection.
	//
	// > This parameter is required when the Defense mode **Mode*	- is set to **blacklist*	- pattern.
	//
	// example:
	//
	// /home/admin/apache.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The file types to exclude from web tamper proofing protection. Separate multiple file types with semicolons (;). Valid values:
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
	// > This parameter is required when the Defense mode **Mode*	- is set to **blacklist*	- pattern.
	//
	// example:
	//
	// jpg
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	// The file to protect.
	//
	// > This parameter is required when the Defense mode **Mode*	- is set to **whitelist*	- pattern.
	//
	// example:
	//
	// /home/admin/test.log
	InclusiveFile *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	// The file types to protect with web tamper proofing. Separate multiple file types with semicolons (;). Valid values:
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
	// > This parameter is required when the Defense mode **Mode*	- is set to **whitelist*	- pattern.
	//
	// example:
	//
	// jpg
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The language type of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The local backup path used for secure backup of the protected directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// /usr/local/backup
	LocalBackupDir *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	// The protection directory mode. Valid values:
	//
	// - **whitelist**: whitelist mode. Protects only the specified directories and file types.
	//
	// - **blacklist**: blacklist mode. Protects all subdirectories, file types, and specified files under the protected directory that are not excluded.
	//
	// example:
	//
	// whitelist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server for which you want to add a protected directory.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the UUID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-12345****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockCreateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockCreateConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigRequest) GetDefenceMode() *string {
	return s.DefenceMode
}

func (s *ModifyWebLockCreateConfigRequest) GetDir() *string {
	return s.Dir
}

func (s *ModifyWebLockCreateConfigRequest) GetExclusiveDir() *string {
	return s.ExclusiveDir
}

func (s *ModifyWebLockCreateConfigRequest) GetExclusiveFile() *string {
	return s.ExclusiveFile
}

func (s *ModifyWebLockCreateConfigRequest) GetExclusiveFileType() *string {
	return s.ExclusiveFileType
}

func (s *ModifyWebLockCreateConfigRequest) GetInclusiveFile() *string {
	return s.InclusiveFile
}

func (s *ModifyWebLockCreateConfigRequest) GetInclusiveFileType() *string {
	return s.InclusiveFileType
}

func (s *ModifyWebLockCreateConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyWebLockCreateConfigRequest) GetLocalBackupDir() *string {
	return s.LocalBackupDir
}

func (s *ModifyWebLockCreateConfigRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyWebLockCreateConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyWebLockCreateConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockCreateConfigRequest) SetDefenceMode(v string) *ModifyWebLockCreateConfigRequest {
	s.DefenceMode = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetDir(v string) *ModifyWebLockCreateConfigRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveDir(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveFile(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveFileType(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetInclusiveFile(v string) *ModifyWebLockCreateConfigRequest {
	s.InclusiveFile = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetInclusiveFileType(v string) *ModifyWebLockCreateConfigRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetLang(v string) *ModifyWebLockCreateConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetLocalBackupDir(v string) *ModifyWebLockCreateConfigRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetMode(v string) *ModifyWebLockCreateConfigRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetSourceIp(v string) *ModifyWebLockCreateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetUuid(v string) *ModifyWebLockCreateConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) Validate() error {
	return dara.Validate(s)
}
