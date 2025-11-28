// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDatabaseRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *DescribeDatabaseRequest
	GetDatabaseName() *string
}

type DescribeDatabaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
}

func (s DescribeDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDatabaseRequest) SetDBInstanceId(v string) *DescribeDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDatabaseRequest) SetDatabaseName(v string) *DescribeDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
