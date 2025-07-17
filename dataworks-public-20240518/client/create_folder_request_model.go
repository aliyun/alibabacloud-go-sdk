// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderPath(v string) *CreateFolderRequest
	GetFolderPath() *string
	SetProjectId(v int64) *CreateFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *CreateFolderRequest
	GetProjectIdentifier() *string
}

type CreateFolderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Business_process/System_Data/MaxCompute/import_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) GetFolderPath() *string {
	return s.FolderPath
}

func (s *CreateFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateFolderRequest) SetFolderPath(v string) *CreateFolderRequest {
	s.FolderPath = &v
	return s
}

func (s *CreateFolderRequest) SetProjectId(v int64) *CreateFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFolderRequest) SetProjectIdentifier(v string) *CreateFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateFolderRequest) Validate() error {
	return dara.Validate(s)
}
