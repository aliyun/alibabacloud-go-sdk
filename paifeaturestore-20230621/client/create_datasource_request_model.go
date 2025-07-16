// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateDatasourceRequest
	GetConfig() *string
	SetName(v string) *CreateDatasourceRequest
	GetName() *string
	SetType(v string) *CreateDatasourceRequest
	GetType() *string
	SetUri(v string) *CreateDatasourceRequest
	GetUri() *string
	SetWorkspaceId(v string) *CreateDatasourceRequest
	GetWorkspaceId() *string
}

type CreateDatasourceRequest struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasourceRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateDatasourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatasourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDatasourceRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateDatasourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasourceRequest) SetConfig(v string) *CreateDatasourceRequest {
	s.Config = &v
	return s
}

func (s *CreateDatasourceRequest) SetName(v string) *CreateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasourceRequest) SetType(v string) *CreateDatasourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDatasourceRequest) SetUri(v string) *CreateDatasourceRequest {
	s.Uri = &v
	return s
}

func (s *CreateDatasourceRequest) SetWorkspaceId(v string) *CreateDatasourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
