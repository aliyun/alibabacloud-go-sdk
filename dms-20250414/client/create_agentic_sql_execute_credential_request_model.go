// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSqlExecuteCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *CreateAgenticSqlExecuteCredentialRequest
	GetCatalogName() *string
	SetCatalogUuid(v string) *CreateAgenticSqlExecuteCredentialRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *CreateAgenticSqlExecuteCredentialRequest
	GetDatabaseUuid() *string
	SetQualifiedName(v string) *CreateAgenticSqlExecuteCredentialRequest
	GetQualifiedName() *string
}

type CreateAgenticSqlExecuteCredentialRequest struct {
	// The name of the instance.
	//
	// example:
	//
	// rm-bp1xxxxxxxx
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// 9f8e7d6c-xxxx-xxxx-xxxx-1a2b3c4d5e6f
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the database. If you specify this parameter, the credential is restricted to this database. You can call the ListDatabases operation to obtain the UUID.
	//
	// example:
	//
	// 6c4e8a2b-xxxx-xxxx-xxxx-0f1d2e3c4b5a
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The fully qualified name of the database. This parameter must be used together with CatalogUuid or CatalogName.
	//
	// - MySQL database: The value is the same as the database name.
	//
	// - PostgreSQL database or other databases with schemas: The value is in the database.schema format.
	//
	// example:
	//
	// test_db
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s CreateAgenticSqlExecuteCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSqlExecuteCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticSqlExecuteCredentialRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateAgenticSqlExecuteCredentialRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *CreateAgenticSqlExecuteCredentialRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *CreateAgenticSqlExecuteCredentialRequest) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *CreateAgenticSqlExecuteCredentialRequest) SetCatalogName(v string) *CreateAgenticSqlExecuteCredentialRequest {
	s.CatalogName = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialRequest) SetCatalogUuid(v string) *CreateAgenticSqlExecuteCredentialRequest {
	s.CatalogUuid = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialRequest) SetDatabaseUuid(v string) *CreateAgenticSqlExecuteCredentialRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialRequest) SetQualifiedName(v string) *CreateAgenticSqlExecuteCredentialRequest {
	s.QualifiedName = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialRequest) Validate() error {
	return dara.Validate(s)
}
