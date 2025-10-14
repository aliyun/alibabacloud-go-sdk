// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGdnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateGdnInstanceRequest
	GetDBInstanceName() *string
	SetDescription(v string) *CreateGdnInstanceRequest
	GetDescription() *string
	SetGdnMode(v string) *CreateGdnInstanceRequest
	GetGdnMode() *string
	SetRegionId(v string) *CreateGdnInstanceRequest
	GetRegionId() *string
	SetRplConflictStrategy(v string) *CreateGdnInstanceRequest
	GetRplConflictStrategy() *string
	SetRplDmlStrategy(v string) *CreateGdnInstanceRequest
	GetRplDmlStrategy() *string
	SetRplSyncDdl(v bool) *CreateGdnInstanceRequest
	GetRplSyncDdl() *bool
}

type CreateGdnInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// FASTJSON 2.0.x has been released, faster and more secure, recommend you upgrade.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GdnMode     *string `json:"GdnMode,omitempty" xml:"GdnMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RplConflictStrategy *string `json:"RplConflictStrategy,omitempty" xml:"RplConflictStrategy,omitempty"`
	RplDmlStrategy      *string `json:"RplDmlStrategy,omitempty" xml:"RplDmlStrategy,omitempty"`
	RplSyncDdl          *bool   `json:"RplSyncDdl,omitempty" xml:"RplSyncDdl,omitempty"`
}

func (s CreateGdnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGdnInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGdnInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateGdnInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGdnInstanceRequest) GetGdnMode() *string {
	return s.GdnMode
}

func (s *CreateGdnInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGdnInstanceRequest) GetRplConflictStrategy() *string {
	return s.RplConflictStrategy
}

func (s *CreateGdnInstanceRequest) GetRplDmlStrategy() *string {
	return s.RplDmlStrategy
}

func (s *CreateGdnInstanceRequest) GetRplSyncDdl() *bool {
	return s.RplSyncDdl
}

func (s *CreateGdnInstanceRequest) SetDBInstanceName(v string) *CreateGdnInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetDescription(v string) *CreateGdnInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetGdnMode(v string) *CreateGdnInstanceRequest {
	s.GdnMode = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetRegionId(v string) *CreateGdnInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetRplConflictStrategy(v string) *CreateGdnInstanceRequest {
	s.RplConflictStrategy = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetRplDmlStrategy(v string) *CreateGdnInstanceRequest {
	s.RplDmlStrategy = &v
	return s
}

func (s *CreateGdnInstanceRequest) SetRplSyncDdl(v bool) *CreateGdnInstanceRequest {
	s.RplSyncDdl = &v
	return s
}

func (s *CreateGdnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
