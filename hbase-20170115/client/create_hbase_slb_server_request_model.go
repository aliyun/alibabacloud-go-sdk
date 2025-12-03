// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseSlbServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateHbaseSlbServerRequest
	GetClusterId() *string
	SetOwnerId(v int64) *CreateHbaseSlbServerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateHbaseSlbServerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateHbaseSlbServerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHbaseSlbServerRequest
	GetResourceOwnerId() *int64
	SetSlbServer(v string) *CreateHbaseSlbServerRequest
	GetSlbServer() *string
	SetZoneId(v string) *CreateHbaseSlbServerRequest
	GetZoneId() *string
}

type CreateHbaseSlbServerRequest struct {
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

func (s CreateHbaseSlbServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseSlbServerRequest) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHbaseSlbServerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHbaseSlbServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHbaseSlbServerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHbaseSlbServerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHbaseSlbServerRequest) GetSlbServer() *string {
	return s.SlbServer
}

func (s *CreateHbaseSlbServerRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateHbaseSlbServerRequest) SetClusterId(v string) *CreateHbaseSlbServerRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetOwnerId(v int64) *CreateHbaseSlbServerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetRegionId(v string) *CreateHbaseSlbServerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetResourceOwnerAccount(v string) *CreateHbaseSlbServerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetResourceOwnerId(v int64) *CreateHbaseSlbServerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetSlbServer(v string) *CreateHbaseSlbServerRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) SetZoneId(v string) *CreateHbaseSlbServerRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateHbaseSlbServerRequest) Validate() error {
	return dara.Validate(s)
}
