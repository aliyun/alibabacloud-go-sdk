// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceCodePublishSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SetWorkspaceCodePublishSettingRequest
	GetConfig() *string
	SetWorkspaceId(v string) *SetWorkspaceCodePublishSettingRequest
	GetWorkspaceId() *string
}

type SetWorkspaceCodePublishSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"repos":[{"repo":"git@xxxx.git", "branch":"master"}], "exclude":["/.dms", "/username"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SetWorkspaceCodePublishSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceCodePublishSettingRequest) GoString() string {
	return s.String()
}

func (s *SetWorkspaceCodePublishSettingRequest) GetConfig() *string {
	return s.Config
}

func (s *SetWorkspaceCodePublishSettingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SetWorkspaceCodePublishSettingRequest) SetConfig(v string) *SetWorkspaceCodePublishSettingRequest {
	s.Config = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingRequest) SetWorkspaceId(v string) *SetWorkspaceCodePublishSettingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SetWorkspaceCodePublishSettingRequest) Validate() error {
	return dara.Validate(s)
}
