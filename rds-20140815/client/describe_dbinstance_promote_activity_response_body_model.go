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
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 22973492**********
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 	- China site: 26842
	//
	// 	- International site: 26888
	//
	// example:
	//
	// 268**
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The activity information about the instance. For more information, see [Instance activities](https://help.aliyun.com/document_detail/2391834.html).
	//
	// example:
	//
	// 1
	IsActivity *string `json:"IsActivity,omitempty" xml:"IsActivity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94CB8D93-017A-5AE7-A118-6E0F89D93C0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
