// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSourcesToTrafficMirrorSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddSourcesToTrafficMirrorSessionRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddSourcesToTrafficMirrorSessionRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AddSourcesToTrafficMirrorSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddSourcesToTrafficMirrorSessionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddSourcesToTrafficMirrorSessionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddSourcesToTrafficMirrorSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSourcesToTrafficMirrorSessionRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorSessionId(v string) *AddSourcesToTrafficMirrorSessionRequest
	GetTrafficMirrorSessionId() *string
	SetTrafficMirrorSourceIds(v []*string) *AddSourcesToTrafficMirrorSessionRequest
	GetTrafficMirrorSourceIds() []*string
}

type AddSourcesToTrafficMirrorSessionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the traffic mirror session belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic mirror session.
	//
	// This parameter is required.
	//
	// example:
	//
	// tms-j6cla50buc44ap8tu****
	TrafficMirrorSessionId *string `json:"TrafficMirrorSessionId,omitempty" xml:"TrafficMirrorSessionId,omitempty"`
	// The ID of the traffic mirror source. You can specify only an elastic network interface (ENI) as the traffic mirror source. The default value of **N*	- is **1**, which indicates that you can add only one traffic mirror source to a traffic mirror session.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-j6ccmrl8z3xkvxgw****
	TrafficMirrorSourceIds []*string `json:"TrafficMirrorSourceIds,omitempty" xml:"TrafficMirrorSourceIds,omitempty" type:"Repeated"`
}

func (s AddSourcesToTrafficMirrorSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSourcesToTrafficMirrorSessionRequest) GoString() string {
	return s.String()
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *AddSourcesToTrafficMirrorSessionRequest) GetTrafficMirrorSourceIds() []*string {
	return s.TrafficMirrorSourceIds
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetClientToken(v string) *AddSourcesToTrafficMirrorSessionRequest {
	s.ClientToken = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetDryRun(v bool) *AddSourcesToTrafficMirrorSessionRequest {
	s.DryRun = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetOwnerAccount(v string) *AddSourcesToTrafficMirrorSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetOwnerId(v int64) *AddSourcesToTrafficMirrorSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetRegionId(v string) *AddSourcesToTrafficMirrorSessionRequest {
	s.RegionId = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetResourceOwnerAccount(v string) *AddSourcesToTrafficMirrorSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetResourceOwnerId(v int64) *AddSourcesToTrafficMirrorSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetTrafficMirrorSessionId(v string) *AddSourcesToTrafficMirrorSessionRequest {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) SetTrafficMirrorSourceIds(v []*string) *AddSourcesToTrafficMirrorSessionRequest {
	s.TrafficMirrorSourceIds = v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionRequest) Validate() error {
	return dara.Validate(s)
}
