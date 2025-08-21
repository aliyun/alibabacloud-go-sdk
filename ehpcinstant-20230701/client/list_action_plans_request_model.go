// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanIds(v []*string) *ListActionPlansRequest
	GetActionPlanIds() []*string
	SetMaxResults(v int32) *ListActionPlansRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionPlansRequest
	GetNextToken() *string
}

type ListActionPlansRequest struct {
	ActionPlanIds []*string `json:"ActionPlanIds,omitempty" xml:"ActionPlanIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListActionPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlansRequest) GoString() string {
	return s.String()
}

func (s *ListActionPlansRequest) GetActionPlanIds() []*string {
	return s.ActionPlanIds
}

func (s *ListActionPlansRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionPlansRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionPlansRequest) SetActionPlanIds(v []*string) *ListActionPlansRequest {
	s.ActionPlanIds = v
	return s
}

func (s *ListActionPlansRequest) SetMaxResults(v int32) *ListActionPlansRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionPlansRequest) SetNextToken(v string) *ListActionPlansRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionPlansRequest) Validate() error {
	return dara.Validate(s)
}
