// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerlessInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteServerlessInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteServerlessInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DeleteServerlessInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteServerlessInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteServerlessInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteServerlessInstanceRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DeleteServerlessInstanceRequest
	GetZoneId() *string
}

type DeleteServerlessInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteServerlessInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerlessInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerlessInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteServerlessInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteServerlessInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteServerlessInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServerlessInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteServerlessInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteServerlessInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteServerlessInstanceRequest) SetClientToken(v string) *DeleteServerlessInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetInstanceId(v string) *DeleteServerlessInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetOwnerId(v int64) *DeleteServerlessInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetRegionId(v string) *DeleteServerlessInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetResourceOwnerAccount(v string) *DeleteServerlessInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetResourceOwnerId(v int64) *DeleteServerlessInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) SetZoneId(v string) *DeleteServerlessInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteServerlessInstanceRequest) Validate() error {
	return dara.Validate(s)
}
