// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAccountAuthorityResponseBodyData) *DescribeAccountAuthorityResponseBody
	GetData() *DescribeAccountAuthorityResponseBodyData
	SetRequestId(v string) *DescribeAccountAuthorityResponseBody
	GetRequestId() *string
}

type DescribeAccountAuthorityResponseBody struct {
	// The returned result.
	Data *DescribeAccountAuthorityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBody) GetData() *DescribeAccountAuthorityResponseBodyData {
	return s.Data
}

func (s *DescribeAccountAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountAuthorityResponseBody) SetData(v *DescribeAccountAuthorityResponseBodyData) *DescribeAccountAuthorityResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetRequestId(v string) *DescribeAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountAuthorityResponseBodyData struct {
	// The name of the database account.
	//
	// example:
	//
	// test1
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The databases on which permissions are granted.
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// The dictionaries on which permissions are granted.
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the DDL permissions are granted to the database account. Valid values:
	//
	// 	- **true**: The account has the permissions to execute DDL statements.
	//
	// 	- **false**: The account does not have the permissions to execute DDL statements.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Indicates whether the DML permissions are granted to the database account. Valid values:
	//
	// 	- 0: The account has the permissions to read data from the database, write data to the database, and modify the settings of the database.
	//
	// 	- 1: The account only has the permissions to read data from the database.
	//
	// 	- 2: The account only has the permissions to read data from the database and modify the settings of the database.
	//
	// example:
	//
	// 0
	DmlAuthority *int32 `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	// All databases.
	TotalDatabases []*string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty" type:"Repeated"`
	// The database.
	TotalDictionaries []*string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty" type:"Repeated"`
}

func (s DescribeAccountAuthorityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *DescribeAccountAuthorityResponseBodyData) GetAllowDatabases() []*string {
	return s.AllowDatabases
}

func (s *DescribeAccountAuthorityResponseBodyData) GetAllowDictionaries() []*string {
	return s.AllowDictionaries
}

func (s *DescribeAccountAuthorityResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAccountAuthorityResponseBodyData) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *DescribeAccountAuthorityResponseBodyData) GetDmlAuthority() *int32 {
	return s.DmlAuthority
}

func (s *DescribeAccountAuthorityResponseBodyData) GetTotalDatabases() []*string {
	return s.TotalDatabases
}

func (s *DescribeAccountAuthorityResponseBodyData) GetTotalDictionaries() []*string {
	return s.TotalDictionaries
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAccount(v string) *DescribeAccountAuthorityResponseBodyData {
	s.Account = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAllowDatabases(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.AllowDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetAllowDictionaries(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.AllowDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDBInstanceId(v string) *DescribeAccountAuthorityResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDdlAuthority(v bool) *DescribeAccountAuthorityResponseBodyData {
	s.DdlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetDmlAuthority(v int32) *DescribeAccountAuthorityResponseBodyData {
	s.DmlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetTotalDatabases(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.TotalDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) SetTotalDictionaries(v []*string) *DescribeAccountAuthorityResponseBodyData {
	s.TotalDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBodyData) Validate() error {
	return dara.Validate(s)
}
