// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCustomAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomAgentsRequest
	GetNextToken() *string
	SetQ(v string) *ListCustomAgentsRequest
	GetQ() *string
	SetVisibility(v []*string) *ListCustomAgentsRequest
	GetVisibility() []*string
}

type ListCustomAgentsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// analysis
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// example:
	//
	// -
	Visibility []*string `json:"Visibility,omitempty" xml:"Visibility,omitempty" type:"Repeated"`
}

func (s ListCustomAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomAgentsRequest) GetQ() *string {
	return s.Q
}

func (s *ListCustomAgentsRequest) GetVisibility() []*string {
	return s.Visibility
}

func (s *ListCustomAgentsRequest) SetMaxResults(v int32) *ListCustomAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsRequest) SetNextToken(v string) *ListCustomAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsRequest) SetQ(v string) *ListCustomAgentsRequest {
	s.Q = &v
	return s
}

func (s *ListCustomAgentsRequest) SetVisibility(v []*string) *ListCustomAgentsRequest {
	s.Visibility = v
	return s
}

func (s *ListCustomAgentsRequest) Validate() error {
	return dara.Validate(s)
}
