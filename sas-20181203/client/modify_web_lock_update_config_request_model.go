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
	// The directory for which you want to enable web tamper proofing.
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
	// The ID of the protected directory for which you want to change the status of web tamper proofing.
	//
	// > You can call the [DescribeWebLockConfigList](~~DescribeWebLockConfigList~~) operation to query the IDs of protected directories.
	//
	// This parameter is required.
	//
	// example:
	//
	// 312077
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The file for which you want to enable web tamper proofing.
	//
	// > If you set **Mode*	- to **whitelist**, you must specify this parameter.
	//
	// example:
	//
	// /home/admin/test.log
	InclusiveFile *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
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
	// example:
	//
	// blacklist
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 36.112.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server on which the protected directory is located.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
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
