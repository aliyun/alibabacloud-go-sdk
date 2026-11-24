// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateLakeStorageRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateLakeStorageRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateLakeStorageRequest
	GetDescription() *string
	SetPermissions(v []*CreateLakeStorageRequestPermissions) *CreateLakeStorageRequest
	GetPermissions() []*CreateLakeStorageRequestPermissions
	SetRegionId(v string) *CreateLakeStorageRequest
	GetRegionId() *string
}

type CreateLakeStorageRequest struct {
	// The client token that is used to ensure request idempotence. If you use the same token for requests with different parameters, the IdempotentParameterMismatch error is returned.
	//
	// example:
	//
	// b2cd****3efg
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the ADB instance to which the lake storage is attached.
	//
	// example:
	//
	// amv-******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the lake storage.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// When you create a lake storage, permissions are automatically granted to the Resource Access Management (RAM) users and the primary account used for the operation. You can increase additional authorizations for the primary account here.
	//
	// example:
	//
	// [{"Account":"1234567890123456","Read":true,"Type":"CUSTOMER","Write":true}]
	Permissions []*CreateLakeStorageRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLakeStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeStorageRequest) GoString() string {
	return s.String()
}

func (s *CreateLakeStorageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLakeStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateLakeStorageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLakeStorageRequest) GetPermissions() []*CreateLakeStorageRequestPermissions {
	return s.Permissions
}

func (s *CreateLakeStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLakeStorageRequest) SetClientToken(v string) *CreateLakeStorageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLakeStorageRequest) SetDBClusterId(v string) *CreateLakeStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateLakeStorageRequest) SetDescription(v string) *CreateLakeStorageRequest {
	s.Description = &v
	return s
}

func (s *CreateLakeStorageRequest) SetPermissions(v []*CreateLakeStorageRequestPermissions) *CreateLakeStorageRequest {
	s.Permissions = v
	return s
}

func (s *CreateLakeStorageRequest) SetRegionId(v string) *CreateLakeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLakeStorageRequest) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLakeStorageRequestPermissions struct {
	// The account ID.
	//
	// example:
	//
	// 1234567890123456
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The read permission.
	//
	// example:
	//
	// true
	Read *bool `json:"Read,omitempty" xml:"Read,omitempty"`
	// The account type.
	//
	// example:
	//
	// CUSTOMER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permission.
	//
	// example:
	//
	// true
	Write *bool `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s CreateLakeStorageRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeStorageRequestPermissions) GoString() string {
	return s.String()
}

func (s *CreateLakeStorageRequestPermissions) GetAccount() *string {
	return s.Account
}

func (s *CreateLakeStorageRequestPermissions) GetRead() *bool {
	return s.Read
}

func (s *CreateLakeStorageRequestPermissions) GetType() *string {
	return s.Type
}

func (s *CreateLakeStorageRequestPermissions) GetWrite() *bool {
	return s.Write
}

func (s *CreateLakeStorageRequestPermissions) SetAccount(v string) *CreateLakeStorageRequestPermissions {
	s.Account = &v
	return s
}

func (s *CreateLakeStorageRequestPermissions) SetRead(v bool) *CreateLakeStorageRequestPermissions {
	s.Read = &v
	return s
}

func (s *CreateLakeStorageRequestPermissions) SetType(v string) *CreateLakeStorageRequestPermissions {
	s.Type = &v
	return s
}

func (s *CreateLakeStorageRequestPermissions) SetWrite(v bool) *CreateLakeStorageRequestPermissions {
	s.Write = &v
	return s
}

func (s *CreateLakeStorageRequestPermissions) Validate() error {
	return dara.Validate(s)
}
