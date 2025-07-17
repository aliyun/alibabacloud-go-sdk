// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *GetFolderRequest
	GetFolderId() *string
	SetFolderPath(v string) *GetFolderRequest
	GetFolderPath() *string
	SetProjectId(v int64) *GetFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *GetFolderRequest
	GetProjectIdentifier() *string
}

type GetFolderRequest struct {
	// example:
	//
	// 273****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// Business_process/my_first_business_process/MaxCompute/ods_layer
	FolderPath *string `json:"FolderPath,omitempty" xml:"FolderPath,omitempty"`
	// example:
	//
	// 1000011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s GetFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetFolderRequest) GetFolderPath() *string {
	return s.FolderPath
}

func (s *GetFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

func (s *GetFolderRequest) SetFolderPath(v string) *GetFolderRequest {
	s.FolderPath = &v
	return s
}

func (s *GetFolderRequest) SetProjectId(v int64) *GetFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *GetFolderRequest) SetProjectIdentifier(v string) *GetFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *GetFolderRequest) Validate() error {
	return dara.Validate(s)
}
