// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRename(v bool) *CopyCdsFileRequest
	GetAutoRename() *bool
	SetCdsId(v string) *CopyCdsFileRequest
	GetCdsId() *string
	SetEndUserId(v string) *CopyCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *CopyCdsFileRequest
	GetFileId() *string
	SetFileReceiverId(v string) *CopyCdsFileRequest
	GetFileReceiverId() *string
	SetFileReceiverType(v string) *CopyCdsFileRequest
	GetFileReceiverType() *string
	SetGroupId(v string) *CopyCdsFileRequest
	GetGroupId() *string
	SetParentFolderId(v string) *CopyCdsFileRequest
	GetParentFolderId() *string
	SetRegionId(v string) *CopyCdsFileRequest
	GetRegionId() *string
}

type CopyCdsFileRequest struct {
	// Specifies whether to automatically rename the file if a file that has the same name exists in the folder to which you want to copy the file. Default value: false.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	AutoRename *bool `json:"AutoRename,omitempty" xml:"AutoRename,omitempty"`
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The user ID that you want to use to access the cloud disk.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call the CreateCdsFile operation to query the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 640985a0ca2f71f489d2497682ca0bf468de****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 目标复制文件所在的个人空间ID（即UserId，您可以在DescribeCloudDriveUsers接口返回的报文中获取。）或者目标复制文件所在的团队空间ID（即GroupId，您可以在DescribeCloudDriveGroups接口返回的报文中获取。）
	//
	// > FileReceiverId和FileReceiverType都为空时，默认复制到文件所在的个人空间。
	//
	// >
	//
	// example:
	//
	// user02
	FileReceiverId *string `json:"FileReceiverId,omitempty" xml:"FileReceiverId,omitempty"`
	// 文件所属的空间类型。
	//
	// example:
	//
	// user
	FileReceiverType *string `json:"FileReceiverType,omitempty" xml:"FileReceiverType,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the parent folder of the folder to which you want to copy the file. If you want to copy the file to the root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CopyCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyCdsFileRequest) GoString() string {
	return s.String()
}

func (s *CopyCdsFileRequest) GetAutoRename() *bool {
	return s.AutoRename
}

func (s *CopyCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CopyCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CopyCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *CopyCdsFileRequest) GetFileReceiverId() *string {
	return s.FileReceiverId
}

func (s *CopyCdsFileRequest) GetFileReceiverType() *string {
	return s.FileReceiverType
}

func (s *CopyCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CopyCdsFileRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CopyCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyCdsFileRequest) SetAutoRename(v bool) *CopyCdsFileRequest {
	s.AutoRename = &v
	return s
}

func (s *CopyCdsFileRequest) SetCdsId(v string) *CopyCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *CopyCdsFileRequest) SetEndUserId(v string) *CopyCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *CopyCdsFileRequest) SetFileId(v string) *CopyCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *CopyCdsFileRequest) SetFileReceiverId(v string) *CopyCdsFileRequest {
	s.FileReceiverId = &v
	return s
}

func (s *CopyCdsFileRequest) SetFileReceiverType(v string) *CopyCdsFileRequest {
	s.FileReceiverType = &v
	return s
}

func (s *CopyCdsFileRequest) SetGroupId(v string) *CopyCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *CopyCdsFileRequest) SetParentFolderId(v string) *CopyCdsFileRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CopyCdsFileRequest) SetRegionId(v string) *CopyCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *CopyCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
