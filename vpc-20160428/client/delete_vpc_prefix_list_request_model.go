// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpcPrefixListRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteVpcPrefixListRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteVpcPrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteVpcPrefixListRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *DeleteVpcPrefixListRequest
	GetPrefixListId() *string
	SetRegionId(v string) *DeleteVpcPrefixListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpcPrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpcPrefixListRequest
	GetResourceOwnerId() *int64
}

type DeleteVpcPrefixListRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without deleting the prefix list. The system checks the required parameters, request syntax, and business restrictions. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned and the prefix list is deleted.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the prefix list to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0b78hw45f****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list to delete.
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

func (s DeleteVpcPrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPrefixListRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcPrefixListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpcPrefixListRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteVpcPrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpcPrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVpcPrefixListRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DeleteVpcPrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpcPrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpcPrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpcPrefixListRequest) SetClientToken(v string) *DeleteVpcPrefixListRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetDryRun(v bool) *DeleteVpcPrefixListRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetOwnerAccount(v string) *DeleteVpcPrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetOwnerId(v int64) *DeleteVpcPrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetPrefixListId(v string) *DeleteVpcPrefixListRequest {
	s.PrefixListId = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetRegionId(v string) *DeleteVpcPrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetResourceOwnerAccount(v string) *DeleteVpcPrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) SetResourceOwnerId(v int64) *DeleteVpcPrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpcPrefixListRequest) Validate() error {
	return dara.Validate(s)
}
