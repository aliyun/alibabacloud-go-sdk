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
	SetProjectId(v int64) *ModifyDatabaseAccountRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyDatabaseAccountRequest
	GetRegionId() *string
}

type ModifyDatabaseAccountRequest struct {
	// The ID of the database account to be modified.
	//
	// > You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The new name of the database account. The name can be up to 128 characters in length.
	//
	// example:
	//
	// aaa
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The new database schema name. This parameter is required when the database specified by DatabaseId is a PostgreSQL or Oracle database.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// The instance ID of the Bastionhost instance where the database account to be modified resides.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
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
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance where the database account to be modified resides.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ModifyDatabaseAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ModifyDatabaseAccountRequest) SetProjectId(v int64) *ModifyDatabaseAccountRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) SetRegionId(v string) *ModifyDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
