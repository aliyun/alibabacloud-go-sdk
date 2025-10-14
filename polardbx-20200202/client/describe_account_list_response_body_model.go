// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAccountListResponseBodyData) *DescribeAccountListResponseBody
	GetData() []*DescribeAccountListResponseBodyData
	SetMessage(v string) *DescribeAccountListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAccountListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAccountListResponseBody
	GetSuccess() *bool
}

type DescribeAccountListResponseBody struct {
	Data []*DescribeAccountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAccountListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBody) GetData() []*DescribeAccountListResponseBodyData {
	return s.Data
}

func (s *DescribeAccountListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAccountListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAccountListResponseBody) SetData(v []*DescribeAccountListResponseBodyData) *DescribeAccountListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountListResponseBody) SetMessage(v string) *DescribeAccountListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAccountListResponseBody) SetRequestId(v string) *DescribeAccountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountListResponseBody) SetSuccess(v bool) *DescribeAccountListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAccountListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountListResponseBodyData struct {
	// example:
	//
	// testaccount desc
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// example:
	//
	// testAccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// example:
	//
	// 0
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 2012-06-08T15:00Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
}

func (s DescribeAccountListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountListResponseBodyData) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *DescribeAccountListResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeAccountListResponseBodyData) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeAccountListResponseBodyData) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeAccountListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAccountListResponseBodyData) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountListResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAccountListResponseBodyData) SetAccountDescription(v string) *DescribeAccountListResponseBodyData {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountName(v string) *DescribeAccountListResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountPrivilege(v string) *DescribeAccountListResponseBodyData {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetAccountType(v string) *DescribeAccountListResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetDBInstanceName(v string) *DescribeAccountListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetDBName(v string) *DescribeAccountListResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) SetGmtCreated(v string) *DescribeAccountListResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAccountListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
