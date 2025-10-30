// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceProjectAddableUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceProjectAddableUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceProjectAddableUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceProjectAddableUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataServiceProjectAddableUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceProjectAddableUsersResponseBody
	GetSuccess() *bool
	SetUserList(v []*GetDataServiceProjectAddableUsersResponseBodyUserList) *GetDataServiceProjectAddableUsersResponseBody
	GetUserList() []*GetDataServiceProjectAddableUsersResponseBodyUserList
}

type GetDataServiceProjectAddableUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	UserList  []*GetDataServiceProjectAddableUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetDataServiceProjectAddableUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceProjectAddableUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceProjectAddableUsersResponseBody) GetUserList() []*GetDataServiceProjectAddableUsersResponseBodyUserList {
	return s.UserList
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetCode(v string) *GetDataServiceProjectAddableUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetHttpStatusCode(v int32) *GetDataServiceProjectAddableUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetMessage(v string) *GetDataServiceProjectAddableUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetRequestId(v string) *GetDataServiceProjectAddableUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetSuccess(v bool) *GetDataServiceProjectAddableUsersResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) SetUserList(v []*GetDataServiceProjectAddableUsersResponseBodyUserList) *GetDataServiceProjectAddableUsersResponseBody {
	s.UserList = v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBody) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataServiceProjectAddableUsersResponseBodyUserList struct {
	// example:
	//
	// xx@aliyuncs.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 30012011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDataServiceProjectAddableUsersResponseBodyUserList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceProjectAddableUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) GetAccountName() *string {
	return s.AccountName
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) GetUserId() *string {
	return s.UserId
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) SetAccountName(v string) *GetDataServiceProjectAddableUsersResponseBodyUserList {
	s.AccountName = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) SetDisplayName(v string) *GetDataServiceProjectAddableUsersResponseBodyUserList {
	s.DisplayName = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) SetUserId(v string) *GetDataServiceProjectAddableUsersResponseBodyUserList {
	s.UserId = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponseBodyUserList) Validate() error {
	return dara.Validate(s)
}
