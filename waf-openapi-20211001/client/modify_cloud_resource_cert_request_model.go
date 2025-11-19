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
	// This parameter is required.
	Certificates    []*ModifyCloudResourceCertRequestCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	CloudResourceId *string                                       `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// Deprecated
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
	// This parameter is required.
	//
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
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
