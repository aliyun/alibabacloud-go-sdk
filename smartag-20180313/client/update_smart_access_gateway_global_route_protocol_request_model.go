// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayGlobalRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetResourceUid() *string
	SetRouteProtocol(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetRouteProtocol() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest
	GetSagSn() *string
}

type UpdateSmartAccessGatewayGlobalRouteProtocolRequest struct {
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
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**
	//
	// 	- **OSPF**
	//
	// 	- **BGP**
	//
	// This parameter is required.
	//
	// example:
	//
	// BGP
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-3manef62evrfr6****
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

func (s UpdateSmartAccessGatewayGlobalRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetRegionId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetRouteProtocol(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.RouteProtocol = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) SetSagSn(v string) *UpdateSmartAccessGatewayGlobalRouteProtocolRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayGlobalRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
