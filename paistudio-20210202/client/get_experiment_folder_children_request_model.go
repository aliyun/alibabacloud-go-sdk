// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentFolderChildrenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetExperimentFolderChildrenRequest
	GetAccessibility() *string
	SetOnlyFolder(v bool) *GetExperimentFolderChildrenRequest
	GetOnlyFolder() *bool
	SetSource(v string) *GetExperimentFolderChildrenRequest
	GetSource() *string
	SetUserId(v string) *GetExperimentFolderChildrenRequest
	GetUserId() *string
	SetWorkspaceId(v string) *GetExperimentFolderChildrenRequest
	GetWorkspaceId() *string
}

type GetExperimentFolderChildrenRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// true
	OnlyFolder *bool `json:"OnlyFolder,omitempty" xml:"OnlyFolder,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 12345******13324
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentFolderChildrenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentFolderChildrenRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentFolderChildrenRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetExperimentFolderChildrenRequest) GetOnlyFolder() *bool {
	return s.OnlyFolder
}

func (s *GetExperimentFolderChildrenRequest) GetSource() *string {
	return s.Source
}

func (s *GetExperimentFolderChildrenRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetExperimentFolderChildrenRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetExperimentFolderChildrenRequest) SetAccessibility(v string) *GetExperimentFolderChildrenRequest {
	s.Accessibility = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetOnlyFolder(v bool) *GetExperimentFolderChildrenRequest {
	s.OnlyFolder = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetSource(v string) *GetExperimentFolderChildrenRequest {
	s.Source = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetUserId(v string) *GetExperimentFolderChildrenRequest {
	s.UserId = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) SetWorkspaceId(v string) *GetExperimentFolderChildrenRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetExperimentFolderChildrenRequest) Validate() error {
	return dara.Validate(s)
}
