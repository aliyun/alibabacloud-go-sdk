// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPolarClawSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SearchPolarClawSkillsRequest
	GetApplicationId() *string
	SetLimit(v int32) *SearchPolarClawSkillsRequest
	GetLimit() *int32
	SetQuery(v string) *SearchPolarClawSkillsRequest
	GetQuery() *string
}

type SearchPolarClawSkillsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The search keyword. If this parameter is not specified, popular or recommended results are returned.
	//
	// example:
	//
	// rds
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s SearchPolarClawSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchPolarClawSkillsRequest) GoString() string {
	return s.String()
}

func (s *SearchPolarClawSkillsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SearchPolarClawSkillsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchPolarClawSkillsRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchPolarClawSkillsRequest) SetApplicationId(v string) *SearchPolarClawSkillsRequest {
	s.ApplicationId = &v
	return s
}

func (s *SearchPolarClawSkillsRequest) SetLimit(v int32) *SearchPolarClawSkillsRequest {
	s.Limit = &v
	return s
}

func (s *SearchPolarClawSkillsRequest) SetQuery(v string) *SearchPolarClawSkillsRequest {
	s.Query = &v
	return s
}

func (s *SearchPolarClawSkillsRequest) Validate() error {
	return dara.Validate(s)
}
