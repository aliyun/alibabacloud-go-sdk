// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseUuid(v string) *GetAgenticDatabaseRequest
	GetDatabaseUuid() *string
}

type GetAgenticDatabaseRequest struct {
	// The UUID of the database. You can obtain this value by calling ListAgenticDatabase or SearchAgenticDatabase.
	//
	// This parameter is required.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
}

func (s GetAgenticDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticDatabaseRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *GetAgenticDatabaseRequest) SetDatabaseUuid(v string) *GetAgenticDatabaseRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *GetAgenticDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
