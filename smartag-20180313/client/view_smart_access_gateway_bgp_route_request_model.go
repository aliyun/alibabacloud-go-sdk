// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayBgpRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayBgpRouteRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayBgpRouteRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayBgpRouteRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayBgpRouteRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayBgpRouteRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayBgpRouteRequest struct {
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

func (s ViewSmartAccessGatewayBgpRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayBgpRouteRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayBgpRouteRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) SetRegionId(v string) *ViewSmartAccessGatewayBgpRouteRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) SetResourceUid(v string) *ViewSmartAccessGatewayBgpRouteRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) SetSagInsId(v string) *ViewSmartAccessGatewayBgpRouteRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) SetSagSn(v string) *ViewSmartAccessGatewayBgpRouteRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteRequest) Validate() error {
	return dara.Validate(s)
}
