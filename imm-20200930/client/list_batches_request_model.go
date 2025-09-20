// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListBatchesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBatchesRequest
	GetNextToken() *string
	SetOrder(v string) *ListBatchesRequest
	GetOrder() *string
	SetProjectName(v string) *ListBatchesRequest
	GetProjectName() *string
	SetSort(v string) *ListBatchesRequest
	GetSort() *string
	SetState(v string) *ListBatchesRequest
	GetState() *string
	SetTagSelector(v string) *ListBatchesRequest
	GetTagSelector() *string
}

type ListBatchesRequest struct {
	// The maximum number of results to return. Valid values: 0 to 100.
	//
	// If you do not specify this parameter or set the parameter to 0, the default value of 100 is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// The pagination token is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter. The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// You do not need to specify this parameter in your initial request.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values:
	//
	// 	- ASC: sorts the results in ascending order. This is the default sort order.
	//
	// 	- DES: sorts the results in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The sort field. Valid values:
	//
	// 	- CreateTime
	//
	// 	- UpdateTime
	//
	// example:
	//
	// 2020-11-11T06:51:17.5Z
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The task status.
	//
	// 	- Ready: The task is newly created and ready.
	//
	// 	- Running: The task is running.
	//
	// 	- Failed: The task failed and cannot be automatically recovered.
	//
	// 	- Suspended: The task is suspended.
	//
	// 	- Succeeded: The task is successful.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The custom tag. You can use this parameter to query tasks that have the specified tag.
	//
	// example:
	//
	// test=val1
	TagSelector *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
}

func (s ListBatchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchesRequest) GoString() string {
	return s.String()
}

func (s *ListBatchesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListBatchesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListBatchesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListBatchesRequest) GetState() *string {
	return s.State
}

func (s *ListBatchesRequest) GetTagSelector() *string {
	return s.TagSelector
}

func (s *ListBatchesRequest) SetMaxResults(v int32) *ListBatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchesRequest) SetNextToken(v string) *ListBatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchesRequest) SetOrder(v string) *ListBatchesRequest {
	s.Order = &v
	return s
}

func (s *ListBatchesRequest) SetProjectName(v string) *ListBatchesRequest {
	s.ProjectName = &v
	return s
}

func (s *ListBatchesRequest) SetSort(v string) *ListBatchesRequest {
	s.Sort = &v
	return s
}

func (s *ListBatchesRequest) SetState(v string) *ListBatchesRequest {
	s.State = &v
	return s
}

func (s *ListBatchesRequest) SetTagSelector(v string) *ListBatchesRequest {
	s.TagSelector = &v
	return s
}

func (s *ListBatchesRequest) Validate() error {
	return dara.Validate(s)
}
