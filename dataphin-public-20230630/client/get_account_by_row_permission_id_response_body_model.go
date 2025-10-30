// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountByRowPermissionIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAccountByRowPermissionIdResponseBody
	GetCode() *string
	SetData(v []*GetAccountByRowPermissionIdResponseBodyData) *GetAccountByRowPermissionIdResponseBody
	GetData() []*GetAccountByRowPermissionIdResponseBodyData
	SetHttpStatusCode(v int32) *GetAccountByRowPermissionIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAccountByRowPermissionIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAccountByRowPermissionIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAccountByRowPermissionIdResponseBody
	GetSuccess() *bool
}

type GetAccountByRowPermissionIdResponseBody struct {
	// example:
	//
	// OK
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetAccountByRowPermissionIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountByRowPermissionIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAccountByRowPermissionIdResponseBody) GetData() []*GetAccountByRowPermissionIdResponseBodyData {
	return s.Data
}

func (s *GetAccountByRowPermissionIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAccountByRowPermissionIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAccountByRowPermissionIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountByRowPermissionIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAccountByRowPermissionIdResponseBody) SetCode(v string) *GetAccountByRowPermissionIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) SetData(v []*GetAccountByRowPermissionIdResponseBodyData) *GetAccountByRowPermissionIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) SetHttpStatusCode(v int32) *GetAccountByRowPermissionIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) SetMessage(v string) *GetAccountByRowPermissionIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) SetRequestId(v string) *GetAccountByRowPermissionIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) SetSuccess(v bool) *GetAccountByRowPermissionIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBody) Validate() error {
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

type GetAccountByRowPermissionIdResponseBodyData struct {
	// example:
	//
	// 300001235
	Id              *int64                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	UserMappingList []*GetAccountByRowPermissionIdResponseBodyDataUserMappingList `json:"UserMappingList,omitempty" xml:"UserMappingList,omitempty" type:"Repeated"`
}

func (s GetAccountByRowPermissionIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetAccountByRowPermissionIdResponseBodyData) GetUserMappingList() []*GetAccountByRowPermissionIdResponseBodyDataUserMappingList {
	return s.UserMappingList
}

func (s *GetAccountByRowPermissionIdResponseBodyData) SetId(v int64) *GetAccountByRowPermissionIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyData) SetUserMappingList(v []*GetAccountByRowPermissionIdResponseBodyDataUserMappingList) *GetAccountByRowPermissionIdResponseBodyData {
	s.UserMappingList = v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyData) Validate() error {
	if s.UserMappingList != nil {
		for _, item := range s.UserMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAccountByRowPermissionIdResponseBodyDataUserMappingList struct {
	// example:
	//
	// PERSONAL
	AccountType *string                                                               `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Accounts    []*GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s GetAccountByRowPermissionIdResponseBodyDataUserMappingList) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdResponseBodyDataUserMappingList) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingList) GetAccountType() *string {
	return s.AccountType
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingList) GetAccounts() []*GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts {
	return s.Accounts
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingList) SetAccountType(v string) *GetAccountByRowPermissionIdResponseBodyDataUserMappingList {
	s.AccountType = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingList) SetAccounts(v []*GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) *GetAccountByRowPermissionIdResponseBodyDataUserMappingList {
	s.Accounts = v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingList) Validate() error {
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

type GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts struct {
	// example:
	//
	// 300901111
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) SetAccountId(v string) *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts {
	s.AccountId = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) SetAccountName(v string) *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts {
	s.AccountName = &v
	return s
}

func (s *GetAccountByRowPermissionIdResponseBodyDataUserMappingListAccounts) Validate() error {
	return dara.Validate(s)
}
