// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccount(v bool) *UpdateSmartAccessGatewayDnsRequest
	GetCrossAccount() *bool
	SetMasterDns(v string) *UpdateSmartAccessGatewayDnsRequest
	GetMasterDns() *string
	SetRegionId(v string) *UpdateSmartAccessGatewayDnsRequest
	GetRegionId() *string
	SetResourceUid(v string) *UpdateSmartAccessGatewayDnsRequest
	GetResourceUid() *string
	SetSagInsId(v string) *UpdateSmartAccessGatewayDnsRequest
	GetSagInsId() *string
	SetSagSn(v string) *UpdateSmartAccessGatewayDnsRequest
	GetSagSn() *string
	SetSlaveDns(v string) *UpdateSmartAccessGatewayDnsRequest
	GetSlaveDns() *string
}

type UpdateSmartAccessGatewayDnsRequest struct {
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
	// The IP address of the primary DNS server.
	//
	// example:
	//
	// 10.10.XX.XX
	MasterDns *string `json:"MasterDns,omitempty" xml:"MasterDns,omitempty"`
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
	// sag-3manef62evrfr6****
	SagInsId *string `json:"SagInsId,omitempty" xml:"SagInsId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sagf4ea****
	SagSn *string `json:"SagSn,omitempty" xml:"SagSn,omitempty"`
	// The IP address of the secondary DNS server.
	//
	// example:
	//
	// 10.10.XX.XX
	SlaveDns *string `json:"SlaveDns,omitempty" xml:"SlaveDns,omitempty"`
}

func (s UpdateSmartAccessGatewayDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayDnsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetCrossAccount() *bool {
	return s.CrossAccount
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetMasterDns() *string {
	return s.MasterDns
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetResourceUid() *string {
	return s.ResourceUid
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetSagInsId() *string {
	return s.SagInsId
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetSagSn() *string {
	return s.SagSn
}

func (s *UpdateSmartAccessGatewayDnsRequest) GetSlaveDns() *string {
	return s.SlaveDns
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetCrossAccount(v bool) *UpdateSmartAccessGatewayDnsRequest {
	s.CrossAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetMasterDns(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.MasterDns = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetRegionId(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetResourceUid(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.ResourceUid = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetSagInsId(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.SagInsId = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetSagSn(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.SagSn = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) SetSlaveDns(v string) *UpdateSmartAccessGatewayDnsRequest {
	s.SlaveDns = &v
	return s
}

func (s *UpdateSmartAccessGatewayDnsRequest) Validate() error {
	return dara.Validate(s)
}
