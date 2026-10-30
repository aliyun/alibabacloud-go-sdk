// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationScope(v string) *ListModelPermissionsRequest
	GetAuthorizationScope() *string
	SetFilter(v *ListModelPermissionsRequestFilter) *ListModelPermissionsRequest
	GetFilter() *ListModelPermissionsRequestFilter
	SetMaxResults(v int32) *ListModelPermissionsRequest
	GetMaxResults() *int32
	SetModelAction(v string) *ListModelPermissionsRequest
	GetModelAction() *string
	SetNextToken(v string) *ListModelPermissionsRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListModelPermissionsRequest
	GetWorkspaceId() *string
}

type ListModelPermissionsRequest struct {
	// The query dimension for the authorization scope.
	//
	// - **AUTHORIZED**: models that are authorized for the specified modelAction. This value must be used together with modelAction.
	//
	// - **AUTHORIZABLE**: all authorizable directories.
	//
	// example:
	//
	// AUTHORIZABLE
	AuthorizationScope *string `json:"authorizationScope,omitempty" xml:"authorizationScope,omitempty"`
	// The filter conditions.
	Filter *ListModelPermissionsRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The maximum number of entries to return on each page. Default value: 20. If the value exceeds the upper limit, the `InvalidParameter.maxResults` error is returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The authorization action dimension.
	//
	// - **INFERENCE**: model invocation authorization.
	//
	// example:
	//
	// INFERENCE
	ModelAction *string `json:"modelAction,omitempty" xml:"modelAction,omitempty"`
	// The pagination token that is used to retrieve the next page. This parameter is not required for the first page.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the workspace. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-32klhjk2312334jkh
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListModelPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsRequest) GetAuthorizationScope() *string {
	return s.AuthorizationScope
}

func (s *ListModelPermissionsRequest) GetFilter() *ListModelPermissionsRequestFilter {
	return s.Filter
}

func (s *ListModelPermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelPermissionsRequest) GetModelAction() *string {
	return s.ModelAction
}

func (s *ListModelPermissionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelPermissionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelPermissionsRequest) SetAuthorizationScope(v string) *ListModelPermissionsRequest {
	s.AuthorizationScope = &v
	return s
}

func (s *ListModelPermissionsRequest) SetFilter(v *ListModelPermissionsRequestFilter) *ListModelPermissionsRequest {
	s.Filter = v
	return s
}

func (s *ListModelPermissionsRequest) SetMaxResults(v int32) *ListModelPermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelPermissionsRequest) SetModelAction(v string) *ListModelPermissionsRequest {
	s.ModelAction = &v
	return s
}

func (s *ListModelPermissionsRequest) SetNextToken(v string) *ListModelPermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelPermissionsRequest) SetWorkspaceId(v string) *ListModelPermissionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelPermissionsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelPermissionsRequestFilter struct {
	// The exact match for a single model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The fuzzy match for the model name.
	//
	// example:
	//
	// OVERLAY
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListModelPermissionsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsRequestFilter) GetModel() *string {
	return s.Model
}

func (s *ListModelPermissionsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListModelPermissionsRequestFilter) SetModel(v string) *ListModelPermissionsRequestFilter {
	s.Model = &v
	return s
}

func (s *ListModelPermissionsRequestFilter) SetName(v string) *ListModelPermissionsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListModelPermissionsRequestFilter) Validate() error {
	return dara.Validate(s)
}
