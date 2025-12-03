// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListWorkitemsRequest
	GetCategory() *string
	SetConditions(v string) *ListWorkitemsRequest
	GetConditions() *string
	SetExtraConditions(v string) *ListWorkitemsRequest
	GetExtraConditions() *string
	SetGroupCondition(v string) *ListWorkitemsRequest
	GetGroupCondition() *string
	SetMaxResults(v string) *ListWorkitemsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListWorkitemsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListWorkitemsRequest
	GetOrderBy() *string
	SetSearchType(v string) *ListWorkitemsRequest
	GetSearchType() *string
	SetSpaceIdentifier(v string) *ListWorkitemsRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *ListWorkitemsRequest
	GetSpaceType() *string
}

type ListWorkitemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Req
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// {"conditionGroups":[]}
	Conditions *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	// example:
	//
	// {"conditionGroups":[]}
	ExtraConditions *string `json:"extraConditions,omitempty" xml:"extraConditions,omitempty"`
	// example:
	//
	// {"fieldIdentifier":"tag","className":"tag","format":"multiList","value":["c76e0e4bf64801cfad73......"],"operator":"EQUALS"}
	GroupCondition *string `json:"groupCondition,omitempty" xml:"groupCondition,omitempty"`
	// example:
	//
	// 20
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// {"fieldIdentifier":"status","format":"list","order":"desc","className":"status"}
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// LIST
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8fb83debd69a6c7c6626......
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListWorkitemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkitemsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListWorkitemsRequest) GetConditions() *string {
	return s.Conditions
}

func (s *ListWorkitemsRequest) GetExtraConditions() *string {
	return s.ExtraConditions
}

func (s *ListWorkitemsRequest) GetGroupCondition() *string {
	return s.GroupCondition
}

func (s *ListWorkitemsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListWorkitemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkitemsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListWorkitemsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListWorkitemsRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListWorkitemsRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListWorkitemsRequest) SetCategory(v string) *ListWorkitemsRequest {
	s.Category = &v
	return s
}

func (s *ListWorkitemsRequest) SetConditions(v string) *ListWorkitemsRequest {
	s.Conditions = &v
	return s
}

func (s *ListWorkitemsRequest) SetExtraConditions(v string) *ListWorkitemsRequest {
	s.ExtraConditions = &v
	return s
}

func (s *ListWorkitemsRequest) SetGroupCondition(v string) *ListWorkitemsRequest {
	s.GroupCondition = &v
	return s
}

func (s *ListWorkitemsRequest) SetMaxResults(v string) *ListWorkitemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkitemsRequest) SetNextToken(v string) *ListWorkitemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkitemsRequest) SetOrderBy(v string) *ListWorkitemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkitemsRequest) SetSearchType(v string) *ListWorkitemsRequest {
	s.SearchType = &v
	return s
}

func (s *ListWorkitemsRequest) SetSpaceIdentifier(v string) *ListWorkitemsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListWorkitemsRequest) SetSpaceType(v string) *ListWorkitemsRequest {
	s.SpaceType = &v
	return s
}

func (s *ListWorkitemsRequest) Validate() error {
	return dara.Validate(s)
}
