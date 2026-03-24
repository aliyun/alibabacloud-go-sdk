// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceDefaultCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *ModifyCloudResourceDefaultCertRequest
	GetCertId() *string
	SetCloudResourceId(v string) *ModifyCloudResourceDefaultCertRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *ModifyCloudResourceDefaultCertRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyCloudResourceDefaultCertRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceDefaultCertRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyCloudResourceDefaultCertRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the connected cloud resource. WAF generates this ID automatically when you connect a cloud resource.
	//
	// > After you connect a resource using [CreateCloudResource](https://help.aliyun.com/document_detail/2839876.html), check the response to get the resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-***-80-clb7
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-kf74****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland
	//
	// - **ap-southeast-1**: outside the Chinese mainland
	//
	// This parameter is required.
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

func (s ModifyCloudResourceDefaultCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceDefaultCertRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceDefaultCertRequest) GetCertId() *string {
	return s.CertId
}

func (s *ModifyCloudResourceDefaultCertRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *ModifyCloudResourceDefaultCertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCloudResourceDefaultCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudResourceDefaultCertRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyCloudResourceDefaultCertRequest) SetCertId(v string) *ModifyCloudResourceDefaultCertRequest {
	s.CertId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertRequest) SetCloudResourceId(v string) *ModifyCloudResourceDefaultCertRequest {
	s.CloudResourceId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertRequest) SetInstanceId(v string) *ModifyCloudResourceDefaultCertRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertRequest) SetRegionId(v string) *ModifyCloudResourceDefaultCertRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertRequest) SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceDefaultCertRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertRequest) Validate() error {
	return dara.Validate(s)
}
