// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateAppGroupRequest
	GetChargeType() *string
	SetName(v string) *CreateAppGroupRequest
	GetName() *string
	SetQuota(v *CreateAppGroupRequestQuota) *CreateAppGroupRequest
	GetQuota() *CreateAppGroupRequestQuota
	SetResourceGroupId(v string) *CreateAppGroupRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateAppGroupRequestTags) *CreateAppGroupRequest
	GetTags() []*CreateAppGroupRequestTags
	SetType(v string) *CreateAppGroupRequest
	GetType() *string
}

type CreateAppGroupRequest struct {
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// jmbon_analyzer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The quota.
	Quota *CreateAppGroupRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm2ij6pwxsvua
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags.
	Tags []*CreateAppGroupRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard
	//
	// 	- enhanced
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAppGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppGroupRequest) GetQuota() *CreateAppGroupRequestQuota {
	return s.Quota
}

func (s *CreateAppGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAppGroupRequest) GetTags() []*CreateAppGroupRequestTags {
	return s.Tags
}

func (s *CreateAppGroupRequest) GetType() *string {
	return s.Type
}

func (s *CreateAppGroupRequest) SetChargeType(v string) *CreateAppGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppGroupRequest) SetName(v string) *CreateAppGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateAppGroupRequest) SetQuota(v *CreateAppGroupRequestQuota) *CreateAppGroupRequest {
	s.Quota = v
	return s
}

func (s *CreateAppGroupRequest) SetResourceGroupId(v string) *CreateAppGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAppGroupRequest) SetTags(v []*CreateAppGroupRequestTags) *CreateAppGroupRequest {
	s.Tags = v
	return s
}

func (s *CreateAppGroupRequest) SetType(v string) *CreateAppGroupRequest {
	s.Type = &v
	return s
}

func (s *CreateAppGroupRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppGroupRequestQuota struct {
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic
	//
	// 	- opensearch.share.common: shared general-purpose
	//
	// 	- opensearch.share.compute: shared computing
	//
	// 	- opensearch.share.storage: shared storage
	//
	// 	- opensearch.private.common: exclusive general-purpose
	//
	// 	- opensearch.private.compute: exclusive computing
	//
	// 	- opensearch.private.storage: exclusive storage
	//
	// example:
	//
	// opensearch.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateAppGroupRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequestQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *CreateAppGroupRequestQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *CreateAppGroupRequestQuota) GetSpec() *string {
	return s.Spec
}

func (s *CreateAppGroupRequestQuota) SetComputeResource(v int32) *CreateAppGroupRequestQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppGroupRequestQuota) SetDocSize(v int32) *CreateAppGroupRequestQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppGroupRequestQuota) SetSpec(v string) *CreateAppGroupRequestQuota {
	s.Spec = &v
	return s
}

func (s *CreateAppGroupRequestQuota) Validate() error {
	return dara.Validate(s)
}

type CreateAppGroupRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// a
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAppGroupRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateAppGroupRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateAppGroupRequestTags) SetKey(v string) *CreateAppGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAppGroupRequestTags) SetValue(v string) *CreateAppGroupRequestTags {
	s.Value = &v
	return s
}

func (s *CreateAppGroupRequestTags) Validate() error {
	return dara.Validate(s)
}
