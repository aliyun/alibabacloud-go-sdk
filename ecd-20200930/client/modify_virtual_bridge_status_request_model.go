// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualBridgeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v string) *ModifyVirtualBridgeStatusRequest
	GetBridgeId() *string
	SetRegionId(v string) *ModifyVirtualBridgeStatusRequest
	GetRegionId() *string
	SetStatus(v string) *ModifyVirtualBridgeStatusRequest
	GetStatus() *string
}

type ModifyVirtualBridgeStatusRequest struct {
	// The virtual bridge ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vb-sjfiahsiufhisda***
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The virtual bridge status.
	//
	// This parameter is required.
	//
	// example:
	//
	// unuse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyVirtualBridgeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualBridgeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVirtualBridgeStatusRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *ModifyVirtualBridgeStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVirtualBridgeStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyVirtualBridgeStatusRequest) SetBridgeId(v string) *ModifyVirtualBridgeStatusRequest {
	s.BridgeId = &v
	return s
}

func (s *ModifyVirtualBridgeStatusRequest) SetRegionId(v string) *ModifyVirtualBridgeStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVirtualBridgeStatusRequest) SetStatus(v string) *ModifyVirtualBridgeStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyVirtualBridgeStatusRequest) Validate() error {
	return dara.Validate(s)
}
