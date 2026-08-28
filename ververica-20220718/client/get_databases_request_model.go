// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *GetDatabasesRequest
	GetDatabaseName() *string
}

type GetDatabasesRequest struct {
	// The database name. If this parameter is left empty, information about all databases is returned.
	//
	// example:
	//
	// paimon-ods
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
}

func (s GetDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatabasesRequest) GoString() string {
	return s.String()
}

func (s *GetDatabasesRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetDatabasesRequest) SetDatabaseName(v string) *GetDatabasesRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
