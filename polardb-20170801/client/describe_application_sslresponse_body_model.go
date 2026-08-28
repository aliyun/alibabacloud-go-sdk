// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertCommonName(v string) *DescribeApplicationSSLResponseBody
	GetCertCommonName() *string
	SetCertExpiredTime(v string) *DescribeApplicationSSLResponseBody
	GetCertExpiredTime() *string
	SetCertFingerprintSha256Der(v string) *DescribeApplicationSSLResponseBody
	GetCertFingerprintSha256Der() *string
	SetCertModifiedTime(v string) *DescribeApplicationSSLResponseBody
	GetCertModifiedTime() *string
	SetCertSource(v string) *DescribeApplicationSSLResponseBody
	GetCertSource() *string
	SetRequestId(v string) *DescribeApplicationSSLResponseBody
	GetRequestId() *string
	SetSSLAutoRotate(v bool) *DescribeApplicationSSLResponseBody
	GetSSLAutoRotate() *bool
	SetSSLEnabled(v bool) *DescribeApplicationSSLResponseBody
	GetSSLEnabled() *bool
}

type DescribeApplicationSSLResponseBody struct {
	// The Common Name of the certificate. This field is empty when SSL is not enabled.
	//
	// example:
	//
	// xxx.polarclaw.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The certificate expiration time in UTC. This field is empty when SSL is not enabled.
	//
	// example:
	//
	// 2027-08-09T10:23:26Z
	CertExpiredTime *string `json:"CertExpiredTime,omitempty" xml:"CertExpiredTime,omitempty"`
	// The SHA-256 (DER) fingerprint of the server certificate in lowercase hex. Use this value for client pinning. This is consistent with openssl -fingerprint -sha256. This field is empty when SSL is not enabled.
	//
	// example:
	//
	// 20769803152bf6a3abed626f6b8cae3a1f0d0f2c3b4a59687776655443322110
	CertFingerprintSha256Der *string `json:"CertFingerprintSha256Der,omitempty" xml:"CertFingerprintSha256Der,omitempty"`
	// The most recent certificate installation time in UTC. This field is empty when SSL is not enabled.
	//
	// example:
	//
	// 2026-08-09T10:23:49Z
	CertModifiedTime *string `json:"CertModifiedTime,omitempty" xml:"CertModifiedTime,omitempty"`
	// The certificate source. Valid values:
	//
	// - ca: issued by the platform.
	//
	// - customer: provided by the user.
	//
	// This field is empty when SSL is not enabled.
	//
	// example:
	//
	// ca
	CertSource *string `json:"CertSource,omitempty" xml:"CertSource,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2F029645-FED9-4FE8-A6D3-488954******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether automatic rotation of platform-issued certificates is enabled.
	//
	// example:
	//
	// true
	SSLAutoRotate *bool `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
	// Indicates whether SSL is enabled.
	//
	// example:
	//
	// true
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s DescribeApplicationSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSSLResponseBody) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeApplicationSSLResponseBody) GetCertExpiredTime() *string {
	return s.CertExpiredTime
}

func (s *DescribeApplicationSSLResponseBody) GetCertFingerprintSha256Der() *string {
	return s.CertFingerprintSha256Der
}

func (s *DescribeApplicationSSLResponseBody) GetCertModifiedTime() *string {
	return s.CertModifiedTime
}

func (s *DescribeApplicationSSLResponseBody) GetCertSource() *string {
	return s.CertSource
}

func (s *DescribeApplicationSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationSSLResponseBody) GetSSLAutoRotate() *bool {
	return s.SSLAutoRotate
}

func (s *DescribeApplicationSSLResponseBody) GetSSLEnabled() *bool {
	return s.SSLEnabled
}

func (s *DescribeApplicationSSLResponseBody) SetCertCommonName(v string) *DescribeApplicationSSLResponseBody {
	s.CertCommonName = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetCertExpiredTime(v string) *DescribeApplicationSSLResponseBody {
	s.CertExpiredTime = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetCertFingerprintSha256Der(v string) *DescribeApplicationSSLResponseBody {
	s.CertFingerprintSha256Der = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetCertModifiedTime(v string) *DescribeApplicationSSLResponseBody {
	s.CertModifiedTime = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetCertSource(v string) *DescribeApplicationSSLResponseBody {
	s.CertSource = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetRequestId(v string) *DescribeApplicationSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetSSLAutoRotate(v bool) *DescribeApplicationSSLResponseBody {
	s.SSLAutoRotate = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) SetSSLEnabled(v bool) *DescribeApplicationSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeApplicationSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
