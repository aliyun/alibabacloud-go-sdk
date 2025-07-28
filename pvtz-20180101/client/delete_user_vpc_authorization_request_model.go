// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserVpcAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DeleteUserVpcAuthorizationRequest
	GetAuthType() *string
	SetAuthorizedUserId(v int64) *DeleteUserVpcAuthorizationRequest
	GetAuthorizedUserId() *int64
}

type DeleteUserVpcAuthorizationRequest struct {
	// The authorization scope. Valid values:
	//
	// 	- NORMAL: general authorization
	//
	// 	- NORMAL: cloud service-related authorization
	//
	// Default value: NORMAL.
	//
	// example:
	//
	// NORMAL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 141339776561****
	AuthorizedUserId *int64 `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
}

func (s DeleteUserVpcAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserVpcAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserVpcAuthorizationRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DeleteUserVpcAuthorizationRequest) GetAuthorizedUserId() *int64 {
	return s.AuthorizedUserId
}

func (s *DeleteUserVpcAuthorizationRequest) SetAuthType(v string) *DeleteUserVpcAuthorizationRequest {
	s.AuthType = &v
	return s
}

func (s *DeleteUserVpcAuthorizationRequest) SetAuthorizedUserId(v int64) *DeleteUserVpcAuthorizationRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DeleteUserVpcAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
