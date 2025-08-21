// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeVSwitchesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeVSwitchesRequest
	GetEnsRegionIds() []*string
	SetNetworkId(v string) *DescribeVSwitchesRequest
	GetNetworkId() *string
	SetPageNumber(v int32) *DescribeVSwitchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesRequest
	GetPageSize() *int32
	SetVSwitchId(v string) *DescribeVSwitchesRequest
	GetVSwitchId() *string
	SetVSwitchIds(v []*string) *DescribeVSwitchesRequest
	GetVSwitchIds() []*string
	SetVSwitchName(v string) *DescribeVSwitchesRequest
	GetVSwitchName() *string
}

type DescribeVSwitchesRequest struct {
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of edge nodes. You can specify 1 to 100 IDs.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-25cdvfeq58pl****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5m9xhlq8oh***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of vSwitches. You can specify 1 to 100 IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The name of the vSwitch.
	//
	// example:
	//
	// testVSwitchName
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeVSwitchesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeVSwitchesRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeVSwitchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeVSwitchesRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesRequest) SetEnsRegionId(v string) *DescribeVSwitchesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetEnsRegionIds(v []*string) *DescribeVSwitchesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeVSwitchesRequest) SetNetworkId(v string) *DescribeVSwitchesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchIds(v []*string) *DescribeVSwitchesRequest {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchName(v string) *DescribeVSwitchesRequest {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
