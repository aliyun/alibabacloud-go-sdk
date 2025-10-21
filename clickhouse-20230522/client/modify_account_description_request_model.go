// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *ModifyAccountDescriptionRequest
	GetAccount() *string
	SetDBInstanceId(v string) *ModifyAccountDescriptionRequest
	GetDBInstanceId() *string
	SetDescription(v string) *ModifyAccountDescriptionRequest
	GetDescription() *string
	SetRegionId(v string) *ModifyAccountDescriptionRequest
	GetRegionId() *string
}

type ModifyAccountDescriptionRequest struct {
	// The name of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testuser
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccount() *string {
	return s.Account
}

func (s *ModifyAccountDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAccountDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountDescriptionRequest) SetAccount(v string) *ModifyAccountDescriptionRequest {
	s.Account = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDescription(v string) *ModifyAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetRegionId(v string) *ModifyAccountDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
