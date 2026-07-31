// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeStorageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateLakeStorageShrinkRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateLakeStorageShrinkRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateLakeStorageShrinkRequest
	GetDescription() *string
	SetPermissionsShrink(v string) *CreateLakeStorageShrinkRequest
	GetPermissionsShrink() *string
	SetRegionId(v string) *CreateLakeStorageShrinkRequest
	GetRegionId() *string
}

type CreateLakeStorageShrinkRequest struct {
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
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLakeStorageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeStorageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLakeStorageShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLakeStorageShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateLakeStorageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLakeStorageShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *CreateLakeStorageShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLakeStorageShrinkRequest) SetClientToken(v string) *CreateLakeStorageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLakeStorageShrinkRequest) SetDBClusterId(v string) *CreateLakeStorageShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateLakeStorageShrinkRequest) SetDescription(v string) *CreateLakeStorageShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateLakeStorageShrinkRequest) SetPermissionsShrink(v string) *CreateLakeStorageShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *CreateLakeStorageShrinkRequest) SetRegionId(v string) *CreateLakeStorageShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLakeStorageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
