// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateExperimentFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *MigrateExperimentFoldersRequest
	GetAccessibility() *string
	SetIsOwner(v bool) *MigrateExperimentFoldersRequest
	GetIsOwner() *bool
	SetWorkspaceId(v string) *MigrateExperimentFoldersRequest
	GetWorkspaceId() *string
}

type MigrateExperimentFoldersRequest struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// true
	IsOwner *bool `json:"IsOwner,omitempty" xml:"IsOwner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MigrateExperimentFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateExperimentFoldersRequest) GoString() string {
	return s.String()
}

func (s *MigrateExperimentFoldersRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *MigrateExperimentFoldersRequest) GetIsOwner() *bool {
	return s.IsOwner
}

func (s *MigrateExperimentFoldersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MigrateExperimentFoldersRequest) SetAccessibility(v string) *MigrateExperimentFoldersRequest {
	s.Accessibility = &v
	return s
}

func (s *MigrateExperimentFoldersRequest) SetIsOwner(v bool) *MigrateExperimentFoldersRequest {
	s.IsOwner = &v
	return s
}

func (s *MigrateExperimentFoldersRequest) SetWorkspaceId(v string) *MigrateExperimentFoldersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MigrateExperimentFoldersRequest) Validate() error {
	return dara.Validate(s)
}
