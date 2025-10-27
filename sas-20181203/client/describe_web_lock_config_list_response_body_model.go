// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockConfigListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*DescribeWebLockConfigListResponseBodyConfigList) *DescribeWebLockConfigListResponseBody
	GetConfigList() []*DescribeWebLockConfigListResponseBodyConfigList
	SetRequestId(v string) *DescribeWebLockConfigListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockConfigListResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockConfigListResponseBody struct {
	// The configurations of web tamper proofing.
	ConfigList []*DescribeWebLockConfigListResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of directories that have web tamper proofing enabled on the server.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockConfigListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponseBody) GetConfigList() []*DescribeWebLockConfigListResponseBodyConfigList {
	return s.ConfigList
}

func (s *DescribeWebLockConfigListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockConfigListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockConfigListResponseBody) SetConfigList(v []*DescribeWebLockConfigListResponseBodyConfigList) *DescribeWebLockConfigListResponseBody {
	s.ConfigList = v
	return s
}

func (s *DescribeWebLockConfigListResponseBody) SetRequestId(v string) *DescribeWebLockConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBody) SetTotalCount(v int32) *DescribeWebLockConfigListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBody) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebLockConfigListResponseBodyConfigList struct {
	// The prevention mode. Valid values:
	//
	// 	- **block**: Interception Mode
	//
	// 	- **audit**: Alert Mode
	//
	// example:
	//
	// audit
	DefenceMode *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	// The directory that has web tamper proofing enabled.
	//
	// example:
	//
	// /www/tmp/
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The directory that has web tamper proofing disabled.
	//
	// > If the value of **Mode*	- is **blacklist**, this parameter is returned.
	//
	// example:
	//
	// /home/admin/tomcat
	ExclusiveDir *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	// The file that has web tamper proofing disabled.
	//
	// > If the value of **Mode*	- is **blacklist**, this parameter is returned.
	//
	// example:
	//
	// /home/admin/tomcat/localhost.log
	ExclusiveFile *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	// The type of the file that has web tamper proofing disabled.
	//
	// > If the value of **Mode*	- is **blacklist**, this parameter is returned.
	//
	// example:
	//
	// *.jpg
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	// The configuration ID of the protected directory.
	//
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The file that has web tamper proofing enabled.
	//
	// > If the value of **Mode*	- is **whitelist**, this parameter is returned.
	//
	// example:
	//
	// /home/admin/tomcat/aaa.log
	InclusiveFile *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	// The type of the file that has web tamper proofing enabled.
	//
	// > If the value of **Mode*	- is **whitelist**, this parameter is returned.
	//
	// example:
	//
	// jpg
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	// The local path to the backup files of the protected directory.
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
	// The UUID of the server that has web tamper proofing enabled.
	//
	// example:
	//
	// 80d2f7d6-31a9-4d7f-8ff4-7ecc42f8****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockConfigListResponseBodyConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockConfigListResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetDefenceMode() *string {
	return s.DefenceMode
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetDir() *string {
	return s.Dir
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetExclusiveDir() *string {
	return s.ExclusiveDir
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetExclusiveFile() *string {
	return s.ExclusiveFile
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetExclusiveFileType() *string {
	return s.ExclusiveFileType
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetId() *string {
	return s.Id
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetInclusiveFile() *string {
	return s.InclusiveFile
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetInclusiveFileType() *string {
	return s.InclusiveFileType
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetLocalBackupDir() *string {
	return s.LocalBackupDir
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetDefenceMode(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.DefenceMode = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Dir = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveDir = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveFile(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveFile = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveFileType(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveFileType = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetId(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Id = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetInclusiveFile(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.InclusiveFile = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetInclusiveFileType(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.InclusiveFileType = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetLocalBackupDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.LocalBackupDir = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetMode(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Mode = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetUuid(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) Validate() error {
	return dara.Validate(s)
}
