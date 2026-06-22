// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenceMode(v string) *ModifyWebLockStartRequest
	GetDefenceMode() *string
	SetDir(v string) *ModifyWebLockStartRequest
	GetDir() *string
	SetExclusiveDir(v string) *ModifyWebLockStartRequest
	GetExclusiveDir() *string
	SetExclusiveFile(v string) *ModifyWebLockStartRequest
	GetExclusiveFile() *string
	SetExclusiveFileType(v string) *ModifyWebLockStartRequest
	GetExclusiveFileType() *string
	SetInclusiveFileType(v string) *ModifyWebLockStartRequest
	GetInclusiveFileType() *string
	SetLocalBackupDir(v string) *ModifyWebLockStartRequest
	GetLocalBackupDir() *string
	SetMode(v string) *ModifyWebLockStartRequest
	GetMode() *string
	SetUuid(v string) *ModifyWebLockStartRequest
	GetUuid() *string
}

type ModifyWebLockStartRequest struct {
	// The defense mode. Valid values:
	//
	// - **block**: block
	//
	// - **audit**: alert.
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	DefenceMode *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	// The protection directories. Separate multiple directories with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/admin/tomcat
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The folder that does not require web tamper proofing protection (excluded folder).
	//
	// > This parameter is required when the Defense mode **Mode*	- is set to the **blacklist*	- pattern.
	//
	// example:
	//
	// /home/admin/java
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The files that do not require web tamper proofing protection (excluded files).
	//
	// > This parameter is required when the Defense mode **Mode*	- is set to the **blacklist*	- pattern.
	//
	// example:
	//
	// /home/admin/tomcat/localhost.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The file types that do not require web tamper proofing protection (excluded file types). Separate multiple file types with commas (,). Valid values:
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
	// > This parameter is required when the Defense mode **Mode*	- is set to the **blacklist*	- pattern.
	//
	// example:
	//
	// jpg
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	// The file types that require web tamper proofing protection. Separate multiple file types with commas (,). Valid values:
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
	// > This parameter is required when the Defense mode **Mode*	- is set to the **whitelist*	- pattern.
	//
	// example:
	//
	// php
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The local backup path used to back up the protection directories. The format of the protection directory path may differ between Linux servers and Windows servers. Make sure that you enter the path in the correct format. The following examples show the directory formats:
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
	// The protection type. Valid values:
	//
	// - **whitelist**: whitelist mode. Protects the specified protection directories and file types.
	//
	// - **blacklist**: blacklist mode. Protects all subdirectories, file types, and specified files in the protection directories that are not excluded.
	//
	// This parameter is required.
	//
	// example:
	//
	// whitelist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The UUID of the server that you want to protect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d2f7d6-31a9-4d7f-8ff4-7ecc42f89ca****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockStartRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStartRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartRequest) GetDefenceMode() *string {
	return s.DefenceMode
}

func (s *ModifyWebLockStartRequest) GetDir() *string {
	return s.Dir
}

func (s *ModifyWebLockStartRequest) GetExclusiveDir() *string {
	return s.ExclusiveDir
}

func (s *ModifyWebLockStartRequest) GetExclusiveFile() *string {
	return s.ExclusiveFile
}

func (s *ModifyWebLockStartRequest) GetExclusiveFileType() *string {
	return s.ExclusiveFileType
}

func (s *ModifyWebLockStartRequest) GetInclusiveFileType() *string {
	return s.InclusiveFileType
}

func (s *ModifyWebLockStartRequest) GetLocalBackupDir() *string {
	return s.LocalBackupDir
}

func (s *ModifyWebLockStartRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyWebLockStartRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockStartRequest) SetDefenceMode(v string) *ModifyWebLockStartRequest {
	s.DefenceMode = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetDir(v string) *ModifyWebLockStartRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveDir(v string) *ModifyWebLockStartRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveFile(v string) *ModifyWebLockStartRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveFileType(v string) *ModifyWebLockStartRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetInclusiveFileType(v string) *ModifyWebLockStartRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetLocalBackupDir(v string) *ModifyWebLockStartRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetMode(v string) *ModifyWebLockStartRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetUuid(v string) *ModifyWebLockStartRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockStartRequest) Validate() error {
	return dara.Validate(s)
}
