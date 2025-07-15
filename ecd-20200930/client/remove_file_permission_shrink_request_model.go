// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFilePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *RemoveFilePermissionShrinkRequest
	GetCdsId() *string
	SetEndUserId(v string) *RemoveFilePermissionShrinkRequest
	GetEndUserId() *string
	SetFileId(v string) *RemoveFilePermissionShrinkRequest
	GetFileId() *string
	SetGroupId(v string) *RemoveFilePermissionShrinkRequest
	GetGroupId() *string
	SetMemberListShrink(v string) *RemoveFilePermissionShrinkRequest
	GetMemberListShrink() *string
	SetRegionId(v string) *RemoveFilePermissionShrinkRequest
	GetRegionId() *string
}

type RemoveFilePermissionShrinkRequest struct {
	// The ID of the cloud disk in Cloud Drive Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-066224****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. The ID is a unique identifier for the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6333e553a133ce21e6f747cf948bb9ef95d7****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The group ID.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The users that you want to authorize.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveFilePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFilePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveFilePermissionShrinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *RemoveFilePermissionShrinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *RemoveFilePermissionShrinkRequest) GetFileId() *string {
	return s.FileId
}

func (s *RemoveFilePermissionShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveFilePermissionShrinkRequest) GetMemberListShrink() *string {
	return s.MemberListShrink
}

func (s *RemoveFilePermissionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveFilePermissionShrinkRequest) SetCdsId(v string) *RemoveFilePermissionShrinkRequest {
	s.CdsId = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) SetEndUserId(v string) *RemoveFilePermissionShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) SetFileId(v string) *RemoveFilePermissionShrinkRequest {
	s.FileId = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) SetGroupId(v string) *RemoveFilePermissionShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) SetMemberListShrink(v string) *RemoveFilePermissionShrinkRequest {
	s.MemberListShrink = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) SetRegionId(v string) *RemoveFilePermissionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveFilePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
