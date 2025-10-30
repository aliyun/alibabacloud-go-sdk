// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDiagnosisRecordsResponseBodyItems) *DescribeDiagnosisRecordsResponseBody
	GetItems() []*DescribeDiagnosisRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeDiagnosisRecordsResponseBody struct {
	// The threshold that determines whether the SQL statement must be truncated. The value is the number of characters.
	Items []*DescribeDiagnosisRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBody) GetItems() []*DescribeDiagnosisRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeDiagnosisRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDiagnosisRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiagnosisRecordsResponseBody) SetItems(v []*DescribeDiagnosisRecordsResponseBodyItems) *DescribeDiagnosisRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetPageNumber(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetRequestId(v string) *DescribeDiagnosisRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) SetTotalCount(v int32) *DescribeDiagnosisRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosisRecordsResponseBodyItems struct {
	// The name of the database.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The execution duration of the query. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the query. It is a unique identifier of the query.
	//
	// example:
	//
	// 2022042612465401000000012903151998970
	QueryID *string `json:"QueryID,omitempty" xml:"QueryID,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM t1,t2 WHERE t1.id=t2.id;
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// Indicates whether the SQL statement needs to be truncated. Valid values:
	//
	// 	- **true**: The SQL statement needs to be truncated.
	//
	// 	- **false**: The SQL statement does not need to be truncated.
	//
	// example:
	//
	// false
	SQLTruncated *bool `json:"SQLTruncated,omitempty" xml:"SQLTruncated,omitempty"`
	// The threshold used to determine whether an SQL statement must be truncated. The value is the number of characters.
	//
	// example:
	//
	// 5120
	SQLTruncatedThreshold *int32 `json:"SQLTruncatedThreshold,omitempty" xml:"SQLTruncatedThreshold,omitempty"`
	// The ID of the session that contains the query.
	//
	// example:
	//
	// 50
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The start time of the query. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1651877940000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution state of the query. Valid values:
	//
	// 	- **running**: The query is being executed.
	//
	// 	- **finished**: The query is complete.
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDiagnosisRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetQueryID() *string {
	return s.QueryID
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetSQLStmt() *string {
	return s.SQLStmt
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetSQLTruncated() *bool {
	return s.SQLTruncated
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetSQLTruncatedThreshold() *int32 {
	return s.SQLTruncatedThreshold
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetSessionID() *string {
	return s.SessionID
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDatabase(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetDuration(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Duration = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetQueryID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.QueryID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncated(v bool) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncated = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSQLTruncatedThreshold(v int32) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SQLTruncatedThreshold = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetSessionID(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStartTime(v int64) *DescribeDiagnosisRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetStatus(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) SetUser(v string) *DescribeDiagnosisRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeDiagnosisRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
