// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceErrorLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstanceErrorLogResponseBodyItems) *DescribeDBInstanceErrorLogResponseBody
	GetItems() []*DescribeDBInstanceErrorLogResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceErrorLogResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDBInstanceErrorLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBInstanceErrorLogResponseBody
	GetTotalCount() *int32
}

type DescribeDBInstanceErrorLogResponseBody struct {
	// The content of the error log.
	Items []*DescribeDBInstanceErrorLogResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBody) GetItems() []*DescribeDBInstanceErrorLogResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceErrorLogResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceErrorLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceErrorLogResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetItems(v []*DescribeDBInstanceErrorLogResponseBodyItems) *DescribeDBInstanceErrorLogResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetPageNumber(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetRequestId(v string) *DescribeDBInstanceErrorLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) SetTotalCount(v int32) *DescribeDBInstanceErrorLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBody) Validate() error {
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

type DescribeDBInstanceErrorLogResponseBodyItems struct {
	// The name of the database.
	//
	// example:
	//
	// adbtest
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// null
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The content of the error log.
	//
	// example:
	//
	// unsupported frontend protocol 2689.28208: server supports 1.0 to 3.0
	LogContext *string `json:"LogContext,omitempty" xml:"LogContext,omitempty"`
	// The level of the queried log.
	//
	// example:
	//
	// FATAL
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// The time when the log was generated. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-24 11:28:14
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// adbpguser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceErrorLogResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetDatabase() *string {
	return s.Database
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetHost() *string {
	return s.Host
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetLogContext() *string {
	return s.LogContext
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetTime() *int64 {
	return s.Time
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) GetUser() *string {
	return s.User
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetDatabase(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetHost(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Host = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogContext(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogContext = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetLogLevel(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.LogLevel = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetTime(v int64) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.Time = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) SetUser(v string) *DescribeDBInstanceErrorLogResponseBodyItems {
	s.User = &v
	return s
}

func (s *DescribeDBInstanceErrorLogResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
