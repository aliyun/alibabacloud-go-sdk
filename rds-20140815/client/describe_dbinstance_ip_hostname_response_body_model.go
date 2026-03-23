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
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	IpHostnameInfos *string `json:"IpHostnameInfos,omitempty" xml:"IpHostnameInfos,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
