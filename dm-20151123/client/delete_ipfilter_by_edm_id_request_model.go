// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpfilterByEdmIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *DeleteIpfilterByEdmIdRequest
	GetFromType() *int32
	SetId(v string) *DeleteIpfilterByEdmIdRequest
	GetId() *string
	SetOwnerId(v int64) *DeleteIpfilterByEdmIdRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteIpfilterByEdmIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpfilterByEdmIdRequest
	GetResourceOwnerId() *int64
}

type DeleteIpfilterByEdmIdRequest struct {
	// Deprecated, kept for historical compatibility.
	//
	// example:
	//
	// 1
	FromType *int32 `json:"FromType,omitempty" xml:"FromType,omitempty"`
	// Record ID
	//
	// example:
	//
	// 10120
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteIpfilterByEdmIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpfilterByEdmIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpfilterByEdmIdRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *DeleteIpfilterByEdmIdRequest) GetId() *string {
	return s.Id
}

func (s *DeleteIpfilterByEdmIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpfilterByEdmIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpfilterByEdmIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpfilterByEdmIdRequest) SetFromType(v int32) *DeleteIpfilterByEdmIdRequest {
	s.FromType = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetId(v string) *DeleteIpfilterByEdmIdRequest {
	s.Id = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetOwnerId(v int64) *DeleteIpfilterByEdmIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetResourceOwnerAccount(v string) *DeleteIpfilterByEdmIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) SetResourceOwnerId(v int64) *DeleteIpfilterByEdmIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpfilterByEdmIdRequest) Validate() error {
	return dara.Validate(s)
}
