// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyAppInstanceShrinkRequest
	GetBranchName() *string
	SetClientToken(v string) *ModifyAppInstanceShrinkRequest
	GetClientToken() *string
	SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest
	GetComponentsShrink() *string
	SetDBInstanceName(v string) *ModifyAppInstanceShrinkRequest
	GetDBInstanceName() *string
	SetInstanceClass(v string) *ModifyAppInstanceShrinkRequest
	GetInstanceClass() *string
	SetInstanceName(v string) *ModifyAppInstanceShrinkRequest
	GetInstanceName() *string
	SetMigrationMode(v string) *ModifyAppInstanceShrinkRequest
	GetMigrationMode() *string
	SetRegionId(v string) *ModifyAppInstanceShrinkRequest
	GetRegionId() *string
}

type ModifyAppInstanceShrinkRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The idempotency parameter.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of modules.
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The instance ID of the ApsaraDB RDS for PostgreSQL database connected to the AI application.
	//
	// <notice>Only **newly purchased empty ApsaraDB RDS for PostgreSQL instances*	- are supported. The major engine version must be **17**, and the minor engine version must be **20250630 or later**.</notice>
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance type. Only **rdsai.supabase.basic*	- is supported.
	//
	// example:
	//
	// rdsai.supabase.basic
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	MigrationMode *string `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceShrinkRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppInstanceShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *ModifyAppInstanceShrinkRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyAppInstanceShrinkRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *ModifyAppInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyAppInstanceShrinkRequest) GetMigrationMode() *string {
	return s.MigrationMode
}

func (s *ModifyAppInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppInstanceShrinkRequest) SetBranchName(v string) *ModifyAppInstanceShrinkRequest {
	s.BranchName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetClientToken(v string) *ModifyAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetDBInstanceName(v string) *ModifyAppInstanceShrinkRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetInstanceClass(v string) *ModifyAppInstanceShrinkRequest {
	s.InstanceClass = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetInstanceName(v string) *ModifyAppInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetMigrationMode(v string) *ModifyAppInstanceShrinkRequest {
	s.MigrationMode = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetRegionId(v string) *ModifyAppInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
