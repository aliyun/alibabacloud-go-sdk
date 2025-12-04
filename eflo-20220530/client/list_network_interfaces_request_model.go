// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnablePage(v bool) *ListNetworkInterfacesRequest
	GetEnablePage() *bool
	SetIp(v string) *ListNetworkInterfacesRequest
	GetIp() *string
	SetNetworkInterfaceId(v string) *ListNetworkInterfacesRequest
	GetNetworkInterfaceId() *string
	SetNodeId(v string) *ListNetworkInterfacesRequest
	GetNodeId() *string
	SetPageNumber(v int32) *ListNetworkInterfacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNetworkInterfacesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListNetworkInterfacesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListNetworkInterfacesRequest
	GetResourceGroupId() *string
	SetSubnetId(v string) *ListNetworkInterfacesRequest
	GetSubnetId() *string
	SetTag(v []*ListNetworkInterfacesRequestTag) *ListNetworkInterfacesRequest
	GetTag() []*ListNetworkInterfacesRequestTag
	SetVpdId(v string) *ListNetworkInterfacesRequest
	GetVpdId() *string
}

type ListNetworkInterfacesRequest struct {
	// Specifies whether pagination is required.
	//
	// example:
	//
	// false
	EnablePage *bool `json:"EnablePage,omitempty" xml:"EnablePage,omitempty"`
	// network interface controller the IP address.
	//
	// example:
	//
	// 203.107.46.227
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Lingjun network interface controller ID.
	//
	// example:
	//
	// lni-bp18exxqa2rvfn45e5pz
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the machine to which the instance belongs.
	//
	// example:
	//
	// r-2ze121o4uhr4np3r5t-db-5
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The current number of pages.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmzzka6bnjvbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the instance to which the Lingjun subnet belongs.
	//
	// example:
	//
	// subnet-anhtskts
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The list of tags
	Tag []*ListNetworkInterfacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPD.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesRequest) GetEnablePage() *bool {
	return s.EnablePage
}

func (s *ListNetworkInterfacesRequest) GetIp() *string {
	return s.Ip
}

func (s *ListNetworkInterfacesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListNetworkInterfacesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNetworkInterfacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNetworkInterfacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNetworkInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNetworkInterfacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNetworkInterfacesRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ListNetworkInterfacesRequest) GetTag() []*ListNetworkInterfacesRequestTag {
	return s.Tag
}

func (s *ListNetworkInterfacesRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *ListNetworkInterfacesRequest) SetEnablePage(v bool) *ListNetworkInterfacesRequest {
	s.EnablePage = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetIp(v string) *ListNetworkInterfacesRequest {
	s.Ip = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetNetworkInterfaceId(v string) *ListNetworkInterfacesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetNodeId(v string) *ListNetworkInterfacesRequest {
	s.NodeId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetPageNumber(v int32) *ListNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetPageSize(v int32) *ListNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetRegionId(v string) *ListNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetResourceGroupId(v string) *ListNetworkInterfacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetSubnetId(v string) *ListNetworkInterfacesRequest {
	s.SubnetId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) SetTag(v []*ListNetworkInterfacesRequestTag) *ListNetworkInterfacesRequest {
	s.Tag = v
	return s
}

func (s *ListNetworkInterfacesRequest) SetVpdId(v string) *ListNetworkInterfacesRequest {
	s.VpdId = &v
	return s
}

func (s *ListNetworkInterfacesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetworkInterfacesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key-test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNetworkInterfacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListNetworkInterfacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListNetworkInterfacesRequestTag) SetKey(v string) *ListNetworkInterfacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListNetworkInterfacesRequestTag) SetValue(v string) *ListNetworkInterfacesRequestTag {
	s.Value = &v
	return s
}

func (s *ListNetworkInterfacesRequestTag) Validate() error {
	return dara.Validate(s)
}
