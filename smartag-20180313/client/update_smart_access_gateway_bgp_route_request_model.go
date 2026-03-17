// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayBgpRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayBgpRouteRequest
	GetCrossAccount() *bool
	SetHoldTime(v int32) *UpdateSmartAccessGatewayBgpRouteRequest
	GetHoldTime() *int32
	SetKeepAlive(v int32) *UpdateSmartAccessGatewayBgpRouteRequest
	GetKeepAlive() *int32
	SetLocalAs(v int64) *UpdateSmartAccessGatewayBgpRouteRequest
	GetLocalAs() *int64
	SetRegionId(v string) *UpdateSmartAccessGatewayBgpRouteRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayBgpRouteRequest
	GetResourceUid() *string
	SetRouterId(v string) *UpdateSmartAccessGatewayBgpRouteRequest
	GetRouterId() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayBgpRouteRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayBgpRouteRequest
	GetSagSn() *string
}

type UpdateSmartAccessGatewayBgpRouteRequest struct {
	// Specifies whether to query only the SAG instances that belong to another Alibaba Cloud account. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// The hold time. Unit: seconds.
	//
	// Valid values: **3*	- to **65535**.
	//
	// > When the SAG device establishes a peering connection with a peer device, the hold time of both devices must be the same. If the SAG device does not receive a keepalive or update message from the peer device within the hold time, the connection between the BGP peers is closed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	HoldTime *int32 `json:"HoldTime,omitempty" xml:"HoldTime,omitempty"`
	// The time interval at which keep-alive packets are sent. Unit: seconds.
	//
	// Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	KeepAlive *int32 `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The autonomous system number (ASN) of the SAG device.
	//
	// Valid values: **1*	- to **4294967295**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65536
	LocalAs *int64 `json:"LocalAs,omitempty" xml:"LocalAs,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 147304382796****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The ID of the BGP router.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-v9un1ccz22owd7****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dk****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s UpdateSmartAccessGatewayBgpRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayBgpRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetHoldTime() *int32 {
	return s.HoldTime
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetKeepAlive() *int32 {
	return s.KeepAlive
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetLocalAs() *int64 {
	return s.LocalAs
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetHoldTime(v int32) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.HoldTime = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetKeepAlive(v int32) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.KeepAlive = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetLocalAs(v int64) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.LocalAs = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetRegionId(v string) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetRouterId(v string) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.RouterId = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) SetSagSn(v string) *UpdateSmartAccessGatewayBgpRouteRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayBgpRouteRequest) Validate() error {
	return dara.Validate(s)
}
