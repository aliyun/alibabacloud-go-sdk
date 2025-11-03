// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilePermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *AddFilePermissionShrinkRequest
	GetCdsId() *string
	SetEndUserId(v string) *AddFilePermissionShrinkRequest
	GetEndUserId() *string
	SetFileId(v string) *AddFilePermissionShrinkRequest
	GetFileId() *string
	SetGroupId(v string) *AddFilePermissionShrinkRequest
	GetGroupId() *string
	SetMemberListShrink(v string) *AddFilePermissionShrinkRequest
	GetMemberListShrink() *string
	SetRegionId(v string) *AddFilePermissionShrinkRequest
	GetRegionId() *string
}

type AddFilePermissionShrinkRequest struct {
	// The ID of the enterprise drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the user who uses the network disk.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The file ID. You can call the [ListCdsFiles](https://help.aliyun.com/document_detail/2247622.html) operation to query the ID of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6333e553a133ce21e6f747cf948bb9ef95d7****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the team space.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The users that you want to authorize to use the cloud disk.
	//
	// This parameter is required.
	MemberListShrink *string `json:"MemberList,omitempty" xml:"MemberList,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddFilePermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFilePermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddFilePermissionShrinkRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *AddFilePermissionShrinkRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *AddFilePermissionShrinkRequest) GetFileId() *string {
	return s.FileId
}

func (s *AddFilePermissionShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddFilePermissionShrinkRequest) GetMemberListShrink() *string {
	return s.MemberListShrink
}

func (s *AddFilePermissionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddFilePermissionShrinkRequest) SetCdsId(v string) *AddFilePermissionShrinkRequest {
	s.CdsId = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) SetEndUserId(v string) *AddFilePermissionShrinkRequest {
	s.EndUserId = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) SetFileId(v string) *AddFilePermissionShrinkRequest {
	s.FileId = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) SetGroupId(v string) *AddFilePermissionShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) SetMemberListShrink(v string) *AddFilePermissionShrinkRequest {
	s.MemberListShrink = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) SetRegionId(v string) *AddFilePermissionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AddFilePermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
