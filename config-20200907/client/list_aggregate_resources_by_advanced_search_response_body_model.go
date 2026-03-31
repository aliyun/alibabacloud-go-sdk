// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourcesByAdvancedSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryResults(v *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) *ListAggregateResourcesByAdvancedSearchResponseBody
	GetQueryResults() *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults
	SetRequestId(v string) *ListAggregateResourcesByAdvancedSearchResponseBody
	GetRequestId() *string
}

type ListAggregateResourcesByAdvancedSearchResponseBody struct {
	// The returned data.
	QueryResults *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults `json:"QueryResults,omitempty" xml:"QueryResults,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B0DBF868-460F-5E7C-8F76-1ACE2FCCE153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateResourcesByAdvancedSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourcesByAdvancedSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBody) GetQueryResults() *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults {
	return s.QueryResults
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBody) SetQueryResults(v *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) *ListAggregateResourcesByAdvancedSearchResponseBody {
	s.QueryResults = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBody) SetRequestId(v string) *ListAggregateResourcesByAdvancedSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBody) Validate() error {
	if s.QueryResults != nil {
		if err := s.QueryResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults struct {
	// The queried resources. A maximum of 1,000 data records can be returned. To view more data, use the download URL of the resource file.
	QueryResultList *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList `json:"QueryResultList,omitempty" xml:"QueryResultList,omitempty" type:"Struct"`
}

func (s ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) GoString() string {
	return s.String()
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) GetQueryResultList() *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	return s.QueryResultList
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) SetQueryResultList(v *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults {
	s.QueryResultList = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResults) Validate() error {
	if s.QueryResultList != nil {
		if err := s.QueryResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList struct {
	// The field names.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The resource data.
	Values []interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GoString() string {
	return s.String()
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GetColumns() []*string {
	return s.Columns
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GetValues() []interface{} {
	return s.Values
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) SetColumns(v []*string) *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	s.Columns = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) SetValues(v []interface{}) *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	s.Values = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) Validate() error {
	return dara.Validate(s)
}
