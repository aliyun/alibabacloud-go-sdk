// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeMachineSpecRequest
	GetChargeType() *string
	SetInstanceTypes(v []*string) *DescribeMachineSpecRequest
	GetInstanceTypes() []*string
	SetResourceType(v string) *DescribeMachineSpecRequest
	GetResourceType() *string
}

type DescribeMachineSpecRequest struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeMachineSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeMachineSpecRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeMachineSpecRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMachineSpecRequest) SetChargeType(v string) *DescribeMachineSpecRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeMachineSpecRequest) SetInstanceTypes(v []*string) *DescribeMachineSpecRequest {
	s.InstanceTypes = v
	return s
}

func (s *DescribeMachineSpecRequest) SetResourceType(v string) *DescribeMachineSpecRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeMachineSpecRequest) Validate() error {
	return dara.Validate(s)
}
