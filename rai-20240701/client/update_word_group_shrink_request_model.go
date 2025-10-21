// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWordGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyDataShrink(v string) *UpdateWordGroupShrinkRequest
	GetBodyDataShrink() *string
	SetCommit(v bool) *UpdateWordGroupShrinkRequest
	GetCommit() *bool
	SetGroupId(v int64) *UpdateWordGroupShrinkRequest
	GetGroupId() *int64
	SetGroupName(v string) *UpdateWordGroupShrinkRequest
	GetGroupName() *string
	SetRegionId(v string) *UpdateWordGroupShrinkRequest
	GetRegionId() *string
	SetWordInfoListModified(v bool) *UpdateWordGroupShrinkRequest
	GetWordInfoListModified() *bool
}

type UpdateWordGroupShrinkRequest struct {
	// Request object.
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
	// Keyword group ID.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Keyword group name.
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
	// Whether the keyword information has been modified.
	//
	// example:
	//
	// true
	WordInfoListModified *bool `json:"WordInfoListModified,omitempty" xml:"WordInfoListModified,omitempty"`
}

func (s UpdateWordGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupShrinkRequest) GetBodyDataShrink() *string {
	return s.BodyDataShrink
}

func (s *UpdateWordGroupShrinkRequest) GetCommit() *bool {
	return s.Commit
}

func (s *UpdateWordGroupShrinkRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateWordGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateWordGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWordGroupShrinkRequest) GetWordInfoListModified() *bool {
	return s.WordInfoListModified
}

func (s *UpdateWordGroupShrinkRequest) SetBodyDataShrink(v string) *UpdateWordGroupShrinkRequest {
	s.BodyDataShrink = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) SetCommit(v bool) *UpdateWordGroupShrinkRequest {
	s.Commit = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) SetGroupId(v int64) *UpdateWordGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) SetGroupName(v string) *UpdateWordGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) SetRegionId(v string) *UpdateWordGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) SetWordInfoListModified(v bool) *UpdateWordGroupShrinkRequest {
	s.WordInfoListModified = &v
	return s
}

func (s *UpdateWordGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
