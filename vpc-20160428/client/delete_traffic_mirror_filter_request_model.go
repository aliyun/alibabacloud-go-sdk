// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTrafficMirrorFilterRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTrafficMirrorFilterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTrafficMirrorFilterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTrafficMirrorFilterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteTrafficMirrorFilterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteTrafficMirrorFilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTrafficMirrorFilterRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorFilterId(v string) *DeleteTrafficMirrorFilterRequest
	GetTrafficMirrorFilterId() *string
}

type DeleteTrafficMirrorFilterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a check request without deleting the traffic mirror filter. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the traffic mirror filter is deleted.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the traffic mirror. You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list. For more information about regions that support traffic mirroring, see [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The instance ID of the traffic mirror filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
}

func (s DeleteTrafficMirrorFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTrafficMirrorFilterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTrafficMirrorFilterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTrafficMirrorFilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTrafficMirrorFilterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTrafficMirrorFilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTrafficMirrorFilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTrafficMirrorFilterRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *DeleteTrafficMirrorFilterRequest) SetClientToken(v string) *DeleteTrafficMirrorFilterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetDryRun(v bool) *DeleteTrafficMirrorFilterRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetOwnerAccount(v string) *DeleteTrafficMirrorFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetOwnerId(v int64) *DeleteTrafficMirrorFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetRegionId(v string) *DeleteTrafficMirrorFilterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetResourceOwnerAccount(v string) *DeleteTrafficMirrorFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetResourceOwnerId(v int64) *DeleteTrafficMirrorFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) SetTrafficMirrorFilterId(v string) *DeleteTrafficMirrorFilterRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *DeleteTrafficMirrorFilterRequest) Validate() error {
	return dara.Validate(s)
}
