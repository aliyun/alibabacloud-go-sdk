// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ModifyCenAttributeRequest
	GetCenId() *string
	SetDescription(v string) *ModifyCenAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyCenAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCenAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCenAttributeRequest
	GetOwnerId() *int64
	SetProtectionLevel(v string) *ModifyCenAttributeRequest
	GetProtectionLevel() *string
	SetResourceOwnerAccount(v string) *ModifyCenAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCenAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyCenAttributeRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the CEN instance.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cen
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the CEN instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The level of CIDR block overlapping.
	//
	// Set the value to **REDUCED*	- (default). This value specifies that CIDR blocks can overlap but cannot be the same.
	//
	// example:
	//
	// REDUCED
	ProtectionLevel      *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenAttributeRequest) GetCenId() *string {
	return s.CenId
}

func (s *ModifyCenAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCenAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCenAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCenAttributeRequest) GetProtectionLevel() *string {
	return s.ProtectionLevel
}

func (s *ModifyCenAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCenAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCenAttributeRequest) SetCenId(v string) *ModifyCenAttributeRequest {
	s.CenId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetDescription(v string) *ModifyCenAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetName(v string) *ModifyCenAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetOwnerAccount(v string) *ModifyCenAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetOwnerId(v int64) *ModifyCenAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetProtectionLevel(v string) *ModifyCenAttributeRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetResourceOwnerAccount(v string) *ModifyCenAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenAttributeRequest) SetResourceOwnerId(v int64) *ModifyCenAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenAttributeRequest) Validate() error {
	return dara.Validate(s)
}
