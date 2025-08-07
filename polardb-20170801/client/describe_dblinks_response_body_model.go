// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLinksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBLinksResponseBody
	GetDBInstanceName() *string
	SetDBLinkInfos(v []*DescribeDBLinksResponseBodyDBLinkInfos) *DescribeDBLinksResponseBody
	GetDBLinkInfos() []*DescribeDBLinksResponseBodyDBLinkInfos
	SetRequestId(v string) *DescribeDBLinksResponseBody
	GetRequestId() *string
}

type DescribeDBLinksResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-a*************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Details about the database links.
	DBLinkInfos []*DescribeDBLinksResponseBodyDBLinkInfos `json:"DBLinkInfos,omitempty" xml:"DBLinkInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 45D24263-7E3A-4140-9472-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBLinksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBLinksResponseBody) GetDBLinkInfos() []*DescribeDBLinksResponseBodyDBLinkInfos {
	return s.DBLinkInfos
}

func (s *DescribeDBLinksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBLinksResponseBody) SetDBInstanceName(v string) *DescribeDBLinksResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBLinksResponseBody) SetDBLinkInfos(v []*DescribeDBLinksResponseBodyDBLinkInfos) *DescribeDBLinksResponseBody {
	s.DBLinkInfos = v
	return s
}

func (s *DescribeDBLinksResponseBody) SetRequestId(v string) *DescribeDBLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBLinksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBLinksResponseBodyDBLinkInfos struct {
	// The ID of the source cluster that the database link connects.
	//
	// example:
	//
	// pc-a*************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the database link.
	//
	// example:
	//
	// dblink_test
	DBLinkName *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
	// The name of the source database of the database link.
	//
	// example:
	//
	// testdb1
	SourceDBName *string `json:"SourceDBName,omitempty" xml:"SourceDBName,omitempty"`
	// The account of the destination database of the database link.
	//
	// example:
	//
	// testacc
	TargetAccount *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
	// The ID of the destination cluster that the database link connects.
	//
	// > If the destination cluster is not a PolarDB for Oracle cluster, the returned value is empty.
	//
	// example:
	//
	// pc-b************
	TargetDBInstanceName *string `json:"TargetDBInstanceName,omitempty" xml:"TargetDBInstanceName,omitempty"`
	// The name of the destination database of the database link.
	//
	// example:
	//
	// testdb2
	TargetDBName *string `json:"TargetDBName,omitempty" xml:"TargetDBName,omitempty"`
}

func (s DescribeDBLinksResponseBodyDBLinkInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLinksResponseBodyDBLinkInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetDBLinkName() *string {
	return s.DBLinkName
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetSourceDBName() *string {
	return s.SourceDBName
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetTargetAccount() *string {
	return s.TargetAccount
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetTargetDBInstanceName() *string {
	return s.TargetDBInstanceName
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) GetTargetDBName() *string {
	return s.TargetDBName
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetDBInstanceName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetDBLinkName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.DBLinkName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetSourceDBName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.SourceDBName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetAccount(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetAccount = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetDBInstanceName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetDBInstanceName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetDBName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetDBName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) Validate() error {
	return dara.Validate(s)
}
