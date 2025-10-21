// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWordGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *CreateWordGroupShrinkRequest
	GetBodyDataShrink() *string
	SetCommit(v bool) *CreateWordGroupShrinkRequest
	GetCommit() *bool
	SetGroupName(v string) *CreateWordGroupShrinkRequest
	GetGroupName() *string
	SetRegionId(v string) *CreateWordGroupShrinkRequest
	GetRegionId() *string
	SetWorkspaceId(v int64) *CreateWordGroupShrinkRequest
	GetWorkspaceId() *int64
}

type CreateWordGroupShrinkRequest struct {
	// Request object
	BodyDataShrink *string `json:"BodyData,omitempty" xml:"BodyData,omitempty"`
	// Whether to commit.
	//
	// false: Not actually saved, only checked
	//
	// true: Commit and save
	//
	// example:
	//
	// true
	Commit *bool `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// Keyword group name
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// 643168
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWordGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWordGroupShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *CreateWordGroupShrinkRequest) GetCommit() *bool {
	return s.Commit
}

func (s *CreateWordGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateWordGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWordGroupShrinkRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateWordGroupShrinkRequest) SetBodyDataShrink(v string) *CreateWordGroupShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *CreateWordGroupShrinkRequest) SetCommit(v bool) *CreateWordGroupShrinkRequest {
	s.Commit = &v
	return s
}

func (s *CreateWordGroupShrinkRequest) SetGroupName(v string) *CreateWordGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateWordGroupShrinkRequest) SetRegionId(v string) *CreateWordGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWordGroupShrinkRequest) SetWorkspaceId(v int64) *CreateWordGroupShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWordGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
