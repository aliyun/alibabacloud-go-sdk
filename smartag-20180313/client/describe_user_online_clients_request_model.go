// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserOnlineClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeUserOnlineClientsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserOnlineClientsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUserOnlineClientsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserOnlineClientsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserOnlineClientsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeUserOnlineClientsRequest
	GetSmartAGId() *string
	SetUserName(v string) *DescribeUserOnlineClientsRequest
	GetUserName() *string
}

type DescribeUserOnlineClientsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the SAG APP instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG APP instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-wfjgn**********
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account. Usernames of client accounts added to the same SAG APP instance are unique.
	//
	// For a client account, if you specify the username, you must also specify the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// doctest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeUserOnlineClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserOnlineClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserOnlineClientsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserOnlineClientsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserOnlineClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserOnlineClientsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserOnlineClientsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserOnlineClientsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeUserOnlineClientsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeUserOnlineClientsRequest) SetOwnerAccount(v string) *DescribeUserOnlineClientsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetOwnerId(v int64) *DescribeUserOnlineClientsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetRegionId(v string) *DescribeUserOnlineClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetResourceOwnerAccount(v string) *DescribeUserOnlineClientsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetResourceOwnerId(v int64) *DescribeUserOnlineClientsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetSmartAGId(v string) *DescribeUserOnlineClientsRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) SetUserName(v string) *DescribeUserOnlineClientsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeUserOnlineClientsRequest) Validate() error {
	return dara.Validate(s)
}
