// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstagramPostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListInstagramPostsRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListInstagramPostsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListInstagramPostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListInstagramPostsRequest
	GetResourceOwnerId() *int64
}

type ListInstagramPostsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListInstagramPostsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstagramPostsRequest) GoString() string {
	return s.String()
}

func (s *ListInstagramPostsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListInstagramPostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListInstagramPostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListInstagramPostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListInstagramPostsRequest) SetCustSpaceId(v string) *ListInstagramPostsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListInstagramPostsRequest) SetOwnerId(v int64) *ListInstagramPostsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListInstagramPostsRequest) SetResourceOwnerAccount(v string) *ListInstagramPostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListInstagramPostsRequest) SetResourceOwnerId(v int64) *ListInstagramPostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListInstagramPostsRequest) Validate() error {
	return dara.Validate(s)
}
