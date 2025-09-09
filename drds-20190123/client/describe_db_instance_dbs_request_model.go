// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeDbInstanceDbsRequest
	GetAccountName() *string
	SetDbInstType(v string) *DescribeDbInstanceDbsRequest
	GetDbInstType() *string
	SetDbInstanceId(v string) *DescribeDbInstanceDbsRequest
	GetDbInstanceId() *string
	SetDrdsInstanceId(v string) *DescribeDbInstanceDbsRequest
	GetDrdsInstanceId() *string
	SetPassword(v string) *DescribeDbInstanceDbsRequest
	GetPassword() *string
}

type DescribeDbInstanceDbsRequest struct {
	// The name of the privileged account of the PolarDB-X 1.0 instance. You do not need to specify this parameter if you have no privileged account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The engine type of the storage-layer databases. Valid values: **POLARDB*	- and **RDS**.
	//
	// example:
	//
	// POLARDB
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the instance in which the storage-layer databases are deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The password of the privileged account. You do not need to specify this parameter if you have no privileged account.
	//
	// example:
	//
	// pwd_111111
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s DescribeDbInstanceDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDbInstanceDbsRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDbInstanceDbsRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDbInstanceDbsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDbInstanceDbsRequest) GetPassword() *string {
	return s.Password
}

func (s *DescribeDbInstanceDbsRequest) SetAccountName(v string) *DescribeDbInstanceDbsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstType(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDbInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetDrdsInstanceId(v string) *DescribeDbInstanceDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) SetPassword(v string) *DescribeDbInstanceDbsRequest {
	s.Password = &v
	return s
}

func (s *DescribeDbInstanceDbsRequest) Validate() error {
	return dara.Validate(s)
}
