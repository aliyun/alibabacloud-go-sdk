// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceTaskRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTaskRequest
	GetNextToken() *string
	SetSearchCondition(v string) *ListServiceTaskRequest
	GetSearchCondition() *string
	SetType(v string) *ListServiceTaskRequest
	GetType() *string
}

type ListServiceTaskRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// d23d8f3f0f0cd1984566b1986c9343122fa0385a05c09694c17fe87709f3eb56d1a7ead56b4a2536
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// {"ip":"10.0.0.1"}
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// example:
	//
	// live_debug_log_probe
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListServiceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTaskRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTaskRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTaskRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTaskRequest) GetSearchCondition() *string {
	return s.SearchCondition
}

func (s *ListServiceTaskRequest) GetType() *string {
	return s.Type
}

func (s *ListServiceTaskRequest) SetMaxResults(v int32) *ListServiceTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTaskRequest) SetNextToken(v string) *ListServiceTaskRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTaskRequest) SetSearchCondition(v string) *ListServiceTaskRequest {
	s.SearchCondition = &v
	return s
}

func (s *ListServiceTaskRequest) SetType(v string) *ListServiceTaskRequest {
	s.Type = &v
	return s
}

func (s *ListServiceTaskRequest) Validate() error {
	return dara.Validate(s)
}
