// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMVInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeIMVInfosRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeIMVInfosRequest
	GetDatabase() *string
	SetMVName(v string) *DescribeIMVInfosRequest
	GetMVName() *string
}

type DescribeIMVInfosRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The name of MV
	//
	// example:
	//
	// public."mv1"
	MVName *string `json:"MVName,omitempty" xml:"MVName,omitempty"`
}

func (s DescribeIMVInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMVInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeIMVInfosRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeIMVInfosRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeIMVInfosRequest) GetMVName() *string {
	return s.MVName
}

func (s *DescribeIMVInfosRequest) SetDBInstanceId(v string) *DescribeIMVInfosRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeIMVInfosRequest) SetDatabase(v string) *DescribeIMVInfosRequest {
	s.Database = &v
	return s
}

func (s *DescribeIMVInfosRequest) SetMVName(v string) *DescribeIMVInfosRequest {
	s.MVName = &v
	return s
}

func (s *DescribeIMVInfosRequest) Validate() error {
	return dara.Validate(s)
}
