// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewayClientUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSmartAccessGatewayClientUsersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSmartAccessGatewayClientUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSmartAccessGatewayClientUsersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayClientUsersRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetSmartAGId() *string
	SetUserMail(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetUserMail() *string
	SetUserName(v string) *DescribeSmartAccessGatewayClientUsersRequest
	GetUserName() *string
}

type DescribeSmartAccessGatewayClientUsersRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-kzo5dvms3dqii3*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The email address of the client account.
	//
	// example:
	//
	// username@example.com
	UserMail *string `json:"UserMail,omitempty" xml:"UserMail,omitempty"`
	// The username of the client account.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSmartAccessGatewayClientUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewayClientUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetUserMail() *string {
	return s.UserMail
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetOwnerAccount(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetOwnerId(v int64) *DescribeSmartAccessGatewayClientUsersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetPageNumber(v int32) *DescribeSmartAccessGatewayClientUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetPageSize(v int32) *DescribeSmartAccessGatewayClientUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetRegionId(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetResourceOwnerId(v int64) *DescribeSmartAccessGatewayClientUsersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetSmartAGId(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetUserMail(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.UserMail = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) SetUserName(v string) *DescribeSmartAccessGatewayClientUsersRequest {
	s.UserName = &v
	return s
}

func (s *DescribeSmartAccessGatewayClientUsersRequest) Validate() error {
	return dara.Validate(s)
}
