// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectPath(v string) *RefreshObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshObjectCachesRequest
	GetObjectType() *string
	SetOwnerAccount(v string) *RefreshObjectCachesRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *RefreshObjectCachesRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *RefreshObjectCachesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *RefreshObjectCachesRequest
	GetResourceOwnerId() *string
}

type RefreshObjectCachesRequest struct {
	// This parameter is required.
	ObjectPath           *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType           *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefreshObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshObjectCachesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RefreshObjectCachesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *RefreshObjectCachesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RefreshObjectCachesRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *RefreshObjectCachesRequest) SetObjectPath(v string) *RefreshObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectType(v string) *RefreshObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetOwnerAccount(v string) *RefreshObjectCachesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetOwnerId(v string) *RefreshObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetResourceOwnerAccount(v string) *RefreshObjectCachesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetResourceOwnerId(v string) *RefreshObjectCachesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
