// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePromoteActivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetAliUid() *string
	SetBid(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetBid() *string
	SetDBInstanceId(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetDBInstanceId() *string
	SetDBInstanceName(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetDBInstanceName() *string
	SetDBType(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetDBType() *string
	SetIsActivity(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetIsActivity() *string
	SetRequestId(v string) *DescribeDBInstancePromoteActivityResponseBody
	GetRequestId() *string
}

type DescribeDBInstancePromoteActivityResponseBody struct {
	AliUid         *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Bid            *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBType         *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	IsActivity     *string `json:"IsActivity,omitempty" xml:"IsActivity,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstancePromoteActivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePromoteActivityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetIsActivity() *string {
	return s.IsActivity
}

func (s *DescribeDBInstancePromoteActivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetAliUid(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetBid(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.Bid = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetDBInstanceName(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetDBType(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetIsActivity(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.IsActivity = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) SetRequestId(v string) *DescribeDBInstancePromoteActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityResponseBody) Validate() error {
	return dara.Validate(s)
}
