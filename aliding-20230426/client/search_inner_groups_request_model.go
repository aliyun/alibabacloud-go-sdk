// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInnerGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *SearchInnerGroupsRequest
	GetMaxResults() *int32
	SetSearchKey(v string) *SearchInnerGroupsRequest
	GetSearchKey() *string
}

type SearchInnerGroupsRequest struct {
	// example:
	//
	// 100
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s SearchInnerGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsRequest) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchInnerGroupsRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchInnerGroupsRequest) SetMaxResults(v int32) *SearchInnerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchInnerGroupsRequest) SetSearchKey(v string) *SearchInnerGroupsRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchInnerGroupsRequest) Validate() error {
	return dara.Validate(s)
}
