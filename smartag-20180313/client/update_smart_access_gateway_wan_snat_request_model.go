// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayWanSnatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayWanSnatRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *UpdateSmartAccessGatewayWanSnatRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayWanSnatRequest
	GetResourceUid() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayWanSnatRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayWanSnatRequest
	GetSagSn() *string
	SetSnat(v string) *UpdateSmartAccessGatewayWanSnatRequest
	GetSnat() *string
}

type UpdateSmartAccessGatewayWanSnatRequest struct {
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
	// sag-jwbtsyzom0ol4v****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4dk****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
	// Specifies whether to enable SNAT. Valid values:
	//
	// 	- **1**: enables SNAT
	//
	// 	- **0**: disables SNAT
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Snat *string `json:"Snat,omitempty" xml:"Snat,omitempty"`
}

func (s UpdateSmartAccessGatewayWanSnatRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayWanSnatRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) GetSnat() *string {
	return s.Snat
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayWanSnatRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetRegionId(v string) *UpdateSmartAccessGatewayWanSnatRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayWanSnatRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayWanSnatRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetSagSn(v string) *UpdateSmartAccessGatewayWanSnatRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) SetSnat(v string) *UpdateSmartAccessGatewayWanSnatRequest {
	s.Snat = &v
	return s
}

func (s *UpdateSmartAccessGatewayWanSnatRequest) Validate() error {
	return dara.Validate(s)
}
