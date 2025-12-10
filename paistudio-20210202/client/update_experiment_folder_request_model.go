// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateExperimentFolderRequest
	GetName() *string
	SetParentFolderId(v string) *UpdateExperimentFolderRequest
	GetParentFolderId() *string
}

type UpdateExperimentFolderRequest struct {
	// example:
	//
	// folder1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// folder-xzf7t785nka4c2334
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s UpdateExperimentFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentFolderRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExperimentFolderRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *UpdateExperimentFolderRequest) SetName(v string) *UpdateExperimentFolderRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentFolderRequest) SetParentFolderId(v string) *UpdateExperimentFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *UpdateExperimentFolderRequest) Validate() error {
	return dara.Validate(s)
}
