// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeHaVipsRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeHaVipsRequest
	GetEnsRegionIds() []*string
	SetHaVipAddress(v string) *DescribeHaVipsRequest
	GetHaVipAddress() *string
	SetHaVipId(v string) *DescribeHaVipsRequest
	GetHaVipId() *string
	SetName(v string) *DescribeHaVipsRequest
	GetName() *string
	SetNetworkId(v string) *DescribeHaVipsRequest
	GetNetworkId() *string
	SetPageNumber(v string) *DescribeHaVipsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeHaVipsRequest
	GetPageSize() *string
	SetStatus(v string) *DescribeHaVipsRequest
	GetStatus() *string
	SetVSwitchId(v string) *DescribeHaVipsRequest
	GetVSwitchId() *string
}

type DescribeHaVipsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of ENS nodes. Array length: 1 to 100.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The IP address of the HaVip instance.
	//
	// example:
	//
	// 10.5.XX.XX
	HaVipAddress *string `json:"HaVipAddress,omitempty" xml:"HaVipAddress,omitempty"`
	// The HaVip instance ID.
	//
	// example:
	//
	// havip-5p14t****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The name of the HaVip instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network ID.
	//
	// example:
	//
	// n-57gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Settings for paging.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The HaVip status. Valid values:
	//
	// - Creating: The HaVip is being created.
	//
	// - Available: The HaVip is available.
	//
	// - InUse: The HaVip is in use.
	//
	// - Deleting: The HaVip is being deleted.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeHaVipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeHaVipsRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeHaVipsRequest) GetHaVipAddress() *string {
	return s.HaVipAddress
}

func (s *DescribeHaVipsRequest) GetHaVipId() *string {
	return s.HaVipId
}

func (s *DescribeHaVipsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeHaVipsRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeHaVipsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeHaVipsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeHaVipsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeHaVipsRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeHaVipsRequest) SetEnsRegionId(v string) *DescribeHaVipsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetEnsRegionIds(v []*string) *DescribeHaVipsRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeHaVipsRequest) SetHaVipAddress(v string) *DescribeHaVipsRequest {
	s.HaVipAddress = &v
	return s
}

func (s *DescribeHaVipsRequest) SetHaVipId(v string) *DescribeHaVipsRequest {
	s.HaVipId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetName(v string) *DescribeHaVipsRequest {
	s.Name = &v
	return s
}

func (s *DescribeHaVipsRequest) SetNetworkId(v string) *DescribeHaVipsRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageNumber(v string) *DescribeHaVipsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageSize(v string) *DescribeHaVipsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHaVipsRequest) SetStatus(v string) *DescribeHaVipsRequest {
	s.Status = &v
	return s
}

func (s *DescribeHaVipsRequest) SetVSwitchId(v string) *DescribeHaVipsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeHaVipsRequest) Validate() error {
	return dara.Validate(s)
}
