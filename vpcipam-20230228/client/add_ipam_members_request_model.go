// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddIpamMembersRequest
	GetClientToken() *string
	SetDryRun(v bool) *AddIpamMembersRequest
	GetDryRun() *bool
	SetMembers(v []*AddIpamMembersRequestMembers) *AddIpamMembersRequest
	GetMembers() []*AddIpamMembersRequestMembers
	SetOwnerAccount(v string) *AddIpamMembersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddIpamMembersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddIpamMembersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddIpamMembersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddIpamMembersRequest
	GetResourceOwnerId() *int64
}

type AddIpamMembersRequest struct {
	// A client token to ensure the idempotence of the request.
	//
	// Generate a unique value from your client for each request. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The request is checked for required parameters, request format, and service limits. Members are not added to the IPAM trusted service. If the check passes, the DryRunOperation error code is returned. If the check fails, an error is returned.
	//
	// - **false*	- (Default): Sends a request. If the request passes the check, the members are added to the IPAM trusted service and an HTTP 2xx status code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The members managed by the IPAM trusted service.
	//
	// This parameter is required.
	Members      []*AddIpamMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	OwnerAccount *string                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the IPAM hosted region.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to get the region ID.
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

func (s AddIpamMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIpamMembersRequest) GoString() string {
	return s.String()
}

func (s *AddIpamMembersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddIpamMembersRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AddIpamMembersRequest) GetMembers() []*AddIpamMembersRequestMembers {
	return s.Members
}

func (s *AddIpamMembersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddIpamMembersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddIpamMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddIpamMembersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddIpamMembersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddIpamMembersRequest) SetClientToken(v string) *AddIpamMembersRequest {
	s.ClientToken = &v
	return s
}

func (s *AddIpamMembersRequest) SetDryRun(v bool) *AddIpamMembersRequest {
	s.DryRun = &v
	return s
}

func (s *AddIpamMembersRequest) SetMembers(v []*AddIpamMembersRequestMembers) *AddIpamMembersRequest {
	s.Members = v
	return s
}

func (s *AddIpamMembersRequest) SetOwnerAccount(v string) *AddIpamMembersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddIpamMembersRequest) SetOwnerId(v int64) *AddIpamMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *AddIpamMembersRequest) SetRegionId(v string) *AddIpamMembersRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpamMembersRequest) SetResourceOwnerAccount(v string) *AddIpamMembersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddIpamMembersRequest) SetResourceOwnerId(v int64) *AddIpamMembersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddIpamMembersRequest) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddIpamMembersRequestMembers struct {
	// The member ID.
	//
	// - **Folder ID**: The ID of the folder.
	//
	// - **Account UID**: The UID of a member account in the resource directory.
	//
	// example:
	//
	// fd-ccccncASqa
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The member type. Valid values:
	//
	// - **Folder**: A folder.
	//
	// - **Account**: A member account in a resource directory.
	//
	// example:
	//
	// Folder
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s AddIpamMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddIpamMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddIpamMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *AddIpamMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *AddIpamMembersRequestMembers) SetMemberId(v string) *AddIpamMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *AddIpamMembersRequestMembers) SetMemberType(v string) *AddIpamMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *AddIpamMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}
