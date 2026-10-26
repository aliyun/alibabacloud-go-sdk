// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInspectionTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInspectionTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInspectionTemplatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInspectionTemplatesRequest
	GetRegionId() *string
}

type ListInspectionTemplatesRequest struct {
	// example:
	//
	// 20
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

func (s ListInspectionTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInspectionTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListInspectionTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInspectionTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInspectionTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInspectionTemplatesRequest) SetMaxResults(v int32) *ListInspectionTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInspectionTemplatesRequest) SetNextToken(v string) *ListInspectionTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInspectionTemplatesRequest) SetRegionId(v string) *ListInspectionTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInspectionTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
