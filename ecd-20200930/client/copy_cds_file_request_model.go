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
	// Specifies whether to automatically rename the file when a file with the same name already exists in the destination folder.
	//
	// example:
	//
	// true
	AutoRename *bool `json:"AutoRename,omitempty" xml:"AutoRename,omitempty"`
	// The enterprise cloud drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user who is logged on to the cloud drive.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) to query the ID of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 640985a0ca2f71f489d2497682ca0bf468de****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the personal drive (which can be obtained from the UserId response parameter of the [DescribeCloudDriveUsers](https://help.aliyun.com/document_detail/2357237.html) operation) or the team space ID (which can be obtained from the GroupId response parameter of the [DescribeCloudDriveGroups](https://help.aliyun.com/document_detail/609896.html) operation) at the destination.
	//
	// > If both FileReceiverId and FileReceiverType are empty, the file is copied to the personal drive where the file currently resides by default.
	//
	// example:
	//
	// user02
	FileReceiverId *string `json:"FileReceiverId,omitempty" xml:"FileReceiverId,omitempty"`
	// The type of the space to which the file belongs.
	//
	// example:
	//
	// user
	FileReceiverType *string `json:"FileReceiverType,omitempty" xml:"FileReceiverType,omitempty"`
	// The team space ID.
	//
	// example:
	//
	// cg-hs3i1w39o68ma****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the parent folder at the destination. You can call [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) to query the ID of the folder. Set this parameter to `root` if you want to copy the file or folder to the root directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
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
