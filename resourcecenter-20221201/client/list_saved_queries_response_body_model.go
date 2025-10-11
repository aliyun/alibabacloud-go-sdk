// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavedQueriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListSavedQueriesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListSavedQueriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSavedQueriesResponseBody
	GetRequestId() *string
	SetSavedQueries(v []*ListSavedQueriesResponseBodySavedQueries) *ListSavedQueriesResponseBody
	GetSavedQueries() []*ListSavedQueriesResponseBodySavedQueries
}

type ListSavedQueriesResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the custom query templates.
	SavedQueries []*ListSavedQueriesResponseBodySavedQueries `json:"SavedQueries,omitempty" xml:"SavedQueries,omitempty" type:"Repeated"`
}

func (s ListSavedQueriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSavedQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListSavedQueriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSavedQueriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSavedQueriesResponseBody) GetSavedQueries() []*ListSavedQueriesResponseBodySavedQueries {
	return s.SavedQueries
}

func (s *ListSavedQueriesResponseBody) SetMaxResults(v string) *ListSavedQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetNextToken(v string) *ListSavedQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetRequestId(v string) *ListSavedQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetSavedQueries(v []*ListSavedQueriesResponseBodySavedQueries) *ListSavedQueriesResponseBody {
	s.SavedQueries = v
	return s
}

func (s *ListSavedQueriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSavedQueriesResponseBodySavedQueries struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-13T05:50:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template name.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-14 10:27:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSavedQueriesResponseBodySavedQueries) String() string {
	return dara.Prettify(s)
}

func (s ListSavedQueriesResponseBodySavedQueries) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBodySavedQueries) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSavedQueriesResponseBodySavedQueries) GetDescription() *string {
	return s.Description
}

func (s *ListSavedQueriesResponseBodySavedQueries) GetName() *string {
	return s.Name
}

func (s *ListSavedQueriesResponseBodySavedQueries) GetQueryId() *string {
	return s.QueryId
}

func (s *ListSavedQueriesResponseBodySavedQueries) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetCreateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.CreateTime = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetDescription(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Description = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetName(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Name = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetQueryId(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.QueryId = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetUpdateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.UpdateTime = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) Validate() error {
	return dara.Validate(s)
}
