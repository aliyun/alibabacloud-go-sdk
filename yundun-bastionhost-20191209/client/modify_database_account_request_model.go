// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountId(v string) *ModifyDatabaseAccountRequest
	GetDatabaseAccountId() *string
	SetDatabaseAccountName(v string) *ModifyDatabaseAccountRequest
	GetDatabaseAccountName() *string
	SetDatabaseSchema(v string) *ModifyDatabaseAccountRequest
	GetDatabaseSchema() *string
	SetInstanceId(v string) *ModifyDatabaseAccountRequest
	GetInstanceId() *string
	SetPassword(v string) *ModifyDatabaseAccountRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyDatabaseAccountRequest
	GetRegionId() *string
}

type ModifyDatabaseAccountRequest struct {
	// The ID of the database account to modify.
	//
	// >  You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to query the database account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The new username of the database account. The username can be up to 128 characters in length.
	//
	// example:
	//
	// aaa
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The new name of the database. This parameter is required if the database engine is PostgreSQL or Oracle.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// The ID of the bastion host that manages the database account to modify.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zpr2zyqx603
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password of the database account.
	//
	// example:
	//
	// 14SZ!******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host that manages the database account to modify.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseAccountRequest) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *ModifyDatabaseAccountRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ModifyDatabaseAccountRequest) GetDatabaseSchema() *string {
	return s.DatabaseSchema
}

func (s *ModifyDatabaseAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDatabaseAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyDatabaseAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabaseAccountRequest) SetDatabaseAccountId(v string) *ModifyDatabaseAccountRequest {
	s.DatabaseAccountId = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetDatabaseAccountName(v string) *ModifyDatabaseAccountRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetDatabaseSchema(v string) *ModifyDatabaseAccountRequest {
	s.DatabaseSchema = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetInstanceId(v string) *ModifyDatabaseAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetPassword(v string) *ModifyDatabaseAccountRequest {
	s.Password = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetRegionId(v string) *ModifyDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
