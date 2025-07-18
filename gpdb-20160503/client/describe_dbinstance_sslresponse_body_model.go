// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody
	GetCertCommonName() *string
	SetDBInstanceId(v string) *DescribeDBInstanceSSLResponseBody
	GetDBInstanceId() *string
	SetDBInstanceName(v string) *DescribeDBInstanceSSLResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *DescribeDBInstanceSSLResponseBody
	GetRequestId() *string
	SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBody
	GetSSLEnabled() *bool
	SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLExpiredTime() *string
}

type DescribeDBInstanceSSLResponseBody struct {
	// The name of the SSL certificate.
	//
	// example:
	//
	// *.gpdbmaster.xxx.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D5FF8636-37F6-4CE0-8002-F8734C62C686
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SSL encryption is enabled.
	//
	// example:
	//
	// true
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The expiration time of the SSL certificate.
	//
	// example:
	//
	// 2023-08-05T09:05:53Z
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeDBInstanceSSLResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceSSLResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLEnabled() *bool {
	return s.SSLEnabled
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLExpiredTime() *string {
	return s.SSLExpiredTime
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceId(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceSSLResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
