// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInstanceInspectionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstanceInspectionsRequest
	GetRegionId() *string
}

type ListInstanceInspectionsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListInstanceInspectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceInspectionsRequest) SetMaxResults(v int32) *ListInstanceInspectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionsRequest) SetNextToken(v string) *ListInstanceInspectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionsRequest) SetRegionId(v string) *ListInstanceInspectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceInspectionsRequest) Validate() error {
	return dara.Validate(s)
}
