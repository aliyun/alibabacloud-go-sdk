// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DescribeAccountAuthorityRequest
	GetAccount() *string
	SetDBInstanceId(v string) *DescribeAccountAuthorityRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeAccountAuthorityRequest
	GetRegionId() *string
}

type DescribeAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityRequest) GetAccount() *string {
	return s.Account
}

func (s *DescribeAccountAuthorityRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountAuthorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountAuthorityRequest) SetAccount(v string) *DescribeAccountAuthorityRequest {
	s.Account = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetDBInstanceId(v string) *DescribeAccountAuthorityRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) SetRegionId(v string) *DescribeAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
