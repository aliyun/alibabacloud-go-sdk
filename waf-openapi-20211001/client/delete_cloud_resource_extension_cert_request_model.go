// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceExtensionCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DeleteCloudResourceExtensionCertRequest
	GetCertId() *string
	SetCloudResourceId(v string) *DeleteCloudResourceExtensionCertRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *DeleteCloudResourceExtensionCertRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteCloudResourceExtensionCertRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteCloudResourceExtensionCertRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteCloudResourceExtensionCertRequest struct {
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
	// waf_v3prepay_public_intl-sg-***
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

func (s DeleteCloudResourceExtensionCertRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceExtensionCertRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceExtensionCertRequest) GetCertId() *string {
	return s.CertId
}

func (s *DeleteCloudResourceExtensionCertRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *DeleteCloudResourceExtensionCertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudResourceExtensionCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudResourceExtensionCertRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteCloudResourceExtensionCertRequest) SetCertId(v string) *DeleteCloudResourceExtensionCertRequest {
	s.CertId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertRequest) SetCloudResourceId(v string) *DeleteCloudResourceExtensionCertRequest {
	s.CloudResourceId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertRequest) SetInstanceId(v string) *DeleteCloudResourceExtensionCertRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertRequest) SetRegionId(v string) *DeleteCloudResourceExtensionCertRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertRequest) SetResourceManagerResourceGroupId(v string) *DeleteCloudResourceExtensionCertRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertRequest) Validate() error {
	return dara.Validate(s)
}
