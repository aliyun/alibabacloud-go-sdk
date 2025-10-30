// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWaitingSQLInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v string) *DescribeWaitingSQLInfoResponseBody
	GetDatabase() *string
	SetItems(v []*DescribeWaitingSQLInfoResponseBodyItems) *DescribeWaitingSQLInfoResponseBody
	GetItems() []*DescribeWaitingSQLInfoResponseBodyItems
	SetRequestId(v string) *DescribeWaitingSQLInfoResponseBody
	GetRequestId() *string
}

type DescribeWaitingSQLInfoResponseBody struct {
	// The name of the database.
	//
	// example:
	//
	// test
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The queried lock-waiting query.
	Items []*DescribeWaitingSQLInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBody) GetDatabase() *string {
	return s.Database
}

func (s *DescribeWaitingSQLInfoResponseBody) GetItems() []*DescribeWaitingSQLInfoResponseBodyItems {
	return s.Items
}

func (s *DescribeWaitingSQLInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWaitingSQLInfoResponseBody) SetDatabase(v string) *DescribeWaitingSQLInfoResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetItems(v []*DescribeWaitingSQLInfoResponseBodyItems) *DescribeWaitingSQLInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) SetRequestId(v string) *DescribeWaitingSQLInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBody) Validate() error {
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

type DescribeWaitingSQLInfoResponseBodyItems struct {
	// The application that sent the query.
	//
	// example:
	//
	// DataGrip 2022.1.5
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// The application that sent the blocking query.
	//
	// example:
	//
	// DataGrip 2022.1.5
	BlockedByApplication *string `json:"BlockedByApplication,omitempty" xml:"BlockedByApplication,omitempty"`
	// The process ID of the blocking query.
	//
	// example:
	//
	// 110
	BlockedByPID *string `json:"BlockedByPID,omitempty" xml:"BlockedByPID,omitempty"`
	// The SQL statement of the blocking query.
	//
	// example:
	//
	// Select 	- from t1;
	BlockedBySQLStmt *string `json:"BlockedBySQLStmt,omitempty" xml:"BlockedBySQLStmt,omitempty"`
	// The database account that is used to perform the blocking query.
	//
	// example:
	//
	// testUser1
	BlockedByUser *string `json:"BlockedByUser,omitempty" xml:"BlockedByUser,omitempty"`
	// The authorized locks.
	//
	// example:
	//
	// ShareLock,AccessExclusiveLock
	GrantLocks *string `json:"GrantLocks,omitempty" xml:"GrantLocks,omitempty"`
	// The unauthorized locks.
	//
	// example:
	//
	// AccessShareLock
	NotGrantLocks *string `json:"NotGrantLocks,omitempty" xml:"NotGrantLocks,omitempty"`
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
	// The database account that is used to perform the query.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeWaitingSQLInfoResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeWaitingSQLInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetApplication() *string {
	return s.Application
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetBlockedByApplication() *string {
	return s.BlockedByApplication
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetBlockedByPID() *string {
	return s.BlockedByPID
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetBlockedBySQLStmt() *string {
	return s.BlockedBySQLStmt
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetBlockedByUser() *string {
	return s.BlockedByUser
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetGrantLocks() *string {
	return s.GrantLocks
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetNotGrantLocks() *string {
	return s.NotGrantLocks
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetPID() *string {
	return s.PID
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetSQLStmt() *string {
	return s.SQLStmt
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.Application = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByApplication(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByApplication = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByPID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedBySQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedBySQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetBlockedByUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.BlockedByUser = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.GrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetNotGrantLocks(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.NotGrantLocks = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetPID(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.PID = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetSQLStmt(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.SQLStmt = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) SetUser(v string) *DescribeWaitingSQLInfoResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeWaitingSQLInfoResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
