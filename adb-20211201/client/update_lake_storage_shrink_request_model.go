// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLakeStorageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpdateLakeStorageShrinkRequest
	GetDBClusterId() *string
	SetDescription(v string) *UpdateLakeStorageShrinkRequest
	GetDescription() *string
	SetLakeStorageId(v string) *UpdateLakeStorageShrinkRequest
	GetLakeStorageId() *string
	SetPermissionsShrink(v string) *UpdateLakeStorageShrinkRequest
	GetPermissionsShrink() *string
	SetRegionId(v string) *UpdateLakeStorageShrinkRequest
	GetRegionId() *string
}

type UpdateLakeStorageShrinkRequest struct {
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
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLakeStorageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateLakeStorageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLakeStorageShrinkRequest) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *UpdateLakeStorageShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *UpdateLakeStorageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLakeStorageShrinkRequest) SetDBClusterId(v string) *UpdateLakeStorageShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpdateLakeStorageShrinkRequest) SetDescription(v string) *UpdateLakeStorageShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLakeStorageShrinkRequest) SetLakeStorageId(v string) *UpdateLakeStorageShrinkRequest {
	s.LakeStorageId = &v
	return s
}

func (s *UpdateLakeStorageShrinkRequest) SetPermissionsShrink(v string) *UpdateLakeStorageShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *UpdateLakeStorageShrinkRequest) SetRegionId(v string) *UpdateLakeStorageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLakeStorageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
