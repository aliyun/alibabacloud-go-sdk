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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully delete the VPC. Valid values:
	//
	// - **true**: yes
	//
	// - **false*	- (default): no
	//
	// You can forcefully delete a VPC in the following scenarios:
	//
	// - Only an IPv4 gateway and routes that point to the IPv4 gateway exist in the VPC.
	//
	// - Only an IPv6 gateway and routes that point to the IPv6 gateway exist in the VPC.
	//
	// example:
	//
	// false
	ForceDelete  *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the VPC is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC that you want to delete.
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
