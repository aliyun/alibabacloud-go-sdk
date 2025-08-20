// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegesShrinkRequest
	GetAccountName() *string
	SetAccountPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest
	GetAccountPrivilegesShrink() *string
	SetDBClusterId(v string) *ModifyAccountPrivilegesShrinkRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ModifyAccountPrivilegesShrinkRequest
	GetRegionId() *string
}

type ModifyAccountPrivilegesShrinkRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions that you want to grant to the database account.
	//
	// This parameter is required.
	AccountPrivilegesShrink *string `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountPrivilegesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetAccountPrivilegesShrink() *string {
	return s.AccountPrivilegesShrink
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountName(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountPrivilegesShrink = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetRegionId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
