// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayPortRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayPortRouteProtocolRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayPortRouteProtocolRequest struct {
	// example:
	//
	// false
	CrossAccount *bool `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sag-sv487b7lno6go5****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sagf4dk****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s ViewSmartAccessGatewayPortRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayPortRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayPortRouteProtocolRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) SetRegionId(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) SetResourceUid(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) SetSagInsId(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) SetSagSn(v string) *ViewSmartAccessGatewayPortRouteProtocolRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
