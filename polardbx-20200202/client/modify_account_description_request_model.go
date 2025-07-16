// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetDBInstanceName(v string) *ModifyAccountDescriptionRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ModifyAccountDescriptionRequest
	GetRegionId() *string
}

type ModifyAccountDescriptionRequest struct {
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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

func (s *ModifyAccountDescriptionRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyAccountDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceName(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetRegionId(v string) *ModifyAccountDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
