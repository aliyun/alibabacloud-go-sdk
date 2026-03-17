// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayGlobalRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayGlobalRouteProtocolRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayGlobalRouteProtocolRequest struct {
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
	// sag-iv408aov6k7xxm****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4ea****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayGlobalRouteProtocolRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) SetRegionId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) SetResourceUid(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) SetSagInsId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) SetSagSn(v string) *ViewSmartAccessGatewayGlobalRouteProtocolRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
