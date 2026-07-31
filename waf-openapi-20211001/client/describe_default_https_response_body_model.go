// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultHttpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultHttps(v *DescribeDefaultHttpsResponseBodyDefaultHttps) *DescribeDefaultHttpsResponseBody
	GetDefaultHttps() *DescribeDefaultHttpsResponseBodyDefaultHttps
	SetRequestId(v string) *DescribeDefaultHttpsResponseBody
	GetRequestId() *string
}

type DescribeDefaultHttpsResponseBody struct {
	// The default SSL/TLS settings.
	DefaultHttps *DescribeDefaultHttpsResponseBodyDefaultHttps `json:"DefaultHttps,omitempty" xml:"DefaultHttps,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefaultHttpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultHttpsResponseBody) GetDefaultHttps() *DescribeDefaultHttpsResponseBodyDefaultHttps {
	return s.DefaultHttps
}

func (s *DescribeDefaultHttpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefaultHttpsResponseBody) SetDefaultHttps(v *DescribeDefaultHttpsResponseBodyDefaultHttps) *DescribeDefaultHttpsResponseBody {
	s.DefaultHttps = v
	return s
}

func (s *DescribeDefaultHttpsResponseBody) SetRequestId(v string) *DescribeDefaultHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBody) Validate() error {
	if s.DefaultHttps != nil {
		if err := s.DefaultHttps.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDefaultHttpsResponseBodyDefaultHttps struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suite. Valid values:
	//
	// - **1**: All cipher suites are added.
	//
	// - **2**: Strong cipher suites are added. This value is available only when TLSVersion is set to tlsv1.2.
	//
	// - **99**: Custom cipher suites are added. This value is available only when TLSVersion is not set to tlsv1.3.
	//
	// example:
	//
	// 1
	CipherSuite *string `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	//
	// example:
	//
	// ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384
	CustomCiphers *string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// - **true**: TLS 1.3 is supported.
	//
	// - **false**: TLS 1.3 is not supported.
	//
	// > This parameter takes effect only when HttpsPorts is not empty, which indicates that the domain name uses the HTTPS protocol. When TLSVersion is set to tlsv1.3, this value must be true.
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// The TLS version. Valid values:
	//
	// - **tlsv1**: TLS 1.0 and later are supported. This value provides the highest compatibility and the lowest security.
	//
	// - **tlsv1.1**: TLS 1.1 and later are supported. This value provides good compatibility and security.
	//
	// - **tlsv1.2**: TLS 1.2 and later are supported. This value provides good compatibility and the highest security.
	//
	// - **tlsv1.3**: Only TLS 1.3 is supported. This value provides the highest security and the lowest compatibility.
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
}

func (s DescribeDefaultHttpsResponseBodyDefaultHttps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultHttpsResponseBodyDefaultHttps) GoString() string {
	return s.String()
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) GetCertId() *string {
	return s.CertId
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) GetCipherSuite() *string {
	return s.CipherSuite
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) GetCustomCiphers() *string {
	return s.CustomCiphers
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) SetCertId(v string) *DescribeDefaultHttpsResponseBodyDefaultHttps {
	s.CertId = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) SetCipherSuite(v string) *DescribeDefaultHttpsResponseBodyDefaultHttps {
	s.CipherSuite = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) SetCustomCiphers(v string) *DescribeDefaultHttpsResponseBodyDefaultHttps {
	s.CustomCiphers = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) SetEnableTLSv3(v bool) *DescribeDefaultHttpsResponseBodyDefaultHttps {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) SetTLSVersion(v string) *DescribeDefaultHttpsResponseBodyDefaultHttps {
	s.TLSVersion = &v
	return s
}

func (s *DescribeDefaultHttpsResponseBodyDefaultHttps) Validate() error {
	return dara.Validate(s)
}
