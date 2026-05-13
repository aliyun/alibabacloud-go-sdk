// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceCodePublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *WorkspaceCodePublishRequest
	GetConfig() *string
	SetWorkspaceId(v string) *WorkspaceCodePublishRequest
	GetWorkspaceId() *string
}

type WorkspaceCodePublishRequest struct {
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

func (s WorkspaceCodePublishRequest) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceCodePublishRequest) GoString() string {
	return s.String()
}

func (s *WorkspaceCodePublishRequest) GetConfig() *string {
	return s.Config
}

func (s *WorkspaceCodePublishRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspaceCodePublishRequest) SetConfig(v string) *WorkspaceCodePublishRequest {
	s.Config = &v
	return s
}

func (s *WorkspaceCodePublishRequest) SetWorkspaceId(v string) *WorkspaceCodePublishRequest {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspaceCodePublishRequest) Validate() error {
	return dara.Validate(s)
}
