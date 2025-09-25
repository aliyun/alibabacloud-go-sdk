// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCAType(v string) *DescribeInstanceSSLResponseBody
	GetCAType() *string
	SetInstanceName(v string) *DescribeInstanceSSLResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *DescribeInstanceSSLResponseBody
	GetRequestId() *string
	SetSSLEnabled(v string) *DescribeInstanceSSLResponseBody
	GetSSLEnabled() *string
	SetServerCert(v string) *DescribeInstanceSSLResponseBody
	GetServerCert() *string
	SetServerKey(v string) *DescribeInstanceSSLResponseBody
	GetServerKey() *string
}

type DescribeInstanceSSLResponseBody struct {
	// example:
	//
	// custom
	CAType *string `json:"CAType,omitempty" xml:"CAType,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 32DEFB4A-861F-5D1D-ADD5-918E4FD7AB8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****QqEP-----END CERTIFICATE-----
	ServerCert *string `json:"ServerCert,omitempty" xml:"ServerCert,omitempty"`
	// example:
	//
	// -----BEGIN PRIVATE KEY-----MIIE****ihfg==-----END PRIVATE KEY-----
	ServerKey *string `json:"ServerKey,omitempty" xml:"ServerKey,omitempty"`
}

func (s DescribeInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLResponseBody) GetCAType() *string {
	return s.CAType
}

func (s *DescribeInstanceSSLResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSSLResponseBody) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeInstanceSSLResponseBody) GetServerCert() *string {
	return s.ServerCert
}

func (s *DescribeInstanceSSLResponseBody) GetServerKey() *string {
	return s.ServerKey
}

func (s *DescribeInstanceSSLResponseBody) SetCAType(v string) *DescribeInstanceSSLResponseBody {
	s.CAType = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetInstanceName(v string) *DescribeInstanceSSLResponseBody {
	s.InstanceName = &v
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

func (s *DescribeInstanceSSLResponseBody) SetServerCert(v string) *DescribeInstanceSSLResponseBody {
	s.ServerCert = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) SetServerKey(v string) *DescribeInstanceSSLResponseBody {
	s.ServerKey = &v
	return s
}

func (s *DescribeInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
