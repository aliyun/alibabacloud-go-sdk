// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPrivilege(v string) *DescribeDatabaseResponseBody
	GetAccessPrivilege() *string
	SetCharacterSetName(v string) *DescribeDatabaseResponseBody
	GetCharacterSetName() *string
	SetCollate(v string) *DescribeDatabaseResponseBody
	GetCollate() *string
	SetConnLimit(v string) *DescribeDatabaseResponseBody
	GetConnLimit() *string
	SetCtype(v string) *DescribeDatabaseResponseBody
	GetCtype() *string
	SetDBInstanceId(v string) *DescribeDatabaseResponseBody
	GetDBInstanceId() *string
	SetDatabaseName(v string) *DescribeDatabaseResponseBody
	GetDatabaseName() *string
	SetDescription(v string) *DescribeDatabaseResponseBody
	GetDescription() *string
	SetOwner(v string) *DescribeDatabaseResponseBody
	GetOwner() *string
	SetRequestId(v string) *DescribeDatabaseResponseBody
	GetRequestId() *string
	SetSize(v string) *DescribeDatabaseResponseBody
	GetSize() *string
	SetTableSpace(v string) *DescribeDatabaseResponseBody
	GetTableSpace() *string
}

type DescribeDatabaseResponseBody struct {
	// example:
	//
	// =Tc/test_001\\ntest_001=CTc/test_001
	AccessPrivilege *string `json:"AccessPrivilege,omitempty" xml:"AccessPrivilege,omitempty"`
	// example:
	//
	// UTF8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// example:
	//
	// en_US.utf8
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// example:
	//
	// -1
	ConnLimit *string `json:"ConnLimit,omitempty" xml:"ConnLimit,omitempty"`
	// example:
	//
	// en_US.utf8
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// test01
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// example:
	//
	// A user management database storing customer profiles.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// testadb
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2000 MB
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// pg_default
	TableSpace *string `json:"TableSpace,omitempty" xml:"TableSpace,omitempty"`
}

func (s DescribeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseResponseBody) GetAccessPrivilege() *string {
	return s.AccessPrivilege
}

func (s *DescribeDatabaseResponseBody) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *DescribeDatabaseResponseBody) GetCollate() *string {
	return s.Collate
}

func (s *DescribeDatabaseResponseBody) GetConnLimit() *string {
	return s.ConnLimit
}

func (s *DescribeDatabaseResponseBody) GetCtype() *string {
	return s.Ctype
}

func (s *DescribeDatabaseResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDatabaseResponseBody) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDatabaseResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDatabaseResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *DescribeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseResponseBody) GetSize() *string {
	return s.Size
}

func (s *DescribeDatabaseResponseBody) GetTableSpace() *string {
	return s.TableSpace
}

func (s *DescribeDatabaseResponseBody) SetAccessPrivilege(v string) *DescribeDatabaseResponseBody {
	s.AccessPrivilege = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetCharacterSetName(v string) *DescribeDatabaseResponseBody {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetCollate(v string) *DescribeDatabaseResponseBody {
	s.Collate = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetConnLimit(v string) *DescribeDatabaseResponseBody {
	s.ConnLimit = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetCtype(v string) *DescribeDatabaseResponseBody {
	s.Ctype = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetDBInstanceId(v string) *DescribeDatabaseResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetDatabaseName(v string) *DescribeDatabaseResponseBody {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetDescription(v string) *DescribeDatabaseResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetOwner(v string) *DescribeDatabaseResponseBody {
	s.Owner = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetRequestId(v string) *DescribeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetSize(v string) *DescribeDatabaseResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeDatabaseResponseBody) SetTableSpace(v string) *DescribeDatabaseResponseBody {
	s.TableSpace = &v
	return s
}

func (s *DescribeDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
