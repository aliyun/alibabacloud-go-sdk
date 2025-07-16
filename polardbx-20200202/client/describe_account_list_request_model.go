// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountListRequest
	GetAccountName() *string
	SetAccountType(v string) *DescribeAccountListRequest
	GetAccountType() *string
	SetDBInstanceName(v string) *DescribeAccountListRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeAccountListRequest
	GetRegionId() *string
}

type DescribeAccountListRequest struct {
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 0
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAccountListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountListRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountListRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAccountListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountListRequest) SetAccountName(v string) *DescribeAccountListRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountListRequest) SetAccountType(v string) *DescribeAccountListRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountListRequest) SetDBInstanceName(v string) *DescribeAccountListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListRequest) SetRegionId(v string) *DescribeAccountListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountListRequest) Validate() error {
	return dara.Validate(s)
}
