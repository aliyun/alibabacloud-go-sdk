// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVpcPrefixListAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RetryVpcPrefixListAssociationRequest
	GetClientToken() *string
	SetDryRun(v bool) *RetryVpcPrefixListAssociationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *RetryVpcPrefixListAssociationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RetryVpcPrefixListAssociationRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *RetryVpcPrefixListAssociationRequest
	GetPrefixListId() *string
	SetRegionId(v string) *RetryVpcPrefixListAssociationRequest
	GetRegionId() *string
	SetResourceId(v string) *RetryVpcPrefixListAssociationRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *RetryVpcPrefixListAssociationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RetryVpcPrefixListAssociationRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *RetryVpcPrefixListAssociationRequest
	GetResourceType() *string
}

type RetryVpcPrefixListAssociationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without associating the prefix list. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and then sends the request. If the check succeeds, an HTTP 2xx status code is returned and the prefix list is associated.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the prefix list to redeliver.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list instance to redeliver.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the associated resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1drpcfz9srr393h****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource associated with the prefix list instance. Valid values:
	//
	// - **vpcRouteTable**: VPC route table.
	//
	// - **trRouteTable**: transit router route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcRouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RetryVpcPrefixListAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryVpcPrefixListAssociationRequest) GoString() string {
	return s.String()
}

func (s *RetryVpcPrefixListAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RetryVpcPrefixListAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RetryVpcPrefixListAssociationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RetryVpcPrefixListAssociationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RetryVpcPrefixListAssociationRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *RetryVpcPrefixListAssociationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetryVpcPrefixListAssociationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RetryVpcPrefixListAssociationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RetryVpcPrefixListAssociationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RetryVpcPrefixListAssociationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RetryVpcPrefixListAssociationRequest) SetClientToken(v string) *RetryVpcPrefixListAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetDryRun(v bool) *RetryVpcPrefixListAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetOwnerAccount(v string) *RetryVpcPrefixListAssociationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetOwnerId(v int64) *RetryVpcPrefixListAssociationRequest {
	s.OwnerId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetPrefixListId(v string) *RetryVpcPrefixListAssociationRequest {
	s.PrefixListId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetRegionId(v string) *RetryVpcPrefixListAssociationRequest {
	s.RegionId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetResourceId(v string) *RetryVpcPrefixListAssociationRequest {
	s.ResourceId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetResourceOwnerAccount(v string) *RetryVpcPrefixListAssociationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetResourceOwnerId(v int64) *RetryVpcPrefixListAssociationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) SetResourceType(v string) *RetryVpcPrefixListAssociationRequest {
	s.ResourceType = &v
	return s
}

func (s *RetryVpcPrefixListAssociationRequest) Validate() error {
	return dara.Validate(s)
}
