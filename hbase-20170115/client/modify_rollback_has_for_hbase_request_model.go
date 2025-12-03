// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRollbackHasForHbaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyRollbackHasForHbaseRequest
	GetClientToken() *string
	SetClusterId(v string) *ModifyRollbackHasForHbaseRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ModifyRollbackHasForHbaseRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRollbackHasForHbaseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRollbackHasForHbaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRollbackHasForHbaseRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyRollbackHasForHbaseRequest
	GetZoneId() *string
}

type ModifyRollbackHasForHbaseRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyRollbackHasForHbaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRollbackHasForHbaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyRollbackHasForHbaseRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyRollbackHasForHbaseRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyRollbackHasForHbaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRollbackHasForHbaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRollbackHasForHbaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRollbackHasForHbaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRollbackHasForHbaseRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyRollbackHasForHbaseRequest) SetClientToken(v string) *ModifyRollbackHasForHbaseRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetClusterId(v string) *ModifyRollbackHasForHbaseRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetOwnerId(v int64) *ModifyRollbackHasForHbaseRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetRegionId(v string) *ModifyRollbackHasForHbaseRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetResourceOwnerAccount(v string) *ModifyRollbackHasForHbaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetResourceOwnerId(v int64) *ModifyRollbackHasForHbaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) SetZoneId(v string) *ModifyRollbackHasForHbaseRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyRollbackHasForHbaseRequest) Validate() error {
	return dara.Validate(s)
}
