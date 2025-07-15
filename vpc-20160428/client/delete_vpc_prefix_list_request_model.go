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
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails to pass the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0b78hw45f****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list.
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
