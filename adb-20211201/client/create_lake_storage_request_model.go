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
	// -
	//
	// example:
	//
	// ******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the ADB instance attached to the lake storage.
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
	// When lake storage is created, permissions are automatically granted to the Resource Access Management (RAM) users performing the operation and the Alibaba Cloud account. You can increase additional Alibaba Cloud account authorizations here.
	//
	// example:
	//
	// -
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
	// -
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The read permission.
	//
	// example:
	//
	// -
	Read *bool `json:"Read,omitempty" xml:"Read,omitempty"`
	// The account type.
	//
	// example:
	//
	// -
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permission.
	//
	// example:
	//
	// -
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
