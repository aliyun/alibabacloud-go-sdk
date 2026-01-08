// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteGroupByIdRequest
	GetGroupId() *string
	SetOwnerId(v int64) *DeleteGroupByIdRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteGroupByIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteGroupByIdRequest
	GetResourceOwnerId() *int64
}

type DeleteGroupByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 112
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteGroupByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupByIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteGroupByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGroupByIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteGroupByIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteGroupByIdRequest) SetGroupId(v string) *DeleteGroupByIdRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupByIdRequest) SetOwnerId(v int64) *DeleteGroupByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGroupByIdRequest) SetResourceOwnerAccount(v string) *DeleteGroupByIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGroupByIdRequest) SetResourceOwnerId(v int64) *DeleteGroupByIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGroupByIdRequest) Validate() error {
	return dara.Validate(s)
}
