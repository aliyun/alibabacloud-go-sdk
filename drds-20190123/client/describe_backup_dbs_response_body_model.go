// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbNames(v *DescribeBackupDbsResponseBodyDbNames) *DescribeBackupDbsResponseBody
	GetDbNames() *DescribeBackupDbsResponseBodyDbNames
	SetRequestId(v string) *DescribeBackupDbsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupDbsResponseBody
	GetSuccess() *bool
}

type DescribeBackupDbsResponseBody struct {
	// The details about a database.
	DbNames *DescribeBackupDbsResponseBodyDbNames `json:"DbNames,omitempty" xml:"DbNames,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 842DFA7F-B09B-42A2-B115-E684AE******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBody) GetDbNames() *DescribeBackupDbsResponseBodyDbNames {
	return s.DbNames
}

func (s *DescribeBackupDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupDbsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupDbsResponseBody) SetDbNames(v *DescribeBackupDbsResponseBodyDbNames) *DescribeBackupDbsResponseBody {
	s.DbNames = v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetRequestId(v string) *DescribeBackupDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) SetSuccess(v bool) *DescribeBackupDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupDbsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupDbsResponseBodyDbNames struct {
	DbName []*string `json:"dbName,omitempty" xml:"dbName,omitempty" type:"Repeated"`
}

func (s DescribeBackupDbsResponseBodyDbNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDbsResponseBodyDbNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupDbsResponseBodyDbNames) GetDbName() []*string {
	return s.DbName
}

func (s *DescribeBackupDbsResponseBodyDbNames) SetDbName(v []*string) *DescribeBackupDbsResponseBodyDbNames {
	s.DbName = v
	return s
}

func (s *DescribeBackupDbsResponseBodyDbNames) Validate() error {
	return dara.Validate(s)
}
