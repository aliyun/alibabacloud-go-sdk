// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeDbListResponseBodyData) *DescribeDbListResponseBody
	GetData() []*DescribeDbListResponseBodyData
	SetMessage(v string) *DescribeDbListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDbListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDbListResponseBody
	GetSuccess() *bool
}

type DescribeDbListResponseBody struct {
	Data []*DescribeDbListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDbListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBody) GetData() []*DescribeDbListResponseBodyData {
	return s.Data
}

func (s *DescribeDbListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDbListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDbListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDbListResponseBody) SetData(v []*DescribeDbListResponseBodyData) *DescribeDbListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDbListResponseBody) SetMessage(v string) *DescribeDbListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDbListResponseBody) SetRequestId(v string) *DescribeDbListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbListResponseBody) SetSuccess(v bool) *DescribeDbListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDbListResponseBody) Validate() error {
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

type DescribeDbListResponseBodyData struct {
	Accounts []*DescribeDbListResponseBodyDataAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// utf8mb4
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// example:
	//
	// test
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// test
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeDbListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyData) GetAccounts() []*DescribeDbListResponseBodyDataAccounts {
	return s.Accounts
}

func (s *DescribeDbListResponseBodyData) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *DescribeDbListResponseBodyData) GetDBDescription() *string {
	return s.DBDescription
}

func (s *DescribeDbListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDbListResponseBodyData) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDbListResponseBodyData) SetAccounts(v []*DescribeDbListResponseBodyDataAccounts) *DescribeDbListResponseBodyData {
	s.Accounts = v
	return s
}

func (s *DescribeDbListResponseBodyData) SetCharacterSetName(v string) *DescribeDbListResponseBodyData {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetDBDescription(v string) *DescribeDbListResponseBodyData {
	s.DBDescription = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetDBInstanceName(v string) *DescribeDbListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) SetDBName(v string) *DescribeDbListResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DescribeDbListResponseBodyData) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDbListResponseBodyDataAccounts struct {
	// example:
	//
	// root4test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s DescribeDbListResponseBodyDataAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbListResponseBodyDataAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponseBodyDataAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDbListResponseBodyDataAccounts) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountName(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountName = &v
	return s
}

func (s *DescribeDbListResponseBodyDataAccounts) SetAccountPrivilege(v string) *DescribeDbListResponseBodyDataAccounts {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDbListResponseBodyDataAccounts) Validate() error {
	return dara.Validate(s)
}
