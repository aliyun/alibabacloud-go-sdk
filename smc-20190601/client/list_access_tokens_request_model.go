// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenId(v []*string) *ListAccessTokensRequest
	GetAccessTokenId() []*string
	SetName(v string) *ListAccessTokensRequest
	GetName() *string
	SetOwnerId(v int64) *ListAccessTokensRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAccessTokensRequest
	GetResourceOwnerAccount() *string
	SetStatus(v string) *ListAccessTokensRequest
	GetStatus() *string
}

type ListAccessTokensRequest struct {
	// The information about activation codes.
	AccessTokenId []*string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty" type:"Repeated"`
	// The name of the activation code.
	//
	// example:
	//
	// test_name
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The status of the activation code. Valid values:
	//
	// 	- activated
	//
	// 	- unactivated
	//
	// 	- expired
	//
	// example:
	//
	// activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAccessTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessTokensRequest) GoString() string {
	return s.String()
}

func (s *ListAccessTokensRequest) GetAccessTokenId() []*string {
	return s.AccessTokenId
}

func (s *ListAccessTokensRequest) GetName() *string {
	return s.Name
}

func (s *ListAccessTokensRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAccessTokensRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAccessTokensRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAccessTokensRequest) SetAccessTokenId(v []*string) *ListAccessTokensRequest {
	s.AccessTokenId = v
	return s
}

func (s *ListAccessTokensRequest) SetName(v string) *ListAccessTokensRequest {
	s.Name = &v
	return s
}

func (s *ListAccessTokensRequest) SetOwnerId(v int64) *ListAccessTokensRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAccessTokensRequest) SetResourceOwnerAccount(v string) *ListAccessTokensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAccessTokensRequest) SetStatus(v string) *ListAccessTokensRequest {
	s.Status = &v
	return s
}

func (s *ListAccessTokensRequest) Validate() error {
	return dara.Validate(s)
}
