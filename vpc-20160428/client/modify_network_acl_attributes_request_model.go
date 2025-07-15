// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkAclAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyNetworkAclAttributesRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyNetworkAclAttributesRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyNetworkAclAttributesRequest
	GetDryRun() *bool
	SetNetworkAclId(v string) *ModifyNetworkAclAttributesRequest
	GetNetworkAclId() *string
	SetNetworkAclName(v string) *ModifyNetworkAclAttributesRequest
	GetNetworkAclName() *string
	SetOwnerAccount(v string) *ModifyNetworkAclAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNetworkAclAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNetworkAclAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNetworkAclAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNetworkAclAttributesRequest
	GetResourceOwnerId() *int64
}

type ModifyNetworkAclAttributesRequest struct {
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
	// The description of the network ACL.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
	// acl-bp1lhl0taikrxxxxxxxx
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ModifyNetworkAclAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkAclAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkAclAttributesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNetworkAclAttributesRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNetworkAclAttributesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyNetworkAclAttributesRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *ModifyNetworkAclAttributesRequest) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *ModifyNetworkAclAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNetworkAclAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNetworkAclAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkAclAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNetworkAclAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNetworkAclAttributesRequest) SetClientToken(v string) *ModifyNetworkAclAttributesRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetDescription(v string) *ModifyNetworkAclAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetDryRun(v bool) *ModifyNetworkAclAttributesRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetNetworkAclId(v string) *ModifyNetworkAclAttributesRequest {
	s.NetworkAclId = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetNetworkAclName(v string) *ModifyNetworkAclAttributesRequest {
	s.NetworkAclName = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetOwnerAccount(v string) *ModifyNetworkAclAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetOwnerId(v int64) *ModifyNetworkAclAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetRegionId(v string) *ModifyNetworkAclAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetResourceOwnerAccount(v string) *ModifyNetworkAclAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) SetResourceOwnerId(v int64) *ModifyNetworkAclAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNetworkAclAttributesRequest) Validate() error {
	return dara.Validate(s)
}
