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
	// The directory that you want to protect.
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
	// /home/admin/test
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The file for which you want to disable web tamper proofing.
	//
	// > If you set **Mode*	- to **blacklist**, you must specify this parameter.
	//
	// example:
	//
	// /home/admin/apache.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The type of file for which you want to disable web tamper proofing. Separate multiple types with semicolons (;). Valid values:
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
	// The file for which you want to enable web tamper proofing.
	//
	// > If you set **Mode*	- to **whitelist**, you must specify this parameter.
	//
	// example:
	//
	// /home/admin/test.log
	InclusiveFile *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	// The type of file for which you want to enable web tamper proofing. Separate multiple types with semicolons (;). Valid values:
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
	// jpg
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The local path to the backup files of the protected directory.
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
	// 	- **blacklist**: In this mode, web tamper proofing is enabled for the unspecified sub-directories, file types, and files in the protected directories.
	//
	// example:
	//
	// whitelist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 39.170.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server for which you want to add a directory to protect.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
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
