// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrivateDnsEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *ModifyPrivateDnsEndpointRequest
	GetAccessInstanceId() *string
	SetAccessInstanceName(v string) *ModifyPrivateDnsEndpointRequest
	GetAccessInstanceName() *string
	SetPrimaryDns(v string) *ModifyPrivateDnsEndpointRequest
	GetPrimaryDns() *string
	SetPrivateDnsType(v string) *ModifyPrivateDnsEndpointRequest
	GetPrivateDnsType() *string
	SetRegionNo(v string) *ModifyPrivateDnsEndpointRequest
	GetRegionNo() *string
	SetStandbyDns(v string) *ModifyPrivateDnsEndpointRequest
	GetStandbyDns() *string
}

type ModifyPrivateDnsEndpointRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// The name of the private instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	AccessInstanceName *string `json:"AccessInstanceName,omitempty" xml:"AccessInstanceName,omitempty"`
	// The primary DNS server.
	//
	// example:
	//
	// 1.1.1.1
	PrimaryDns *string `json:"PrimaryDns,omitempty" xml:"PrimaryDns,omitempty"`
	// The private DNS type. Valid values:
	//
	// - **PrivateZone**
	//
	// - **Custom**
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	PrivateDnsType *string `json:"PrivateDnsType,omitempty" xml:"PrivateDnsType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The secondary DNS server.
	//
	// example:
	//
	// 1.1.1.2
	StandbyDns *string `json:"StandbyDns,omitempty" xml:"StandbyDns,omitempty"`
}

func (s ModifyPrivateDnsEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrivateDnsEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrivateDnsEndpointRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *ModifyPrivateDnsEndpointRequest) GetAccessInstanceName() *string {
	return s.AccessInstanceName
}

func (s *ModifyPrivateDnsEndpointRequest) GetPrimaryDns() *string {
	return s.PrimaryDns
}

func (s *ModifyPrivateDnsEndpointRequest) GetPrivateDnsType() *string {
	return s.PrivateDnsType
}

func (s *ModifyPrivateDnsEndpointRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *ModifyPrivateDnsEndpointRequest) GetStandbyDns() *string {
	return s.StandbyDns
}

func (s *ModifyPrivateDnsEndpointRequest) SetAccessInstanceId(v string) *ModifyPrivateDnsEndpointRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) SetAccessInstanceName(v string) *ModifyPrivateDnsEndpointRequest {
	s.AccessInstanceName = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) SetPrimaryDns(v string) *ModifyPrivateDnsEndpointRequest {
	s.PrimaryDns = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) SetPrivateDnsType(v string) *ModifyPrivateDnsEndpointRequest {
	s.PrivateDnsType = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) SetRegionNo(v string) *ModifyPrivateDnsEndpointRequest {
	s.RegionNo = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) SetStandbyDns(v string) *ModifyPrivateDnsEndpointRequest {
	s.StandbyDns = &v
	return s
}

func (s *ModifyPrivateDnsEndpointRequest) Validate() error {
	return dara.Validate(s)
}
