// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWordGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *UpdateWordGroupRequestBodyData) *UpdateWordGroupRequest
	GetBodyData() *UpdateWordGroupRequestBodyData
	SetCommit(v bool) *UpdateWordGroupRequest
	GetCommit() *bool
	SetGroupId(v int64) *UpdateWordGroupRequest
	GetGroupId() *int64
	SetGroupName(v string) *UpdateWordGroupRequest
	GetGroupName() *string
	SetRegionId(v string) *UpdateWordGroupRequest
	GetRegionId() *string
	SetWordInfoListModified(v bool) *UpdateWordGroupRequest
	GetWordInfoListModified() *bool
}

type UpdateWordGroupRequest struct {
	// Request object.
	BodyData *UpdateWordGroupRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s UpdateWordGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupRequest) GetBodyData() *UpdateWordGroupRequestBodyData {
	return s.BodyData
}

func (s *UpdateWordGroupRequest) GetCommit() *bool {
	return s.Commit
}

func (s *UpdateWordGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateWordGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateWordGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWordGroupRequest) GetWordInfoListModified() *bool {
	return s.WordInfoListModified
}

func (s *UpdateWordGroupRequest) SetBodyData(v *UpdateWordGroupRequestBodyData) *UpdateWordGroupRequest {
	s.BodyData = v
	return s
}

func (s *UpdateWordGroupRequest) SetCommit(v bool) *UpdateWordGroupRequest {
	s.Commit = &v
	return s
}

func (s *UpdateWordGroupRequest) SetGroupId(v int64) *UpdateWordGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateWordGroupRequest) SetGroupName(v string) *UpdateWordGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateWordGroupRequest) SetRegionId(v string) *UpdateWordGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWordGroupRequest) SetWordInfoListModified(v bool) *UpdateWordGroupRequest {
	s.WordInfoListModified = &v
	return s
}

func (s *UpdateWordGroupRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWordGroupRequestBodyData struct {
	// List of keyword groups.
	WordInfoList []*UpdateWordGroupRequestBodyDataWordInfoList `json:"WordInfoList,omitempty" xml:"WordInfoList,omitempty" type:"Repeated"`
}

func (s UpdateWordGroupRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupRequestBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupRequestBodyData) GetWordInfoList() []*UpdateWordGroupRequestBodyDataWordInfoList {
	return s.WordInfoList
}

func (s *UpdateWordGroupRequestBodyData) SetWordInfoList(v []*UpdateWordGroupRequestBodyDataWordInfoList) *UpdateWordGroupRequestBodyData {
	s.WordInfoList = v
	return s
}

func (s *UpdateWordGroupRequestBodyData) Validate() error {
	if s.WordInfoList != nil {
		for _, item := range s.WordInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWordGroupRequestBodyDataWordInfoList struct {
	// Label.
	//
	// example:
	//
	// Buss.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Keyword.
	//
	// example:
	//
	// Inv.
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s UpdateWordGroupRequestBodyDataWordInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupRequestBodyDataWordInfoList) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupRequestBodyDataWordInfoList) GetLabel() *string {
	return s.Label
}

func (s *UpdateWordGroupRequestBodyDataWordInfoList) GetWord() *string {
	return s.Word
}

func (s *UpdateWordGroupRequestBodyDataWordInfoList) SetLabel(v string) *UpdateWordGroupRequestBodyDataWordInfoList {
	s.Label = &v
	return s
}

func (s *UpdateWordGroupRequestBodyDataWordInfoList) SetWord(v string) *UpdateWordGroupRequestBodyDataWordInfoList {
	s.Word = &v
	return s
}

func (s *UpdateWordGroupRequestBodyDataWordInfoList) Validate() error {
	return dara.Validate(s)
}
