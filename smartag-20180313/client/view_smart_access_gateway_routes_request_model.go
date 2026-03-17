// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayRoutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *ViewSmartAccessGatewayRoutesRequest
	GetCrossAccount() *bool
	SetRegionId(v string) *ViewSmartAccessGatewayRoutesRequest
	GetRegionId() *string
	SetResourceUid(v string) *ViewSmartAccessGatewayRoutesRequest
	GetResourceUid() *string
	SetSagInsId(v string) *ViewSmartAccessGatewayRoutesRequest
	GetSagInsId() *string
	SetSagSn(v string) *ViewSmartAccessGatewayRoutesRequest
	GetSagSn() *string
}

type ViewSmartAccessGatewayRoutesRequest struct {
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
	// 147304382796****
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

func (s ViewSmartAccessGatewayRoutesRequest) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayRoutesRequest) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayRoutesRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *ViewSmartAccessGatewayRoutesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ViewSmartAccessGatewayRoutesRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *ViewSmartAccessGatewayRoutesRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *ViewSmartAccessGatewayRoutesRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *ViewSmartAccessGatewayRoutesRequest) SetCrossAccount(v bool) *ViewSmartAccessGatewayRoutesRequest {
	s.CrossAccount = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesRequest) SetRegionId(v string) *ViewSmartAccessGatewayRoutesRequest {
	s.RegionId = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesRequest) SetResourceUid(v string) *ViewSmartAccessGatewayRoutesRequest {
	s.ResourceUid = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesRequest) SetSagInsId(v string) *ViewSmartAccessGatewayRoutesRequest {
	s.SagInsId = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesRequest) SetSagSn(v string) *ViewSmartAccessGatewayRoutesRequest {
	s.SagSn = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesRequest) Validate() error {
	return dara.Validate(s)
}
