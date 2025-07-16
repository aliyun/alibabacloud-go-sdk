// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceSSLResponseBodyData) *DescribeDBInstanceSSLResponseBody
	GetData() *DescribeDBInstanceSSLResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceSSLResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceSSLResponseBody struct {
	Data *DescribeDBInstanceSSLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) GetData() *DescribeDBInstanceSSLResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSSLResponseBody) SetData(v *DescribeDBInstanceSSLResponseBodyData) *DescribeDBInstanceSSLResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceSSLResponseBodyData struct {
	// example:
	//
	// pxc-sddddddcym7g7w****.polarx.singapore.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// example:
	//
	// false
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// example:
	//
	// 2022-11-04T09:39:07Z
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBodyData) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeDBInstanceSSLResponseBodyData) GetSSLEnabled() *bool {
	return s.SSLEnabled
}

func (s *DescribeDBInstanceSSLResponseBodyData) GetSSLExpiredTime() *string {
	return s.SSLExpiredTime
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetCertCommonName(v string) *DescribeDBInstanceSSLResponseBodyData {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponseBodyData {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBodyData) Validate() error {
	return dara.Validate(s)
}
