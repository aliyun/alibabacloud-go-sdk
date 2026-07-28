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
	// The language of the response. Valid values:
	//
	// - **zh-CN*	- (default): Chinese.
	//
	// - **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The filter information. You can specify key-value pairs to filter the query results.
	Filter       []*ListEnhanhcedNatGatewayAvailableZonesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This example queries the zones available for NAT gateway resources in the UAE (Dubai) region.
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
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnhanhcedNatGatewayAvailableZonesRequestFilter struct {
	// The filter condition. Currently, only **PrivateLinkEnabled*	- is supported.
	//
	// example:
	//
	// PrivateLinkEnabled
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value that corresponds to the filter condition.
	//
	// > If the filter condition is **PrivateLinkEnabled**, you must specify a filter value. Valid values: **true*	- and **false**.
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
