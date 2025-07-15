// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTrafficMirrorFilterRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTrafficMirrorFilterRulesRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTrafficMirrorFilterRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTrafficMirrorFilterRulesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTrafficMirrorFilterRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTrafficMirrorFilterRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTrafficMirrorFilterRulesRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorFilterId(v string) *DeleteTrafficMirrorFilterRulesRequest
	GetTrafficMirrorFilterId() *string
	SetTrafficMirrorFilterRuleIds(v []*string) *DeleteTrafficMirrorFilterRulesRequest
	GetTrafficMirrorFilterRuleIds() []*string
}

type DeleteTrafficMirrorFilterRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the API request without performing the operation. The system checks the required parameters, request format, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The ID of the inbound or outbound rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmr-j6cbmubn323k7jlq3****
	TrafficMirrorFilterRuleIds []*string `json:"TrafficMirrorFilterRuleIds,omitempty" xml:"TrafficMirrorFilterRuleIds,omitempty" type:"Repeated"`
}

func (s DeleteTrafficMirrorFilterRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *DeleteTrafficMirrorFilterRulesRequest) GetTrafficMirrorFilterRuleIds() []*string {
	return s.TrafficMirrorFilterRuleIds
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetClientToken(v string) *DeleteTrafficMirrorFilterRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetDryRun(v bool) *DeleteTrafficMirrorFilterRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetOwnerAccount(v string) *DeleteTrafficMirrorFilterRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetOwnerId(v int64) *DeleteTrafficMirrorFilterRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetRegionId(v string) *DeleteTrafficMirrorFilterRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetResourceOwnerAccount(v string) *DeleteTrafficMirrorFilterRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetResourceOwnerId(v int64) *DeleteTrafficMirrorFilterRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetTrafficMirrorFilterId(v string) *DeleteTrafficMirrorFilterRulesRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) SetTrafficMirrorFilterRuleIds(v []*string) *DeleteTrafficMirrorFilterRulesRequest {
	s.TrafficMirrorFilterRuleIds = v
	return s
}

func (s *DeleteTrafficMirrorFilterRulesRequest) Validate() error {
	return dara.Validate(s)
}
