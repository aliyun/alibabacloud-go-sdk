// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *GetGroupRequest
	GetGroupName() *string
	SetProjectName(v string) *GetGroupRequest
	GetProjectName() *string
}

type GetGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetGroupRequest) SetGroupName(v string) *GetGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetGroupRequest) SetProjectName(v string) *GetGroupRequest {
	s.ProjectName = &v
	return s
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
