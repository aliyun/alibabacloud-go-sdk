// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveCdsFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *MoveCdsFileRequest
	GetCdsId() *string
	SetConflictPolicy(v string) *MoveCdsFileRequest
	GetConflictPolicy() *string
	SetEndUserId(v string) *MoveCdsFileRequest
	GetEndUserId() *string
	SetFileId(v string) *MoveCdsFileRequest
	GetFileId() *string
	SetGroupId(v string) *MoveCdsFileRequest
	GetGroupId() *string
	SetParentFolderId(v string) *MoveCdsFileRequest
	GetParentFolderId() *string
	SetRegionId(v string) *MoveCdsFileRequest
	GetRegionId() *string
}

type MoveCdsFileRequest struct {
	// The ID of the cloud disk.
	//
	// example:
	//
	// cn-hangzhou+cds-346063****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The processing mode of files that have the same name.
	//
	// Valid values:
	//
	// 	- <!-- -->
	//
	//     refuse
	//
	//     <!-- -->
	//
	//     : If you want to create a file that uses the same name as an existing file in the cloud, the system denies your request and returns the details of the existing file.
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     auto_rename
	//
	//     <!-- -->
	//
	//     : If you want to create a file that uses the same name as an existing file in the cloud, the system renames the file that you want to create by appending the current time point.
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     ignore
	//
	//     <!-- -->
	//
	//     : The system allows you to create a file that uses the same name as an existing file in the cloud.
	//
	//     <!-- -->
	//
	// 	- <!-- -->
	//
	//     over_write
	//
	//     <!-- -->
	//
	//     : After you create a file that uses the same name as an existing file in the cloud, the new file overwrites the existing file.
	//
	//     <!-- -->
	//
	// example:
	//
	// ignore
	ConflictPolicy *string `json:"ConflictPolicy,omitempty" xml:"ConflictPolicy,omitempty"`
	// The user ID that you want to use to access the cloud disk.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the file.
	//
	// example:
	//
	// 63f3257b68b018170b194d87b875512d108f****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the parent folder that you want to move. If you want to remove the root folder, set the value to root.
	//
	// example:
	//
	// 6409848a6da91d6240604e7ba7337d85ba8a1****
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

func (s MoveCdsFileRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveCdsFileRequest) GoString() string {
	return s.String()
}

func (s *MoveCdsFileRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *MoveCdsFileRequest) GetConflictPolicy() *string {
	return s.ConflictPolicy
}

func (s *MoveCdsFileRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *MoveCdsFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *MoveCdsFileRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *MoveCdsFileRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *MoveCdsFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MoveCdsFileRequest) SetCdsId(v string) *MoveCdsFileRequest {
	s.CdsId = &v
	return s
}

func (s *MoveCdsFileRequest) SetConflictPolicy(v string) *MoveCdsFileRequest {
	s.ConflictPolicy = &v
	return s
}

func (s *MoveCdsFileRequest) SetEndUserId(v string) *MoveCdsFileRequest {
	s.EndUserId = &v
	return s
}

func (s *MoveCdsFileRequest) SetFileId(v string) *MoveCdsFileRequest {
	s.FileId = &v
	return s
}

func (s *MoveCdsFileRequest) SetGroupId(v string) *MoveCdsFileRequest {
	s.GroupId = &v
	return s
}

func (s *MoveCdsFileRequest) SetParentFolderId(v string) *MoveCdsFileRequest {
	s.ParentFolderId = &v
	return s
}

func (s *MoveCdsFileRequest) SetRegionId(v string) *MoveCdsFileRequest {
	s.RegionId = &v
	return s
}

func (s *MoveCdsFileRequest) Validate() error {
	return dara.Validate(s)
}
