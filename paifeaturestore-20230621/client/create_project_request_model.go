// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
	SetOfflineDatasourceId(v string) *CreateProjectRequest
	GetOfflineDatasourceId() *string
	SetOfflineLifeCycle(v int32) *CreateProjectRequest
	GetOfflineLifeCycle() *int32
	SetOnlineDatasourceId(v string) *CreateProjectRequest
	GetOnlineDatasourceId() *string
	SetWorkspaceId(v string) *CreateProjectRequest
	GetWorkspaceId() *string
}

type CreateProjectRequest struct {
	// The description of the project. This description is displayed in the PAI console.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the project. The name must be unique within the instance. It must be 2 to 18 characters long, begin with a letter, and contain only letters, digits, and underscores (_). Regex: ^[a-zA-Z][a-zA-Z0-9_]+$.
	//
	// This parameter is required.
	//
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the offline data source, which must be of the **MaxCompute*	- type. You can call the ListDatasources operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	OfflineDatasourceId *string `json:"OfflineDatasourceId,omitempty" xml:"OfflineDatasourceId,omitempty"`
	// The retention period in days for offline tables created by FeatureStore. This setting does not affect existing tables registered with the RegisterTable operation. A default value of 0 means the tables are never automatically deleted.
	//
	// example:
	//
	// 90
	OfflineLifeCycle *int32 `json:"OfflineLifeCycle,omitempty" xml:"OfflineLifeCycle,omitempty"`
	// The ID of the online data source. The data source type must be **Hologres**, **Tablestore**, or **FeatureDB**. You can call the ListDatasources operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	OnlineDatasourceId *string `json:"OnlineDatasourceId,omitempty" xml:"OnlineDatasourceId,omitempty"`
	// The ID of the PAI workspace. You can call the [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 324
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) GetOfflineDatasourceId() *string {
	return s.OfflineDatasourceId
}

func (s *CreateProjectRequest) GetOfflineLifeCycle() *int32 {
	return s.OfflineLifeCycle
}

func (s *CreateProjectRequest) GetOnlineDatasourceId() *string {
	return s.OnlineDatasourceId
}

func (s *CreateProjectRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetOfflineDatasourceId(v string) *CreateProjectRequest {
	s.OfflineDatasourceId = &v
	return s
}

func (s *CreateProjectRequest) SetOfflineLifeCycle(v int32) *CreateProjectRequest {
	s.OfflineLifeCycle = &v
	return s
}

func (s *CreateProjectRequest) SetOnlineDatasourceId(v string) *CreateProjectRequest {
	s.OnlineDatasourceId = &v
	return s
}

func (s *CreateProjectRequest) SetWorkspaceId(v string) *CreateProjectRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
