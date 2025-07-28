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
	// The default SSL and TLS settings.
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
	return dara.Validate(s)
}

type DescribeDefaultHttpsResponseBodyDefaultHttps struct {
	// The certificate ID.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suites. Valid values:
	//
	// 	- **1**: all cipher suites.
	//
	// 	- **2**: strong cipher suites.
	//
	// 	- **99**: custom cipher suites.
	//
	// example:
	//
	// 1
	CipherSuite *string `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suite.
	//
	// example:
	//
	// ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384
	CustomCiphers *string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// The version of the TLS protocol. Valid values:
	//
	// 	- **tlsv1**
	//
	// 	- **tlsv1.1**
	//
	// 	- **tlsv1.2**
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
