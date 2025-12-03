// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodAddComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *MultimodAddComponentsRequest
	GetClusterId() *string
	SetComponents(v string) *MultimodAddComponentsRequest
	GetComponents() *string
	SetOwnerId(v int64) *MultimodAddComponentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MultimodAddComponentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *MultimodAddComponentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MultimodAddComponentsRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *MultimodAddComponentsRequest
	GetZoneId() *string
}

type MultimodAddComponentsRequest struct {
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

func (s MultimodAddComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s MultimodAddComponentsRequest) GoString() string {
	return s.String()
}

func (s *MultimodAddComponentsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *MultimodAddComponentsRequest) GetComponents() *string {
	return s.Components
}

func (s *MultimodAddComponentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MultimodAddComponentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MultimodAddComponentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MultimodAddComponentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MultimodAddComponentsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MultimodAddComponentsRequest) SetClusterId(v string) *MultimodAddComponentsRequest {
	s.ClusterId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetComponents(v string) *MultimodAddComponentsRequest {
	s.Components = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetOwnerId(v int64) *MultimodAddComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetRegionId(v string) *MultimodAddComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetResourceOwnerAccount(v string) *MultimodAddComponentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetResourceOwnerId(v int64) *MultimodAddComponentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MultimodAddComponentsRequest) SetZoneId(v string) *MultimodAddComponentsRequest {
	s.ZoneId = &v
	return s
}

func (s *MultimodAddComponentsRequest) Validate() error {
	return dara.Validate(s)
}
