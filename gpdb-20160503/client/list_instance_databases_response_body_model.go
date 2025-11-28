// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListInstanceDatabasesResponseBody
	GetDBInstanceId() *string
	SetDatabases(v []*ListInstanceDatabasesResponseBodyDatabases) *ListInstanceDatabasesResponseBody
	GetDatabases() []*ListInstanceDatabasesResponseBodyDatabases
	SetRequestId(v string) *ListInstanceDatabasesResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListInstanceDatabasesResponseBody
	GetTotalCount() *string
}

type ListInstanceDatabasesResponseBody struct {
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Databases    []*ListInstanceDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceDatabasesResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListInstanceDatabasesResponseBody) GetDatabases() []*ListInstanceDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *ListInstanceDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceDatabasesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListInstanceDatabasesResponseBody) SetDBInstanceId(v string) *ListInstanceDatabasesResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ListInstanceDatabasesResponseBody) SetDatabases(v []*ListInstanceDatabasesResponseBodyDatabases) *ListInstanceDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListInstanceDatabasesResponseBody) SetRequestId(v string) *ListInstanceDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceDatabasesResponseBody) SetTotalCount(v string) *ListInstanceDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceDatabasesResponseBody) Validate() error {
	if s.Databases != nil {
		for _, item := range s.Databases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceDatabasesResponseBodyDatabases struct {
	// example:
	//
	// testdatabase
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListInstanceDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListInstanceDatabasesResponseBodyDatabases) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListInstanceDatabasesResponseBodyDatabases) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceDatabasesResponseBodyDatabases) SetDatabaseName(v string) *ListInstanceDatabasesResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *ListInstanceDatabasesResponseBodyDatabases) SetDescription(v string) *ListInstanceDatabasesResponseBodyDatabases {
	s.Description = &v
	return s
}

func (s *ListInstanceDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
