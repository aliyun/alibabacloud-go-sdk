// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseUuid(v string) *GetAgenticTableIndexRequest
	GetDatabaseUuid() *string
	SetTableQualifiedName(v string) *GetAgenticTableIndexRequest
	GetTableQualifiedName() *string
}

type GetAgenticTableIndexRequest struct {
	// The UUID of the database to which the table belongs. You can obtain this value by calling ListAgenticDatabase or SearchAgenticDatabase.
	//
	// This parameter is required.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The fully qualified name of the table. You can obtain this value from the QualifiedName parameter returned by ListAgenticTableBaseInfo. You must specify the full name.
	//
	// This parameter is required.
	//
	// example:
	//
	// orders
	TableQualifiedName *string `json:"TableQualifiedName,omitempty" xml:"TableQualifiedName,omitempty"`
}

func (s GetAgenticTableIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableIndexRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticTableIndexRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *GetAgenticTableIndexRequest) GetTableQualifiedName() *string {
	return s.TableQualifiedName
}

func (s *GetAgenticTableIndexRequest) SetDatabaseUuid(v string) *GetAgenticTableIndexRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *GetAgenticTableIndexRequest) SetTableQualifiedName(v string) *GetAgenticTableIndexRequest {
	s.TableQualifiedName = &v
	return s
}

func (s *GetAgenticTableIndexRequest) Validate() error {
	return dara.Validate(s)
}
