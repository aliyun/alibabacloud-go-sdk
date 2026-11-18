// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseUserAllowOrgCreationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowOrgCreation(v bool) *ModifyLangfuseUserAllowOrgCreationRequest
	GetAllowOrgCreation() *bool
	SetDBInstanceId(v string) *ModifyLangfuseUserAllowOrgCreationRequest
	GetDBInstanceId() *string
	SetEmail(v string) *ModifyLangfuseUserAllowOrgCreationRequest
	GetEmail() *string
	SetRegionId(v string) *ModifyLangfuseUserAllowOrgCreationRequest
	GetRegionId() *string
}

type ModifyLangfuseUserAllowOrgCreationRequest struct {
	// Specifies whether the user is allowed to create organizations.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AllowOrgCreation *bool `json:"AllowOrgCreation,omitempty" xml:"AllowOrgCreation,omitempty"`
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-2ze1n52*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyLangfuseUserAllowOrgCreationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseUserAllowOrgCreationRequest) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) GetAllowOrgCreation() *bool {
	return s.AllowOrgCreation
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) SetAllowOrgCreation(v bool) *ModifyLangfuseUserAllowOrgCreationRequest {
	s.AllowOrgCreation = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) SetDBInstanceId(v string) *ModifyLangfuseUserAllowOrgCreationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) SetEmail(v string) *ModifyLangfuseUserAllowOrgCreationRequest {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) SetRegionId(v string) *ModifyLangfuseUserAllowOrgCreationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationRequest) Validate() error {
	return dara.Validate(s)
}
