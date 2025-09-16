// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticResourceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyElasticResourceSpecShrinkRequest
	GetInstanceId() *string
	SetRegion(v string) *ModifyElasticResourceSpecShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *ModifyElasticResourceSpecShrinkRequest
	GetResourceSpecShrink() *string
}

type ModifyElasticResourceSpecShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5sess
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ModifyElasticResourceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticResourceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticResourceSpecShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticResourceSpecShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyElasticResourceSpecShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *ModifyElasticResourceSpecShrinkRequest) SetInstanceId(v string) *ModifyElasticResourceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticResourceSpecShrinkRequest) SetRegion(v string) *ModifyElasticResourceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyElasticResourceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyElasticResourceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *ModifyElasticResourceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
