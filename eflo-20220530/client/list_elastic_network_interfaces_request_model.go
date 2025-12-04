// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElasticNetworkInterfaceId(v string) *ListElasticNetworkInterfacesRequest
	GetElasticNetworkInterfaceId() *string
	SetIp(v string) *ListElasticNetworkInterfacesRequest
	GetIp() *string
	SetNetworkType(v string) *ListElasticNetworkInterfacesRequest
	GetNetworkType() *string
	SetNodeId(v string) *ListElasticNetworkInterfacesRequest
	GetNodeId() *string
	SetPageNumber(v int32) *ListElasticNetworkInterfacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListElasticNetworkInterfacesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListElasticNetworkInterfacesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListElasticNetworkInterfacesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListElasticNetworkInterfacesRequest
	GetStatus() *string
	SetTag(v []*ListElasticNetworkInterfacesRequestTag) *ListElasticNetworkInterfacesRequest
	GetTag() []*ListElasticNetworkInterfacesRequestTag
	SetType(v string) *ListElasticNetworkInterfacesRequest
	GetType() *string
	SetVSwitchId(v string) *ListElasticNetworkInterfacesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ListElasticNetworkInterfacesRequest
	GetVpcId() *string
	SetZoneId(v string) *ListElasticNetworkInterfacesRequest
	GetZoneId() *string
}

type ListElasticNetworkInterfacesRequest struct {
	// Lingjun Elastic Network Interface ID
	//
	// example:
	//
	// leni-1234****
	ElasticNetworkInterfaceId *string `json:"ElasticNetworkInterfaceId,omitempty" xml:"ElasticNetworkInterfaceId,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 10.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The network type.
	//
	// Valid value:
	//
	// 	- Tenant: Tenant.
	//
	// 	- VPC
	//
	// example:
	//
	// tenant
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// e01-cn-lbj3aej****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzmcpv7odnta
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the enterprise-level snapshot policy.
	//
	// Valid value:
	//
	// 	- Create Failed: the creation failure.
	//
	// 	- Delete Failed: the that failed to be deleted.
	//
	// 	- Executing
	//
	// 	- Available: The template is available.
	//
	// 	- Deleting
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of Tags
	Tag []*ListElasticNetworkInterfacesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the variable.
	//
	// Valid value:
	//
	// 	- CUSTOM: custom type.
	//
	// 	- DEFAULT: system type.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6u8473r84e9****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-uf6aa4ddo97fr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListElasticNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesRequest) GetElasticNetworkInterfaceId() *string {
	return s.ElasticNetworkInterfaceId
}

func (s *ListElasticNetworkInterfacesRequest) GetIp() *string {
	return s.Ip
}

func (s *ListElasticNetworkInterfacesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListElasticNetworkInterfacesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListElasticNetworkInterfacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListElasticNetworkInterfacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListElasticNetworkInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListElasticNetworkInterfacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListElasticNetworkInterfacesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListElasticNetworkInterfacesRequest) GetTag() []*ListElasticNetworkInterfacesRequestTag {
	return s.Tag
}

func (s *ListElasticNetworkInterfacesRequest) GetType() *string {
	return s.Type
}

func (s *ListElasticNetworkInterfacesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListElasticNetworkInterfacesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListElasticNetworkInterfacesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListElasticNetworkInterfacesRequest) SetElasticNetworkInterfaceId(v string) *ListElasticNetworkInterfacesRequest {
	s.ElasticNetworkInterfaceId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetIp(v string) *ListElasticNetworkInterfacesRequest {
	s.Ip = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetNetworkType(v string) *ListElasticNetworkInterfacesRequest {
	s.NetworkType = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetNodeId(v string) *ListElasticNetworkInterfacesRequest {
	s.NodeId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetPageNumber(v int32) *ListElasticNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetPageSize(v int32) *ListElasticNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetRegionId(v string) *ListElasticNetworkInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetResourceGroupId(v string) *ListElasticNetworkInterfacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetStatus(v string) *ListElasticNetworkInterfacesRequest {
	s.Status = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetTag(v []*ListElasticNetworkInterfacesRequestTag) *ListElasticNetworkInterfacesRequest {
	s.Tag = v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetType(v string) *ListElasticNetworkInterfacesRequest {
	s.Type = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetVSwitchId(v string) *ListElasticNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetVpcId(v string) *ListElasticNetworkInterfacesRequest {
	s.VpcId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) SetZoneId(v string) *ListElasticNetworkInterfacesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequest) Validate() error {
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

type ListElasticNetworkInterfacesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key-test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// key-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListElasticNetworkInterfacesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListElasticNetworkInterfacesRequestTag) GoString() string {
	return s.String()
}

func (s *ListElasticNetworkInterfacesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListElasticNetworkInterfacesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListElasticNetworkInterfacesRequestTag) SetKey(v string) *ListElasticNetworkInterfacesRequestTag {
	s.Key = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequestTag) SetValue(v string) *ListElasticNetworkInterfacesRequestTag {
	s.Value = &v
	return s
}

func (s *ListElasticNetworkInterfacesRequestTag) Validate() error {
	return dara.Validate(s)
}
