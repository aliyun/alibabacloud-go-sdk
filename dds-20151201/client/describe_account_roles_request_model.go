// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAccountRolesRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeAccountRolesRequest
	GetDBName() *string
}

type DescribeAccountRolesRequest struct {
	// The ID of the ApsaraDB for MongoDB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-uf6e9433e955****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database to which the custom role belongs.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeAccountRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRolesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountRolesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountRolesRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountRolesRequest) SetDBInstanceId(v string) *DescribeAccountRolesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountRolesRequest) SetDBName(v string) *DescribeAccountRolesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAccountRolesRequest) Validate() error {
	return dara.Validate(s)
}
