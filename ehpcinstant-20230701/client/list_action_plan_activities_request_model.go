// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlanActivitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanId(v string) *ListActionPlanActivitiesRequest
	GetActionPlanId() *string
	SetMaxResults(v int32) *ListActionPlanActivitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionPlanActivitiesRequest
	GetNextToken() *string
}

type ListActionPlanActivitiesRequest struct {
	// The ID of the execution plan.
	//
	// example:
	//
	// ap-hz036ubmx2qmw93k****
	ActionPlanId *string `json:"ActionPlanId,omitempty" xml:"ActionPlanId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListActionPlanActivitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlanActivitiesRequest) GoString() string {
	return s.String()
}

func (s *ListActionPlanActivitiesRequest) GetActionPlanId() *string {
	return s.ActionPlanId
}

func (s *ListActionPlanActivitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionPlanActivitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionPlanActivitiesRequest) SetActionPlanId(v string) *ListActionPlanActivitiesRequest {
	s.ActionPlanId = &v
	return s
}

func (s *ListActionPlanActivitiesRequest) SetMaxResults(v int32) *ListActionPlanActivitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionPlanActivitiesRequest) SetNextToken(v string) *ListActionPlanActivitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionPlanActivitiesRequest) Validate() error {
	return dara.Validate(s)
}
