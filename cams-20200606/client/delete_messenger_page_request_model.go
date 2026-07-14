// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessengerPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteMessengerPageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DeleteMessengerPageRequest
	GetOwnerId() *int64
	SetPageId(v string) *DeleteMessengerPageRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *DeleteMessengerPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMessengerPageRequest
	GetResourceOwnerId() *int64
}

type DeleteMessengerPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e01-in-71cc15xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 154332432432
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMessengerPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessengerPageRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessengerPageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteMessengerPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMessengerPageRequest) GetPageId() *string {
	return s.PageId
}

func (s *DeleteMessengerPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMessengerPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMessengerPageRequest) SetInstanceId(v string) *DeleteMessengerPageRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMessengerPageRequest) SetOwnerId(v int64) *DeleteMessengerPageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMessengerPageRequest) SetPageId(v string) *DeleteMessengerPageRequest {
	s.PageId = &v
	return s
}

func (s *DeleteMessengerPageRequest) SetResourceOwnerAccount(v string) *DeleteMessengerPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMessengerPageRequest) SetResourceOwnerId(v int64) *DeleteMessengerPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMessengerPageRequest) Validate() error {
	return dara.Validate(s)
}
