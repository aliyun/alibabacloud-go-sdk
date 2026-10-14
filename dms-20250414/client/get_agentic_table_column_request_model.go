// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticTableColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseUuid(v string) *GetAgenticTableColumnRequest
	GetDatabaseUuid() *string
	SetTableQualifiedName(v string) *GetAgenticTableColumnRequest
	GetTableQualifiedName() *string
}

type GetAgenticTableColumnRequest struct {
	// The UUID of the database to which the table belongs. You can obtain the UUID by calling the ListAgenticDatabase or SearchAgenticDatabase operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The fully qualified name of the table. You can obtain the name from the QualifiedName parameter returned by the ListAgenticTableBaseInfo operation. You must specify the complete name.
	//
	// This parameter is required.
	//
	// example:
	//
	// orders
	TableQualifiedName *string `json:"TableQualifiedName,omitempty" xml:"TableQualifiedName,omitempty"`
}

func (s GetAgenticTableColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticTableColumnRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticTableColumnRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *GetAgenticTableColumnRequest) GetTableQualifiedName() *string {
	return s.TableQualifiedName
}

func (s *GetAgenticTableColumnRequest) SetDatabaseUuid(v string) *GetAgenticTableColumnRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *GetAgenticTableColumnRequest) SetTableQualifiedName(v string) *GetAgenticTableColumnRequest {
	s.TableQualifiedName = &v
	return s
}

func (s *GetAgenticTableColumnRequest) Validate() error {
	return dara.Validate(s)
}
