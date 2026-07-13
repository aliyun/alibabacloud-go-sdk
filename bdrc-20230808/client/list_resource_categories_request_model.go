// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceCategoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceCategoriesRequest
	GetNextToken() *string
	SetResourceCategoryId(v string) *ListResourceCategoriesRequest
	GetResourceCategoryId() *string
}

type ListResourceCategoriesRequest struct {
	// The maximum number of records to return in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If more entries are to be returned on the next page, a pagination token is returned. Note: If this parameter returns data, it indicates there is a next page. You can use the returned NextToken as a request parameter to obtain the next page of data until it returns Null, indicating all data has been retrieved.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rc-123****7890
	ResourceCategoryId *string `json:"ResourceCategoryId,omitempty" xml:"ResourceCategoryId,omitempty"`
}

func (s ListResourceCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceCategoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceCategoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceCategoriesRequest) GetResourceCategoryId() *string {
	return s.ResourceCategoryId
}

func (s *ListResourceCategoriesRequest) SetMaxResults(v int32) *ListResourceCategoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceCategoriesRequest) SetNextToken(v string) *ListResourceCategoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceCategoriesRequest) SetResourceCategoryId(v string) *ListResourceCategoriesRequest {
	s.ResourceCategoryId = &v
	return s
}

func (s *ListResourceCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
