// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupQueryUserListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DsgUserGroupQueryUserListResponseBodyData) *DsgUserGroupQueryUserListResponseBody
	GetData() []*DsgUserGroupQueryUserListResponseBodyData
	SetErrorCode(v string) *DsgUserGroupQueryUserListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUserGroupQueryUserListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUserGroupQueryUserListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgUserGroupQueryUserListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUserGroupQueryUserListResponseBody
	GetSuccess() *bool
}

type DsgUserGroupQueryUserListResponseBody struct {
	// The returned result.
	Data []*DsgUserGroupQueryUserListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgUserGroupQueryUserListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryUserListResponseBody) GetData() []*DsgUserGroupQueryUserListResponseBodyData {
	return s.Data
}

func (s *DsgUserGroupQueryUserListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUserGroupQueryUserListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUserGroupQueryUserListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUserGroupQueryUserListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUserGroupQueryUserListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUserGroupQueryUserListResponseBody) SetData(v []*DsgUserGroupQueryUserListResponseBodyData) *DsgUserGroupQueryUserListResponseBody {
	s.Data = v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) SetErrorCode(v string) *DsgUserGroupQueryUserListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) SetErrorMessage(v string) *DsgUserGroupQueryUserListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) SetHttpStatusCode(v int32) *DsgUserGroupQueryUserListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) SetRequestId(v string) *DsgUserGroupQueryUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) SetSuccess(v bool) *DsgUserGroupQueryUserListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DsgUserGroupQueryUserListResponseBodyData struct {
	// The name of the user.
	//
	// example:
	//
	// user1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the user. Valid values:
	//
	// 	- 1 to 5: Alibaba Cloud account
	//
	// 	- 6: RAM role
	//
	// example:
	//
	// 2
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The user ID or role ID.
	//
	// example:
	//
	// 234
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// The parent user ID. This parameter is required if a RAM user is used.
	//
	// example:
	//
	// 123
	ParentAccountId *string `json:"ParentAccountId,omitempty" xml:"ParentAccountId,omitempty"`
	// The name of the Alibaba Cloud account or RAM role. The return value of this parameter must be used if the user group is created by using an Alibaba Cloud account or a RAM role.
	//
	// example:
	//
	// user:user1
	YunAccount *string `json:"YunAccount,omitempty" xml:"YunAccount,omitempty"`
}

func (s DsgUserGroupQueryUserListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryUserListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryUserListResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DsgUserGroupQueryUserListResponseBodyData) GetAccountType() *int32 {
	return s.AccountType
}

func (s *DsgUserGroupQueryUserListResponseBodyData) GetBaseId() *string {
	return s.BaseId
}

func (s *DsgUserGroupQueryUserListResponseBodyData) GetParentAccountId() *string {
	return s.ParentAccountId
}

func (s *DsgUserGroupQueryUserListResponseBodyData) GetYunAccount() *string {
	return s.YunAccount
}

func (s *DsgUserGroupQueryUserListResponseBodyData) SetAccountName(v string) *DsgUserGroupQueryUserListResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBodyData) SetAccountType(v int32) *DsgUserGroupQueryUserListResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBodyData) SetBaseId(v string) *DsgUserGroupQueryUserListResponseBodyData {
	s.BaseId = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBodyData) SetParentAccountId(v string) *DsgUserGroupQueryUserListResponseBodyData {
	s.ParentAccountId = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBodyData) SetYunAccount(v string) *DsgUserGroupQueryUserListResponseBodyData {
	s.YunAccount = &v
	return s
}

func (s *DsgUserGroupQueryUserListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
