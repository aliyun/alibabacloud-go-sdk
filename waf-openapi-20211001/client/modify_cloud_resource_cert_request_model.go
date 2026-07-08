// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificates(v []*ModifyCloudResourceCertRequestCertificates) *ModifyCloudResourceCertRequest
	GetCertificates() []*ModifyCloudResourceCertRequestCertificates
	SetCloudResourceId(v string) *ModifyCloudResourceCertRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *ModifyCloudResourceCertRequest
	GetInstanceId() *string
	SetPort(v int32) *ModifyCloudResourceCertRequest
	GetPort() *int32
	SetRegionId(v string) *ModifyCloudResourceCertRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *ModifyCloudResourceCertRequest
	GetResourceInstanceId() *string
	SetResourceProduct(v string) *ModifyCloudResourceCertRequest
	GetResourceProduct() *string
}

type ModifyCloudResourceCertRequest struct {
	// The list of certificates.
	//
	// > Enter all certificate IDs. This includes the default certificate and all additional certificates. After you submit the request, WAF compares the submitted IDs with the existing ones. WAF adds new certificates and deletes certificates that are not in your list. Deleting a certificate may affect related services.
	//
	// This parameter is required.
	Certificates []*ModifyCloudResourceCertRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The ID of the resource that is added to WAF. WAF automatically generates this ID when you add the resource in cloud native mode.
	//
	// > Call the [CreateCloudResource](https://help.aliyun.com/document_detail/2839876.html) operation to add a resource. Then, view the resource ID in the response.
	//
	// example:
	//
	// lb-bp*********k5uj2p-443-clb7
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// The port of the cloud product that is added to WAF.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the cloud product instance.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// Deprecated
	//
	// The type of the cloud product. Valid values:
	//
	// - **ecs**: Elastic Compute Service (ECS).
	//
	// - **clb4**: Layer 4 Classic Load Balancer (CLB).
	//
	// - **nlb**: Network Load Balancer (NLB).
	//
	// example:
	//
	// clb4
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s ModifyCloudResourceCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceCertRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceCertRequest) GetCertificates() []*ModifyCloudResourceCertRequestCertificates {
	return s.Certificates
}

func (s *ModifyCloudResourceCertRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *ModifyCloudResourceCertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCloudResourceCertRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyCloudResourceCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudResourceCertRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *ModifyCloudResourceCertRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *ModifyCloudResourceCertRequest) SetCertificates(v []*ModifyCloudResourceCertRequestCertificates) *ModifyCloudResourceCertRequest {
	s.Certificates = v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetCloudResourceId(v string) *ModifyCloudResourceCertRequest {
	s.CloudResourceId = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetInstanceId(v string) *ModifyCloudResourceCertRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetPort(v int32) *ModifyCloudResourceCertRequest {
	s.Port = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetRegionId(v string) *ModifyCloudResourceCertRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetResourceInstanceId(v string) *ModifyCloudResourceCertRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) SetResourceProduct(v string) *ModifyCloudResourceCertRequest {
	s.ResourceProduct = &v
	return s
}

func (s *ModifyCloudResourceCertRequest) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCloudResourceCertRequestCertificates struct {
	// The type of the certificate for the HTTPS protocol. Valid values:
	//
	// - **default**: the default certificate.
	//
	// - **extension**: the additional certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 232***-cn-hangzhou
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s ModifyCloudResourceCertRequestCertificates) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceCertRequestCertificates) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceCertRequestCertificates) GetAppliedType() *string {
	return s.AppliedType
}

func (s *ModifyCloudResourceCertRequestCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ModifyCloudResourceCertRequestCertificates) SetAppliedType(v string) *ModifyCloudResourceCertRequestCertificates {
	s.AppliedType = &v
	return s
}

func (s *ModifyCloudResourceCertRequestCertificates) SetCertificateId(v string) *ModifyCloudResourceCertRequestCertificates {
	s.CertificateId = &v
	return s
}

func (s *ModifyCloudResourceCertRequestCertificates) Validate() error {
	return dara.Validate(s)
}
