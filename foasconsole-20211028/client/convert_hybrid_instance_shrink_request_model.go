// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHybridInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConvertHybridInstanceShrinkRequest
	GetInstanceId() *string
	SetRegion(v string) *ConvertHybridInstanceShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *ConvertHybridInstanceShrinkRequest
	GetResourceSpecShrink() *string
}

type ConvertHybridInstanceShrinkRequest struct {
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
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ConvertHybridInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertHybridInstanceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertHybridInstanceShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *ConvertHybridInstanceShrinkRequest) SetInstanceId(v string) *ConvertHybridInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertHybridInstanceShrinkRequest) SetRegion(v string) *ConvertHybridInstanceShrinkRequest {
	s.Region = &v
	return s
}

func (s *ConvertHybridInstanceShrinkRequest) SetResourceSpecShrink(v string) *ConvertHybridInstanceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *ConvertHybridInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
