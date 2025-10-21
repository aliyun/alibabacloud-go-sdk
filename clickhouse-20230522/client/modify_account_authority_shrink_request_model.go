// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyAccountAuthorityShrinkRequest
	GetAccount() *string
	SetDBInstanceId(v string) *ModifyAccountAuthorityShrinkRequest
	GetDBInstanceId() *string
	SetDmlAuthSettingShrink(v string) *ModifyAccountAuthorityShrinkRequest
	GetDmlAuthSettingShrink() *string
	SetRegionId(v string) *ModifyAccountAuthorityShrinkRequest
	GetRegionId() *string
}

type ModifyAccountAuthorityShrinkRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The information about permissions.
	//
	// This parameter is required.
	DmlAuthSettingShrink *string `json:"DmlAuthSetting,omitempty" xml:"DmlAuthSetting,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountAuthorityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityShrinkRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyAccountAuthorityShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountAuthorityShrinkRequest) GetDmlAuthSettingShrink() *string {
	return s.DmlAuthSettingShrink
}

func (s *ModifyAccountAuthorityShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountAuthorityShrinkRequest) SetAccount(v string) *ModifyAccountAuthorityShrinkRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetDBInstanceId(v string) *ModifyAccountAuthorityShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetDmlAuthSettingShrink(v string) *ModifyAccountAuthorityShrinkRequest {
	s.DmlAuthSettingShrink = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) SetRegionId(v string) *ModifyAccountAuthorityShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountAuthorityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
