// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindProduceAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BindProduceAuthorizationResponseBodyData) *BindProduceAuthorizationResponseBody
	GetData() *BindProduceAuthorizationResponseBodyData
	SetErrorCode(v string) *BindProduceAuthorizationResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *BindProduceAuthorizationResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *BindProduceAuthorizationResponseBody
	GetRequestId() *string
}

type BindProduceAuthorizationResponseBody struct {
	Data *BindProduceAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// AC492C5D-29D0-5103-9271-2C3A9D99F5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindProduceAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBody) GetData() *BindProduceAuthorizationResponseBodyData {
	return s.Data
}

func (s *BindProduceAuthorizationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BindProduceAuthorizationResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BindProduceAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindProduceAuthorizationResponseBody) SetData(v *BindProduceAuthorizationResponseBodyData) *BindProduceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetErrorCode(v string) *BindProduceAuthorizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetErrorMsg(v string) *BindProduceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BindProduceAuthorizationResponseBody) SetRequestId(v string) *BindProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindProduceAuthorizationResponseBody) Validate() error {
	return dara.Validate(s)
}

type BindProduceAuthorizationResponseBodyData struct {
	AuthorizedUserList []*BindProduceAuthorizationResponseBodyDataAuthorizedUserList `json:"AuthorizedUserList,omitempty" xml:"AuthorizedUserList,omitempty" type:"Repeated"`
	Message            *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindProduceAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindProduceAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBodyData) GetAuthorizedUserList() []*BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	return s.AuthorizedUserList
}

func (s *BindProduceAuthorizationResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *BindProduceAuthorizationResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *BindProduceAuthorizationResponseBodyData) SetAuthorizedUserList(v []*BindProduceAuthorizationResponseBodyDataAuthorizedUserList) *BindProduceAuthorizationResponseBodyData {
	s.AuthorizedUserList = v
	return s
}

func (s *BindProduceAuthorizationResponseBodyData) SetMessage(v string) *BindProduceAuthorizationResponseBodyData {
	s.Message = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyData) SetSuccess(v bool) *BindProduceAuthorizationResponseBodyData {
	s.Success = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type BindProduceAuthorizationResponseBodyDataAuthorizedUserList struct {
	// example:
	//
	// 2
	AccountValidType *int32 `json:"AccountValidType,omitempty" xml:"AccountValidType,omitempty"`
	// example:
	//
	// 1219541161213058
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// test@alibaba-inc.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BindProduceAuthorizationResponseBodyDataAuthorizedUserList) String() string {
	return dara.Prettify(s)
}

func (s BindProduceAuthorizationResponseBodyDataAuthorizedUserList) GoString() string {
	return s.String()
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) GetAccountValidType() *int32 {
	return s.AccountValidType
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) GetUserId() *string {
	return s.UserId
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) GetUserName() *string {
	return s.UserName
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetAccountValidType(v int32) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.AccountValidType = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetUserId(v string) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.UserId = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) SetUserName(v string) *BindProduceAuthorizationResponseBodyDataAuthorizedUserList {
	s.UserName = &v
	return s
}

func (s *BindProduceAuthorizationResponseBodyDataAuthorizedUserList) Validate() error {
	return dara.Validate(s)
}
