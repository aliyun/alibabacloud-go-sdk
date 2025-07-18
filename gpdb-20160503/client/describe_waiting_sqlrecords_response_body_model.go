// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeWaitingSQLRecordsResponseBodyItems) *DescribeWaitingSQLRecordsResponseBody
	GetItems() []*DescribeWaitingSQLRecordsResponseBodyItems
	SetPageNumber(v int32) *DescribeWaitingSQLRecordsResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeWaitingSQLRecordsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWaitingSQLRecordsResponseBody
	GetTotalCount() *int32
}

type DescribeWaitingSQLRecordsResponseBody struct {
	// The list of lock diagnostics records.
	Items []*DescribeWaitingSQLRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBody) GetItems() []*DescribeWaitingSQLRecordsResponseBodyItems {
	return s.Items
}

func (s *DescribeWaitingSQLRecordsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWaitingSQLRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWaitingSQLRecordsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetItems(v []*DescribeWaitingSQLRecordsResponseBodyItems) *DescribeWaitingSQLRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetPageNumber(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetRequestId(v string) *DescribeWaitingSQLRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) SetTotalCount(v int32) *DescribeWaitingSQLRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWaitingSQLRecordsResponseBodyItems struct {
	// The name of the database.
	//
	// example:
	//
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The ID of the process that uniquely identifies the query.
	//
	// example:
	//
	// 100
	PID *string `json:"PID,omitempty" xml:"PID,omitempty"`
	// The SQL statement of the query.
	//
	// example:
	//
	// Select 	- from t1,t2 where t1.id=t2.id;
	SQLStmt *string `json:"SQLStmt,omitempty" xml:"SQLStmt,omitempty"`
	// The ID of the session that contains the query.
	//
	// example:
	//
	// 50
	SessionID *string `json:"SessionID,omitempty" xml:"SessionID,omitempty"`
	// The start time of the query. This value is in the timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1660902033374
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The waiting state of the query. Valid values:
	//
	// 	- **LockWaiting**
	//
	// 	- **ResourceWaiting**
	//
	// example:
	//
	// LockWaiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The database account that is used to perform the query.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The waiting period of the query. Unit: milliseconds.
	//
	// example:
	//
	// 26911000
	WaitingTime *int64 `json:"WaitingTime,omitempty" xml:"WaitingTime,omitempty"`
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetDatabase() *string {
	return s.Database
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetPID() *string {
	return s.PID
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetSQLStmt() *string {
	return s.SQLStmt
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetSessionID() *string {
	return s.SessionID
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) GetWaitingTime() *int64 {
	return s.WaitingTime
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetDatabase(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetPID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetSessionID(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.SessionID = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStartTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetStatus(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetUser(v string) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) SetWaitingTime(v int64) *DescribeWaitingSQLRecordsResponseBodyItems {
	s.WaitingTime = &v
	return s
}

func (s *DescribeWaitingSQLRecordsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
