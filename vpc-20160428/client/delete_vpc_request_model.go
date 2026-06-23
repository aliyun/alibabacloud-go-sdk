// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVpcRequest
	GetDryRun() *bool
	SetForceDelete(v bool) *DeleteVpcRequest
	GetForceDelete() *bool
	SetOwnerAccount(v string) *DeleteVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DeleteVpcRequest
	GetVpcId() *string
}

type DeleteVpcRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without actually deleting the VPC. The system checks whether the required parameters are set, the request format is valid, and business restrictions are met. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned, and the VPC is directly deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the VPC. Valid values:
	//
	// - **true**: forcefully deletes the VPC.
	//
	// - **false*	- (default): does not forcefully delete the VPC.
	//
	// The VPC can be forcefully deleted only when the following resources exist in the VPC:
	//
	// - The VPC contains only an IPv4 gateway and routes pointing to the IPv4 gateway.
	//
	// - The VPC contains only an IPv6 gateway and routes pointing to the IPv6 gateway.
	//
	// example:
	//
	// false
	ForceDelete  *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC to be deleted.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1m7v25emi1h5mtc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVpcRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DeleteVpcRequest) SetClientToken(v string) *DeleteVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcRequest) SetDryRun(v bool) *DeleteVpcRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcRequest) SetForceDelete(v bool) *DeleteVpcRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteVpcRequest) SetOwnerAccount(v string) *DeleteVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpcRequest) SetOwnerId(v int64) *DeleteVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpcRequest) SetRegionId(v string) *DeleteVpcRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcRequest) SetResourceOwnerAccount(v string) *DeleteVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpcRequest) SetResourceOwnerId(v int64) *DeleteVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpcRequest) SetVpcId(v string) *DeleteVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteVpcRequest) Validate() error {
	return dara.Validate(s)
}
