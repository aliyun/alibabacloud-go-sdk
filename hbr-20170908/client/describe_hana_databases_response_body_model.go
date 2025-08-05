// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHanaDatabasesResponseBody
	GetCode() *string
	SetHanaDatabases(v *DescribeHanaDatabasesResponseBodyHanaDatabases) *DescribeHanaDatabasesResponseBody
	GetHanaDatabases() *DescribeHanaDatabasesResponseBodyHanaDatabases
	SetMessage(v string) *DescribeHanaDatabasesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHanaDatabasesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaDatabasesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHanaDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHanaDatabasesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeHanaDatabasesResponseBody
	GetTotalCount() *int64
}

type DescribeHanaDatabasesResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about SAP HANA databases.
	HanaDatabases *DescribeHanaDatabasesResponseBodyHanaDatabases `json:"HanaDatabases,omitempty" xml:"HanaDatabases,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DAAB6A29-34EB-5F56-962F-D5BDBFE8A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHanaDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHanaDatabasesResponseBody) GetHanaDatabases() *DescribeHanaDatabasesResponseBodyHanaDatabases {
	return s.HanaDatabases
}

func (s *DescribeHanaDatabasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHanaDatabasesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaDatabasesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHanaDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHanaDatabasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHanaDatabasesResponseBody) SetCode(v string) *DescribeHanaDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetHanaDatabases(v *DescribeHanaDatabasesResponseBodyHanaDatabases) *DescribeHanaDatabasesResponseBody {
	s.HanaDatabases = v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetMessage(v string) *DescribeHanaDatabasesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetPageNumber(v int32) *DescribeHanaDatabasesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetPageSize(v int32) *DescribeHanaDatabasesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetRequestId(v string) *DescribeHanaDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetSuccess(v bool) *DescribeHanaDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) SetTotalCount(v int64) *DescribeHanaDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHanaDatabasesResponseBodyHanaDatabases struct {
	HanaDatabase []*DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase `json:"HanaDatabase,omitempty" xml:"HanaDatabase,omitempty" type:"Repeated"`
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabases) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabases) GetHanaDatabase() []*DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	return s.HanaDatabase
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabases) SetHanaDatabase(v []*DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) *DescribeHanaDatabasesResponseBodyHanaDatabases {
	s.HanaDatabase = v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabases) Validate() error {
	return dara.Validate(s)
}

type DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase struct {
	// Indicates whether the database is started. Valid values:
	//
	// 	- **YES**: The database is started.
	//
	// 	- **NO**: The database is not started.
	//
	// example:
	//
	// YES
	ActiveStatus *string `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	// The database name.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The detailed information.
	//
	// example:
	//
	// master
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The hostname.
	//
	// example:
	//
	// izbp1jbf3zy******antqmz
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The service name.
	//
	// example:
	//
	// indexserver
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The port number.
	//
	// example:
	//
	// 30013
	SqlPort *int32 `json:"SqlPort,omitempty" xml:"SqlPort,omitempty"`
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetActiveStatus() *string {
	return s.ActiveStatus
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetDetail() *string {
	return s.Detail
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetHost() *string {
	return s.Host
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) GetSqlPort() *int32 {
	return s.SqlPort
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetActiveStatus(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.ActiveStatus = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetDatabaseName(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetDetail(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.Detail = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetHost(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.Host = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetServiceName(v string) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.ServiceName = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) SetSqlPort(v int32) *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase {
	s.SqlPort = &v
	return s
}

func (s *DescribeHanaDatabasesResponseBodyHanaDatabasesHanaDatabase) Validate() error {
	return dara.Validate(s)
}
