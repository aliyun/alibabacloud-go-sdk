// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplateCacheRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplateCacheRequest
	GetNextToken() *string
	SetStatus(v string) *ListTemplateCacheRequest
	GetStatus() *string
	SetTeamID(v string) *ListTemplateCacheRequest
	GetTeamID() *string
	SetTemplateID(v string) *ListTemplateCacheRequest
	GetTemplateID() *string
}

type ListTemplateCacheRequest struct {
	// The maximum number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned from the previous response.
	//
	// example:
	//
	// cae5f900-8b1d-4c0e-9c2a-1a2b3c4d5e6f
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters by cache status. Valid values:
	//
	// - InProgress
	//
	// - Success
	//
	// - Failed
	//
	// - Deleting
	//
	// - Evicted
	//
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The team ID.
	//
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The unique identifier of the template.
	//
	// example:
	//
	// us7dxqaezw5uu7aa2cm5
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s ListTemplateCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateCacheRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateCacheRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplateCacheRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplateCacheRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTemplateCacheRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListTemplateCacheRequest) GetTemplateID() *string {
	return s.TemplateID
}

func (s *ListTemplateCacheRequest) SetMaxResults(v int32) *ListTemplateCacheRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateCacheRequest) SetNextToken(v string) *ListTemplateCacheRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateCacheRequest) SetStatus(v string) *ListTemplateCacheRequest {
	s.Status = &v
	return s
}

func (s *ListTemplateCacheRequest) SetTeamID(v string) *ListTemplateCacheRequest {
	s.TeamID = &v
	return s
}

func (s *ListTemplateCacheRequest) SetTemplateID(v string) *ListTemplateCacheRequest {
	s.TemplateID = &v
	return s
}

func (s *ListTemplateCacheRequest) Validate() error {
	return dara.Validate(s)
}
