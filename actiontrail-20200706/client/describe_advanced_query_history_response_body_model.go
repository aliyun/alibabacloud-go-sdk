// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedQueryHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryHistoryList(v []*DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) *DescribeAdvancedQueryHistoryResponseBody
	GetQueryHistoryList() []*DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList
	SetRequestId(v string) *DescribeAdvancedQueryHistoryResponseBody
	GetRequestId() *string
}

type DescribeAdvancedQueryHistoryResponseBody struct {
	QueryHistoryList []*DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList `json:"QueryHistoryList,omitempty" xml:"QueryHistoryList,omitempty" type:"Repeated"`
	// example:
	//
	// 19F032B7-5FD8-5AC9-97FD-ACF54371****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvancedQueryHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryHistoryResponseBody) GetQueryHistoryList() []*DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList {
	return s.QueryHistoryList
}

func (s *DescribeAdvancedQueryHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvancedQueryHistoryResponseBody) SetQueryHistoryList(v []*DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) *DescribeAdvancedQueryHistoryResponseBody {
	s.QueryHistoryList = v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBody) SetRequestId(v string) *DescribeAdvancedQueryHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBody) Validate() error {
	if s.QueryHistoryList != nil {
		for _, item := range s.QueryHistoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList struct {
	// example:
	//
	// query-uIkIvLiVSuCKqg0yoa****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// event.userIdentity.accessKeyId: *
	QuerySql *string `json:"QuerySql,omitempty" xml:"QuerySql,omitempty"`
	// example:
	//
	// false
	SimpleQuery *bool `json:"SimpleQuery,omitempty" xml:"SimpleQuery,omitempty"`
	// example:
	//
	// 1753695874000
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) GetQuerySql() *string {
	return s.QuerySql
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) GetSimpleQuery() *bool {
	return s.SimpleQuery
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) SetQueryId(v string) *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList {
	s.QueryId = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) SetQuerySql(v string) *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList {
	s.QuerySql = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) SetSimpleQuery(v bool) *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList {
	s.SimpleQuery = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) SetTimeStamp(v string) *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList {
	s.TimeStamp = &v
	return s
}

func (s *DescribeAdvancedQueryHistoryResponseBodyQueryHistoryList) Validate() error {
	return dara.Validate(s)
}
