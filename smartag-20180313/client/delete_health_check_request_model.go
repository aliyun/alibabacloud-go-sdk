// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHcInstanceId(v string) *DeleteHealthCheckRequest
	GetHcInstanceId() *string
	SetOwnerAccount(v string) *DeleteHealthCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteHealthCheckRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteHealthCheckRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteHealthCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteHealthCheckRequest
	GetResourceOwnerId() *int64
}

type DeleteHealthCheckRequest struct {
	// The ID of the health check instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc-vfgyz7dv07lthr****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the health check instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckRequest) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DeleteHealthCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteHealthCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteHealthCheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHealthCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteHealthCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteHealthCheckRequest) SetHcInstanceId(v string) *DeleteHealthCheckRequest {
	s.HcInstanceId = &v
	return s
}

func (s *DeleteHealthCheckRequest) SetOwnerAccount(v string) *DeleteHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteHealthCheckRequest) SetOwnerId(v int64) *DeleteHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteHealthCheckRequest) SetRegionId(v string) *DeleteHealthCheckRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHealthCheckRequest) SetResourceOwnerAccount(v string) *DeleteHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteHealthCheckRequest) SetResourceOwnerId(v int64) *DeleteHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
