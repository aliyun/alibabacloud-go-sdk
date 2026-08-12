// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNamespaceSpecV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticResourceSpecShrink(v string) *ModifyNamespaceSpecV2ShrinkRequest
	GetElasticResourceSpecShrink() *string
	SetGuaranteedResourceSpecShrink(v string) *ModifyNamespaceSpecV2ShrinkRequest
	GetGuaranteedResourceSpecShrink() *string
	SetHa(v bool) *ModifyNamespaceSpecV2ShrinkRequest
	GetHa() *bool
	SetInstanceId(v string) *ModifyNamespaceSpecV2ShrinkRequest
	GetInstanceId() *string
	SetNamespace(v string) *ModifyNamespaceSpecV2ShrinkRequest
	GetNamespace() *string
	SetRegion(v string) *ModifyNamespaceSpecV2ShrinkRequest
	GetRegion() *string
}

type ModifyNamespaceSpecV2ShrinkRequest struct {
	// The upper limit of pay-as-you-go resources allocated to the project namespace.
	ElasticResourceSpecShrink *string `json:"ElasticResourceSpec,omitempty" xml:"ElasticResourceSpec,omitempty"`
	// The size of subscription resources allocated to the project namespace.
	GuaranteedResourceSpecShrink *string `json:"GuaranteedResourceSpec,omitempty" xml:"GuaranteedResourceSpec,omitempty"`
	// Specifies whether the project namespace uses zone-disaster recovery.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// The order instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// di-5934394438****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyNamespaceSpecV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetElasticResourceSpecShrink() *string {
	return s.ElasticResourceSpecShrink
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetGuaranteedResourceSpecShrink() *string {
	return s.GuaranteedResourceSpecShrink
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetElasticResourceSpecShrink(v string) *ModifyNamespaceSpecV2ShrinkRequest {
	s.ElasticResourceSpecShrink = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetGuaranteedResourceSpecShrink(v string) *ModifyNamespaceSpecV2ShrinkRequest {
	s.GuaranteedResourceSpecShrink = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetHa(v bool) *ModifyNamespaceSpecV2ShrinkRequest {
	s.Ha = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetInstanceId(v string) *ModifyNamespaceSpecV2ShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetNamespace(v string) *ModifyNamespaceSpecV2ShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) SetRegion(v string) *ModifyNamespaceSpecV2ShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyNamespaceSpecV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
