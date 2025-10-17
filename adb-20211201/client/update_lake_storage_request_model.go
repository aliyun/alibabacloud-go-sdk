// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLakeStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateLakeStorageRequest
	GetDBClusterId() *string
	SetDescription(v string) *UpdateLakeStorageRequest
	GetDescription() *string
	SetLakeStorageId(v string) *UpdateLakeStorageRequest
	GetLakeStorageId() *string
	SetPermissions(v []*UpdateLakeStorageRequestPermissions) *UpdateLakeStorageRequest
	GetPermissions() []*UpdateLakeStorageRequestPermissions
	SetRegionId(v string) *UpdateLakeStorageRequest
	GetRegionId() *string
}

type UpdateLakeStorageRequest struct {
	// The ID of the AnalyticDB for MySQL cluster that is associated with the lake storage.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the lake storage.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the lake storage.
	//
	// example:
	//
	// -
	LakeStorageId *string `json:"LakeStorageId,omitempty" xml:"LakeStorageId,omitempty"`
	// The permissions on the lake storage.
	//
	// example:
	//
	// -
	Permissions []*UpdateLakeStorageRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLakeStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageRequest) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateLakeStorageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLakeStorageRequest) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *UpdateLakeStorageRequest) GetPermissions() []*UpdateLakeStorageRequestPermissions {
	return s.Permissions
}

func (s *UpdateLakeStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLakeStorageRequest) SetDBClusterId(v string) *UpdateLakeStorageRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateLakeStorageRequest) SetDescription(v string) *UpdateLakeStorageRequest {
	s.Description = &v
	return s
}

func (s *UpdateLakeStorageRequest) SetLakeStorageId(v string) *UpdateLakeStorageRequest {
	s.LakeStorageId = &v
	return s
}

func (s *UpdateLakeStorageRequest) SetPermissions(v []*UpdateLakeStorageRequestPermissions) *UpdateLakeStorageRequest {
	s.Permissions = v
	return s
}

func (s *UpdateLakeStorageRequest) SetRegionId(v string) *UpdateLakeStorageRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLakeStorageRequest) Validate() error {
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

type UpdateLakeStorageRequestPermissions struct {
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The read permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Read *bool `json:"Read,omitempty" xml:"Read,omitempty"`
	// The account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Write *bool `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s UpdateLakeStorageRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageRequestPermissions) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageRequestPermissions) GetAccount() *string {
	return s.Account
}

func (s *UpdateLakeStorageRequestPermissions) GetRead() *bool {
	return s.Read
}

func (s *UpdateLakeStorageRequestPermissions) GetType() *string {
	return s.Type
}

func (s *UpdateLakeStorageRequestPermissions) GetWrite() *bool {
	return s.Write
}

func (s *UpdateLakeStorageRequestPermissions) SetAccount(v string) *UpdateLakeStorageRequestPermissions {
	s.Account = &v
	return s
}

func (s *UpdateLakeStorageRequestPermissions) SetRead(v bool) *UpdateLakeStorageRequestPermissions {
	s.Read = &v
	return s
}

func (s *UpdateLakeStorageRequestPermissions) SetType(v string) *UpdateLakeStorageRequestPermissions {
	s.Type = &v
	return s
}

func (s *UpdateLakeStorageRequestPermissions) SetWrite(v bool) *UpdateLakeStorageRequestPermissions {
	s.Write = &v
	return s
}

func (s *UpdateLakeStorageRequestPermissions) Validate() error {
	return dara.Validate(s)
}
