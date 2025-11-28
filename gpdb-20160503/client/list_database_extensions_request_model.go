// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListDatabaseExtensionsRequest
	GetDBInstanceId() *string
	SetDatabaseName(v string) *ListDatabaseExtensionsRequest
	GetDatabaseName() *string
}

type ListDatabaseExtensionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
}

func (s ListDatabaseExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseExtensionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatabaseExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListDatabaseExtensionsRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabaseExtensionsRequest) SetDBInstanceId(v string) *ListDatabaseExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListDatabaseExtensionsRequest) SetDatabaseName(v string) *ListDatabaseExtensionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabaseExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
