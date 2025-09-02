// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderName(v string) *CreateDataServiceFolderRequest
	GetFolderName() *string
	SetGroupId(v string) *CreateDataServiceFolderRequest
	GetGroupId() *string
	SetParentId(v int64) *CreateDataServiceFolderRequest
	GetParentId() *int64
	SetProjectId(v int64) *CreateDataServiceFolderRequest
	GetProjectId() *int64
	SetTenantId(v int64) *CreateDataServiceFolderRequest
	GetTenantId() *int64
}

type CreateDataServiceFolderRequest struct {
	// The name of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test folder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the desired workflow to which the folder belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000abcd
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the desired parent folder of the folder. The ID of the root folder in a workflow is 0. The ID of the folder created by users in a workflow is greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateDataServiceFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceFolderRequest) GetFolderName() *string {
	return s.FolderName
}

func (s *CreateDataServiceFolderRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDataServiceFolderRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataServiceFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataServiceFolderRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *CreateDataServiceFolderRequest) SetFolderName(v string) *CreateDataServiceFolderRequest {
	s.FolderName = &v
	return s
}

func (s *CreateDataServiceFolderRequest) SetGroupId(v string) *CreateDataServiceFolderRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDataServiceFolderRequest) SetParentId(v int64) *CreateDataServiceFolderRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataServiceFolderRequest) SetProjectId(v int64) *CreateDataServiceFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataServiceFolderRequest) SetTenantId(v int64) *CreateDataServiceFolderRequest {
	s.TenantId = &v
	return s
}

func (s *CreateDataServiceFolderRequest) Validate() error {
	return dara.Validate(s)
}
