// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayDnsRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayDnsRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayDnsRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayDnsRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayDnsRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayDnsRequest struct {
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
	// sag-sv487b7lno6go5****
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

func (s ViewSmartAccessGatewayDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayDnsRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayDnsRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayDnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayDnsRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayDnsRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayDnsRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayDnsRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayDnsRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsRequest) SetRegionId(v string) *ViewSmartAccessGatewayDnsRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsRequest) SetResourceUid(v string) *ViewSmartAccessGatewayDnsRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsRequest) SetSagInsId(v string) *ViewSmartAccessGatewayDnsRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsRequest) SetSagSn(v string) *ViewSmartAccessGatewayDnsRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayDnsRequest) Validate() error {
	return dara.Validate(s)
}
