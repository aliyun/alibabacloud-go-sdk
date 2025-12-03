// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultimodeCmsUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetMultimodeCmsUrlRequest
	GetClusterId() *string
	SetOwnerId(v int64) *GetMultimodeCmsUrlRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetMultimodeCmsUrlRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetMultimodeCmsUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMultimodeCmsUrlRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *GetMultimodeCmsUrlRequest
	GetZoneId() *string
}

type GetMultimodeCmsUrlRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultimodeCmsUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultimodeCmsUrlRequest) GoString() string {
	return s.String()
}

func (s *GetMultimodeCmsUrlRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMultimodeCmsUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMultimodeCmsUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMultimodeCmsUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMultimodeCmsUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMultimodeCmsUrlRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetMultimodeCmsUrlRequest) SetClusterId(v string) *GetMultimodeCmsUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetOwnerId(v int64) *GetMultimodeCmsUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetRegionId(v string) *GetMultimodeCmsUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetResourceOwnerAccount(v string) *GetMultimodeCmsUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetResourceOwnerId(v int64) *GetMultimodeCmsUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) SetZoneId(v string) *GetMultimodeCmsUrlRequest {
	s.ZoneId = &v
	return s
}

func (s *GetMultimodeCmsUrlRequest) Validate() error {
	return dara.Validate(s)
}
