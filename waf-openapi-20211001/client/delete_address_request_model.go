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
	// The list of addresses to delete.
	//
	// This parameter is required.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-wwo36****0i
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the address book.
	//
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
