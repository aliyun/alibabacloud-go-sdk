// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayOspfRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayOspfRouteRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayOspfRouteRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayOspfRouteRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayOspfRouteRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayOspfRouteRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayOspfRouteRequest struct {
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
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 109790620697****
	ResourceUid *string `json:"ResourceUid,omitempty" xml:"ResourceUid,omitempty"`
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

func (s ViewSmartAccessGatewayOspfRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayOspfRouteRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayOspfRouteRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) SetRegionId(v string) *ViewSmartAccessGatewayOspfRouteRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) SetResourceUid(v string) *ViewSmartAccessGatewayOspfRouteRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) SetSagInsId(v string) *ViewSmartAccessGatewayOspfRouteRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) SetSagSn(v string) *ViewSmartAccessGatewayOspfRouteRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayOspfRouteRequest) Validate() error {
	return dara.Validate(s)
}
