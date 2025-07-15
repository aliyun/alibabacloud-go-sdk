// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchCidrReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVSwitchCidrReservationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVSwitchCidrReservationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteVSwitchCidrReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVSwitchCidrReservationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVSwitchCidrReservationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVSwitchCidrReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVSwitchCidrReservationRequest
	GetResourceOwnerId() *int64
	SetVSwitchCidrReservationId(v string) *DeleteVSwitchCidrReservationRequest
	GetVSwitchCidrReservationId() *string
}

type DeleteVSwitchCidrReservationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run. The system checks the required parameters, request format, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the vSwitch is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the reserved CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
}

func (s DeleteVSwitchCidrReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchCidrReservationRequest) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchCidrReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVSwitchCidrReservationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVSwitchCidrReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVSwitchCidrReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVSwitchCidrReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVSwitchCidrReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVSwitchCidrReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVSwitchCidrReservationRequest) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *DeleteVSwitchCidrReservationRequest) SetClientToken(v string) *DeleteVSwitchCidrReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetDryRun(v bool) *DeleteVSwitchCidrReservationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetOwnerAccount(v string) *DeleteVSwitchCidrReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetOwnerId(v int64) *DeleteVSwitchCidrReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetRegionId(v string) *DeleteVSwitchCidrReservationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetResourceOwnerAccount(v string) *DeleteVSwitchCidrReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetResourceOwnerId(v int64) *DeleteVSwitchCidrReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) SetVSwitchCidrReservationId(v string) *DeleteVSwitchCidrReservationRequest {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *DeleteVSwitchCidrReservationRequest) Validate() error {
	return dara.Validate(s)
}
