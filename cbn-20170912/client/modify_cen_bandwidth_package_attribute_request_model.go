// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageAttributeRequest
	GetCenBandwidthPackageId() *string
	SetDescription(v string) *ModifyCenBandwidthPackageAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyCenBandwidthPackageAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyCenBandwidthPackageAttributeRequest struct {
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5fx****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The new description of the bandwidth plan.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// Bandwidth Plans
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the bandwidth plan.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// test
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenBandwidthPackageAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCenBandwidthPackageAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetDescription(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetName(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageAttributeRequest) Validate() error {
	return dara.Validate(s)
}
