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
	SetRequestId(v string) *DescribeDBInstanceSSLResponseBody
	GetRequestId() *string
	SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLExpiredTime() *string
	SetSSLStatus(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLStatus() *string
}

type DescribeDBInstanceSSLResponseBody struct {
	// The name of the SSL certificate.
	//
	// example:
	//
	// dds-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36BB1BC2-789C-4BBA-A519-C5B388E4B0D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the SSL certificate expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2020-03-11T02:28:25Z
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
	// The status of the SSL feature. Valid values:
	//
	// 	- **Open**: The SSL feature is enabled.
	//
	// 	- **Closed**: The SSL feature is disabled.
	//
	// example:
	//
	// Open
	SSLStatus *string `json:"SSLStatus,omitempty" xml:"SSLStatus,omitempty"`
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

func (s *DescribeDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLExpiredTime() *string {
	return s.SSLExpiredTime
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLStatus() *string {
	return s.SSLStatus
}

func (s *DescribeDBInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLStatus(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLStatus = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
