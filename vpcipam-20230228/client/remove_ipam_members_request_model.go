// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIpamMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveIpamMembersRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveIpamMembersRequest
	GetDryRun() *bool
	SetMembers(v []*RemoveIpamMembersRequestMembers) *RemoveIpamMembersRequest
	GetMembers() []*RemoveIpamMembersRequestMembers
	SetOwnerAccount(v string) *RemoveIpamMembersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveIpamMembersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveIpamMembersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveIpamMembersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveIpamMembersRequest
	GetResourceOwnerId() *int64
}

type RemoveIpamMembersRequest struct {
	// A client token to ensure the idempotence of the request. Generate a value from your client to make sure that the value is unique among different requests. The ClientToken parameter can contain only ASCII characters. Note: If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run. The system checks the required parameters, the request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned. The members are not removed.
	//
	// - **false*	- (default): Sends a request. If the request passes the check, a 2xx HTTP status code is returned, and the members are removed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The members managed by the IPAM trusted service.
	//
	// This parameter is required.
	Members      []*RemoveIpamMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	OwnerAccount *string                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM is hosted. To get the region ID, call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation.
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

func (s RemoveIpamMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpamMembersRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpamMembersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveIpamMembersRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveIpamMembersRequest) GetMembers() []*RemoveIpamMembersRequestMembers {
	return s.Members
}

func (s *RemoveIpamMembersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveIpamMembersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveIpamMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveIpamMembersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveIpamMembersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveIpamMembersRequest) SetClientToken(v string) *RemoveIpamMembersRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetDryRun(v bool) *RemoveIpamMembersRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetMembers(v []*RemoveIpamMembersRequestMembers) *RemoveIpamMembersRequest {
	s.Members = v
	return s
}

func (s *RemoveIpamMembersRequest) SetOwnerAccount(v string) *RemoveIpamMembersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetOwnerId(v int64) *RemoveIpamMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetRegionId(v string) *RemoveIpamMembersRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetResourceOwnerAccount(v string) *RemoveIpamMembersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveIpamMembersRequest) SetResourceOwnerId(v int64) *RemoveIpamMembersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveIpamMembersRequest) Validate() error {
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

type RemoveIpamMembersRequestMembers struct {
	// The member ID.
	//
	// - **Folder ID**: The ID of the folder.
	//
	// - **Account UID**: The UID of the member account in the resource directory.
	//
	// example:
	//
	// Folder
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The type of the member. Valid values:
	//
	// - **Folder**: The member is a folder.
	//
	// - **Account**: The member is a member account in the resource directory.
	//
	// example:
	//
	// fd-ccccncASqa
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s RemoveIpamMembersRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s RemoveIpamMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *RemoveIpamMembersRequestMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *RemoveIpamMembersRequestMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *RemoveIpamMembersRequestMembers) SetMemberId(v string) *RemoveIpamMembersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *RemoveIpamMembersRequestMembers) SetMemberType(v string) *RemoveIpamMembersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *RemoveIpamMembersRequestMembers) Validate() error {
	return dara.Validate(s)
}
