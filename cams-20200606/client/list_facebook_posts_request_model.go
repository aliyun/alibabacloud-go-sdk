// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFacebookPostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListFacebookPostsRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListFacebookPostsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListFacebookPostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFacebookPostsRequest
	GetResourceOwnerId() *int64
}

type ListFacebookPostsRequest struct {
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

func (s ListFacebookPostsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFacebookPostsRequest) GoString() string {
	return s.String()
}

func (s *ListFacebookPostsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListFacebookPostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFacebookPostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFacebookPostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFacebookPostsRequest) SetCustSpaceId(v string) *ListFacebookPostsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListFacebookPostsRequest) SetOwnerId(v int64) *ListFacebookPostsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFacebookPostsRequest) SetResourceOwnerAccount(v string) *ListFacebookPostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFacebookPostsRequest) SetResourceOwnerId(v int64) *ListFacebookPostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFacebookPostsRequest) Validate() error {
	return dara.Validate(s)
}
