// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertCommonName(v string) *DescribeInstanceSSLResponseBody
	GetCertCommonName() *string
	SetCertDownloadURL(v string) *DescribeInstanceSSLResponseBody
	GetCertDownloadURL() *string
	SetInstanceId(v string) *DescribeInstanceSSLResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *DescribeInstanceSSLResponseBody
	GetRequestId() *string
	SetSSLEnabled(v string) *DescribeInstanceSSLResponseBody
	GetSSLEnabled() *string
	SetSSLExpiredTime(v string) *DescribeInstanceSSLResponseBody
	GetSSLExpiredTime() *string
}

type DescribeInstanceSSLResponseBody struct {
	// The common name of the CA certificate. The default value is the internal endpoint of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The download URL of the CA certificate.
	//
	// example:
	//
	// https://apsaradb-public.oss-ap-sout****-1.aliy****.com/ApsaraDB-CA-Chain.zip
	CertDownloadURL *string `json:"CertDownloadURL,omitempty" xml:"CertDownloadURL,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 02260F96-913E-4655-9BA5-A3651CAF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the TLS (SSL) encryption feature. Valid values:
	//
	// 	- **Enable**: SSL encryption is enabled.
	//
	// 	- **Disable**: SSL encryption is disabled.
	//
	// example:
	//
	// Enable
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The time when the CA certificate expires.
	//
	// example:
	//
	// 2020-08-05T09:05:53Z
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLResponseBody) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeInstanceSSLResponseBody) GetCertDownloadURL() *string {
	return s.CertDownloadURL
}

func (s *DescribeInstanceSSLResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSSLResponseBody) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeInstanceSSLResponseBody) GetSSLExpiredTime() *string {
	return s.SSLExpiredTime
}

func (s *DescribeInstanceSSLResponseBody) SetCertCommonName(v string) *DescribeInstanceSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetCertDownloadURL(v string) *DescribeInstanceSSLResponseBody {
	s.CertDownloadURL = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetInstanceId(v string) *DescribeInstanceSSLResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetRequestId(v string) *DescribeInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetSSLEnabled(v string) *DescribeInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetSSLExpiredTime(v string) *DescribeInstanceSSLResponseBody {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
