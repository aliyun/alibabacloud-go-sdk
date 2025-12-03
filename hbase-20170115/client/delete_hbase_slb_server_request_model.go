// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseSlbServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteHbaseSlbServerRequest
	GetClusterId() *string
	SetOwnerId(v int64) *DeleteHbaseSlbServerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteHbaseSlbServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteHbaseSlbServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteHbaseSlbServerRequest
	GetResourceOwnerId() *int64
	SetSlbServer(v string) *DeleteHbaseSlbServerRequest
	GetSlbServer() *string
	SetZoneId(v string) *DeleteHbaseSlbServerRequest
	GetZoneId() *string
}

type DeleteHbaseSlbServerRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SlbServer *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteHbaseSlbServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteHbaseSlbServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteHbaseSlbServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHbaseSlbServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteHbaseSlbServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteHbaseSlbServerRequest) GetSlbServer() *string {
	return s.SlbServer
}

func (s *DeleteHbaseSlbServerRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteHbaseSlbServerRequest) SetClusterId(v string) *DeleteHbaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetOwnerId(v int64) *DeleteHbaseSlbServerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetRegionId(v string) *DeleteHbaseSlbServerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetResourceOwnerAccount(v string) *DeleteHbaseSlbServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetResourceOwnerId(v int64) *DeleteHbaseSlbServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetSlbServer(v string) *DeleteHbaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) SetZoneId(v string) *DeleteHbaseSlbServerRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteHbaseSlbServerRequest) Validate() error {
	return dara.Validate(s)
}
