// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeProcessListResponseBodyData) *DescribeProcessListResponseBody
	GetData() *DescribeProcessListResponseBodyData
	SetRequestId(v string) *DescribeProcessListResponseBody
	GetRequestId() *string
}

type DescribeProcessListResponseBody struct {
	// The data returned.
	Data *DescribeProcessListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBody) GetData() *DescribeProcessListResponseBodyData {
	return s.Data
}

func (s *DescribeProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProcessListResponseBody) SetData(v *DescribeProcessListResponseBodyData) *DescribeProcessListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProcessListResponseBody) SetRequestId(v string) *DescribeProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProcessListResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-xxxx
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The result sets.
	ResultSet []*DescribeProcessListResponseBodyDataResultSet `json:"ResultSet,omitempty" xml:"ResultSet,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProcessListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeProcessListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeProcessListResponseBodyData) GetResultSet() []*DescribeProcessListResponseBodyDataResultSet {
	return s.ResultSet
}

func (s *DescribeProcessListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProcessListResponseBodyData) SetDBInstanceID(v int32) *DescribeProcessListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetDBInstanceName(v string) *DescribeProcessListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetResultSet(v []*DescribeProcessListResponseBodyDataResultSet) *DescribeProcessListResponseBodyData {
	s.ResultSet = v
	return s
}

func (s *DescribeProcessListResponseBodyData) SetTotalCount(v int32) *DescribeProcessListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeProcessListResponseBodyData) Validate() error {
	if s.ResultSet != nil {
		for _, item := range s.ResultSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProcessListResponseBodyDataResultSet struct {
	// The address to which the query statement is sent.
	//
	// example:
	//
	// 0:0:0:0:0:ffff:1edd65ea
	InitialAddress *string `json:"InitialAddress,omitempty" xml:"InitialAddress,omitempty"`
	// The query ID.
	//
	// example:
	//
	// \\"79f7e40b-87e2-4ef4-b6df-21889a3a030e\\"
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	// The user who executes the query statement.
	//
	// example:
	//
	// bany
	InitialUser *string `json:"InitialUser,omitempty" xml:"InitialUser,omitempty"`
	// The query statement that is running.
	//
	// example:
	//
	// select 	- from test
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The minimum query duration. Minimum value: **1000**. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	QueryDurationMs *int64 `json:"QueryDurationMs,omitempty" xml:"QueryDurationMs,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-07-23T10:13:23Z
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
}

func (s DescribeProcessListResponseBodyDataResultSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponseBodyDataResultSet) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetInitialAddress() *string {
	return s.InitialAddress
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetInitialUser() *string {
	return s.InitialUser
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetQuery() *string {
	return s.Query
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetQueryDurationMs() *int64 {
	return s.QueryDurationMs
}

func (s *DescribeProcessListResponseBodyDataResultSet) GetQueryStartTime() *string {
	return s.QueryStartTime
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialAddress(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialAddress = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialQueryId(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialQueryId = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetInitialUser(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.InitialUser = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQuery(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.Query = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQueryDurationMs(v int64) *DescribeProcessListResponseBodyDataResultSet {
	s.QueryDurationMs = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) SetQueryStartTime(v string) *DescribeProcessListResponseBodyDataResultSet {
	s.QueryStartTime = &v
	return s
}

func (s *DescribeProcessListResponseBodyDataResultSet) Validate() error {
	return dara.Validate(s)
}
