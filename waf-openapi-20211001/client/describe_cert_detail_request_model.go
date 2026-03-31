// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DescribeCertDetailRequest
	GetCertIdentifier() *string
	SetInstanceId(v string) *DescribeCertDetailRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeCertDetailRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCertDetailRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCertDetailRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-uax****3k0e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCertDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCertDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCertDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCertDetailRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCertDetailRequest) SetCertIdentifier(v string) *DescribeCertDetailRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCertDetailRequest) SetInstanceId(v string) *DescribeCertDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCertDetailRequest) SetRegionId(v string) *DescribeCertDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCertDetailRequest) SetResourceManagerResourceGroupId(v string) *DescribeCertDetailRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCertDetailRequest) Validate() error {
	return dara.Validate(s)
}
