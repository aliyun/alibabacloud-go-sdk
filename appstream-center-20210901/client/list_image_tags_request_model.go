// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListImageTagsRequest
	GetCategory() *string
	SetMaxResults(v int32) *ListImageTagsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListImageTagsRequest
	GetNextToken() *string
}

type ListImageTagsRequest struct {
	// The label categorization to query.
	//
	// example:
	//
	// Gpu
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// A compatibility field. The current version does not use this field to specify the number of returned entries.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to determine whether a next query is available. This parameter is not currently supported.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kEmoGlTNe3ekUrwnesWV1Vn+tWdc+4WoaoMP/kUNF7Gg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListImageTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageTagsRequest) GoString() string {
	return s.String()
}

func (s *ListImageTagsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListImageTagsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImageTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImageTagsRequest) SetCategory(v string) *ListImageTagsRequest {
	s.Category = &v
	return s
}

func (s *ListImageTagsRequest) SetMaxResults(v int32) *ListImageTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImageTagsRequest) SetNextToken(v string) *ListImageTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListImageTagsRequest) Validate() error {
	return dara.Validate(s)
}
