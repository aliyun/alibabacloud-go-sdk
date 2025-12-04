// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMinorVersionGreadeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyMinorVersionGreadeTypeRequest
	GetDBClusterId() *string
	SetMaintainAutoType(v bool) *ModifyMinorVersionGreadeTypeRequest
	GetMaintainAutoType() *bool
	SetOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyMinorVersionGreadeTypeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The update type. If you set the parameter to **false**, you perform the manual update.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	MaintainAutoType     *bool   `json:"MaintainAutoType,omitempty" xml:"MaintainAutoType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyMinorVersionGreadeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMinorVersionGreadeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetMaintainAutoType() *bool {
	return s.MaintainAutoType
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMinorVersionGreadeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetDBClusterId(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetMaintainAutoType(v bool) *ModifyMinorVersionGreadeTypeRequest {
	s.MaintainAutoType = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetResourceOwnerAccount(v string) *ModifyMinorVersionGreadeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) SetResourceOwnerId(v int64) *ModifyMinorVersionGreadeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMinorVersionGreadeTypeRequest) Validate() error {
	return dara.Validate(s)
}
