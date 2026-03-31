// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceExtensionCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *CreateCloudResourceExtensionCertRequest
	GetCertId() *string
	SetCloudResourceId(v string) *CreateCloudResourceExtensionCertRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *CreateCloudResourceExtensionCertRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateCloudResourceExtensionCertRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateCloudResourceExtensionCertRequest
	GetResourceManagerResourceGroupId() *string
}

type CreateCloudResourceExtensionCertRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// lb-bp*********k5uj2p-443-clb7
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s CreateCloudResourceExtensionCertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceExtensionCertRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceExtensionCertRequest) GetCertId() *string {
	return s.CertId
}

func (s *CreateCloudResourceExtensionCertRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *CreateCloudResourceExtensionCertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudResourceExtensionCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudResourceExtensionCertRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateCloudResourceExtensionCertRequest) SetCertId(v string) *CreateCloudResourceExtensionCertRequest {
	s.CertId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertRequest) SetCloudResourceId(v string) *CreateCloudResourceExtensionCertRequest {
	s.CloudResourceId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertRequest) SetInstanceId(v string) *CreateCloudResourceExtensionCertRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertRequest) SetRegionId(v string) *CreateCloudResourceExtensionCertRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertRequest) SetResourceManagerResourceGroupId(v string) *CreateCloudResourceExtensionCertRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertRequest) Validate() error {
	return dara.Validate(s)
}
