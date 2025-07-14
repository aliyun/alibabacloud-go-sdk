// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiDatasourceParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ModifyApiDatasourceParametersRequest
	GetApiId() *string
	SetParameters(v string) *ModifyApiDatasourceParametersRequest
	GetParameters() *string
	SetWorkspaceId(v string) *ModifyApiDatasourceParametersRequest
	GetWorkspaceId() *string
}

type ModifyApiDatasourceParametersRequest struct {
	// The ID of the API data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// b66a66de51f24d149116c17718138194
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The configuration of API data parameters in the JSONArray format. You can modify a maximum of 10 parameters.
	//
	// 	- name: the name of a common parameter or a parameter in a query statement
	//
	// 	- value: the value of a common parameter or a parameter in a query statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"token","value":"xxxxxxxxxxxx"},{"name":"pageSize","value":100}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyApiDatasourceParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiDatasourceParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiDatasourceParametersRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ModifyApiDatasourceParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyApiDatasourceParametersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyApiDatasourceParametersRequest) SetApiId(v string) *ModifyApiDatasourceParametersRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApiDatasourceParametersRequest) SetParameters(v string) *ModifyApiDatasourceParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyApiDatasourceParametersRequest) SetWorkspaceId(v string) *ModifyApiDatasourceParametersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyApiDatasourceParametersRequest) Validate() error {
	return dara.Validate(s)
}
