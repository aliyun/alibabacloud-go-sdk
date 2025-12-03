// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateGlobalResourceRequest
	GetClusterId() *string
	SetOwnerId(v int64) *CreateGlobalResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateGlobalResourceRequest
	GetRegionId() *string
	SetResourceName(v string) *CreateGlobalResourceRequest
	GetResourceName() *string
	SetResourceOwnerAccount(v string) *CreateGlobalResourceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGlobalResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *CreateGlobalResourceRequest
	GetResourceType() *string
	SetZoneId(v string) *CreateGlobalResourceRequest
	GetZoneId() *string
}

type CreateGlobalResourceRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateGlobalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateGlobalResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGlobalResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGlobalResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateGlobalResourceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGlobalResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGlobalResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateGlobalResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateGlobalResourceRequest) SetClusterId(v string) *CreateGlobalResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetOwnerId(v int64) *CreateGlobalResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetRegionId(v string) *CreateGlobalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceName(v string) *CreateGlobalResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceOwnerAccount(v string) *CreateGlobalResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceOwnerId(v int64) *CreateGlobalResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetResourceType(v string) *CreateGlobalResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateGlobalResourceRequest) SetZoneId(v string) *CreateGlobalResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateGlobalResourceRequest) Validate() error {
	return dara.Validate(s)
}
