// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndexName(v string) *ListIndicesRequest
	GetIndexName() *string
	SetPageNumber(v string) *ListIndicesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListIndicesRequest
	GetPageSize() *string
}

type ListIndicesRequest struct {
	// The name of the knowledge base. You can use this parameter to search for a knowledge base by name. The name must be 1 to 20 characters in length and can contain characters classified as letters in Unicode (including English letters, Chinese characters, and digits). The name can also contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Default value: empty, which queries all knowledge bases in the specified workspace.
	//
	// example:
	//
	// idx_status_score
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of knowledge bases to display per page in a paging query. No maximum limit.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIndicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListIndicesRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *ListIndicesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListIndicesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListIndicesRequest) SetIndexName(v string) *ListIndicesRequest {
	s.IndexName = &v
	return s
}

func (s *ListIndicesRequest) SetPageNumber(v string) *ListIndicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIndicesRequest) SetPageSize(v string) *ListIndicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIndicesRequest) Validate() error {
	return dara.Validate(s)
}
