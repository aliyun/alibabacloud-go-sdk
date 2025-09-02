// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v int64) *GetDataServiceFolderRequest
	GetFolderId() *int64
	SetProjectId(v int64) *GetDataServiceFolderRequest
	GetProjectId() *int64
	SetTenantId(v int64) *GetDataServiceFolderRequest
	GetTenantId() *int64
}

type GetDataServiceFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	FolderId *int64 `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. This parameter is deprecated. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10003
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataServiceFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceFolderRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceFolderRequest) GetFolderId() *int64 {
	return s.FolderId
}

func (s *GetDataServiceFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataServiceFolderRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataServiceFolderRequest) SetFolderId(v int64) *GetDataServiceFolderRequest {
	s.FolderId = &v
	return s
}

func (s *GetDataServiceFolderRequest) SetProjectId(v int64) *GetDataServiceFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceFolderRequest) SetTenantId(v int64) *GetDataServiceFolderRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataServiceFolderRequest) Validate() error {
	return dara.Validate(s)
}
