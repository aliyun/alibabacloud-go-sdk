// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIpHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceIpHostnameResponseBody
	GetDBInstanceId() *string
	SetIpHostnameInfos(v string) *DescribeDBInstanceIpHostnameResponseBody
	GetIpHostnameInfos() *string
	SetRequestId(v string) *DescribeDBInstanceIpHostnameResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceIpHostnameResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The internal IP addresses and hostnames of the underlying ECS instances for the ApsaraDB RDS for SQL Server instance, including the primary and secondary instances. Format: `ip1,hostname1;ip2,hostname2`.
	//
	// example:
	//
	// 172.16.xx.xx,sd****B;172.16.xx.xx,sd****A
	IpHostnameInfos *string `json:"IpHostnameInfos,omitempty" xml:"IpHostnameInfos,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67CD4719-51E3-4A76-A38C-02F45FAE7E36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceIpHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIpHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIpHostnameResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceIpHostnameResponseBody) GetIpHostnameInfos() *string {
	return s.IpHostnameInfos
}

func (s *DescribeDBInstanceIpHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceIpHostnameResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceIpHostnameResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponseBody) SetIpHostnameInfos(v string) *DescribeDBInstanceIpHostnameResponseBody {
	s.IpHostnameInfos = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponseBody) SetRequestId(v string) *DescribeDBInstanceIpHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIpHostnameResponseBody) Validate() error {
	return dara.Validate(s)
}
