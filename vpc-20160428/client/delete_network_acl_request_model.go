// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteNetworkAclRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteNetworkAclRequest
	GetDryRun() *bool
	SetNetworkAclId(v string) *DeleteNetworkAclRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *DeleteNetworkAclRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNetworkAclRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNetworkAclRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNetworkAclRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNetworkAclRequest
	GetResourceOwnerId() *int64
}

type DeleteNetworkAclRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 223e4867-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-bp1lhl0taikrbgnh****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
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
}

func (s DeleteNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNetworkAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DeleteNetworkAclRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNetworkAclRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNetworkAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkAclRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNetworkAclRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNetworkAclRequest) SetClientToken(v string) *DeleteNetworkAclRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetDryRun(v bool) *DeleteNetworkAclRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetNetworkAclId(v string) *DeleteNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetOwnerAccount(v string) *DeleteNetworkAclRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetOwnerId(v int64) *DeleteNetworkAclRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetRegionId(v string) *DeleteNetworkAclRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetResourceOwnerAccount(v string) *DeleteNetworkAclRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNetworkAclRequest) SetResourceOwnerId(v int64) *DeleteNetworkAclRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}
