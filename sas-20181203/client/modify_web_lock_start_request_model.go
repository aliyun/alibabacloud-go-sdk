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
	// The prevention mode. Valid values:
	//
	// 	- **block**: Interception Mode
	//
	// 	- **audit**: Alert Mode
	//
	// This parameter is required.
	//
	// example:
	//
	// block
	DefenceMode *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	// The directory for which you want to enable web tamper proofing. Separate multiple directories with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/admin/tomcat
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The directory for which you want to disable web tamper proofing.
	//
	// > If you set **Mode*	- to **blacklist**, you must specify this parameter.
	//
	// example:
	//
	// /home/admin/java
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The file for which you want to disable web tamper proofing.
	//
	// > If you set **Mode*	- to **blacklist**, you must specify this parameter.
	//
	// example:
	//
	// /home/admin/tomcat/localhost.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The type of the file for which you want to disable web tamper proofing. Separate multiple types with semicolons (;). Valid values:
	//
	// 	- php
	//
	// 	- jsp
	//
	// 	- asp
	//
	// 	- aspx
	//
	// 	- js
	//
	// 	- cgi
	//
	// 	- html
	//
	// 	- htm
	//
	// 	- xml
	//
	// 	- shtml
	//
	// 	- shtm
	//
	// 	- jpg
	//
	// 	- gif
	//
	// 	- png
	//
	// > If you set **Mode*	- to **blacklist**, you must specify this parameter.
	//
	// example:
	//
	// jpg
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	// The type of the file for which you want to enable web tamper proofing. Separate multiple types with semicolons (;). Valid values:
	//
	// 	- php
	//
	// 	- jsp
	//
	// 	- asp
	//
	// 	- aspx
	//
	// 	- js
	//
	// 	- cgi
	//
	// 	- html
	//
	// 	- htm
	//
	// 	- xml
	//
	// 	- shtml
	//
	// 	- shtm
	//
	// 	- jpg
	//
	// 	- gif
	//
	// 	- png
	//
	// > If you set **Mode*	- to **whitelist**, you must specify this parameter.
	//
	// example:
	//
	// php
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The local path to the backup files of the protected directory.\\
	//
	// The directory format of a Linux server is different from that of a Windows server. You must enter the directory in the required format based on your operating system. Examples:
	//
	// 	- Linux server: /usr/local/aegis/bak
	//
	// 	- Windows server: C:\\Program Files (x86)\\Alibaba\\Aegis\\bak
	//
	// This parameter is required.
	//
	// example:
	//
	// /usr/local/backup
	LocalBackupDir *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	// The protection mode of web tamper proofing. Valid values:
	//
	// 	- **whitelist**: In this mode, web tamper proofing is enabled for the specified directories and file types.
	//
	// 	- **blacklist**: In this mode, web tamper proofing is enabled for the unspecified subdirectories, file types, and files in the protected directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// whitelist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The UUID of the server for which you want to enable web tamper proofing.
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
