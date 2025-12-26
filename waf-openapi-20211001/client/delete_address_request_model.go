// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressList(v []*string) *DeleteAddressRequest
	GetAddressList() []*string
	SetInstanceId(v string) *DeleteAddressRequest
	GetInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteAddressRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DeleteAddressRequest
	GetRuleId() *int64
}

type DeleteAddressRequest struct {
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-wwo36****0i
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

func (s DeleteAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddressRequest) GetAddressList() []*string {
	return s.AddressList
}

func (s *DeleteAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAddressRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteAddressRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteAddressRequest) SetAddressList(v []*string) *DeleteAddressRequest {
	s.AddressList = v
	return s
}

func (s *DeleteAddressRequest) SetInstanceId(v string) *DeleteAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAddressRequest) SetResourceManagerResourceGroupId(v string) *DeleteAddressRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteAddressRequest) SetRuleId(v int64) *DeleteAddressRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAddressRequest) Validate() error {
	return dara.Validate(s)
}
