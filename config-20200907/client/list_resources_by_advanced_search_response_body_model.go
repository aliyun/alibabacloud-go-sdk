// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByAdvancedSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryResults(v *ListResourcesByAdvancedSearchResponseBodyQueryResults) *ListResourcesByAdvancedSearchResponseBody
	GetQueryResults() *ListResourcesByAdvancedSearchResponseBodyQueryResults
	SetRequestId(v string) *ListResourcesByAdvancedSearchResponseBody
	GetRequestId() *string
}

type ListResourcesByAdvancedSearchResponseBody struct {
	QueryResults *ListResourcesByAdvancedSearchResponseBodyQueryResults `json:"QueryResults,omitempty" xml:"QueryResults,omitempty" type:"Struct"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesByAdvancedSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByAdvancedSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesByAdvancedSearchResponseBody) GetQueryResults() *ListResourcesByAdvancedSearchResponseBodyQueryResults {
	return s.QueryResults
}

func (s *ListResourcesByAdvancedSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesByAdvancedSearchResponseBody) SetQueryResults(v *ListResourcesByAdvancedSearchResponseBodyQueryResults) *ListResourcesByAdvancedSearchResponseBody {
	s.QueryResults = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponseBody) SetRequestId(v string) *ListResourcesByAdvancedSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesByAdvancedSearchResponseBody) Validate() error {
	if s.QueryResults != nil {
		if err := s.QueryResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesByAdvancedSearchResponseBodyQueryResults struct {
	QueryResultList *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList `json:"QueryResultList,omitempty" xml:"QueryResultList,omitempty" type:"Struct"`
}

func (s ListResourcesByAdvancedSearchResponseBodyQueryResults) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByAdvancedSearchResponseBodyQueryResults) GoString() string {
	return s.String()
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResults) GetQueryResultList() *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	return s.QueryResultList
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResults) SetQueryResultList(v *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) *ListResourcesByAdvancedSearchResponseBodyQueryResults {
	s.QueryResultList = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResults) Validate() error {
	if s.QueryResultList != nil {
		if err := s.QueryResultList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList struct {
	Columns []*string     `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Values  []interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GoString() string {
	return s.String()
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GetColumns() []*string {
	return s.Columns
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) GetValues() []interface{} {
	return s.Values
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) SetColumns(v []*string) *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	s.Columns = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) SetValues(v []interface{}) *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList {
	s.Values = v
	return s
}

func (s *ListResourcesByAdvancedSearchResponseBodyQueryResultsQueryResultList) Validate() error {
	return dara.Validate(s)
}
