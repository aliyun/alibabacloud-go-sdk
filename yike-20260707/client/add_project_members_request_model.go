// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *AddProjectMembersRequest
	GetProjectId() *string
	SetYikeUserIds(v string) *AddProjectMembersRequest
	GetYikeUserIds() *string
}

type AddProjectMembersRequest struct {
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The Yike user IDs to be added. Separate multiple IDs with commas. The users must belong to the current Alibaba Cloud account. By default, the Edit permission is granted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001,10002
	YikeUserIds *string `json:"YikeUserIds,omitempty" xml:"YikeUserIds,omitempty"`
}

func (s AddProjectMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMembersRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *AddProjectMembersRequest) GetYikeUserIds() *string {
	return s.YikeUserIds
}

func (s *AddProjectMembersRequest) SetProjectId(v string) *AddProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *AddProjectMembersRequest) SetYikeUserIds(v string) *AddProjectMembersRequest {
	s.YikeUserIds = &v
	return s
}

func (s *AddProjectMembersRequest) Validate() error {
	return dara.Validate(s)
}
