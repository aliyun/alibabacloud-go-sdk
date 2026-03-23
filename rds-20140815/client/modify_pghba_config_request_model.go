// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPGHbaConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyPGHbaConfigRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyPGHbaConfigRequest
	GetDBInstanceId() *string
	SetHbaItem(v []*ModifyPGHbaConfigRequestHbaItem) *ModifyPGHbaConfigRequest
	GetHbaItem() []*ModifyPGHbaConfigRequestHbaItem
	SetOpsType(v string) *ModifyPGHbaConfigRequest
	GetOpsType() *string
	SetOwnerAccount(v string) *ModifyPGHbaConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPGHbaConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyPGHbaConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPGHbaConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyPGHbaConfigRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	HbaItem []*ModifyPGHbaConfigRequestHbaItem `json:"HbaItem,omitempty" xml:"HbaItem,omitempty" type:"Repeated"`
	// This parameter is required.
	OpsType              *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPGHbaConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPGHbaConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyPGHbaConfigRequest) GetHbaItem() []*ModifyPGHbaConfigRequestHbaItem {
	return s.HbaItem
}

func (s *ModifyPGHbaConfigRequest) GetOpsType() *string {
	return s.OpsType
}

func (s *ModifyPGHbaConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPGHbaConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPGHbaConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPGHbaConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPGHbaConfigRequest) SetClientToken(v string) *ModifyPGHbaConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetDBInstanceId(v string) *ModifyPGHbaConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetHbaItem(v []*ModifyPGHbaConfigRequestHbaItem) *ModifyPGHbaConfigRequest {
	s.HbaItem = v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOpsType(v string) *ModifyPGHbaConfigRequest {
	s.OpsType = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOwnerAccount(v string) *ModifyPGHbaConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetOwnerId(v int64) *ModifyPGHbaConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetResourceOwnerAccount(v string) *ModifyPGHbaConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) SetResourceOwnerId(v int64) *ModifyPGHbaConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPGHbaConfigRequest) Validate() error {
	if s.HbaItem != nil {
		for _, item := range s.HbaItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPGHbaConfigRequestHbaItem struct {
	// This parameter is required.
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// This parameter is required.
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Mask     *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// This parameter is required.
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// This parameter is required.
	PriorityId *int32 `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ModifyPGHbaConfigRequestHbaItem) String() string {
	return dara.Prettify(s)
}

func (s ModifyPGHbaConfigRequestHbaItem) GoString() string {
	return s.String()
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetAddress() *string {
	return s.Address
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetDatabase() *string {
	return s.Database
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetMask() *string {
	return s.Mask
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetMethod() *string {
	return s.Method
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetOption() *string {
	return s.Option
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetPriorityId() *int32 {
	return s.PriorityId
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetType() *string {
	return s.Type
}

func (s *ModifyPGHbaConfigRequestHbaItem) GetUser() *string {
	return s.User
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetAddress(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Address = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetDatabase(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Database = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetMask(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Mask = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetMethod(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Method = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetOption(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Option = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetPriorityId(v int32) *ModifyPGHbaConfigRequestHbaItem {
	s.PriorityId = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetType(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.Type = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) SetUser(v string) *ModifyPGHbaConfigRequestHbaItem {
	s.User = &v
	return s
}

func (s *ModifyPGHbaConfigRequestHbaItem) Validate() error {
	return dara.Validate(s)
}
