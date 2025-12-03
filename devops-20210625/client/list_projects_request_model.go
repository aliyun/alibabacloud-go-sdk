// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListProjectsRequest
	GetCategory() *string
	SetConditions(v string) *ListProjectsRequest
	GetConditions() *string
	SetExtraConditions(v string) *ListProjectsRequest
	GetExtraConditions() *string
	SetMaxResults(v int64) *ListProjectsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListProjectsRequest
	GetNextToken() *string
	SetScope(v string) *ListProjectsRequest
	GetScope() *string
}

type ListProjectsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Project
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// {"conditionGroups":[[]]}
	Conditions      *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	ExtraConditions *string `json:"extraConditions,omitempty" xml:"extraConditions,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListProjectsRequest) GetConditions() *string {
	return s.Conditions
}

func (s *ListProjectsRequest) GetExtraConditions() *string {
	return s.ExtraConditions
}

func (s *ListProjectsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListProjectsRequest) SetCategory(v string) *ListProjectsRequest {
	s.Category = &v
	return s
}

func (s *ListProjectsRequest) SetConditions(v string) *ListProjectsRequest {
	s.Conditions = &v
	return s
}

func (s *ListProjectsRequest) SetExtraConditions(v string) *ListProjectsRequest {
	s.ExtraConditions = &v
	return s
}

func (s *ListProjectsRequest) SetMaxResults(v int64) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetScope(v string) *ListProjectsRequest {
	s.Scope = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
