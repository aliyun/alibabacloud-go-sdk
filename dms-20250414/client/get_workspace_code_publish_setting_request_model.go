// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceCodePublishSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetWorkspaceCodePublishSettingRequest
	GetWorkspaceId() *string
}

type GetWorkspaceCodePublishSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceCodePublishSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceCodePublishSettingRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceCodePublishSettingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceCodePublishSettingRequest) SetWorkspaceId(v string) *GetWorkspaceCodePublishSettingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceCodePublishSettingRequest) Validate() error {
	return dara.Validate(s)
}
