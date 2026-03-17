// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayWanSnatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayWanSnatRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayWanSnatRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayWanSnatRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayWanSnatRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayWanSnatRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayWanSnatRequest struct {
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
	// example:
	//
	// sag-iv408aov6k7xxm****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sagf4ea****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
}

func (s ViewSmartAccessGatewayWanSnatRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayWanSnatRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayWanSnatRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayWanSnatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayWanSnatRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayWanSnatRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayWanSnatRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayWanSnatRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayWanSnatRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatRequest) SetRegionId(v string) *ViewSmartAccessGatewayWanSnatRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatRequest) SetResourceUid(v string) *ViewSmartAccessGatewayWanSnatRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatRequest) SetSagInsId(v string) *ViewSmartAccessGatewayWanSnatRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatRequest) SetSagSn(v string) *ViewSmartAccessGatewayWanSnatRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatRequest) Validate() error {
	return dara.Validate(s)
}
