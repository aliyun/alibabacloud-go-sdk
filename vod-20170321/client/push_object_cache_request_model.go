// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushObjectCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectPath(v string) *PushObjectCacheRequest
	GetObjectPath() *string
	SetOwnerAccount(v string) *PushObjectCacheRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *PushObjectCacheRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *PushObjectCacheRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *PushObjectCacheRequest
	GetResourceOwnerId() *string
}

type PushObjectCacheRequest struct {
	// This parameter is required.
	ObjectPath           *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PushObjectCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s PushObjectCacheRequest) GoString() string {
	return s.String()
}

func (s *PushObjectCacheRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *PushObjectCacheRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PushObjectCacheRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *PushObjectCacheRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PushObjectCacheRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *PushObjectCacheRequest) SetObjectPath(v string) *PushObjectCacheRequest {
	s.ObjectPath = &v
	return s
}

func (s *PushObjectCacheRequest) SetOwnerAccount(v string) *PushObjectCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PushObjectCacheRequest) SetOwnerId(v string) *PushObjectCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *PushObjectCacheRequest) SetResourceOwnerAccount(v string) *PushObjectCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PushObjectCacheRequest) SetResourceOwnerId(v string) *PushObjectCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PushObjectCacheRequest) Validate() error {
	return dara.Validate(s)
}
