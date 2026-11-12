// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterSetName(v string) *CreateDatabaseRequest
	GetCharacterSetName() *string
	SetCollate(v string) *CreateDatabaseRequest
	GetCollate() *string
	SetCtype(v string) *CreateDatabaseRequest
	GetCtype() *string
	SetDBInstanceId(v string) *CreateDatabaseRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *CreateDatabaseRequest
	GetDatabaseName() *string
	SetDescription(v string) *CreateDatabaseRequest
	GetDescription() *string
	SetOwner(v string) *CreateDatabaseRequest
	GetOwner() *string
}

type CreateDatabaseRequest struct {
	// The character set.
	//
	// For valid values, see References [https://www.postgresql.org/docs/current/multibyte.html](url).
	//
	// Default value: UTF8.
	//
	// example:
	//
	// UTF8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// The collation parameter of the database, which specifies the string comparison and sorting rules.
	//
	// >
	//
	// > - The locale must be compatible with the character set specified by the CharacterSetName parameter.
	//
	// > - Valid values for Collate: You can run the `SELECT DISTINCT collname FROM pg_collation;` command to query the valid values. If this parameter is not specified, the default value is en_US.utf8.
	//
	// example:
	//
	// en_US.utf8
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// The character type parameter of the database, which specifies the character classification and case conversion rules.
	//
	// >
	//
	// > - The locale must be compatible with the character set specified by the CharacterSetName parameter.
	//
	// > - You can run the `SELECT DISTINCT collctype FROM pg_collation;` command to query the valid values. If this parameter is not specified, the default value is en_US.utf8.
	//
	// example:
	//
	// en_US.utf8
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database. The name must meet the following requirements:
	//
	// - It can contain only letters, digits, and underscores (_).
	//
	// - It must start with a letter.
	//
	// - It can be up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The description of the database.
	//
	// - It must start with a Chinese character or a letter.
	//
	// - It cannot start with http:// or https://.
	//
	// - It can contain Chinese characters, letters, underscores (_), hyphens (-), and digits.
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// A user management database storing customer profiles.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The owner of the table.
	//
	// - It can contain only lowercase letters, digits, and underscores (_).
	//
	// - It must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// - It cannot start with gp.
	//
	// - It must be 2 to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testadb
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *CreateDatabaseRequest) GetCollate() *string {
	return s.Collate
}

func (s *CreateDatabaseRequest) GetCtype() *string {
	return s.Ctype
}

func (s *CreateDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateDatabaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatabaseRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDatabaseRequest) SetCharacterSetName(v string) *CreateDatabaseRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseRequest) SetCollate(v string) *CreateDatabaseRequest {
	s.Collate = &v
	return s
}

func (s *CreateDatabaseRequest) SetCtype(v string) *CreateDatabaseRequest {
	s.Ctype = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBInstanceId(v string) *CreateDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseName(v string) *CreateDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDescription(v string) *CreateDatabaseRequest {
	s.Description = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwner(v string) *CreateDatabaseRequest {
	s.Owner = &v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
