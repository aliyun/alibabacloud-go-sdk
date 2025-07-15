// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTrafficMirrorSessionRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTrafficMirrorSessionRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTrafficMirrorSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTrafficMirrorSessionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTrafficMirrorSessionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTrafficMirrorSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTrafficMirrorSessionRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorSessionId(v string) *DeleteTrafficMirrorSessionRequest
	GetTrafficMirrorSessionId() *string
}

type DeleteTrafficMirrorSessionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the traffic mirror session belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
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
}

func (s DeleteTrafficMirrorSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorSessionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTrafficMirrorSessionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTrafficMirrorSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTrafficMirrorSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTrafficMirrorSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTrafficMirrorSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTrafficMirrorSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTrafficMirrorSessionRequest) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *DeleteTrafficMirrorSessionRequest) SetClientToken(v string) *DeleteTrafficMirrorSessionRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetDryRun(v bool) *DeleteTrafficMirrorSessionRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetOwnerAccount(v string) *DeleteTrafficMirrorSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetOwnerId(v int64) *DeleteTrafficMirrorSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetRegionId(v string) *DeleteTrafficMirrorSessionRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetResourceOwnerAccount(v string) *DeleteTrafficMirrorSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetResourceOwnerId(v int64) *DeleteTrafficMirrorSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) SetTrafficMirrorSessionId(v string) *DeleteTrafficMirrorSessionRequest {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *DeleteTrafficMirrorSessionRequest) Validate() error {
	return dara.Validate(s)
}
