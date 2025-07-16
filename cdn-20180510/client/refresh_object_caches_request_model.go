// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *RefreshObjectCachesRequest
	GetForce() *bool
	SetObjectPath(v string) *RefreshObjectCachesRequest
	GetObjectPath() *string
	SetObjectType(v string) *RefreshObjectCachesRequest
	GetObjectType() *string
	SetOwnerId(v int64) *RefreshObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *RefreshObjectCachesRequest
	GetSecurityToken() *string
}

type RefreshObjectCachesRequest struct {
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshObjectCachesRequest) GetForce() *bool {
	return s.Force
}

func (s *RefreshObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *RefreshObjectCachesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *RefreshObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RefreshObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RefreshObjectCachesRequest) SetForce(v bool) *RefreshObjectCachesRequest {
	s.Force = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectPath(v string) *RefreshObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetObjectType(v string) *RefreshObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetOwnerId(v int64) *RefreshObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshObjectCachesRequest) SetSecurityToken(v string) *RefreshObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RefreshObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
