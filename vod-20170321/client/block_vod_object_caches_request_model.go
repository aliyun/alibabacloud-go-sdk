// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockVodObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxage(v int32) *BlockVodObjectCachesRequest
	GetMaxage() *int32
	SetObjectPath(v string) *BlockVodObjectCachesRequest
	GetObjectPath() *string
	SetOwnerId(v int64) *BlockVodObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BlockVodObjectCachesRequest
	GetSecurityToken() *string
	SetType(v string) *BlockVodObjectCachesRequest
	GetType() *string
}

type BlockVodObjectCachesRequest struct {
	Maxage *int32 `json:"Maxage,omitempty" xml:"Maxage,omitempty"`
	// This parameter is required.
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BlockVodObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s BlockVodObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *BlockVodObjectCachesRequest) GetMaxage() *int32 {
	return s.Maxage
}

func (s *BlockVodObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *BlockVodObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BlockVodObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BlockVodObjectCachesRequest) GetType() *string {
	return s.Type
}

func (s *BlockVodObjectCachesRequest) SetMaxage(v int32) *BlockVodObjectCachesRequest {
	s.Maxage = &v
	return s
}

func (s *BlockVodObjectCachesRequest) SetObjectPath(v string) *BlockVodObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *BlockVodObjectCachesRequest) SetOwnerId(v int64) *BlockVodObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *BlockVodObjectCachesRequest) SetSecurityToken(v string) *BlockVodObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *BlockVodObjectCachesRequest) SetType(v string) *BlockVodObjectCachesRequest {
	s.Type = &v
	return s
}

func (s *BlockVodObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
