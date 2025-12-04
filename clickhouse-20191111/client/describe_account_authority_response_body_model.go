// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DescribeAccountAuthorityResponseBody
	GetAccountName() *string
	SetAllowDatabases(v []*string) *DescribeAccountAuthorityResponseBody
	GetAllowDatabases() []*string
	SetAllowDictionaries(v []*string) *DescribeAccountAuthorityResponseBody
	GetAllowDictionaries() []*string
	SetDdlAuthority(v bool) *DescribeAccountAuthorityResponseBody
	GetDdlAuthority() *bool
	SetDmlAuthority(v string) *DescribeAccountAuthorityResponseBody
	GetDmlAuthority() *string
	SetRequestId(v string) *DescribeAccountAuthorityResponseBody
	GetRequestId() *string
	SetTotalDatabases(v []*string) *DescribeAccountAuthorityResponseBody
	GetTotalDatabases() []*string
	SetTotalDictionaries(v []*string) *DescribeAccountAuthorityResponseBody
	GetTotalDictionaries() []*string
}

type DescribeAccountAuthorityResponseBody struct {
	// The name of the database account.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Databases to which permissions have been granted.
	AllowDatabases []*string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty" type:"Repeated"`
	// Dictionaries to which permissions have been granted.
	AllowDictionaries []*string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty" type:"Repeated"`
	// Indicates whether the database account has DDL permissions. Valid values:
	//
	// 	- **true**: has DDL permissions.
	//
	// 	- **false**: does not have DDL permissions.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Indicates whether the database account has DML permissions. Valid values:
	//
	// 	- **all**
	//
	// 	- **readOnly,modify**
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// All databases.
	TotalDatabases []*string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty" type:"Repeated"`
	// All dictionaries.
	TotalDictionaries []*string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty" type:"Repeated"`
}

func (s DescribeAccountAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponseBody) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountAuthorityResponseBody) GetAllowDatabases() []*string {
	return s.AllowDatabases
}

func (s *DescribeAccountAuthorityResponseBody) GetAllowDictionaries() []*string {
	return s.AllowDictionaries
}

func (s *DescribeAccountAuthorityResponseBody) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *DescribeAccountAuthorityResponseBody) GetDmlAuthority() *string {
	return s.DmlAuthority
}

func (s *DescribeAccountAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountAuthorityResponseBody) GetTotalDatabases() []*string {
	return s.TotalDatabases
}

func (s *DescribeAccountAuthorityResponseBody) GetTotalDictionaries() []*string {
	return s.TotalDictionaries
}

func (s *DescribeAccountAuthorityResponseBody) SetAccountName(v string) *DescribeAccountAuthorityResponseBody {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetAllowDatabases(v []*string) *DescribeAccountAuthorityResponseBody {
	s.AllowDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetAllowDictionaries(v []*string) *DescribeAccountAuthorityResponseBody {
	s.AllowDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetDdlAuthority(v bool) *DescribeAccountAuthorityResponseBody {
	s.DdlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetDmlAuthority(v string) *DescribeAccountAuthorityResponseBody {
	s.DmlAuthority = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetRequestId(v string) *DescribeAccountAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetTotalDatabases(v []*string) *DescribeAccountAuthorityResponseBody {
	s.TotalDatabases = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) SetTotalDictionaries(v []*string) *DescribeAccountAuthorityResponseBody {
	s.TotalDictionaries = v
	return s
}

func (s *DescribeAccountAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
