// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccountName(v string) *CreateDatabaseAccountRequest
	GetDatabaseAccountName() *string
	SetDatabaseId(v string) *CreateDatabaseAccountRequest
	GetDatabaseId() *string
	SetDatabaseSchema(v string) *CreateDatabaseAccountRequest
	GetDatabaseSchema() *string
	SetInstanceId(v string) *CreateDatabaseAccountRequest
	GetInstanceId() *string
	SetLoginAttribute(v string) *CreateDatabaseAccountRequest
	GetLoginAttribute() *string
	SetPassword(v string) *CreateDatabaseAccountRequest
	GetPassword() *string
	SetRegionId(v string) *CreateDatabaseAccountRequest
	GetRegionId() *string
}

type CreateDatabaseAccountRequest struct {
	// The username of the database account to be created. The username can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The ID of the database for which you want to create a database account.
	//
	// >  You can call the [ListDatabaseAccounts](https://help.aliyun.com/document_detail/2758839.html) operation to query the database account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database. This parameter is required for PostgreSQL and Oracle databases.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// The ID of the bastion host for which you want to create a database account.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-5yd34ol020a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The logon attribute. This parameter is required for Oracle databases. Valid values:
	//
	// 	- SERVICENAME
	//
	// 	- SID
	//
	// example:
	//
	// SID
	LoginAttribute *string `json:"LoginAttribute,omitempty" xml:"LoginAttribute,omitempty"`
	// The password of the database account to be created.
	//
	// example:
	//
	// MCQ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host for which you want to create a database account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDatabaseAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseAccountRequest) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *CreateDatabaseAccountRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *CreateDatabaseAccountRequest) GetDatabaseSchema() *string {
	return s.DatabaseSchema
}

func (s *CreateDatabaseAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDatabaseAccountRequest) GetLoginAttribute() *string {
	return s.LoginAttribute
}

func (s *CreateDatabaseAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDatabaseAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDatabaseAccountRequest) SetDatabaseAccountName(v string) *CreateDatabaseAccountRequest {
	s.DatabaseAccountName = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetDatabaseId(v string) *CreateDatabaseAccountRequest {
	s.DatabaseId = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetDatabaseSchema(v string) *CreateDatabaseAccountRequest {
	s.DatabaseSchema = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetInstanceId(v string) *CreateDatabaseAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetLoginAttribute(v string) *CreateDatabaseAccountRequest {
	s.LoginAttribute = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetPassword(v string) *CreateDatabaseAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateDatabaseAccountRequest) SetRegionId(v string) *CreateDatabaseAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDatabaseAccountRequest) Validate() error {
	return dara.Validate(s)
}
