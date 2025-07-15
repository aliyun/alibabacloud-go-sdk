// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnhanhcedNatGatewayAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetAcceptLanguage() *string
	SetFilter(v []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetFilter() []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter
	SetOwnerAccount(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListEnhanhcedNatGatewayAvailableZonesRequest
	GetResourceOwnerId() *int64
}

type ListEnhanhcedNatGatewayAvailableZonesRequest struct {
	// The language to display the results. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The filter information. You can specify a filter key and a filter value.
	Filter       []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// In this example, zones that support NAT gateways in the UAE (Dubai) region are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// me-east-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListEnhanhcedNatGatewayAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnhanhcedNatGatewayAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetFilter() []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter {
	return s.Filter
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetAcceptLanguage(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetFilter(v []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.Filter = v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetOwnerAccount(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetOwnerId(v int64) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetRegionId(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetResourceOwnerAccount(v string) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) SetResourceOwnerId(v int64) *ListEnhanhcedNatGatewayAvailableZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}

type ListEnhanhcedNatGatewayAvailableZonesRequestFilter struct {
	// The filter key. Only **PrivateLinkEnabled*	- is supported.
	//
	// example:
	//
	// PrivateLinkEnabled
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key.
	//
	// >  If the filter key is set to **PrivateLinkEnabled**, you must specify a filter value. Valid values: **true*	- and **false**.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEnhanhcedNatGatewayAvailableZonesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListEnhanhcedNatGatewayAvailableZonesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequestFilter) SetKey(v string) *ListEnhanhcedNatGatewayAvailableZonesRequestFilter {
	s.Key = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequestFilter) SetValue(v string) *ListEnhanhcedNatGatewayAvailableZonesRequestFilter {
	s.Value = &v
	return s
}

func (s *ListEnhanhcedNatGatewayAvailableZonesRequestFilter) Validate() error {
	return dara.Validate(s)
}
