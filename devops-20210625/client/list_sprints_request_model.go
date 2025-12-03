// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSprintsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListSprintsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListSprintsRequest
	GetNextToken() *string
	SetSpaceIdentifier(v string) *ListSprintsRequest
	GetSpaceIdentifier() *string
	SetSpaceType(v string) *ListSprintsRequest
	GetSpaceType() *string
}

type ListSprintsRequest struct {
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Project
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s ListSprintsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSprintsRequest) GoString() string {
	return s.String()
}

func (s *ListSprintsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListSprintsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSprintsRequest) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListSprintsRequest) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListSprintsRequest) SetMaxResults(v int64) *ListSprintsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSprintsRequest) SetNextToken(v string) *ListSprintsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSprintsRequest) SetSpaceIdentifier(v string) *ListSprintsRequest {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListSprintsRequest) SetSpaceType(v string) *ListSprintsRequest {
	s.SpaceType = &v
	return s
}

func (s *ListSprintsRequest) Validate() error {
	return dara.Validate(s)
}
