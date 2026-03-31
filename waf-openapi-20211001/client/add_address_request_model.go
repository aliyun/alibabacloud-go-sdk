// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressList(v []*string) *AddAddressRequest
	GetAddressList() []*string
	SetInstanceId(v string) *AddAddressRequest
	GetInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *AddAddressRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *AddAddressRequest
	GetRuleId() *int64
}

type AddAddressRequest struct {
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRequest) GoString() string {
	return s.String()
}

func (s *AddAddressRequest) GetAddressList() []*string {
	return s.AddressList
}

func (s *AddAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddAddressRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *AddAddressRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *AddAddressRequest) SetAddressList(v []*string) *AddAddressRequest {
	s.AddressList = v
	return s
}

func (s *AddAddressRequest) SetInstanceId(v string) *AddAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *AddAddressRequest) SetResourceManagerResourceGroupId(v string) *AddAddressRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *AddAddressRequest) SetRuleId(v int64) *AddAddressRequest {
	s.RuleId = &v
	return s
}

func (s *AddAddressRequest) Validate() error {
	return dara.Validate(s)
}
