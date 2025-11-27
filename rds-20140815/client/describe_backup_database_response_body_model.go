// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseNames(v string) *DescribeBackupDatabaseResponseBody
	GetDatabaseNames() *string
	SetRequestId(v string) *DescribeBackupDatabaseResponseBody
	GetRequestId() *string
}

type DescribeBackupDatabaseResponseBody struct {
	// The name of the database. Format: "db1,db2".
	//
	// example:
	//
	// db1,db2
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 08A3B71B-FE08-4B03-974F-CC7EA6DB1828
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDatabaseResponseBody) GetDatabaseNames() *string {
	return s.DatabaseNames
}

func (s *DescribeBackupDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupDatabaseResponseBody) SetDatabaseNames(v string) *DescribeBackupDatabaseResponseBody {
	s.DatabaseNames = &v
	return s
}

func (s *DescribeBackupDatabaseResponseBody) SetRequestId(v string) *DescribeBackupDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
