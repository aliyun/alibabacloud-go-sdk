// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsEipAddressAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *ModifyEnsEipAddressAttributeRequest
	GetAllocationId() *string
	SetBandwidth(v int32) *ModifyEnsEipAddressAttributeRequest
	GetBandwidth() *int32
	SetDescription(v string) *ModifyEnsEipAddressAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyEnsEipAddressAttributeRequest
	GetName() *string
}

type ModifyEnsEipAddressAttributeRequest struct {
	// The ID of the EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-5sw5dxzgi6umq4uexxkt8wpma
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The peak bandwidth of the EIP. Default value: 5. Valid values: **5*	- to **10000**. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The new description of the EIP. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the EIP. The name must be 2 to 128 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// test-api-modify
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyEnsEipAddressAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsEipAddressAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEnsEipAddressAttributeRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *ModifyEnsEipAddressAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyEnsEipAddressAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyEnsEipAddressAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyEnsEipAddressAttributeRequest) SetAllocationId(v string) *ModifyEnsEipAddressAttributeRequest {
	s.AllocationId = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetBandwidth(v int32) *ModifyEnsEipAddressAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetDescription(v string) *ModifyEnsEipAddressAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) SetName(v string) *ModifyEnsEipAddressAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyEnsEipAddressAttributeRequest) Validate() error {
	return dara.Validate(s)
}
