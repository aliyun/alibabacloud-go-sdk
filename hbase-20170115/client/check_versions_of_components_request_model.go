// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVersionsOfComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CheckVersionsOfComponentsRequest
	GetClusterId() *string
	SetComponents(v string) *CheckVersionsOfComponentsRequest
	GetComponents() *string
	SetOwnerId(v int64) *CheckVersionsOfComponentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckVersionsOfComponentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckVersionsOfComponentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckVersionsOfComponentsRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *CheckVersionsOfComponentsRequest
	GetZoneId() *string
}

type CheckVersionsOfComponentsRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckVersionsOfComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckVersionsOfComponentsRequest) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CheckVersionsOfComponentsRequest) GetComponents() *string {
	return s.Components
}

func (s *CheckVersionsOfComponentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckVersionsOfComponentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckVersionsOfComponentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckVersionsOfComponentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckVersionsOfComponentsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckVersionsOfComponentsRequest) SetClusterId(v string) *CheckVersionsOfComponentsRequest {
	s.ClusterId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetComponents(v string) *CheckVersionsOfComponentsRequest {
	s.Components = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetOwnerId(v int64) *CheckVersionsOfComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetRegionId(v string) *CheckVersionsOfComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetResourceOwnerAccount(v string) *CheckVersionsOfComponentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetResourceOwnerId(v int64) *CheckVersionsOfComponentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) SetZoneId(v string) *CheckVersionsOfComponentsRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckVersionsOfComponentsRequest) Validate() error {
	return dara.Validate(s)
}
