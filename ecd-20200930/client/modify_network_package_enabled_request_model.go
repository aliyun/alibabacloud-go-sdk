// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkPackageEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ModifyNetworkPackageEnabledRequest
	GetEnabled() *bool
	SetNetworkPackageId(v string) *ModifyNetworkPackageEnabledRequest
	GetNetworkPackageId() *string
	SetRegionId(v string) *ModifyNetworkPackageEnabledRequest
	GetRegionId() *string
}

type ModifyNetworkPackageEnabledRequest struct {
	// Specifies whether to restore the premium bandwidth plan of the cloud computer.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the premium bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-cfedn7r2pe48g****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyNetworkPackageEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkPackageEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkPackageEnabledRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyNetworkPackageEnabledRequest) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *ModifyNetworkPackageEnabledRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkPackageEnabledRequest) SetEnabled(v bool) *ModifyNetworkPackageEnabledRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetNetworkPackageId(v string) *ModifyNetworkPackageEnabledRequest {
	s.NetworkPackageId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) SetRegionId(v string) *ModifyNetworkPackageEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkPackageEnabledRequest) Validate() error {
	return dara.Validate(s)
}
