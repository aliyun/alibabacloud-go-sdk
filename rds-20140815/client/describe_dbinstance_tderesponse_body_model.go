// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDEResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *DescribeDBInstanceTDEResponseBodyDatabases) *DescribeDBInstanceTDEResponseBody
	GetDatabases() *DescribeDBInstanceTDEResponseBodyDatabases
	SetEncryptionKey(v string) *DescribeDBInstanceTDEResponseBody
	GetEncryptionKey() *string
	SetRequestId(v string) *DescribeDBInstanceTDEResponseBody
	GetRequestId() *string
	SetTDEMode(v string) *DescribeDBInstanceTDEResponseBody
	GetTDEMode() *string
	SetTDEStatus(v string) *DescribeDBInstanceTDEResponseBody
	GetTDEStatus() *string
}

type DescribeDBInstanceTDEResponseBody struct {
	// The TDE status at the database level.
	//
	// >  If your instance runs SQL Server 2019 SE or SQL Server EE, you can specify whether to enable TDE at the database level when you enable TDE at the instance level.
	Databases *DescribeDBInstanceTDEResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The ID of the custom key.
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C816A4BF-A6EC-4722-95F9-2055859CCFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The method that is used to generate the key for TDE at the instance level. Valid values:
	//
	// 	- **Aliyun_Generate_Key**
	//
	// 	- **Customer_Provided_Key**
	//
	// 	- **Unknown**
	//
	// example:
	//
	// Aliyun_Generate_Key
	TDEMode *string `json:"TDEMode,omitempty" xml:"TDEMode,omitempty"`
	// The TDE status of the instance. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBody) GetDatabases() *DescribeDBInstanceTDEResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeDBInstanceTDEResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBInstanceTDEResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceTDEResponseBody) GetTDEMode() *string {
	return s.TDEMode
}

func (s *DescribeDBInstanceTDEResponseBody) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeDBInstanceTDEResponseBody) SetDatabases(v *DescribeDBInstanceTDEResponseBodyDatabases) *DescribeDBInstanceTDEResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetEncryptionKey(v string) *DescribeDBInstanceTDEResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetRequestId(v string) *DescribeDBInstanceTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetTDEMode(v string) *DescribeDBInstanceTDEResponseBody {
	s.TDEMode = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) SetTDEStatus(v string) *DescribeDBInstanceTDEResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBody) Validate() error {
	if s.Databases != nil {
		if err := s.Databases.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceTDEResponseBodyDatabases struct {
	Database []*DescribeDBInstanceTDEResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceTDEResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBodyDatabases) GetDatabase() []*DescribeDBInstanceTDEResponseBodyDatabasesDatabase {
	return s.Database
}

func (s *DescribeDBInstanceTDEResponseBodyDatabases) SetDatabase(v []*DescribeDBInstanceTDEResponseBodyDatabasesDatabase) *DescribeDBInstanceTDEResponseBodyDatabases {
	s.Database = v
	return s
}

func (s *DescribeDBInstanceTDEResponseBodyDatabases) Validate() error {
	if s.Database != nil {
		for _, item := range s.Database {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceTDEResponseBodyDatabasesDatabase struct {
	// The name of the database.
	//
	// example:
	//
	// test02
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The TDE status at the database level. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s DescribeDBInstanceTDEResponseBodyDatabasesDatabase) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDEResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDEResponseBodyDatabasesDatabase) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDBInstanceTDEResponseBodyDatabasesDatabase) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *DescribeDBInstanceTDEResponseBodyDatabasesDatabase) SetDBName(v string) *DescribeDBInstanceTDEResponseBodyDatabasesDatabase {
	s.DBName = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBodyDatabasesDatabase) SetTDEStatus(v string) *DescribeDBInstanceTDEResponseBodyDatabasesDatabase {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBInstanceTDEResponseBodyDatabasesDatabase) Validate() error {
	return dara.Validate(s)
}
