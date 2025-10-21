// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWordGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyData(v *CreateWordGroupRequestBodyData) *CreateWordGroupRequest
	GetBodyData() *CreateWordGroupRequestBodyData
	SetCommit(v bool) *CreateWordGroupRequest
	GetCommit() *bool
	SetGroupName(v string) *CreateWordGroupRequest
	GetGroupName() *string
	SetRegionId(v string) *CreateWordGroupRequest
	GetRegionId() *string
	SetWorkspaceId(v int64) *CreateWordGroupRequest
	GetWorkspaceId() *int64
}

type CreateWordGroupRequest struct {
	// Request object
	BodyData *CreateWordGroupRequestBodyData `json:"BodyData,omitempty" xml:"BodyData,omitempty" type:"Struct"`
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

func (s CreateWordGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateWordGroupRequest) GetBodyData() *CreateWordGroupRequestBodyData {
	return s.BodyData
}

func (s *CreateWordGroupRequest) GetCommit() *bool {
	return s.Commit
}

func (s *CreateWordGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateWordGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWordGroupRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *CreateWordGroupRequest) SetBodyData(v *CreateWordGroupRequestBodyData) *CreateWordGroupRequest {
	s.BodyData = v
	return s
}

func (s *CreateWordGroupRequest) SetCommit(v bool) *CreateWordGroupRequest {
	s.Commit = &v
	return s
}

func (s *CreateWordGroupRequest) SetGroupName(v string) *CreateWordGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateWordGroupRequest) SetRegionId(v string) *CreateWordGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWordGroupRequest) SetWorkspaceId(v int64) *CreateWordGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWordGroupRequest) Validate() error {
	if s.BodyData != nil {
		if err := s.BodyData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWordGroupRequestBodyData struct {
	// Keyword group list
	WordInfoList []*CreateWordGroupRequestBodyDataWordInfoList `json:"WordInfoList,omitempty" xml:"WordInfoList,omitempty" type:"Repeated"`
}

func (s CreateWordGroupRequestBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupRequestBodyData) GoString() string {
	return s.String()
}

func (s *CreateWordGroupRequestBodyData) GetWordInfoList() []*CreateWordGroupRequestBodyDataWordInfoList {
	return s.WordInfoList
}

func (s *CreateWordGroupRequestBodyData) SetWordInfoList(v []*CreateWordGroupRequestBodyDataWordInfoList) *CreateWordGroupRequestBodyData {
	s.WordInfoList = v
	return s
}

func (s *CreateWordGroupRequestBodyData) Validate() error {
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

type CreateWordGroupRequestBodyDataWordInfoList struct {
	// Label
	//
	// example:
	//
	// Buss.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Keyword
	//
	// example:
	//
	// Inv.
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s CreateWordGroupRequestBodyDataWordInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupRequestBodyDataWordInfoList) GoString() string {
	return s.String()
}

func (s *CreateWordGroupRequestBodyDataWordInfoList) GetLabel() *string {
	return s.Label
}

func (s *CreateWordGroupRequestBodyDataWordInfoList) GetWord() *string {
	return s.Word
}

func (s *CreateWordGroupRequestBodyDataWordInfoList) SetLabel(v string) *CreateWordGroupRequestBodyDataWordInfoList {
	s.Label = &v
	return s
}

func (s *CreateWordGroupRequestBodyDataWordInfoList) SetWord(v string) *CreateWordGroupRequestBodyDataWordInfoList {
	s.Word = &v
	return s
}

func (s *CreateWordGroupRequestBodyDataWordInfoList) Validate() error {
	return dara.Validate(s)
}
