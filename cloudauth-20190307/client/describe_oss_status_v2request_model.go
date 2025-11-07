// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCode(v string) *DescribeOssStatusV2Request
	GetServiceCode() *string
	SetSourceIp(v string) *DescribeOssStatusV2Request
	GetSourceIp() *string
}

type DescribeOssStatusV2Request struct {
	// ServiceCode for Real Person Cloud products:
	//
	// - **antcloudauth**: Financial-grade real person authentication
	//
	// - **cloudauthst (discontinued)**: Enhanced real person authentication
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Visitor\\"s source IP address.
	//
	// example:
	//
	// 120.25.41.25
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeOssStatusV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusV2Request) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusV2Request) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeOssStatusV2Request) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOssStatusV2Request) SetServiceCode(v string) *DescribeOssStatusV2Request {
	s.ServiceCode = &v
	return s
}

func (s *DescribeOssStatusV2Request) SetSourceIp(v string) *DescribeOssStatusV2Request {
	s.SourceIp = &v
	return s
}

func (s *DescribeOssStatusV2Request) Validate() error {
	return dara.Validate(s)
}
